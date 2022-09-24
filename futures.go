package ftx

import (
	"log"

	"github.com/xiagemini/ftx-go/structs"
)

type Futures structs.Futures

func (client *FtxClient) ListAllFutures() (Futures, error) {
	var futures Futures
	resp, err := client._get("futures", []byte(""))
	if err != nil {
		log.Fatal(err)
		return futures, err
	}
	err = _processResponse(resp, &futures)
	return futures, err
}
