package main

type CryptoSlice struct {
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
		past1m  [18]Price
		past5m  [18]Price
		past15m [18]Price
		past1h  [18]Price
		past4h  [18]Price
		past1d  [18]Price
		past1w  [18]Price
	}
	ActualMeter struct {
		actual1m  [18]Price
		actual5m  [18]Price
		actual15m [18]Price
		actual1h  [18]Price
		actual4h  [18]Price
		actual1d  [18]Price
		actual1w  [18]Price
	}
}

type CryptoMatrix struct {
	symbol    string
	PastMeter struct {
		last [7][17]Price
	}
	ActualMeter struct {
		actual [7][17]Price
	}
}

type CryptoSliceMatrix struct {
	symbol     string
	priceMeter [][]Price
}

type Price struct {
	open   float64
	close  float64
	actual float64
}
