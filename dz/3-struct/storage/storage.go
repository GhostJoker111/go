package storage

type Storage struct {
	Bins Bin[]
}

func (storage *Storage) WriteFile(content []byte) {}

func (storage *Storage) Save(content []byte) error {

}
