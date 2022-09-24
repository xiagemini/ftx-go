package structs

type Fills struct {
	Success bool `json:"success"`
	Result  []struct {
		Fee           float32 `json:"fee,omitempty"`
		FeeCurrency   string  `json:"feeCurrency,omitempty"`
		FeeRat        float64 `json:"feeRate,omitempty"`
		Future        string  `json:"future,omitempty"`
		Id            int     `json:"id,omitempty"`
		Liquidity     string  `json:"liquidity,omitempty"`
		Market        string  `json:"market,omitempty"`
		BaseCurrency  string  `json:"baseCurrency,omitempty"`
		QuoteCurrency string  `json:"quoteCurrency,omitempty"`
		OrderId       int     `json:"orderId,omitempty"`
		TradeId       int     `json:"tradeId,omitempty"`
		Price         float64 `json:"price,omitempty"`
		Side          string  `json:"side,omitempty"`
		Size          float64 `json:"size,omitempty"`
		Time          string  `json:"time,omitempty"`
		Type          string  `json:"type,omitempty"`
	} `json:"result,omitempty"`
}
