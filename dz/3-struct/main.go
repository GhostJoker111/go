package main

import "time"

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
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
