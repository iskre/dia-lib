package core

import (
	"strings"

	"github.com/iskre/lib/helper"
)

type Table struct {
	Headers []string
	Body    [][]string
}

// generates a table
//
//	    i := core.Iskre{}
//		fmt.Println(i.Table(&core.Table{
//		Headers: []string{"First", "Second", "Third"},
//		Body: [][]string{
//			{"Test1", "Test2", "Test3"},
//			{"Test1", "Test2", "Test3"},
//		},
//		}))
//		┌─────────────┬──────────────┬─────────────┐
//		│    First    │    Second    │    Third    │
//		├─────────────┼──────────────┼─────────────┤
//		│ Test1       │ Test2        │ Test3       │
//		│ Test1       │ Test2        │ Test3       │
//		└─────────────┴──────────────┴─────────────┘
func (i Iskre) Table(d *Table) string {
	// TODO: add check for body length
	// m = max length for full row
	m := 0
	b := strings.Builder{}

	var headerLengths []int

	b.WriteRune(helper.TABLE_TOP_LEFT)

	// write table header
	for i, e := range d.Headers {
		ml := len(e) + 8
		m += ml
		b.WriteString(strings.Repeat(string(helper.TABLE_HORIZONTAL), ml))

		headerLengths = append(headerLengths, ml)

		if i != len(d.Headers)-1 {
			b.WriteRune(helper.TABLE_INTERSECTION_TOP)
		} else {
			b.WriteRune(helper.TABLE_TOP_RIGHT)
			b.WriteRune('\n')
		}
	}

	for i, e := range d.Headers {
		b.WriteRune(helper.TABLE_VERTICAL)
		b.WriteString("    " + e + "    ")
		if i == len(d.Headers)-1 {
			b.WriteRune(helper.TABLE_VERTICAL)
		}
	}

	b.WriteString("\n")
	b.WriteRune(helper.TABLE_INTERSECTION_LEFT)
	b.WriteRune(helper.TABLE_HORIZONTAL)

	// write horizontal line to close table header
	for i, e := range d.Headers {
		ml := len(e) + 8

		if i == 0 {
			ml--
		}

		b.WriteString(strings.Repeat(string(helper.TABLE_HORIZONTAL), ml))

		if i == len(d.Headers)-1 {
			b.WriteRune(helper.TABLE_INTERSECTION_RIGHT)
		} else {
			b.WriteRune(helper.TABLE_INTERSECTION)
		}
	}

	for i, e := range d.Body {
		b.WriteString("\n")
		b.WriteRune(helper.TABLE_VERTICAL)
		for y, s := range e {
			bodyString := " " + s + " "
			b.WriteString(" " + s + " ")

			if len(bodyString) != headerLengths[y] {
				b.WriteString(strings.Repeat(" ", headerLengths[y]-len(bodyString)))
				if i == 0 {
					b.WriteRune(helper.TABLE_VERTICAL)
				} else {
					b.WriteRune(helper.TABLE_VERTICAL)
				}
			}
		}

		// write last line
		if i == len(d.Body)-1 {
			b.WriteString("\n")
			b.WriteRune(helper.TABLE_BOTTOM_LEFT)

			for y, e := range headerLengths {
				b.WriteString(strings.Repeat(string(helper.TABLE_HORIZONTAL), e))

				if y == len(headerLengths)-1 {
					b.WriteRune(helper.TABLE_BOTTOM_RIGHT)
				} else {
					b.WriteRune(helper.TABLE_INTERSECTION_BOTTOM)
				}
			}

		}
	}

	return b.String()
}
