package core

import (
	"strings"

	"github.com/iskre/lib/helper"
)

type BarChartEntry struct {
	Caption string
	Data    int64
}

func (i Iskre) HorizontalBarChart(d ...*BarChartEntry) string {
	b := strings.Builder{}
	m := 0

	for _, e := range d {
		if len(e.Caption) > m {
			m = len(e.Caption)
		}
	}

	for _, e := range d {
		captionString := "\n" + e.Caption + "     |"
		if len(e.Caption) < m {
			captionString = "\n" + e.Caption + strings.Repeat(" ", m-len(e.Caption)+4) + " |"
		}

		bar := ""

		bar_chunks, remainder := divmod(e.Data, 8)
		bar = bar + strings.Repeat(string(helper.BLOCK_LEFT_FULL), int(bar_chunks))

		if remainder > 0 {
			bar = bar + string(rune('█'+(8-remainder)))
		}

		b.WriteString(captionString + " " + bar)
	}
	return b.String()
}

func divmod(numerator, denominator int64) (quotient, remainder int64) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}
