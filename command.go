package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Delete int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add a new Todo specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a new Todo by index & specify a New title")
	flag.IntVar(&cf.Delete, "del", -1, "specify todo by index to delete ")
	flag.IntVar(&cf.Toggle, "toggle", -1, "specify todo by index to toggle")
	flag.BoolVar(&cf.List, "list", false, "list all todos")

	flag.Parse()
	return &cf
}

func (cf *CmdFlags) Execute (todos *Todos) {
	switch{
		case cf.List : 
			todos.print()
		case cf.Add != "" :
			todos.add(cf.Add)
		case cf.Edit != "" :
			parts := strings.SplitN(cf.Edit,":",2)
			if len(parts) != 2 {
				fmt.Println("invalid edit format, use id:New Title")
				os.Exit(1)
			}
			index, err := strconv.Atoi(parts[0])

			if err != nil {
				fmt.Println("Error: invalid index for edit")
				os.Exit(1)
			}

			todos.edit(index, parts[1])
		case cf.Toggle != -1 :
			todos.toggle(cf.Toggle)
		case cf.Delete != -1 :
			todos.delete(cf.Delete)
		default : 
			fmt.Println("Invalid command")
	}
}