package structs

type Orderbook struct {
	Success bool `json:"success"`
	Result  struct {
		Bids [][]float32 `json:"bids"`
		Asks [][]float32 `json:"asks"`
	} `json:"result"`
}
