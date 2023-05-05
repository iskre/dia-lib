package main

import (
	"github.com/iskre/lib/core"
)

func main() {
	i := core.Isrke{}
	i.Pie(&core.PieEntry{
		Caption: "1",
		Data:    33,
	}, &core.PieEntry{
		Caption: "2",
		Data:    33,
	}, &core.PieEntry{
		Caption: "3",
		Data:    50,
	})
}
