package main

type CryptoDinamic struct {
	symbol    string
	PastMeter struct {
		past1m  []Price
		past5m  []Price
		past15m []Price
		past1h  []Price
		past4h  []Price
		past1d  []Price
		past1w  []Price
	}
	ActualMeter struct {
		actual1m  []Price
		actual5m  []Price
		actual15m []Price
		actual1h  []Price
		actual4h  []Price
		actual1d  []Price
		actual1w  []Price
	}
}

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
