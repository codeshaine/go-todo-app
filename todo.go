package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/alexeyco/simpletable"
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

func (t *Todos) Compelete(index int) error {
	if index <= 0 || index > len(*t) {
		return errors.New("invalid  index")
	}
	(*t)[index-1].CompletedAt = time.Now()
	(*t)[index-1].Done = true
	return nil
}
func (t *Todos) Incompelete(index int) error {
	if index <= 0 || index > len(*t) {
		return errors.New("invalid  index")
	}
	(*t)[index-1].CompletedAt = time.Time{}
	(*t)[index-1].Done = false
	return nil
}

func (t *Todos) Delete(index int) error {

	if index <= 0 || index > len(*t) {
		return errors.New("invalid  index")
	}
	*t = append((*t)[:index-1], (*t)[index:]...)
	return nil
}

func (t *Todos) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return err
	}
	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (t *Todos) Print() {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignCenter, Text: "CreatedAt"},
			{Align: simpletable.AlignCenter, Text: "CompletedAt"},
		},
	}

	for index, item := range *t {
		task := blue(item.Task)
		done := red("No")
		createdAt := grey(item.CreatedAt.Format("02 Jan 2006 15:04"))
		completedAt := grey("-")
		footer := red(fmt.Sprintf("You’ve got %d tasks left. Keep going!", t.CountPendingTodos()))
		if item.Done {
			done = green("Yes")
			completedAt = grey(item.CompletedAt.Format("02 Jan 2006 15:04"))
		}
		if t.CountPendingTodos() == 0 {
			footer = green("Todo list is empty! Time to relax!")
		}
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: strconv.Itoa(index + 1)},
			{Align: simpletable.AlignLeft, Text: task},
			{Align: simpletable.AlignCenter, Text: done},
			{Align: simpletable.AlignCenter, Text: createdAt},
			{Align: simpletable.AlignCenter, Text: completedAt},
		}
		table.Body.Cells = append(table.Body.Cells, r)
		table.Footer = &simpletable.Footer{
			Cells: []*simpletable.Cell{

				{Align: simpletable.AlignCenter, Text: footer, Span: 5},
			},
		}
	}
	table.SetStyle(simpletable.StyleUnicode)
	fmt.Println(table.String())
}

func (t *Todos) CountPendingTodos() int {
	count := 0
	for _, item := range *t {
		if !item.Done {
			count++
		}
	}
	return count
}
