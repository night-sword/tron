//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package tron

import (
	"github.com/go-resty/resty/v2"
	"github.com/google/wire"
)

func NewGrid(client *resty.Client) *Grid {
	panic(
		wire.Build(
			newApiKeys,
			newGrid,
		),
	)
}
