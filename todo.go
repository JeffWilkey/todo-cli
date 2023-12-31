package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

// Method to add a new item to the list
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

// Complete method marks a todo item as completed
// by setting the Done field to true and CompletedAt to the current time
func (l *List) Complete(i int) error {
	ls := *l
	if i < 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	// Adjust index for 0 based index
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

// Delete method removes an item from the list
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	// Adjust index for 0 based index
	*l = append(ls[:i-1], ls[i:]...)

	return nil
}

// Save method encodes the List as JSON and saves it
// using the provided file name
func (l *List) Save(filename string) error {
	jsonData, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonData, 0o644)
}

// Get method decodes and returns a List from a JSON file
func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}

// String prints out a formatted list
// Implements the fmt.Stringer interface
func (l *List) String() string {
	formatted := ""

	for k, t := range *l {
		prefix := "[ ] "
		if t.Done {
			prefix = "[X] "
		}
		// Adjust the item number k to print numbers starting from 1 instead of 0
		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
	}

	return formatted
}
