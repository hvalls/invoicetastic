package invoice

type Line struct {
	Name      string  `yaml:"name"`
	UnitPrice float64 `yaml:"unitPrice"`
	Qty       float64 `yaml:"qty"`
	Total     float64 `yaml:"-"` // Computed
}

func (l *Line) getTotal() float64 {
	return l.UnitPrice * l.Qty
}
