package main

import (
	"fmt"

	"github.com/iskre/lib/core"
)

func main() {
	i := core.Iskre{}
	fmt.Println("Scatter example")
	fmt.Println(i.Scatter(core.ScatterChart{
		CaptionX: "CaptionX",
		CaptionY: "CaptionY",
		Title:    "Title",
		Data: [][]float64{
			{0, 0},
			{1, 1},
			{2, 2},
			{3, 3},
			{4, 4},
			{5, 5},
			{6, 6},
			{7, 7},
			{8, 8},
			{9, 9},
			{10, 10},
		},
	}))

	// 	fmt.Println("Box example: ")
	// 	r, _ := helper.Box("title\nsubtitle")
	// 	fmt.Println(r)

	// 	fmt.Println("Repetition example: ")
	// 	fmt.Println(i.Repetition("While loop conditiontest", "test\n test\n  test\ntest"))

	// 	fmt.Println("Sequence example:")
	// 	fmt.Println(i.Sequence("test short", "test very very very very very very very very very very very very very very very very long", "..."))

	// 	fmt.Println("Pie chart example:")
	// 	fmt.Println(i.Pie(&core.PieEntry{
	// 		Caption: "1",
	// 		Data:    33,
	// 	}, &core.PieEntry{

	// 		Caption: "2",
	// 		Data:    33,
	// 	}, &core.PieEntry{

	// 		Caption: "3",
	// 		Data:    50,
	// 	}))

	// fmt.Println("Table example:")
	//
	//	fmt.Println(i.Table(&core.Table{
	//		Headers: []string{"THE FIRST IS VERY LONG", "Second", "THE THIRD ALSO"},
	//		Body: [][]string{
	//			{"Test1", "Test2", "Test3"},
	//			{"Test1", "Test2", "Test3"},
	//		},
	//	}))
}
