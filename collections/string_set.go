package collections

type StringSet struct {
	data map[string]bool
}

func NewStringSet() StringSet {
	return StringSet{
		data: make(map[string]bool),
	}
}

func (s *StringSet) Has(str string) bool {
	if _, ok := s.data[str]; ok {
		return true
	}

	return false
}

func (s *StringSet) Add(str string) {
	s.data[str] = true
}

func (s *StringSet) GetData() []string {
	var result []string

	for str := range s.data {
		result = append(result, str)
	}

	return result
}
