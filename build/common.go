package build

import "os"

const (
	ServerURL    = "http://54.251.11.180:8080"
	OPSepolia    = "op-sepolia"
	OPBNBTestnet = "opbnb-testnet"
	BNBTestnet   = "bnb-testnet"
	LocalAnvil   = "local-anvil"
)

func CheckChain() string {
	ct := os.Getenv("CHAIN_TYPE")
	if ct == "" {
		panic("please set env 'CHAIN_TYPE' to 'op-sepolia', 'opbnb-testnet' or 'bnb-testnet'")
	}
	switch ct {
	case OPSepolia, OPBNBTestnet, BNBTestnet:
		return ct
	default:
		panic("please set env 'CHAIN_TYPE' to 'op-sepolia', 'opbnb-testnet' or 'bnb-testnet'")
	}
}
