package core

import (
	"strings"
)

type Table struct {
	Headers []string
	Body    [][]string
}

func (i Isrke) Table(d *Table) string {
	// m = max length for full row
	m := 0
	b := strings.Builder{}

	var headerLengths []int

	b.WriteRune(TABLE_TOP_LEFT)

	// write table header
	for i, e := range d.Headers {
		ml := len(e) + 8
		m += ml
		b.WriteString(strings.Repeat(string(TABLE_HORIZONTAL), ml))

		headerLengths = append(headerLengths, ml)

		if i != len(d.Headers)-1 {
			b.WriteRune(TABLE_INTERSECTION_TOP)
		} else {
			b.WriteRune(TABLE_TOP_RIGHT)
			b.WriteRune('\n')
		}
	}

	for i, e := range d.Headers {
		b.WriteRune(TABLE_VERTICAL)
		b.WriteString("    " + e + "    ")
		if i == len(d.Headers)-1 {
			b.WriteRune(TABLE_VERTICAL)
		}
	}

	b.WriteString("\n")
	b.WriteRune(TABLE_INTERSECTION_LEFT)
	b.WriteRune(TABLE_HORIZONTAL)

	// write horizontal line to close table header
	for i, e := range d.Headers {
		ml := len(e) + 8

		if i == 0 {
			ml--
		}

		b.WriteString(strings.Repeat(string(TABLE_HORIZONTAL), ml))

		if i == len(d.Headers)-1 {
			b.WriteRune(TABLE_INTERSECTION_RIGHT)
		} else {
			b.WriteRune(TABLE_INTERSECTION)
		}
	}

	for i, e := range d.Body {
		b.WriteString("\n")
		b.WriteRune(TABLE_VERTICAL)
		for y, s := range e {
			bodyString := " " + s + " "
			b.WriteString(" " + s + " ")

			if len(bodyString) != headerLengths[y] {
				b.WriteString(strings.Repeat(" ", headerLengths[y]-len(bodyString)))
				if i == 0 {
					b.WriteRune(TABLE_VERTICAL)
				} else {
					b.WriteRune(TABLE_VERTICAL)
				}
			}
		}

		// write last line
		if i == len(d.Body)-1 {
			b.WriteString("\n")
			b.WriteRune(TABLE_BOTTOM_LEFT)

			for y, e := range headerLengths {
				b.WriteString(strings.Repeat(string(TABLE_HORIZONTAL), e))

				if y == len(headerLengths)-1 {
					b.WriteRune(TABLE_BOTTOM_RIGHT)
				} else {
					b.WriteRune(TABLE_INTERSECTION_BOTTOM)
				}
			}

		}
	}

	return b.String()
}
