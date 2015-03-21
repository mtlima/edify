package edifact

type Segment struct {
	Name     string
	Elements []*Element
}

func (s *Segment) String() string {
	elementsStr := ""
	for _, e := range s.Elements {
		elementsStr += "\t\t" + e.String() + "\n"
	}
	return s.Name + "\n" + elementsStr
}

func (s *Segment) AddElement(element *Element) {
	s.Elements = append(s.Elements, element)
}

func NewSegment(name string) *Segment {
	return &Segment{name, []*Element{}}
}
