package files

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadFile(name string) ([]byte, error) {
	if filepath.Ext(name) != ".json" {
		return nil, fmt.Errorf("Файл не имеет json расширение")
	}
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func WriteFile(data []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}
