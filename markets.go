package ftx

import (
	"log"
	"strconv"

	"github.com/xiagemini/ftx-go/structs"
)

type HistoricalPrices structs.HistoricalPrices
type Trades structs.Trades
type Markets structs.Markets

func (client *FtxClient) GetHistoricalPrices(market string, resolution int64,
	limit int64, startTime int64, endTime int64) (HistoricalPrices, error) {
	var historicalPrices HistoricalPrices
	resp, err := client._get(
		"markets/"+market+
			"/candles?resolution="+strconv.FormatInt(resolution, 10)+
			"&limit="+strconv.FormatInt(limit, 10)+
			"&start_time="+strconv.FormatInt(startTime, 10)+
			"&end_time="+strconv.FormatInt(endTime, 10),
		[]byte(""))
	if err != nil {
		log.Fatal(err)
		return historicalPrices, err
	}
	err = _processResponse(resp, &historicalPrices)
	return historicalPrices, err
}

func (client *FtxClient) GetTrades(market string, limit int64, startTime int64, endTime int64) (Trades, error) {
	var trades Trades
	resp, err := client._get(
		"markets/"+market+"/trades?"+
			"&limit="+strconv.FormatInt(limit, 10)+
			"&start_time="+strconv.FormatInt(startTime, 10)+
			"&end_time="+strconv.FormatInt(endTime, 10),
		[]byte(""))
	if err != nil {
		log.Fatal(err)
		return trades, err
	}
	err = _processResponse(resp, &trades)
	return trades, err
}

func (client *FtxClient) ListAllMarkets() (Markets, error) {
	var markets Markets
	resp, err := client._get("markets", []byte(""))
	if err != nil {
		log.Fatal(err)
		return markets, err
	}
	err = _processResponse(resp, &markets)
	return markets, err
}
