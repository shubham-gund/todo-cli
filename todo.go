package main

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string){
	todo := Todo{
		Title: title,
		Completed: false,
		CreatedAt: time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(Index int) error{
	if Index < 0 || Index >= len(*todos) {
		err:= errors.New("invalid index");
		return err;
	}
	return nil
}

func(todos *Todos) delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil{
		return err
	}

	*todos = append(t[:index], t[index+1:]...) // append till index excluding index to index+1 and the rest of the todos making it removing the todo at index passed as parameter
	return nil
}

func (todos *Todos) toggle (index int) error{
	t := *todos

	if err := t.validateIndex(index); err != nil{
		return err
	}

	isCompleted := t[index].Completed

	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}

	t[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) edit (index int,title string) error{
	t := *todos

	if err := t.validateIndex(index); err != nil{
		return err
	} 

	t[index].Title = title

	return nil
}

func (todos *Todos) print(){
	table:=table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#","Title", "Completed", "Created At","Completed At")
	for index, t := range *todos{
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}