package structs

import "time"

type Option struct {
	Expiry     time.Time `json:"expiry"`
	Strike     float64   `json:"strike"`
	Type       string    `json:"type"`
	Underlying string    `json:"underlying"`
}

type PostOptionRequest struct {
	Underlying string  `json:"underlying"`
	Type       string  `json:"type"`
	Strike     int     `json:"strike"`
	Expiry     int64   `json:"expiry"`
	Side       string  `json:"side"`
	Size       float32 `json:"size"`
}

type CreateQuoteRequest struct {
	Success bool `json:"success"`
	Result  struct {
		HideLimitPrice bool   `json:"hideLimitPrice"`
		ID             string `json:"id"`
		Option         struct {
			Expiry     string  `json:"expiry"`
			Strike     float32 `json:"strike"`
			Type       string  `json:"type"`
			Underlying string  `json:"underlying"`
			ID         int     `json:"id"`
		} `json:"option"`
		RequestExpiry string  `json:"requestExpiry"`
		Side          string  `json:"side"`
		Size          float32 `json:"size"`
		Status        string  `json:"status"`
		Time          string  `json:"time"`
	}
}

type CreatQuote struct {
	Success bool `json:"success"`
	Result  struct {
		Collateral  float64   `json:"collateral"`
		ID          int       `json:"id"`
		Option      Option    `json:"option"`
		Price       float64   `json:"price"`
		QuoteExpiry time.Time `json:"quoteExpiry"`
		QuoterSide  string    `json:"quoterSide"`
		RequestID   int       `json:"requestId"`
		RequestSide string    `json:"requestSide"`
		Size        float64   `json:"size"`
		Status      string    `json:"status"`
		Time        time.Time `json:"time"`
	}
}

type GetQuoteRequests struct {
	Success bool `json:"success"`
	Result  []struct {
		ID     string `json:"id"`
		Option struct {
			Underlying string    `json:"underlying"`
			Type       string    `json:"type"`
			Strike     float32   `json:"strike"`
			Expiry     time.Time `json:"expiry"`
			ID         int       `json:"id"`
		} `json:"option"`
		Side           string    `json:"side"`
		Size           float64   `json:"size"`
		Time           time.Time `json:"time"`
		RequestExpiry  time.Time `json:"requestExpiry"`
		Status         string    `json:"status"`
		HideLimitPrice bool      `json:"hideLimitPrice"`
		Quotes         []struct {
			Collateral  float64   `json:"collateral"`
			ID          string    `json:"id"`
			Price       float64   `json:"price"`
			QuoteExpiry time.Time `json:"quoteExpiry"`
			Status      string    `json:"status"`
			Time        time.Time `json:"time"`
		} `json:"quotes"`
	}
}

type GetQuote struct {
	Success bool `json:"success"`
	Result  []struct {
		Collateral float64 `json:"collateral"`
		ID         string  `json:"id"`
		Option     struct {
			Expiry     time.Time `json:"expiry"`
			Strike     float64   `json:"strike"`
			Type       string    `json:"type"`
			Underlying string    `json:"underlying"`
			ID         int       `json:"id"`
		} `json:"option"`
		Price       float64     `json:"price"`
		QuoteExpiry interface{} `json:"quoteExpiry"`
		QuoterSide  string      `json:"quoterSide"`
		RequestID   string      `json:"requestId"`
		RequestSide string      `json:"requestSide"`
		Size        float64     `json:"size"`
		Status      string      `json:"status"`
		Time        time.Time   `json:"time"`
	} `json:"result"`
}
