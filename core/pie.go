package core

type PieEntry struct {
	Caption string
	Data    int
}

func (i Isrke) Pie(d ...PieEntry) string {
	total := 0
	for _, e := range d {
		total += e.Data
	}
	return ""
}
