package todoApp

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alexeyco/simpletable"
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

	ls[index-1].CompletedAt = time.Now()
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

func (t *Todos) Load(filename string) (err error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
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

func (t *Todos) Store(filename string) (err error) {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (t *Todos) Print() {
	ls := *t
	//for i, item := range sl {
	//	i++
	//	fmt.Printf("%d - %s\n", i, item.Task)
	//}
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
			{Align: simpletable.AlignRight, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range ls {
		idx++
		task := blue(item.Task)
		done := red("no")
		createdAt := gray(item.CreatedAt.Format(time.RFC822))
		completedAt := gray(item.CompletedAt.Format(time.RFC822))

		if item.Done {
			task = green(fmt.Sprintf("\u2705 %s", item.Task))
			done = green(fmt.Sprintf("yes"))
		}

		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: fmt.Sprintf("%s", done)},
			{Text: fmt.Sprintf("%s", createdAt)},
			{Text: fmt.Sprintf("%s", completedAt)},
		})
	}
	table.Body = &simpletable.Body{Cells: cells}
	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("You have %d pending todos!", ls.CountPending()))},
	}}
	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}

func (t *Todos) CountPending() int {
	ls := *t
	total := 0
	for _, item := range ls {
		if !item.Done {
			total++
		}
	}

	return total
}
