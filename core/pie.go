package core

import (
	"fmt"
	"github.com/iskre/lib/helper"
	"strings"
)

type PieEntry struct {
	Caption    string
	Data       float64
	proportion float64
}

// calculates the proportions each datapoint occupies and returns a pie chart with a legend accordingly
func (i Iskre) Pie(d ...*PieEntry) string {
	b := strings.Builder{}
	b.WriteString("Legend: \n")
	total := 0.0
	for _, e := range d {
		total += e.Data
		b.WriteString(fmt.Sprintf("\n\t\t\t%s: %.2f", e.Caption, e.Data))
	}
	b.WriteString(fmt.Sprintf("\n\n\tTotal: %.2f", total))

	legend, _ := helper.Box(b.String())

	for _, e := range d {
		// calculates the proportion of each data entry
		// https://sciencing.com/work-out-percentages-pie-chart-4777890.html
		e.proportion = e.Data / total * 360
	}
	// b.Reset()

	// return b.String()
	return legend
}
