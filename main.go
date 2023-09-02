package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"flag"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/01node/batch-validators-deposit/contracts/MultiDeposit"
	util "github.com/01node/batch-validators-deposit/internal/utils"
	"github.com/shopspring/decimal"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var pk string
var rpc string
var chainid int64
var file string
var gasprice int64
var gaslimit int64
var chunksize int
var gasmultiplier float64
var multideposit string
var multiDepositAddress common.Address
var privateKey *ecdsa.PrivateKey
var walletAddress common.Address
var client *ethclient.Client
var multiDepositContract *MultiDeposit.MultiDeposit
var txs []string

func init() {
	flag.StringVar(&pk, "private-key", "", "Private key")
	flag.StringVar(&rpc, "rpc", "", "rpc")
	flag.Int64Var(&chainid, "chain-id", 0, "chainid")
	flag.StringVar(&file, "file", "", "File")
	flag.Int64Var(&gasprice, "gas-price", 0, "Gas price")
	flag.Int64Var(&gaslimit, "gas-limit", 0, "Gas limit")
	flag.Float64Var(&gasmultiplier, "gas-multiplier", 0, "Gas multiplier")
	flag.StringVar(&multideposit, "multi-deposit", "", "MultiDeposit address")
	flag.IntVar(&chunksize, "chunk-size", 0, "Chunk size")
	flag.Parse()

	if pk == "" {
		log.Fatal("private-key is required")
	}

	if rpc == "" {
		log.Fatal("rpc is required")
	}

	if chainid == 0 {
		log.Fatal("chain-id is required")
	}

	if file == "" {
		log.Fatal("file is required")
	}

	if multideposit == "" {
		log.Fatal("multi-deposit is required")
	}

	if chunksize == 0 {
		chunksize = 100
	}
}

func main() {
	log.Println("Starting service ... ")

	var err error

	log.Println("Loading private key ... ")

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	walletAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Print("Wallet Address: ", walletAddress.Hex())

	multiDepositAddress = common.HexToAddress(multideposit)
	log.Print("MultiDeposit Address: ", multiDepositAddress.Hex())

	log.Print("Loading file: ", file)
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatalln("Error reading file", err)
	}

	log.Println("Initializing ethereum client ...")

	client, err = ethclient.Dial(rpc)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Ethereum client initialized.")

	// Instantiate the contract
	multiDepositContract, err = MultiDeposit.NewMultiDeposit(multiDepositAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate MultiDeposit contract: %v", err)
	}

	log.Println("MultiDeposit contract loaded.")

	// define struct for file contents
	type DepositData struct {
		Pubkey                string `json:"pubkey"`
		WithdrawalCredentials string `json:"withdrawal_credentials"`
		Amount                int64  `json:"amount"`
		Signature             string `json:"signature"`
		DepositMessageRoot    string `json:"deposit_message_root"`
		DepositDataRoot       string `json:"deposit_data_root"`
		ForkVersion           string `json:"fork_version"`
		NetworkName           string `json:"network_name"`
		DepositCliVersion     string `json:"deposit_cli_version"`
	}

	var fileContents []DepositData

	err = json.Unmarshal(content, &fileContents)
	if err != nil {
		log.Fatalln("Error reading file", err)
	}

	log.Println("File loaded in memory.")

	var chunks [][]DepositData

	if len(content) > chunksize {
		log.Printf("File contains more than %v elements, splitting ...", chunksize)

		for i := 0; i < len(fileContents); i += chunksize {
			end := i + chunksize

			if end > len(fileContents) {
				end = len(fileContents)
			}

			chunks = append(chunks, fileContents[i:end])
		}
	} else {
		chunks = append(chunks, fileContents)
	}

	for index, chunk := range chunks {
		log.Printf("Processing chunk %v out of %v\n", index, len(chunks))

		var publicKeys [][]byte = make([][]byte, len(chunk))
		var withdrawalCredentials [][]byte = make([][]byte, len(chunk))
		var signatures [][]byte = make([][]byte, len(chunk))
		var depositDataRoots [][32]byte = make([][32]byte, len(chunk))

		for i, deposit := range chunk {
			publicKeys[i] = common.Hex2Bytes(deposit.Pubkey)
			withdrawalCredentials[i] = common.Hex2Bytes(deposit.WithdrawalCredentials)
			signatures[i] = common.Hex2Bytes(deposit.Signature)
			depositDataRoots[i] = [32]byte(common.Hex2Bytes(deposit.DepositDataRoot))
		}

		balance, err := client.BalanceAt(context.Background(), walletAddress, nil)
		if err != nil {
			log.Fatal(err)
		}

		log.Print("Wallet Balance: ", balance)

		nonce, err := client.PendingNonceAt(context.Background(), walletAddress)
		if err != nil {
			log.Fatal(err)
		}

		log.Print("Nonce: ", nonce)

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		log.Print("GasPrice Suggestion: ", gasPrice)

		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainid))
		if err != nil {
			log.Fatal(err)
		}

		auth.Nonce = big.NewInt(int64(nonce))
		auth.GasLimit = uint64(15000000) // in units
		auth.GasPrice = gasPrice

		auth.Value = util.ToWei(decimal.NewFromBigInt(big.NewInt(int64(32*len(publicKeys))), 0), 18)

		log.Print("Value:", auth.Value)
		log.Print("PublicKeys:", len(publicKeys))

		if gasmultiplier != 0 {
			auth.GasPrice = big.NewInt(int64(float64(gasPrice.Int64()) * gasmultiplier))
		}

		if gasprice != 0 {
			auth.GasPrice = big.NewInt(gasprice)
		}

		if gaslimit != 0 {
			auth.GasLimit = uint64(gaslimit)
		}

		log.Print("GasPrice: ", auth.GasPrice)
		log.Print("GasLimit: ", auth.GasLimit)

		parsedABI, err := abi.JSON(strings.NewReader(MultiDeposit.MultiDepositMetaData.ABI))
		if err != nil {
			log.Fatal(err)
		}
		inputData, err := parsedABI.Pack("batchDeposit", publicKeys, withdrawalCredentials, signatures, depositDataRoots)

		if err != nil {
			log.Fatal(err)
		}

		// Estimate gas
		msg := ethereum.CallMsg{
			From:  auth.From,
			To:    &multiDepositAddress,
			Data:  inputData,
			Value: auth.Value,
		}

		estimatedGas, err := client.EstimateGas(context.Background(), msg)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Estimated gas: %v\n", estimatedGas)
		auth.GasLimit = estimatedGas

		if gasmultiplier != 0 {
			auth.GasPrice = big.NewInt(int64(float64(gasPrice.Int64()) * gasmultiplier))
			auth.GasLimit = uint64(float64(estimatedGas) * gasmultiplier)
		}

		if gasprice != 0 {
			auth.GasPrice = big.NewInt(gasprice)
		}

		if gaslimit != 0 {
			auth.GasLimit = uint64(gaslimit)
		}

		log.Print("GasPrice: ", auth.GasPrice)
		log.Print("GasLimit: ", auth.GasLimit)

		tx, err := multiDepositContract.BatchDeposit(auth, publicKeys, withdrawalCredentials, signatures, depositDataRoots)

		if err != nil {
			log.Fatal(err)
		}

		log.Println("Tx pending: ", tx.Hash().Hex())

		txs = append(txs, tx.Hash().Hex())

		receipt, err := bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Transaction mined: ", receipt.Status, receipt.GasUsed)

	}
}
