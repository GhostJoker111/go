package files

import (
	"fmt"
	"os"
	"path/filepath"
)

func readFile(name string) ([]byte, error) {
	if filepath.Ext(name) != ".json" {
		return nil, fmt.Errorf("Файл не имеет json расширение")
	}
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	return data, nil
}
