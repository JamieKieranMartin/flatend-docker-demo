package main

import (
	"os"
	"os/signal"
	"strconv"
	"sync/atomic"

	"github.com/lithdew/flatend"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	counter := uint64(0)

	count := func(ctx *flatend.Context) {
		current := atomic.AddUint64(&counter, 1) - 1
		ctx.Write(strconv.AppendUint(nil, current, 10))
	}

	node := &flatend.Node{
		Services: map[string]flatend.Handler{
			"count": count,
		},
	}

	defer node.Shutdown()

	check(node.Start("127.0.0.1:9000"))

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
