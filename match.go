package the_platinum_searcher

type match struct {
	pattern  []byte
	path     string
	lines    []line
	encoding int
}

type line struct {
	num  int
	text string
}

func (m *match) add(num int, text string) {
	m.lines = append(m.lines, line{num, text})
}
