package ftx

import (
	"log"
	"strconv"

	"github.com/xiagemini/ftx-go/structs"
)

type Fills structs.Fills

func (client *FtxClient) GetFills(market string, startTime int, endTime int) (Fills, error) {
	var fills Fills
	// requestBody, _ := json.Marshal(map[string]interface{}{
	// 	"start_time": startTime,
	// 	"end_time":   endTime,
	// })
	// resp, err := client._get("fills?market="+market, requestBody)
	resp, err := client._get("fills?market="+market+"&start_time="+strconv.Itoa(startTime)+"&end_time="+strconv.Itoa(endTime), []byte(""))
	if err != nil {
		log.Fatal(err)
		return fills, err
	}
	err = _processResponse(resp, &fills)
	return fills, err
}
