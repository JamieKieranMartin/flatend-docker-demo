package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/lithdew/flatend"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()

	node := &flatend.Node{
		SecretKey: flatend.GenerateSecretKey(),
	}
	defer node.Shutdown()

	node.PublicAddr = "10.104.0.3:9000"
	check(node.Start())

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
