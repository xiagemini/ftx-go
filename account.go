package ftx

import (
	"log"

	"github.com/xiagemini/ftx-go/structs"
)

type Positions structs.Positions

func (client *FtxClient) GetPositions(showAvgPrice bool) (Positions, error) {
	var positions Positions
	resp, err := client._get("positions", []byte(""))
	if err != nil {
		log.Fatal(err)
		return positions, err
	}
	err = _processResponse(resp, &positions)
	return positions, err
}

func (client *FtxClient) GetPosition(symbol string) (float64, error) {
	var positions Positions
	position := 0.0
	resp, err := client._get("positions", []byte(""))
	if err != nil {
		log.Fatal(err)
		return 9999, err
	}
	err = _processResponse(resp, &positions)
	for _, r := range positions.Result {
		if r.Future == symbol {
			position = float64(r.NetSize)
		}
	}
	return position, err
}
