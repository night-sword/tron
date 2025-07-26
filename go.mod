module github.com/night-sword/tron

go 1.23.0

require (
	github.com/ethereum/go-ethereum v1.15.6
	github.com/fbsobreira/gotron-sdk v0.0.0-20230907131216-1e824406fe8c
	github.com/go-resty/resty/v2 v2.16.5
	github.com/google/go-querystring v1.1.0
	github.com/google/wire v0.6.0
	github.com/pkg/errors v0.9.1
	github.com/samber/lo v1.51.0
	github.com/shockerli/cvt v0.2.8
	github.com/shopspring/decimal v1.4.0
	google.golang.org/grpc v1.74.2
	google.golang.org/protobuf v1.36.6
)

replace github.com/fbsobreira/gotron-sdk v0.0.0-20230907131216-1e824406fe8c => github.com/night-sword/gotron-sdk v1.1.2

require (
	github.com/btcsuite/btcd/btcec/v2 v2.3.4 // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/holiman/uint256 v1.3.2 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/rjeczalik/notify v0.9.3 // indirect
	github.com/shengdoushi/base58 v1.0.0 // indirect
	github.com/tyler-smith/go-bip39 v1.1.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/crypto v0.38.0 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250528174236-200df99c418a // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250528174236-200df99c418a // indirect
)
