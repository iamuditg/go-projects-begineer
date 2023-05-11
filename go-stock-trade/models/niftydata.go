package models

type Nifty50Data struct {
	Chart struct {
		Result []struct {
			Meta struct {
				Currency             string  `json:"currency"`
				Symbol               string  `json:"symbol"`
				ExchangeName         string  `json:"exchangeName"`
				InstrumentType       string  `json:"instrumentType"`
				FirstTradeDate       int     `json:"firstTradeDate"`
				RegularMarketTime    int     `json:"regularMarketTime"`
				Gmtoffset            int     `json:"gmtoffset"`
				Timezone             string  `json:"timezone"`
				ExchangeTimezoneName string  `json:"exchangeTimezoneName"`
				RegularMarketPrice   float64 `json:"regularMarketPrice"`
				ChartPreviousClose   float64 `json:"chartPreviousClose"`
				PreviousClose        float64 `json:"previousClose"`
				Scale                int     `json:"scale"`
				PriceHint            int     `json:"priceHint"`
				CurrentTradingPeriod struct {
					Pre struct {
						Timezone  string `json:"timezone"`
						Start     int    `json:"start"`
						End       int    `json:"end"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"pre"`
					Regular struct {
						Timezone  string `json:"timezone"`
						Start     int    `json:"start"`
						End       int    `json:"end"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"regular"`
					Post struct {
						Timezone  string `json:"timezone"`
						Start     int    `json:"start"`
						End       int    `json:"end"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"post"`
				} `json:"currentTradingPeriod"`
				TradingPeriods [][]struct {
					Timezone  string `json:"timezone"`
					Start     int    `json:"start"`
					End       int    `json:"end"`
					Gmtoffset int    `json:"gmtoffset"`
				} `json:"tradingPeriods"`
				DataGranularity string   `json:"dataGranularity"`
				Range           string   `json:"range"`
				ValidRanges     []string `json:"validRanges"`
			} `json:"meta"`
			Timestamp  []int `json:"timestamp"`
			Indicators struct {
				Quote []struct {
					Open   []float64 `json:"open"`
					Close  []float64 `json:"close"`
					Volume []int     `json:"volume"`
					High   []float64 `json:"high"`
					Low    []float64 `json:"low"`
				} `json:"quote"`
			} `json:"indicators"`
		} `json:"result"`
	} `json:"chart"`
}

type YahooData struct {
	QuoteResponse quoteResponse `json:"quoteResponse"`
}

type quoteResponse struct {
	Result []struct {
		ShortName                     string  `json:"shortName"`
		RegularMarketPrice            float64 `json:"regularMarketPrice"`
		RegularMarketChangePercent    float64 `json:"regularMarketChangePercent"`
		RegularMarketChange           float64 `json:"regularMarketChange"`
		FinancialCurrency             string  `json:"financialCurrency"`
		RegularMarketOpen             float64 `json:"regularMarketOpen"`
		AverageDailyVolume3Month      int     `json:"averageDailyVolume3Month"`
		AverageDailyVolume10Day       int     `json:"averageDailyVolume10Day"`
		FiftyTwoWeekLowChange         float64 `json:"fiftyTwoWeekLowChange"`
		FiftyTwoWeekLowChangePct      float64 `json:"fiftyTwoWeekLowChangePercent"`
		FiftyTwoWeekRange             string  `json:"fiftyTwoWeekRange"`
		FiftyTwoWeekHighChange        float64 `json:"fiftyTwoWeekHighChange"`
		FiftyTwoWeekHighChangePct     float64 `json:"fiftyTwoWeekHighChangePercent"`
		FiftyTwoWeekLow               float64 `json:"fiftyTwoWeekLow"`
		FiftyTwoWeekHigh              float64 `json:"fiftyTwoWeekHigh"`
		EarningsTimestamp             int64   `json:"earningsTimestamp"`
		EarningsTimestampStart        int64   `json:"earningsTimestampStart"`
		EarningsTimestampEnd          int64   `json:"earningsTimestampEnd"`
		TrailingAnnualDividendRate    float64 `json:"trailingAnnualDividendRate"`
		TrailingPE                    float64 `json:"trailingPE"`
		TrailingAnnualDividendYld     float64 `json:"trailingAnnualDividendYield"`
		EpsTrailingTwelveMonths       float64 `json:"epsTrailingTwelveMonths"`
		EpsForward                    float64 `json:"epsForward"`
		EpsCurrentYear                float64 `json:"epsCurrentYear"`
		Exchange                      string  `json:"exchange"`
		LongName                      string  `json:"longName"`
		MessageBoardId                string  `json:"messageBoardId"`
		ExchangeTimezoneName          string  `json:"exchangeTimezoneName"`
		ExchangeTimezoneShortName     string  `json:"exchangeTimezoneShortName"`
		GmtOffSetMilliseconds         int     `json:"gmtOffSetMilliseconds"`
		Market                        string  `json:"market"`
		FirstTradeDateMilliseconds    int64   `json:"firstTradeDateMilliseconds"`
		PriceHint                     int     `json:"priceHint"`
		MarketState                   string  `json:"marketState"`
		PriceEpsCurrentYear           float64 `json:"priceEpsCurrentYear"`
		SharesOutstanding             int64   `json:"sharesOutstanding"`
		BookValue                     float64 `json:"bookValue"`
		FiftyDayAverage               float64 `json:"fiftyDayAverage"`
		FiftyDayAverageChange         float64 `json:"fiftyDayAverageChange"`
		FiftyDayAverageChangePct      float64 `json:"fiftyDayAverageChangePercent"`
		TwoHundredDayAverage          float64 `json:"twoHundredDayAverage"`
		TwoHundredDayAverageChange    float64 `json:"twoHundredDayAverageChange"`
		TwoHundredDayAverageChangePct float64 `json:"twoHundredDayAverageChangePercent"`
	} `json:"result"`
}
