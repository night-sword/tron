package tron

import (
	"crypto/sha256"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

type transaction struct {
	client *client.GrpcClient
	keys   *privateKey
}

func newTransaction(client *client.GrpcClient, keys *privateKey) *transaction {
	return &transaction{
		client: client,
		keys:   keys,
	}
}

func (inst *transaction) BroadcastWithSign(tx *api.TransactionExtention, operator Address, permissionId int32) (err error) {
	if permissionId > 0 {
		tx.Transaction.RawData.Contract[0].PermissionId = permissionId
		_ = inst.client.UpdateHash(tx)
	}

	err = inst.AppendSign(tx, operator)
	if err != nil {
		err = errors.Wrap(err, "sign fail")
		return
	}

	err = inst.Broadcast(tx)
	return
}

func (inst *transaction) Broadcast(tx *api.TransactionExtention) (err error) {
	result, err := inst.client.Broadcast(tx.Transaction)
	if err != nil {
		return
	}

	if !result.GetResult() {
		err = errors.New(string(result.GetMessage()))
		return
	}

	return
}

func (inst *transaction) Sign(tx *api.TransactionExtention, operator Address) (signature []byte, err error) {
	rawData, err := proto.Marshal(tx.Transaction.GetRawData())
	if err != nil {
		return
	}

	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)

	key, err := inst.keys.Get(operator)
	if err != nil {
		return
	}

	// btcec.PrivKeyFromBytes only returns a secret key and public key
	sk, _ := btcec.PrivKeyFromBytes(key)
	signature, err = crypto.Sign(hash, sk.ToECDSA())
	return
}

func (inst *transaction) AppendSign(tx *api.TransactionExtention, operator Address) (err error) {
	signature, err := inst.Sign(tx, operator)
	if err != nil {
		return
	}

	tx.Transaction.Signature = append(tx.Transaction.Signature, signature)
	return
}
