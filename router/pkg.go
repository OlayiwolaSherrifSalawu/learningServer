package router

import (
	"encoding/json"
	"os"
)

func (s *Store) ReadJson(fileName string) ([]byte, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}
func (s *Store) Unmarsha(slices []byte, toMarshal *[]Task) error {
	newStore := make(Store)
	err := json.Unmarshal(slices, toMarshal)
	if err != nil {
		return err
	}
	for _, val := range *toMarshal {
		(newStore)[val.Id] = val
	}
	(*s) = newStore
	return nil
}

func (s Store) WriteJson(fileName string) error {
	var theStore []Task
	for _, val := range s {
		theStore = append(theStore, val)
	}

	byts, err := json.Marshal(theStore)
	if err != nil {
		return err
	}
	err = os.WriteFile(fileName, byts, 0644)
	if err != nil {
		return err
	}
	return nil
}
