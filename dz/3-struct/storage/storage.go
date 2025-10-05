package storage

import (
	"dz-3/struct/bins"
	"dz-3/struct/files"
	"encoding/json"

	"github.com/fatih/color"
)

type Storage struct {
	Bins bins.BinList
}

func (storage *Storage) ToBytes() ([]byte, error) {
	fileToBytes, err := json.Marshal(storage)
	if err != nil {
		return nil, err
	}
	return fileToBytes, nil
}

func (storage *Storage) Read() *Storage {
	_, err := storage.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}
	file, _ := files.ReadFile("data.json")
	unmarshalErr := json.Unmarshal(file, &storage)
	if unmarshalErr != nil {
		color.Red("Не удалось разобрать файл data.json")
		return &Storage{
			Bins: bins.BinList{},
		}
	}

	return storage
}

func (storage *Storage) Save() {
	data, err := storage.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}
	files.WriteFile(data, "data.json")
}
