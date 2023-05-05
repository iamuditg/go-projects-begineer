package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Quote struct {
	RegularMarketPrice         float64 `json:"regularMarketPrice"`
	RegularMarketChangePercent float64 `json:"regularMarketChangePercent"`
}

type QuoteResponse struct {
	Result []Quote `json:"result"`
}

func fetchNifty50Data() (QuoteResponse, error) {
	url := "https://query1.finance.yahoo.com/v7/finance/quote?symbols=^NSEI"
	resp, err := http.Get(url)
	if err != nil {
		return QuoteResponse{}, err
	}
	defer resp.Body.Close()

	var quoteResponse QuoteResponse
	err = json.NewDecoder(resp.Body).Decode(&quoteResponse)
	if err != nil {
		return QuoteResponse{}, err
	}
	fmt.Println(quoteResponse)
	return quoteResponse, nil
}

func main() {
	nifty50Data, err := fetchNifty50Data()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	nifty50Quote := nifty50Data.Result[0]
	fmt.Printf("Nifty 50 Live Data:\nPrice: %.2f\nChange Percent: %.2f%%\n", nifty50Quote.RegularMarketPrice, nifty50Quote.RegularMarketChangePercent)
}
