package endpoint

import (
	"github.com/ethereum/go-ethereum/ethclient"

	"brave-ox-web/common"
	"brave-ox-web/conf"
	"brave-ox-web/config"
)

var clients = make(map[string]*ethclient.Client)

func Init() {
	n := config.AllChainNode()
	for chain, node := range n {
		dial(chain, node.Rpc)
	}
}

func dial(name string, endpoint string) {
	cli, err := ethclient.Dial(endpoint)
	common.MustNoError(err)
	clients[name] = cli
}

func GetMainChainClient() *ethclient.Client {
	return clients[conf.MainChain()]
}
