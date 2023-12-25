package tron

import (
	"strconv"
	"strings"

	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

const (
	methodIDLength = 8  // method length
	addressLength  = 64 // address length
	valueLength    = 64 // Amount length
	valueLength2   = 62 // Amount length Another unexpected value length

	valueStartIndex = methodIDLength + addressLength               // method + address value start position && Length without value (zero trc20 rotated)
	trc20Length     = methodIDLength + addressLength + valueLength // full length
)

var Parser = newParser()

type parser struct{}

func newParser() *parser {
	return &parser{}
}

func (inst *parser) SmartContractEvent(transaction *core.Transaction) (contractEvent *SmartContractEvent, err error) {
	contracts := transaction.GetRawData().GetContract()
	if len(contracts) == 0 {
		err = errors.New("tx has no contract")
		return
	}

	params := contracts[0].GetParameter()

	message, err := anypb.UnmarshalNew(params, proto.UnmarshalOptions{})
	if err != nil {
		err = errors.Errorf("Error unmarshaling Any message: %v", err)
		return
	}

	tsc, ok := message.(*core.TriggerSmartContract)
	if !ok {
		err = errors.New("msg is not TriggerSmartContract")
		return
	}

	method, to, value, err := inst.ContractInput(common.Bytes2Hex(tsc.GetData()))
	if err != nil {
		err = errors.New("parse contract input fail")
		return
	}

	owner := address.HexToAddress(common.Bytes2Hex(tsc.GetOwnerAddress())).String()
	contract := address.HexToAddress(common.Bytes2Hex(tsc.GetContractAddress())).String()

	contractEvent = &SmartContractEvent{
		Contract: Address(contract),
		Owner:    Address(owner),
		To:       to,
		Amount:   value,
		Method:   method,
	}

	return
}

func (inst *parser) ContractInput(data string) (method ContractMethod, to Address, value SUN, err error) {
	dataLen := len(data)

	if dataLen <= methodIDLength {
		err = errors.New("data length to less")
		return
	}

	method, err = ContractMethodFromStr(data[:methodIDLength])
	if err != nil {
		method, err = ContractMethod(0), nil
	}

	if dataLen >= methodIDLength+addressLength+valueLength {
		addr, v, e := inst.getAddressValueFromData(
			data[methodIDLength:methodIDLength+addressLength],
			data[methodIDLength+addressLength:methodIDLength+addressLength+valueLength],
		)

		to, value, err = Address(addr), SUN(v), e
		return
	}

	if dataLen == methodIDLength+addressLength+valueLength-2 {
		addr, v, e := inst.getAddressValueFromData(
			data[methodIDLength:methodIDLength+addressLength],
			data[methodIDLength+addressLength:]+`00`,
		)

		to, value, err = Address(addr), SUN(v), e
		return
	}
	// Collect other encoding information.
	if dataLen != methodIDLength+addressLength+valueLength &&
		dataLen != methodIDLength+addressLength+valueLength+2 && // Unknown reason for the extra length.
		dataLen != methodIDLength+addressLength+valueLength+4 && // Unknown reason for the extra length.
		dataLen != methodIDLength+addressLength+valueLength+6 && // Unknown reason for the extra length.
		dataLen != methodIDLength+addressLength+valueLength+8 && // Unknown reason for the extra length.
		dataLen != methodIDLength+42 { // Short amount: Method + Address + Amount (missing 2 characters).
		// localLog.Error(`UnpackTransfer original data encoding length not in parsing range`, zap.Int(`length`, dataLen), zap.String(`original data`, data))
		return ContractMethod(0), ``, 0, errors.New("UnpackTransfer original data encoding length error")
	}
	return
}

func (inst *parser) getAddressValueFromData(addressData, valueData string) (address string, value int64, err error) {
	if address, err = inst.getAddressFromData(addressData); err != nil {
		return "", 0, err
	}
	if value, err = inst.getValueFromData(valueData); err != nil {
		return "", 0, err
	}
	return
}

// Parse address from raw data.
func (inst *parser) getAddressFromData(s string) (string, error) {
	s = "41" + s[24:] // Must start with "41".
	addr, err := common.Hex2Bytes(s)
	if err != nil {
		return "", errors.Wrapf(err, "Error parsing address from Hex[%s]", s)
	}
	s = common.EncodeCheck(addr)
	return s, nil
}

// Parse value from raw data.
func (inst *parser) getValueFromData(s string) (int64, error) {
	amount := strings.TrimLeft(s, "0")
	if amount == `` {
		return 0, nil
	}
	value, err := strconv.ParseInt(amount, 16, 64)
	if err != nil {
		return 0, errors.Wrapf(err, "Error parsing value[%s]", s)
	}
	return value, nil
}
