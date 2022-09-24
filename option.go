package ftx

import (
	"encoding/json"
	"log"
	"time"

	"github.com/xiagemini/ftx-go/structs"
)

type CreateQuoteRequest structs.CreateQuoteRequest
type OptionResp structs.Response
type GetQuoteRequests structs.GetQuoteRequests
type GetQuote structs.GetQuote

func (client *FtxClient) CreateOpitionRequest(optionType string, strikePrice int, expiry string, side string, size float32) (CreateQuoteRequest, error) {
	var optionRequest CreateQuoteRequest
	loc, _ := time.LoadLocation("UTC")
	strikTime, _ := time.ParseInLocation("2006-01-02 15:04:05", expiry+" 03:00:00", loc)
	unixTime := strikTime.Unix()
	requestBody, err := json.Marshal(structs.PostOptionRequest{
		Underlying: "BTC",
		Type:       optionType,
		Strike:     strikePrice,
		Expiry:     unixTime,
		Side:       side,
		Size:       size})
	if err != nil {
		log.Fatal(err)
		return optionRequest, err
	}
	resp, err := client._post("options/requests", requestBody)
	if err != nil {
		log.Fatal(err)
		return optionRequest, err
	}
	err = _processResponse(resp, &optionRequest)
	return optionRequest, err
}

func (client *FtxClient) GetOpitionRequests() (GetQuoteRequests, error) {
	var myOptionRequests GetQuoteRequests
	resp, err := client._get("options/my_requests", []byte(""))
	if err != nil {
		log.Fatal(err)
		return myOptionRequests, err
	}
	err = _processResponse(resp, &myOptionRequests)
	return myOptionRequests, err
}

func (client *FtxClient) GetOpitionQuote(id string) (GetQuote, error) {
	var myOptionQuote GetQuote
	resp, err := client._get("options/requests/"+id+"/quotes", []byte(""))
	// resp, err := client._get("options/requests/"+strconv.Itoa(id)+"/quotes", []byte(""))
	if err != nil {
		log.Fatal(err)
		return myOptionQuote, err
	}
	err = _processResponse(resp, &myOptionQuote)
	return myOptionQuote, err
}
