package main

import (
	"context"
	"fmt"
	"os"
)

type msgSpec interface {
	putMsg()
}

type send struct{}

func (*send) putMsg(msg string) {
	fmt.Println(msg)
}

type newSend struct{}

func (*newSend) putMsg(msg string) {
	fmt.Printf("%s\n", msg)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	app, err := initializeBaz(ctx)
	if err != nil {
		fmt.Println("error")
	}

	fmt.Println(app)

	if len(os.Args) > 1 && os.Args[1] == "1" {
		sendMsg := new(send)
		sendMsg.putMsg("send:hihi")
	} else {
		sendMsg := new(newSend)
		sendMsg.putMsg("newSend:hihi")
	}

	cancel()
}
