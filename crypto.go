package main

type CryptoArray struct {
	symbol    string
	PastMeter struct {
		past1m  [3]Price
		past5m  [3]Price
		past15m [3]Price
		past1h  [3]Price
		past4h  [3]Price
		past1d  [3]Price
		past1w  [3]Price
	}
	ActualMeter struct {
		actual1m  [3]Price
		actual5m  [3]Price
		actual15m [3]Price
		actual1h  [3]Price
		actual4h  [3]Price
		actual1d  [3]Price
		actual1w  [3]Price
	}
}

type CryptoMatrix struct {
	PastMeter struct {
		last [7][17]Price
	}
	ActualMeter struct {
		actual [7][17]Price
	}
}

type Price struct {
	open   float64
	close  float64
	actual float64
}
