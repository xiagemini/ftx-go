package ftx

import (
	"log"

	"github.com/xiagemini/ftx-go/structs"
)

type Futures structs.Futures
type Future structs.Future
type FutureStats structs.FutureStats
type MoveStats structs.MoveStats

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

func (client *FtxClient) ListFutureStats(futureName string) (FutureStats, error) {
	var futureStats FutureStats
	resp, err := client._get("futures/"+futureName+"/stats", []byte(""))
	if err != nil {
		log.Fatal(err)
		return futureStats, err
	}
	err = _processResponse(resp, &futureStats)
	return futureStats, err
}

func (client *FtxClient) ListMoveStats(futureName string) (MoveStats, error) {
	var moveStats MoveStats
	resp, err := client._get("futures/"+futureName+"/stats", []byte(""))
	if err != nil {
		log.Fatal(err)
		return moveStats, err
	}
	err = _processResponse(resp, &moveStats)
	return moveStats, err
}
