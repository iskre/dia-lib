package main

import (
	"fmt"

	"github.com/iskre/lib/core"
)

func main() {
	i := core.Iskre{}
	fmt.Println(i.Repetition("While loop conditiontest", "test\n test\n  test\ntest"))
	fmt.Println(i.Sequence("test short", "test very very very very very very very very very very very very very very very very long", "..."))

	//	i.Pie(&core.PieEntry{
	//		Caption: "1",
	//		Data:    33,
	//	}, &core.PieEntry{
	//
	//		Caption: "2",
	//		Data:    33,
	//	}, &core.PieEntry{
	//
	//		Caption: "3",
	//		Data:    50,
	//	})
}
