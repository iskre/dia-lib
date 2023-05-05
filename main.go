package main

import (
	"fmt"

	"github.com/iskre/lib/core"
)

func main() {
	i := core.Isrke{}
	fmt.Println(i.Sequence("First", "Second", "Third", "massive long Sequence"))
	fmt.Println(i.Table(&core.Table{
		Headers: []string{"First", "Second", "Third"},
		Body: [][]string{
			{"Test1", "Test2", "Test3"},
			{"Test1", "Test2", "Test3"},
		},
	}))

	// i.Pie(&core.PieEntry{
	// 	Caption: "1",
	// 	Data:    33,
	// }, &core.PieEntry{
	// 	Caption: "2",
	// 	Data:    33,
	// }, &core.PieEntry{
	// 	Caption: "3",
	// 	Data:    50,
	// })
}
