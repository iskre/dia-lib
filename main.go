package main

import (
	"fmt"

	"github.com/iskre/lib/core"
)

func main() {
	i := core.Iskre{}
	fmt.Println(i.Sequence("First", "Second", "Third", "massive long Sequence"))
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
