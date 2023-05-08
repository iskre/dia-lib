package helper

import (
	"strings"
)

func RepeatRune(c rune, n int) string {
	b := strings.Builder{}
	for i := 0; i < n; i++ {
		b.WriteRune(c)
	}
	return b.String()
}

// builds a box around the given content, supports linebreaks, tabs, etc. Returns the box and the length of the longest row
//
//	 r, _ := helper.Box("title\nsubtitle")
//	 fmt.Println(r)
//
//		┌────────────┐
//		│  title     │
//		│  subtitle  │
//		└────────────┘
func Box(content string) (string, int) {
	b := strings.Builder{}
	content = strings.ReplaceAll(content, "\t", "  ")
	d := strings.Split(content, "\n")
	m := 0

	for _, e := range d {
		if len(e) > m {
			m = len(e)
		}
	}
	m += 4

	b.WriteRune(TABLE_TOP_LEFT)
	b.WriteString(strings.Repeat(string(TABLE_HORIZONTAL), m))
	b.WriteRune(TABLE_TOP_RIGHT)

	for _, e := range d {
		b.WriteRune('\n')
		b.WriteRune(TABLE_VERTICAL)
		b.WriteString("  ")
		b.WriteString(e)
		b.WriteString(strings.Repeat(" ", m-(len(e)+2)))
		b.WriteRune(TABLE_VERTICAL)
	}

	b.WriteRune('\n')
	b.WriteRune(TABLE_BOTTOM_LEFT)
	b.WriteString(strings.Repeat(string(TABLE_HORIZONTAL), m))
	b.WriteRune(TABLE_BOTTOM_RIGHT)
	return b.String(), m
}
