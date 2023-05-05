// contains control flow diagrams
package core

import "strings"

// generates a Repetition diagram:
//
//		i := core.Iskre{}
//		fmt.Println(i.Repetition("While loop conditiontest", "test\n test\n  test\ntest"))
//
//	    ┌───────────────────────────────┐
//	    │  While loop conditiontest     │
//	    │  ┌────────────────────────────┤
//	    │  │  test                      │
//	    │  │   test                     │
//	    │  │    test                    │
//	    │  │  test                      │
//	    └──┴────────────────────────────┘
func (i Iskre) Repetition(head, body string) string {
	// keep track of the longest string, to resize the tableccordingly
	m := 0
	if len(head) > len(body) {
		m = len(head)
	} else {
		if strings.ContainsRune(body, '\n') {
			for _, e := range strings.Split(body, "\n") {
				if len(e) > m {
					m = len(e)
				}
			}
		} else {
			m = len(body)
		}
	}

	// add padding before and after
	m += 4

	b := strings.Builder{}

	// top
	b.WriteRune(TABLE_TOP_LEFT)
	b.WriteString(strings.Repeat(string(TABLE_HORIZONTAL), m+3))
	b.WriteRune(TABLE_TOP_RIGHT)

	// head
	b.WriteRune('\n')
	b.WriteRune(TABLE_VERTICAL)
	b.WriteString("  ")
	b.WriteString(head)
	b.WriteString(strings.Repeat(" ", m-(len(head)-1)))
	b.WriteRune(TABLE_VERTICAL)

	// bar between head and body
	b.WriteRune('\n')
	b.WriteRune(TABLE_VERTICAL)
	b.WriteString("  ")
	b.WriteRune(TABLE_TOP_LEFT)
	b.WriteString(strings.Repeat(string(TABLE_HORIZONTAL), m))
	b.WriteRune(TABLE_INTERSECTION_RIGHT)

	// body
	if strings.ContainsRune(body, '\n') {
		for _, e := range strings.Split(body, "\n") {
			b.WriteRune('\n')
			b.WriteRune(TABLE_VERTICAL)
			b.WriteString("  ")
			b.WriteRune(TABLE_VERTICAL)
			b.WriteString("  ")
			b.WriteString(e)
			b.WriteString(strings.Repeat(" ", m-(len(e)+2)))
			b.WriteRune(TABLE_VERTICAL)
		}
	} else {
		b.WriteRune('\n')
		b.WriteRune(TABLE_VERTICAL)
		b.WriteString("  ")
		b.WriteRune(TABLE_VERTICAL)
		b.WriteString("  ")
		b.WriteString(body)
		b.WriteString(strings.Repeat(" ", m-(len(body)+2)))
		b.WriteRune(TABLE_VERTICAL)
	}

	// bottom
	b.WriteRune('\n')
	b.WriteRune(TABLE_BOTTOM_LEFT)
	b.WriteRune(TABLE_HORIZONTAL)
	b.WriteRune(TABLE_HORIZONTAL)
	b.WriteRune(TABLE_INTERSECTION_BOTTOM)
	b.WriteString(strings.Repeat(string(TABLE_HORIZONTAL), m))
	b.WriteRune(TABLE_BOTTOM_RIGHT)
	return b.String()
}

// generates a sequence table like diagram:
//
//	 i := core.Iskre{}
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
func (i Iskre) Sequence(d ...string) string {
	// keep track of the longest string, to resize the table accordingly
	m := 0
	for _, e := range d {
		if len(e) > m {
			m = len(e)
		}
	}
	// add padding before and after
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
