package analysis

type State struct {
	// documentURI -> text
	Documents map[string]string
}

func NewState() State {
	return State{
		Documents: make(map[string]string),
	}
}

func (s *State) OpenDocument(documentURI string, text string) {
	s.Documents[documentURI] = text
}

func (s *State) UpdateDocument(documentURI string, text string) {
	s.Documents[documentURI] = text
}
