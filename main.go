/*
CryptoGo
Universidade de Brasília
Departamento de Ciência da Computação
Linguagens de Programação - CIC0093 - 2020/2 B
Prof. Dr. Marcelo Ladeira

Desenvolvido por
	Bruno Sanguinetti Regadas de Barros - 18/0046063
	Caio Bernardon N. K. Massucato - 16/0115001
	Gabriel Nardelli Aprá - 18/0046322
	Gabriel Rodrigues Pacheco - 17/0058280
	João Marcos Melo Monteiro -13/0143031


*/
package main

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
)

var keepAlive bool = true

func main() {
	var (
		apiKey    = "123"
		secretKey = "123"
	)
	client := binance.NewClient(apiKey, secretKey)
	bitcoin := CryptoArray{}
	ticker(client, "ETHUSDT")
	go streamCandle("BTCUSDT", "1m", bitcoin)
	go streamTicker("ETHUSDT")
	commSwitch()
}

// ticker(client, symbol) Recebe como argumento um cliente autenticado e simbolo de tipo string
// e um tempo de tipo string no formato 1m , 5m , 15m, 30m, 1h, 2h, 4h...
// e inicia uma conexão webSocket com o servidor HTTP da Binance, recebendo
// WsKlineEvent struct em formato json por 2000ms
func ticker(client *binance.Client, symbol string) {

	prices, err := client.NewListPricesService().Symbol(symbol).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, p := range prices {
		fmt.Println(p)
	}
}

// streamCandle(symbol, time) Recebe como argumento um simbolo de tipo string
// e um tempo de tipo string no formato 1m , 5m , 15m, 30m, 1h, 2h, 4h...
// e inicia uma conexão webSocket com o servidor HTTP da Binance, recebendo
// WsKlineEvent struct em formato json por 2000ms
func streamCandle(symbol string, time string, coin CryptoArray) {

	streamedPrice := Price{}
	// Handler de sucesso, receberá o stream do servidor em json
	wsKlineHandler := func(event *binance.WsKlineEvent) {
		streamedPrice.open = event.Kline.Open
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}

	// Inicia a conexão webSocket com o servidor HTTP da Binance,
	// chamando a função da biblioteca que conecta ao endPoint de Candle.
	doneC, _, err := binance.WsKlineServe(symbol, time, wsKlineHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

// streamTicker(symbol) Recebe como argumento um simbolo de tipo string
// e inicia uma conexão webSocket com o servidor HTTP da Binance,
// recebendo WsMarketStatEvent struct em formato json por 1000ms
func streamTicker(symbol string) {

	// Handler de sucesso, receberá o stream do servidor em json
	wsMarketStatEvent := func(event *binance.WsMarketStatEvent) {
		fmt.Println("\n", event)
		fmt.Println("\n-----------------------------------------------------------------------------------------------------------------------")
	}
	// Handler de erro, receberá o stream do servidor com json
	errHandler := func(err error) {
		fmt.Println(err)
	}

	// Inicia a conexão webSocket com o servidor HTTP da Binance,
	// chamando a função da biblioteca que conecta ao endPoint de ticker.
	doneC, _, err := binance.WsMarketStatServe(symbol, wsMarketStatEvent, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

// Mantem a procedure principal viva.
func commSwitch() {
	for keepAlive {
		continue
	}
}
