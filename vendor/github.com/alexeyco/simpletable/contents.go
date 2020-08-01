package simpletable

import "sort"

// Header is table header
type Header struct {
	Cells []*Cell
}

// Body is table body
type Body struct {
	Cells     [][]*Cell
	sortField int
}

// Footer is table footer
type Footer struct {
	Cells []*Cell
}

func (b *Body) Len() int {
	return len(b.Cells)
}

func (b *Body) Less(i, j int) bool {
	return b.Cells[i][b.sortField].Text < b.Cells[j][b.sortField].Text
}

func (b *Body) Swap(i, j int) {
	b.Cells[i], b.Cells[j] = b.Cells[j], b.Cells[i]
}

func (b *Body) SortByField(fieldNum int) {
	b.sortField = fieldNum
	sort.Sort(b)
}
