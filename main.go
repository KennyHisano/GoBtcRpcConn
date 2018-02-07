package main

import (
	"fmt"

	rpc "github.com/btcsuite/btcd/rpcclient"
	"flag"
)


var (
	flagset		= flag.NewFlagSet("",flag.ExitOnError)
	connectFlag = flagset.String("s","localhost","host[:port] of rpc server")
	rpcuserFlag = flagset.String("rpcuser","","username for wallet rpc authentication")
	rpcpassFlag = flagset.String("rpcpass","","password for rpc authentication")
	testnetFlag = flagset.Bool("testnet", false, "use testnet network")
)

func main() {

	connConfig := &rpc.ConnConfig{
		Host:		connect,
		User:		*rpcuserFlag,
		Pass:		*rpcpassFlag,
		DisableTLS: true,


	}

}
