package main

import (
	"fmt"
	"flag"


	rpc "github.com/btcsuite/btcd/rpcclient"

	"net"
	"github.com/btcsuite/btcd/chaincfg"
)



var (
	chainParams = &chaincfg.MainNetParams
)

var (
	flagset		= flag.NewFlagSet("",flag.ExitOnError)
	connectFlag = flagset.String("s","localhost","host[:port] of rpc server")
	rpcuserFlag = flagset.String("rpcuser","","username for wallet rpc authentication")
	rpcpassFlag = flagset.String("rpcpass","","password for rpc authentication")
	testnetFlag = flagset.Bool("testnet", false, "use testnet network")
)

func run() (err error) {

	connect, err :=normalizeAddress(*connectFlag, walletPort(chainParams))
	connConfig := &rpc.ConnConfig{
		Host:		connect,
		User:		*rpcuserFlag,
		Pass:		*rpcpassFlag,
		DisableTLS: true,
		HTTPPostMode: true,

	}
	client, err := rpc.New(connConfig,nil)
	if err != nil {
		return fmt.Errorf("rpc no connect")
		}

	defer func(){
		client.Shutdown()
		client.WaitForShutdown()
	}()

	err = fmt.Errorf("success")
	return err
}

func normalizeAddress(addr string, defaultPort string) (hostport string, err error) {
	host, port, origErr := net.SplitHostPort(addr)
	if origErr == nil {
		return net.JoinHostPort(host, port), nil
	}
	addr = net.JoinHostPort(addr, defaultPort)
	_, _, err = net.SplitHostPort(addr)
	if err != nil {
		return "", origErr
	}
	return addr, nil
}

func walletPort(params *chaincfg.Params) string {
	switch params {
	case &chaincfg.MainNetParams:
		return "8332"
	case &chaincfg.TestNet3Params:
		return "18332"
	default:
		return ""
	}
}
