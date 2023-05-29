package core

import (
	"fmt"
	"strings"

	"github.com/iskre/lib/helper"
)

type ScatterChart struct {
	CaptionX string
	CaptionY string
	Title    string
	Data     [][]float64
}

func (i *Iskre) Scatter(d ScatterChart) string {
	b := strings.Builder{}

	r, _ := helper.Box(d.Title)
	b.WriteString(r + "\n")

	mX := 0.0
	mY := 0.0

	for _, e := range d.Data {
		if e[0] > mX {
			mX = e[0]
		}
		if e[1] > mY {
			mY = e[1]
		}
	}

	mX += 5
	mY += 5
	m := make([][]string, int(mX))
	empty := fmt.Sprintf("%c ", helper.EMPTY_SLOT)
	filled := fmt.Sprintf("%c ", helper.FILLED_SLOT)

	b.WriteString("\n" + d.CaptionX)
	for i := 0; i < int(mX); i++ {
		m[i] = make([]string, int(mY))
		for j := 0; j < len(m[i]); j++ {
			r := empty
			if i == 0 {
				if j%2 == 0 {
					r = fmt.Sprintf("%02d ", j)
				} else {
					r = fmt.Sprintf("%2c ", ' ')
				}
			} else if j == 0 {
				if i%2 == 0 {
					r = fmt.Sprintf("%02d ", i)
				} else {
					r = fmt.Sprintf("%2c ", ' ')
				}
			}
			m[i][j] = r
		}
	}
	for _, e := range d.Data {
		m[int(e[0])+2][int(e[1])+2] = filled
	}

	b.WriteRune('\n')
	for i := 0; i < int(mX); i++ {
		for j := 0; j < int(mY); j++ {
			b.WriteString(m[i][j])
		}
		b.WriteRune('\n')
	}

	b.WriteString(d.CaptionY)

	return b.String()
}
