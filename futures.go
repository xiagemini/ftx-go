package ftx

import (
	"log"

	"github.com/xiagemini/ftx-go/structs"
)

type Futures structs.Futures
type Future structs.Future

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

func (client *FtxClient) ListFuture(futureName string) (Future, error) {
	var future Future
	resp, err := client._get("futures/"+futureName, []byte(""))
	if err != nil {
		log.Fatal(err)
		return future, err
	}
	err = _processResponse(resp, &future)
	return future, err
}
