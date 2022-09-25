package structs

import "time"

type Future struct {
	Success bool `json:"success"`
	Result  struct {
		Name                  string  `json:"name"`
		Underlying            string  `json:"underlying"`
		Description           string  `json:"description"`
		Type                  string  `json:"type"`
		Expiry                string  `json:"expiry"`
		Perpetual             bool    `json:"perpetual"`
		Expired               bool    `json:"expired"`
		Enabled               bool    `json:"enabled"`
		PostOnly              bool    `json:"postOnly"`
		PriceIncrement        float64 `json:"priceIncrement"`
		SizeIncrement         float64 `json:"sizeIncrement"`
		Last                  float64 `json:"last"`
		Bid                   float64 `json:"bid"`
		Ask                   float64 `json:"ask"`
		Index                 float64 `json:"index"`
		Mark                  float64 `json:"mark"`
		ImfFactor             float64 `json:"imfFactor"`
		LowerBound            float64 `json:"lowerBound"`
		UpperBound            float64 `json:"upperBound"`
		UnderlyingDescription string  `json:"underlyingDescription"`
		ExpiryDescription     string  `json:"expiryDescription"`
		MoveStart             string  `json:"moveStart"`
		MarginPrice           float64 `json:"marginPrice"`
		ImfWeight             float64 `json:"imfWeight"`
		MmfWeight             float64 `json:"mmfWeight"`
		PositionLimitWeight   float64 `json:"positionLimitWeight"`
		Group                 string  `json:"group"`
		CloseOnly             bool    `json:"closeOnly"`
		Change1H              float64 `json:"change1h"`
		Change24H             float64 `json:"change24h"`
		ChangeBod             float64 `json:"changeBod"`
		VolumeUsd24H          float64 `json:"volumeUsd24h"`
		Volume                float64 `json:"volume"`
		OpenInterest          float64 `json:"openInterest"`
		OpenInterestUsd       float64 `json:"openInterestUsd"`
	} `json:"result"`
}

type Futures struct {
	Success bool `json:"success"`
	Result  []struct {
		Name                  string  `json:"name"`
		Underlying            string  `json:"underlying"`
		Description           string  `json:"description"`
		Type                  string  `json:"type"`
		Expiry                string  `json:"expiry"`
		Perpetual             bool    `json:"perpetual"`
		Expired               bool    `json:"expired"`
		Enabled               bool    `json:"enabled"`
		PostOnly              bool    `json:"postOnly"`
		PriceIncrement        float64 `json:"priceIncrement"`
		SizeIncrement         float64 `json:"sizeIncrement"`
		Last                  float64 `json:"last"`
		Bid                   float64 `json:"bid"`
		Ask                   float64 `json:"ask"`
		Index                 float64 `json:"index"`
		Mark                  float64 `json:"mark"`
		ImfFactor             float64 `json:"imfFactor"`
		LowerBound            float64 `json:"lowerBound"`
		UpperBound            float64 `json:"upperBound"`
		UnderlyingDescription string  `json:"underlyingDescription"`
		ExpiryDescription     string  `json:"expiryDescription"`
		MoveStart             string  `json:"moveStart"`
		MarginPrice           float64 `json:"marginPrice"`
		ImfWeight             float64 `json:"imfWeight"`
		MmfWeight             float64 `json:"mmfWeight"`
		PositionLimitWeight   float64 `json:"positionLimitWeight"`
		Group                 string  `json:"group"`
		CloseOnly             bool    `json:"closeOnly"`
		Change1H              float64 `json:"change1h"`
		Change24H             float64 `json:"change24h"`
		ChangeBod             float64 `json:"changeBod"`
		VolumeUsd24H          float64 `json:"volumeUsd24h"`
		Volume                float64 `json:"volume"`
		OpenInterest          float64 `json:"openInterest"`
		OpenInterestUsd       float64 `json:"openInterestUsd"`
	} `json:"result"`
}

type FutureStatus struct {
	Success bool `json:"success"`
	Result  struct {
		Volume          float64   `json:"volume"`
		NextFundingRate float64   `json:"nextFundingRate"`
		NextFundingTime time.Time `json:"nextFundingTime"`
		OpenInterest    float64   `json:"openInterest"`
	} `json:"result"`
}

type FutureStats struct {
	Success bool `json:"success"`
	Result  struct {
		Volume          float64   `json:"volume"`
		NextFundingRate float64   `json:"nextFundingRate"`
		NextFundingTime time.Time `json:"nextFundingTime"`
		OpenInterest    float64   `json:"openInterest"`
	} `json:"result"`
}

type FutureRates struct {
	Success bool `json:"success"`
	Result  []struct {
		Future string    `json:"future"`
		Rate   float64   `json:"rate"`
		Time   time.Time `json:"time"`
	} `json:"result"`
}
