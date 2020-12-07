package bags

type Bag struct {
	Color    string
	Contents map[string]int
}

func (b *Bag) Contains(color string) bool {
	_, ok := b.Contents[color]
	return ok
}
