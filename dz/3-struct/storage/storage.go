package storage

import (
	"dz-3/struct/files"
	"encoding/json"

	"github.com/fatih/color"
)

type Storage struct {
	Bins []Bin
}

func (storage *Storage) ToBytes() ([]byte, error) {
	fileToBytes, err := json.Marshal(storage)
	if err != nil {
		return nil, err
	}
	return fileToBytes, nil
}

func (storage *Storage) save() {
	data, err := storage.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}
	files.WriteFile(data, "data.json")
}
