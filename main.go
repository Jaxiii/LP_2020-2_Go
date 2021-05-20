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
	"strconv"

	"github.com/adshao/go-binance/v2"
)

var (
	keepAlive bool   = true
	apiKey    string = "123"
	secretKey string = "123"
)

func main() {
	//client := binance.NewClient(apiKey, secretKey)
	//ticker(client, "ETHUSDT")
	//go streamCandle("BTCUSDT", "1m")
	go streamCandle("ETHUSDT", "1m")
	//go streamTicker("ETHUSDT")
	commSwitch()
}

//Inicializa os objetos (scructs) de cada moeda
func setCoins() CryptoArray {
	fmt.Println("Coin struct configurado e alocado!")
	coin := CryptoArray{}
	return coin
}

// Abre um canal de comunicação (channel) para que go routines
// tenham acesso simultaneo aos objetos (sctructs) das moedas
// retornando um canal com tipos CryptoArray (chan CryptoArray)
func chanCoins() chan CryptoArray {
	fmt.Println("Abrindo canal para coin.")
	coin := make(chan CryptoArray)
	go func() {
		coin <- setCoins()
	}()
	fmt.Println("Canal aberto com sucesso!")
	return coin
}

// Retorna o objeto que esta no canal (channel) como um
// objeto propriamente do tipo CryptoArray.
func returnCoins() CryptoArray {
	fmt.Println("Sinalizando abertura de canal...")
	return <-chanCoins()
}

func actualMean(actualPrices []Price) float64 {
	var actualMean float64 = 0
	var sumActual float64 = 0
	for i := 0; i > len(actualPrices); i++ {
		sumActual = actualPrices[i].open
	}
	actualMean = sumActual / float64(len(actualPrices))
	return actualMean
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
func streamCandle(symbol string, time string) {

	streamedCoin := returnCoins()
	streamedCoin.symbol = symbol
	var (
		//pastMean           float64 = 0
		actualPriceCounter int  = 0
		pastPriceCounter   int  = 0
		firstIteration     bool = true
	)
	// Handler de sucesso, receberá o stream do servidor em json
	wsKlineHandler := func(event *binance.WsKlineEvent) {
		a, err := strconv.ParseFloat(event.Kline.Open, 64)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			streamedPrice := Price{open: a}
			//streamedCoin.ActualMeter.actual1m[0].open = actualMean(streamedCoin.ActualMeter.actual1m)
			if streamedPrice != streamedCoin.ActualMeter.actual1m[actualPriceCounter] {
				if actualPriceCounter == 2 {
					firstIteration = false
					actualPriceCounter = 0
				}
				actualPriceCounter++
				if firstIteration {
					streamedCoin.ActualMeter.actual1m[actualPriceCounter] = streamedPrice
					fmt.Println("_______________________________________________________________")
					fmt.Println(streamedCoin.symbol)
					fmt.Printf("%d - ", actualPriceCounter)
					fmt.Print("Preço Atual: ")
					fmt.Println(streamedCoin.ActualMeter.actual1m[actualPriceCounter])
					fmt.Printf("Local de Memória do Objeto: %p\n", &streamedCoin)
				} else {
					pastPriceCounter++
					if pastPriceCounter <= 2 {
						streamedCoin.PastMeter.past1m[pastPriceCounter] = streamedCoin.ActualMeter.actual1m[pastPriceCounter]
						streamedCoin.ActualMeter.actual1m[actualPriceCounter] = streamedPrice
						fmt.Println("_______________________________________________________________")
						fmt.Println(streamedCoin.symbol)
						fmt.Printf("%d - ", pastPriceCounter)
						fmt.Print("Preço Passado: ")
						fmt.Println(streamedCoin.PastMeter.past1m)
						fmt.Printf("%d - ", actualPriceCounter)
						fmt.Print("Preço Atual: ")
						fmt.Println(streamedCoin.ActualMeter.actual1m)
						fmt.Printf("Local de Memória do Objeto: %p\n", &streamedCoin)
						if pastPriceCounter == 2 {
							pastPriceCounter = 0
						}
					}
				}
			} else {
				streamedCoin.ActualMeter.actual1m[actualPriceCounter] = streamedPrice
			}
		}
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

	streamedCoin := returnCoins()
	streamedCoin.symbol = symbol
	// Handler de sucesso, receberá o stream do servidor em json
	wsMarketStatEvent := func(event *binance.WsMarketStatEvent) {
		a, err := strconv.ParseFloat(event.OpenPrice, 64)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("_______________________________________________________________")
			fmt.Println(streamedCoin.symbol)
			streamedPrice := Price{open: a}
			streamedCoin.ActualMeter.actual1m[0] = streamedPrice
			fmt.Println(streamedCoin.ActualMeter.actual1m[0])
			fmt.Printf("Local de Memória do Objeto: %p\n", &streamedCoin)
		}
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
