package tron

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
)

var from = Address(os.Getenv("ADDR_FROM"))
var to = Address(os.Getenv("ADDR_TO"))
var key = os.Getenv("PRIVATE_KEY")
var blockNum = os.Getenv("BLOCK_NUM")

var USDT_CONTRACT = Address("TG3XXyExBkPp9nzdajDZsozEu4BkaSJozs")
//var USDT_CONTRACT = Address("TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t")

func Test_Transfer(t *testing.T) {
	wallet := getFullNode()

	txId, err := wallet.Transfer.TransferTRX(from, to, 1, from, 0)
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(txId)
}

func Test_ResourceDelegate(t *testing.T) {
	wallet := getFullNode()

	params := &DelegateResourceParams{
		Owner:    from,
		Receiver: to,
		Resource: Resource_Energy,
		Balance:  1 * SUN_VALUE,
	}

	txId, err := wallet.Resource.Delegate(params, from)
	if err != nil {
		fmt.Println(txId, err)
		return
	}

	fmt.Println(txId)
}

func Test_ResourceUnDelegate(t *testing.T) {
	wallet := getFullNode()

	params := &UnDelegateResourceParams{
		Owner:    from,
		Receiver: to,
		Resource: Resource_Energy,
		Balance:  1 * SUN_VALUE,
	}

	txId, err := wallet.Resource.UnDelegate(params, from)
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(txId)
}

func Test_GetAccountBalance(t *testing.T) {
	wallet := getFullNode()

	balance, err := wallet.Account.GetAccountBalance(from)
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(balance)
	j, _ := json.Marshal(balance)
	fmt.Println(string(j))
}

func Test_Approve(t *testing.T) {
	wallet := getFullNode()

	balance, err := wallet.Transfer.Approve(from, to, USDT_CONTRACT, 10*SUN_VALUE, from, 20*SUN_VALUE, 0)
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(balance)
	j, _ := json.Marshal(balance)
	fmt.Println(string(j))
}

func Test_IsActive(t *testing.T) {
	wallet := getFullNode()

	is, err := wallet.Account.IsAccountActive(from)
	fmt.Println(is, err)
}

func Test_Active(t *testing.T) {
	wallet := getFullNode()

	txId, err := wallet.Account.Active(to, from, from, 0)
	fmt.Println(txId, err)
}

func Test_GetByAccount(t *testing.T) {
	wallet := getFullNode()
	ac, err := wallet.Resource.GetByAccount(from)
	fmt.Println(ac, err)
}

func Test_CalcEnergyFromTRXStake(t *testing.T) {
	wallet := getFullNode()
	ac, err := wallet.Resource.CalcEnergyFromTRXStake(10000)
	fmt.Println(ac, err)
}

func Test_CalcTRXStakeForEnergy(t *testing.T) {
	wallet := getFullNode()
	ac, err := wallet.Resource.CalcTRXStakeForEnergy(65000)
	fmt.Println(ac, err)
}

func Test_Parser_SmartContractEvent(t *testing.T) {
	wallet := getFullNode()
	num, err := strconv.ParseInt(blockNum, 10, 64)
	if err != nil {
		return
	}
	got, err := wallet.Network.GetBlockByNum(num)
	if err != nil {
		return
	}

	sce, err := Parser.SmartContractEvent(got.GetTransactions()[1].GetTransaction())
	fmt.Println(sce, err)
}

func Test_GetCurrentEnergyPrice(t *testing.T) {
	wallet := getFullNode()
	price, err := wallet.Network.GetCurrentEnergyPrice(context.Background())
	fmt.Println(price, err)
}

func Test_GetCurrentBandwidthPrice(t *testing.T) {
	wallet := getFullNode()
	price, err := wallet.Network.GetCurrentBandwidthPrice(context.Background())
	fmt.Println(price, err)
}

func Test_GetTransactionByID(t *testing.T) {
	wallet := getFullNode()
	id := ""
	tx, err := wallet.Network.GetTransactionByID(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(tx.Ret) == 0 {
		fmt.Println("fail")
		return
	}

	if tx.Ret[0].ContractRet == core.Transaction_Result_SUCCESS {
		fmt.Println("success")
		return
	}
	fmt.Println("fail", tx.Ret[0].ContractRet)
}

// GetTransactionById
func Test_Solidity_IsTxConfirmed(t *testing.T) {
	node := getFullNode()
	ok, err := node.Solidity.IsTxConfirmed(context.Background(), "")
	fmt.Println(ok, err)
}

func Test_Account_GetUSDTBalance(t *testing.T) {
	node := getFullNode()
	balance, err := node.Account.GetUSDTBalance(to, USDT_CONTRACT)
	fmt.Println(balance, err)
}

func Test_Transfer_TransferUSDT(t *testing.T) {
	node := getFullNode()
	txID, err := node.Transfer.TransferUSDT(from, to, USDT_CONTRACT, 1000000, from, 100*SUN_VALUE, 0)
	fmt.Println(txID, err)
}

func getFullNode() *FullNode {
	grpcClient, err := NewGrpcClient("grpc.trongrid.io:50051")
	//grpcClient, err := NewGrpcClient("grpc.shasta.trongrid.io:50051") // shasta
	if err != nil {
		panic(err)
	}

	httpClient := NewHttpClient("https://api.trongrid.io")
	//httpClient := NewHttpClient("https://api.shasta.trongrid.io") // shasta

	w := NewFullNode(grpcClient, httpClient)
	err = w.PrivateKey.Append(from, key)
	if err != nil {
		panic(err)
	}

	return w
}
