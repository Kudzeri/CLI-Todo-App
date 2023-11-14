package CLI_Todo_App

import (
	"errors"
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

func validateIndex(index, length int) error {
	if index <= 0 || index > length {
		return errors.New("INVALID INDEX")
	}
	return nil
}
