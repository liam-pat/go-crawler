package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string, rpcService interface{}) error {

	err := rpc.Register(rpcService)

	if err != nil {
		return err
	}

	listener, err := net.Listen("tcp", host)

	if err != nil {
		return err
	}
	log.Printf("Listening on %s", host)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept err : %v ", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}

	return nil
}

func NewClient(host string) (*rpc.Client, error) {
	connect, err := net.Dial("tcp", host)

	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(connect), nil
}
