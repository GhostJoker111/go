package main

import "time"

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

func newBin(id string, private bool, createdAt time.Time, name string) (*Bin, error) {
	newBin := &Bin{Id: id, Private: private, CreatedAt: createdAt, Name: name}
	return newBin, nil
}

type BinList []Bin

func createBin(len int) BinList {
	binList := make(BinList, len)
	return binList
}

func main() {

}
