package main

import (
	"bufio"
	"fmt"
	"github.com/pknetwork"
	"net"
	"os"
	"strings"
)

type client struct {
	conn        net.Conn
	inputReader *bufio.Reader
}

// 处理函数
func process(conn net.Conn) {
	defer conn.Close()
	client := client{conn: conn, inputReader: bufio.NewReader(os.Stdin)}
	fmt.Println("What is your name?")
	step(client)
	client.play()
}

func (client client) play() {
	for {
		fmt.Println("press f to draw a fortuneCard")
		step(client)

		fmt.Println("press r to start roll")
		step(client)

		fmt.Println("press s to score this turn")
		step(client)
	}

}

func step(client client) error {
	input, _ := client.inputReader.ReadString('\n') // 读取用户输入
	inputInfo := strings.Trim(input, "\r\n")

	pknetwork.WriteMessage(client.conn, inputInfo)
	msg, err := pknetwork.ReadMessage(client.conn)
	if err != nil {
		return err
	}
	fmt.Println(msg)
	return nil
}

func main() {
	pknetwork.StartClient(process)
}
