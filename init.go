package grset

func NewGrSet() *grSet {
	return &grSet{
		set: set{
			item: make(map[interface{}]setItem),
		},
	}
}
