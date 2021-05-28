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
	apiKey    string = "123"
	secretKey string = "123"
)

var keepAlive bool = true

func main() {
	//client := binance.NewClient(apiKey, secretKey)
	//ticker(client, "ETHUSDT")
	//go streamTicker("ETHUSDT")
	//go streamCandle("BTCUSDT", "5m")
	go streamCandle("ETHUSDT", "1m")
	commSwitch()
}

//Inicializa os objetos (scructs) de cada moeda
func setCoin() CryptoArray {
	coin := CryptoArray{}
	fmt.Println("Coin struct configurado e alocado!")
	return coin
}

// Abre um canal de comunicação (channel) para que go routines
// tenham acesso simultaneo aos objetos (sctructs) das moedas
// retornando um canal com tipos CryptoArray (chan CryptoArray)
func chanCoin() chan CryptoArray {
	fmt.Println("Abrindo canal para coin.")
	coin := make(chan CryptoArray)
	fmt.Println("Canal aberto com sucesso!")
	go func() {
		coin <- setCoin()
		fmt.Println("Coin inserido no canal!")
	}()
	return coin
}

// Retorna o objeto que esta no canal (channel) como um
// objeto propriamente do tipo CryptoArray.
func returnCoin() CryptoArray {
	fmt.Println("Retornando Objeto no canal...")
	return <-chanCoin()
}

//Inicializa os objetos (scructs) de cada moeda
func setSliceMatrixCoin() CryptoSliceMatrix {
	coin := CryptoSliceMatrix{}
	fmt.Println("Coin struct configurado e alocado!")
	return coin
}

// Abre um canal de comunicação (channel) para que go routines
// tenham acesso simultaneo aos objetos (sctructs) das moedas
// retornando um canal com tipos CryptoSliceMatrix (chan CryptoSliceMatrix)
func chanSliceMatrixCoin() chan CryptoSliceMatrix {
	fmt.Println("Abrindo canal para coin.")
	coin := make(chan CryptoSliceMatrix)
	fmt.Println("Canal aberto com sucesso!")
	go func() {
		coin <- setSliceMatrixCoin()
		fmt.Println("Coin inserido no canal!")
	}()
	return coin
}

// Retorna o objeto que esta no canal (channel) como um
// objeto propriamente do tipo CryptoSliceMatrix.
func returnSliceMatrixCoin() CryptoSliceMatrix {
	fmt.Println("Retornando Objeto no canal...")
	return <-chanSliceMatrixCoin()
}

//Inicializa os objetos (scructs) de cada moeda
func setMean() []float64 {
	mean := []float64{}
	fmt.Println("Mean slice configurado e alocado!")
	return mean
}

// Abre um canal de comunicação (channel) para que go routines
// tenham acesso simultaneo aos objetos (sctructs) das moedas
// retornando um canal com tipos CryptoArray (chan CryptoArray)
func chanMean() chan []float64 {
	fmt.Println("Abrindo canal para mean.")
	meanChannel := make(chan []float64)
	fmt.Println("Canal aberto com sucesso!")
	go func() {
		meanChannel <- setMean()
		fmt.Println("Mean inserido no canal!")
	}()
	return meanChannel
}

// Retorna o objeto que esta no canal (channel) como um
// objeto propriamente do tipo CryptoArray.
func returnMean() []float64 {
	fmt.Println("Retornando Objeto no mean channel...")
	return <-chanMean()
}

func priceMean(actualPrice [18]Price, pastPrice [18]Price, meanChannel chan [2]float64) {
	var (
		actualMean float64 = 0
		actualSum  float64 = 0
		pastMean   float64 = 0
		pastSum    float64 = 0
	)
	for i := 0; i < len(actualPrice); i++ {
		actualSum += actualPrice[i].open

	}
	for i := 0; i < len(pastPrice); i++ {
		pastSum += pastPrice[i].open

	}
	actualMean = actualSum / float64(len(actualPrice))
	actualPrice[0].open = actualMean

	pastMean = pastSum / float64(len(pastPrice))
	pastPrice[0].open = pastMean
	var means = [2]float64{actualMean, pastMean}
	meanChannel <- means
}

func printActual(symbol string, actualPriceCounter int, actualPrice [18]Price) {
	fmt.Println("_______________________________________________________________")
	fmt.Println(symbol)
	fmt.Printf("%d - ", actualPriceCounter)
	fmt.Print("Preço Atual: ")
	fmt.Println(actualPrice[actualPriceCounter])
}

func printPast(pastPriceCounter int, pastPrice [18]Price) {
	fmt.Printf("%d - ", pastPriceCounter)
	fmt.Print("Preço Passado: ")
	fmt.Println(pastPrice[pastPriceCounter])
}

func sliceMatrixAssembleKline(event chan *binance.WsKlineEvent, symbol string) {

	var (
		//pastMean           float64 = 0
		streamedCoin  CryptoSliceMatrix
		streamedEvent *binance.WsKlineEvent
		//actualPriceCounter int  = 0
		//pastPriceCounter   int  = 0
		//firstIteration     bool = true
		//means                   = [2]float64{0.0, 0.0}
		streamedCandle  Candle
		actualPriceList []Candle
		//pastPriceList   []Candle
		coinChannel = chanSliceMatrixCoin()
		//meanChannel             = chanMean()
	)

	//meanChannel := make(chan [2]float64)
	streamedCoin = <-coinChannel
	streamedCoin.symbol = symbol

	for keepAlive {
		streamedEvent = <-event
		open, err := strconv.ParseFloat(streamedEvent.Kline.Open, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		close, err := strconv.ParseFloat(streamedEvent.Kline.Close, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		high, err := strconv.ParseFloat(streamedEvent.Kline.High, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		low, err := strconv.ParseFloat(streamedEvent.Kline.Low, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		//volume, err := strconv.ParseFloat(streamedEvent.Kline.Volume, 64)
		//streamedCoin.priceMeter = append(streamedCoin.priceMeter, Price{streamedPrice})
		if streamedCandle == (Candle{}) {
			streamedCandle = Candle{open: open, close: close, high: high, low: low}
		}
		if streamedCandle.open != open {
			actualPriceList = append(actualPriceList, streamedCandle)
			fmt.Println(actualPriceList)
		}
		streamedCandle = Candle{open: open, close: close, high: high, low: low}

	}
	/*if actualPriceCounter == 17 {
					firstIteration = false
					actualPriceCounter = 0
				}
				actualPriceCounter++
				if firstIteration {
					streamedCoin.ActualMeter.actual1m[actualPriceCounter] = streamedPrice
					printActual(symbol, actualPriceCounter, streamedCoin.ActualMeter.actual1m)
					fmt.Printf("Local de Memória do Objeto: %p\n", &streamedCoin)
				} else {
					pastPriceCounter++
					if pastPriceCounter <= 17 {
						streamedCoin.PastMeter.past1m[pastPriceCounter] = streamedCoin.ActualMeter.actual1m[pastPriceCounter]
						streamedCoin.ActualMeter.actual1m[actualPriceCounter] = streamedPrice
						printActual(symbol, actualPriceCounter, streamedCoin.ActualMeter.actual1m)
						printPast(actualPriceCounter, streamedCoin.PastMeter.past1m)
						fmt.Printf("Local de Memória do Objeto: %p\n", &streamedCoin)
						if pastPriceCounter == 17 {
							pastPriceCounter = 0
							go priceMean(streamedCoin.ActualMeter.actual1m, streamedCoin.PastMeter.past1m, meanChannel)
							means = <-meanChannel
							streamedCoin.ActualMeter.actual1m[0].open = means[0]
							streamedCoin.PastMeter.past1m[0].open = means[1]
						}
					}
				}
			} else {
				streamedCoin.ActualMeter.actual1m[actualPriceCounter] = streamedPrice
			}
		}
	}*/
}

func assembleKline(event chan *binance.WsKlineEvent, symbol string) {

	var (
		//pastMean           float64 = 0
		streamedCoin       CryptoArray
		actualPriceCounter int  = 0
		pastPriceCounter   int  = 0
		firstIteration     bool = true
		means                   = [2]float64{0.0, 0.0}
		coinChannel             = chanCoin()
	)

	meanChannel := make(chan [2]float64)
	streamedCoin = <-coinChannel
	streamedCoin.symbol = symbol

	for keepAlive {
		var streamedEvent *binance.WsKlineEvent
		streamedEvent = <-event
		a, err := strconv.ParseFloat(streamedEvent.Kline.Open, 64)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			streamedPrice := Price{open: a}
			if streamedPrice != streamedCoin.ActualMeter.actual1m[actualPriceCounter] {
				if actualPriceCounter == 17 {
					firstIteration = false
					actualPriceCounter = 0
				}
				actualPriceCounter++
				if firstIteration {
					streamedCoin.ActualMeter.actual1m[actualPriceCounter] = streamedPrice
					printActual(symbol, actualPriceCounter, streamedCoin.ActualMeter.actual1m)
					fmt.Printf("Local de Memória do Objeto: %p\n", &streamedCoin)
				} else {
					pastPriceCounter++
					if pastPriceCounter <= 17 {
						streamedCoin.PastMeter.past1m[pastPriceCounter] = streamedCoin.ActualMeter.actual1m[pastPriceCounter]
						streamedCoin.ActualMeter.actual1m[actualPriceCounter] = streamedPrice
						printActual(symbol, actualPriceCounter, streamedCoin.ActualMeter.actual1m)
						printPast(actualPriceCounter, streamedCoin.PastMeter.past1m)
						fmt.Printf("Local de Memória do Objeto: %p\n", &streamedCoin)
						if pastPriceCounter == 17 {
							pastPriceCounter = 0
							go priceMean(streamedCoin.ActualMeter.actual1m, streamedCoin.PastMeter.past1m, meanChannel)
							means = <-meanChannel
							streamedCoin.ActualMeter.actual1m[0].open = means[0]
							streamedCoin.PastMeter.past1m[0].open = means[1]
						}
					}
				}
			} else {
				streamedCoin.ActualMeter.actual1m[actualPriceCounter] = streamedPrice
			}
		}
	}

}

// streamCandle(symbol, time) Recebe como argumento um simbolo de tipo string
// e um tempo de tipo string no formato 1m , 5m , 15m, 30m, 1h, 2h, 4h...
// e inicia uma conexão webSocket com o servidor HTTP da Binance, recebendo
// WsKlineEvent struct em formato json por 2000ms
func streamCandle(symbol string, time string) {

	streamedEvent := make(chan *binance.WsKlineEvent)
	//go assembleKline(streamedEvent, symbol)
	go sliceMatrixAssembleKline(streamedEvent, symbol)
	// Handler de sucesso, receberá o stream do servidor em json
	wsKlineHandler := func(event *binance.WsKlineEvent) {
		streamedEvent <- event
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

// streamTicker(symbol) Recebe como argumento um simbolo de tipo string
// e inicia uma conexão webSocket com o servidor HTTP da Binance,
// recebendo WsMarketStatEvent struct em formato json por 1000ms
func streamTicker(symbol string) {

	streamedCoin := returnCoin()
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
