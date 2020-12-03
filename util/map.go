package util

type Map struct {
	Width  int
	Height int
	Rows   [][]int
}

func (m *Map) SymbolAtWrappedRight(x int, y int) int {
	return m.Rows[y][x%m.Width]
}
