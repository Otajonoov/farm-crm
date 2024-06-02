package domain

type Section struct {
	name string
}

func (s *Section) Name() string {
	return s.name
}

func (s *Section) SetName(name string) {
	s.name = name
}
