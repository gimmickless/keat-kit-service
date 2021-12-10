package enum

type SortDirection int

const (
	Asc SortDirection = iota
	Desc
)

func (d SortDirection) String() string {
	return [...]string{"Asc", "Desc"}[d]
}
