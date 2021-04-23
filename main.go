package main

import (
	"fmt"
	"github.com/adshao/go-binance/v2"
)

var routines bool = true

func main() {
	go getBTC()
	go getETH()
	should_finish()
}

func getBTC() {
	wsKlineHandler := func(event *binance.WsKlineEvent) {
		fmt.Println("\n",event)
		fmt.Println("\n-----------------------------------------------------------------------------------------------------------------------")

	}
	errHandler := func(err error) {
		fmt.Println(err)
	}

	_, _, err := binance.WsKlineServe("BTCBRL", "1m", wsKlineHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func getETH() {
	wsKlineHandler := func(event *binance.WsKlineEvent) {
		fmt.Println("\n",event)
		fmt.Println("\n-----------------------------------------------------------------------------------------------------------------------")

	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	_, _, err := binance.WsKlineServe("ETHBRL", "1m", wsKlineHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func should_finish() {
	for routines {
		continue
	}
}