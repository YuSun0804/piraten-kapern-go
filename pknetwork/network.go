package pknetwork

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func StartClient(callback ClientCallback) {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	callback(conn)
}

type ClientCallback func(conn net.Conn)

func WriteMessage(conn net.Conn, message string) error {
	data, err := encode(message)
	_, err = conn.Write(data) // 发送数据
	if err != nil {
		fmt.Println("encode msg failed, err:", err)
		return err
	}
	return nil
}

func ReadMessage(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	msg, err := decode(reader)
	if err == io.EOF {
		fmt.Println("connection is closed")
		return "", err
	}
	if err != nil {
		fmt.Println("decode msg failed, err:", err)
		return "", err
	}
	return msg, nil
}

type ServerCallback func(conn net.Conn, playerIndex int)

func StartServer(callback ServerCallback) {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	playerIndex := 0
	for {
		conn, err := listen.Accept() // 建立连接
		playerIndex++
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}

		go callback(conn, playerIndex) // 启动一个goroutine处理连接

	}
}
