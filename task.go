package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Tasks []item

func (t *Tasks) Add(task string) {
	toBeAddedTask := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, toBeAddedTask)
}

// index = position - 1
func (t *Tasks) Complete(position int) error {
	tasks := *t

	if position < 0 || position > len(tasks) {
		return errors.New("index is out of range")
	}

	tasks[position-1].CompletedAt = time.Now()
	tasks[position-1].Done = true

	return nil
}

// index = position - 1
func (t *Tasks) Delete(position int) error {
	tasks := *t
	if position < 0 || position > len(tasks) {
		return errors.New("index is out of range")
	}
	*t = append(tasks[:position-1], tasks[position:]...)
	return nil
}

func (t *Tasks) Load(fileName string) error {
	file, err := os.ReadFile(fileName)
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

func (t *Tasks) Store(fileName string) error {
	data, err := json.Marshal(t)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func (t *Tasks) GetPending() int {
	count := 0
	for _, item := range *t {
		if !item.Done {
			count++
		}
	}
	return count
}

func (t *Tasks) Print() {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#ID"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignRight, Text: "Created at"},
			{Align: simpletable.AlignRight, Text: "Completed at"},
		},
	}

	var cells [][]*simpletable.Cell

	for i, item := range *t {
		id := Gray(fmt.Sprintf("%d", i+1))
		isDone := Blue(fmt.Sprintf("%t", item.Done))
		createdAt := Red(item.CreatedAt.Format(time.RFC822))
		completedAt := Red(item.CompletedAt.Format(time.RFC822))
		task := Red(item.Task)
		if item.Done {
			task = Green(fmt.Sprintf("\u2705 %s", item.Task))
		}
		cells = append(cells, []*simpletable.Cell{
			{Text: id},
			{Text: task},
			{Text: isDone},
			{Text: createdAt},
			{Text: completedAt},
		})
	}

	table.Body = &simpletable.Body{
		Cells: cells,
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 5, Text: Red(fmt.Sprintf("You have %d incomplete tasks!", t.GetPending()))},
		},
	}

	table.SetStyle(simpletable.StyleUnicode)
	table.Print()
}
