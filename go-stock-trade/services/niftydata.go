package services

import (
	"encoding/json"
	"fmt"
	"github.com/iamuditg/models"
	"io"
	"net/http"
	"strings"
)

type Nifty struct {
}

type NiftyInterface interface {
	FetchNiftyData() (models.Nifty50Data, error)
	FetchNiftyCompaniesData() (models.YahooData, error)
}

func (n *Nifty) FetchNiftyData() (models.Nifty50Data, error) {
	resp, err := http.Get("https://query1.finance.yahoo.com/v8/finance/chart/%5ENSEI")
	if err != nil {
		return models.Nifty50Data{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Nifty50Data{}, err
	}

	var data models.Nifty50Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		return models.Nifty50Data{}, err
	}

	return data, nil
}

func (n *Nifty) FetchNiftyCompaniesData() (models.YahooData, error) {
	// Define list of Nifty 50 stocks
	nifty50Stocks := []string{
		"BAJFINANCE.NS",
		"ADANIPORTS.NS",
		"ASIANPAINT.NS",
		"AXISBANK.NS",
		"BAJAJ-AUTO.NS",
		"BAJAJFINSV.NS",
		"BAJAJHLDNG.NS",
		"BHARTIARTL.NS",
		"BPCL.NS",
		"BRITANNIA.NS",
		"CIPLA.NS",
		"COALINDIA.NS",
		"DIVISLAB.NS",
		"DRREDDY.NS",
		"EICHERMOT.NS",
		"GAIL.NS",
		"GRASIM.NS",
		"HCLTECH.NS",
		"HDFC.NS",
		"HDFCBANK.NS",
		"HDFCLIFE.NS",
		"HEROMOTOCO.NS",
		"HINDALCO.NS",
		"HINDUNILVR.NS",
		"ICICIBANK.NS",
		"INDUSINDBK.NS",
		"INFY.NS",
		"IOC.NS",
		"ITC.NS",
		"JSWSTEEL.NS",
		"KOTAKBANK.NS",
		"LT.NS",
		"M&M.NS",
		"MARUTI.NS",
		"NESTLEIND.NS",
		"NTPC.NS",
		"ONGC.NS",
		"POWERGRID.NS",
		"RELIANCE.NS",
		"SBILIFE.NS",
		"SBIN.NS",
		"SHREECEM.NS",
		"SUNPHARMA.NS",
		"TATACONSUM.NS",
		"TATAMOTORS.NS",
		"TATASTEEL.NS",
		"TCS.NS",
		"TECHM.NS",
		"TITAN.NS",
		"ULTRACEMCO.NS",
		"UPL.NS",
		"WIPRO.NS",
	}

	// Construct URL for Yahoo Finance API
	symbols := strings.Join(nifty50Stocks, ",")
	url := fmt.Sprintf("https://query1.finance.yahoo.com/v6/finance/quote?symbols=%s", symbols)

	// Send GET request to Yahoo Finance API
	resp, err := http.Get(url)
	if err != nil {
		return models.YahooData{}, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.YahooData{}, err
	}

	// Parse JSON response
	var data models.YahooData
	if err := json.Unmarshal(body, &data); err != nil {
		return models.YahooData{}, err
	}

	return data, nil
}
