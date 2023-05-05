package core

type PieEntry struct {
	Caption    string
	Data       float64
	proportion float64
}

// calculates the proportions each datapoint occupies and returns a pie chart accordingly
func (i Iskre) Pie(d ...*PieEntry) string {
	total := 0.0
	for _, e := range d {
		total += e.Data
	}

	for _, e := range d {
		// calculates the proportion of each data entry
		// https://sciencing.com/work-out-percentages-pie-chart-4777890.html
		e.proportion = e.Data / total * 360
	}

	return ""
}
