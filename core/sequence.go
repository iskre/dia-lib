package core

import "strings"

// generates a sequence table like diagram:
//
//	 i := core.Isrke{}
//	 fmt.Println(i.Sequence("First", "Second", "Third", "massive long Sequence"))
//
//	┌─────────────────────────┐
//	│  First                  │
//	├─────────────────────────┤
//	│  Second                 │
//	├─────────────────────────┤
//	│  Third                  │
//	├─────────────────────────┤
//	│  massive long Sequence  │
//	└─────────────────────────┘
func (i Isrke) Sequence(d ...string) string {
	m := 0
	for _, e := range d {
		if len(e) > m {
			m = len(e)
		}
	}
	m += 4

	b := strings.Builder{}

	b.WriteRune(TABLE_TOP_LEFT)
	b.WriteString(strings.Repeat(string(TABLE_HORIZONTAL), m))
	b.WriteRune(TABLE_TOP_RIGHT)

	for i, e := range d {
		b.WriteRune('\n')
		b.WriteRune(TABLE_VERTICAL)
		b.WriteString("  ")
		b.WriteString(e)
		b.WriteString(strings.Repeat(" ", m-(len(e)+2)))
		b.WriteRune(TABLE_VERTICAL)
		if i+1 < len(d) {
			b.WriteRune('\n')
			b.WriteRune(TABLE_INTERSECTION_LEFT)
			b.WriteString(strings.Repeat(string(TABLE_HORIZONTAL), m))
			b.WriteRune(TABLE_INTERSECTION_RIGHT)
		}
	}

	b.WriteRune('\n')
	b.WriteRune(TABLE_BOTTOM_LEFT)
	b.WriteString(strings.Repeat(string(TABLE_HORIZONTAL), m))
	b.WriteRune(TABLE_BOTTOM_RIGHT)

	return b.String()
}
