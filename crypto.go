package main

type CryptoArray struct {
	PastMeter struct {
		last1m  []Price
		last5m  []Price
		last15m []Price
		last1h  []Price
		last4h  []Price
		last1d  []Price
		last1w  []Price
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

type CryptoMatrix struct {
	PastMeter struct {
		last [7][17]Price
	}
	ActualMeter struct {
		actual [7][17]Price
	}
}

type Price struct {
	open  int
	close int
	mean  int
}
