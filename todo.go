package CLI_Todo_App

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) (err error) {
	ls := *t
	if err := validateIndex(index, len(ls)); err != nil {
		return err
	}

	ls[index-1].CreatedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

func (t *Todos) Delete(index int) (err error) {
	ls := *t
	if err := validateIndex(index, len(ls)); err != nil {
		return err
	}

	*t = append(ls[:index-1], ls[index:]...)

	return nil
}

func validateIndex(index, length int) (err error) {
	if index <= 0 || index > length {
		return errors.New("INVALID INDEX!")
	}
	return nil
}

func (t Todos) Load(filename string) (err error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return errors.New("FILE NOT EXISTS!")
		}
		return err
	}

	if len(file) == 0 {
		return errors.New("FILE IS EMPTY!")
	}
	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t Todos) Store(filename string) (err error) {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
