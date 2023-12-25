package tron

import (
	"math"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/shockerli/cvt"
)

type Address string // tron address
type SUN uint64     // unit:SUN
type TRX float64    // unit:TRX

// ---- Address ---- //

func (inst Address) String() string {
	return string(inst)
}

func (inst Address) IsValid() (valid bool) {
	return IsAddressValid(inst.String())
}

func AddressFromStr(addr string) (address Address, err error) {
	address = Address(addr)
	if !address.IsValid() {
		err = errors.BadRequest("ERR_PARAM", "address invalid")
	}

	return
}

// ---- SUN ---- //

func (inst SUN) TRX() TRX {
	return TRX(inst) / TRX(SUN_VALUE)
}

func (inst SUN) Int64() int64 {
	return int64(inst)
}

func (inst SUN) Int64E() (i int64, err error) {
	if i, err = cvt.Int64E(inst); err != nil {
		err = errors.BadRequest("ERR_PARAM", "amount error").WithCause(err)
	}
	return
}

func (inst SUN) Uint64() uint64 {
	return uint64(inst)
}

func (inst SUN) Uint64E() (u uint64, err error) {
	if u, err = cvt.Uint64E(inst); err != nil {
		err = errors.BadRequest("ERR_PARAM", "amount error").WithCause(err)
	}
	return
}

// ---- TRX ---- //

func (inst TRX) Float64() float64 {
	return float64(inst)
}

func (inst TRX) Float64E() (f float64, err error) {
	if f, err = cvt.Float64E(inst); err != nil {
		err = errors.BadRequest("ERR_PARAM", "amount error").WithCause(err)
	}
	return
}

func (inst TRX) SUN() SUN {
	return SUN(inst * SUN_VALUE)
}

func (inst TRX) CeilInt64() int64 {
	return int64(math.Ceil(inst.Float64()))
}

func (inst TRX) CeilUint64() uint64 {
	return uint64(math.Ceil(inst.Float64()))
}

func (inst TRX) Ceil() TRX {
	return TRX(math.Ceil(inst.Float64()))
}
