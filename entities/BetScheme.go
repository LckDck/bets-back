package entities

// BetScheme is a structure of the current bet
type BetScheme struct {
	ID     int
	Name   string
	Fields []Field
	Items  []Items
}

// AddField adds a field to the scheme
func (s BetScheme) AddField(f Field) {
	s.Fields = append(s.Fields, f)
}

// RemoveField removes a field by id from the scheme
func (s BetScheme) RemoveField(id int) {
	for i, v := range s.Fields {
		if v.ID == id {
			s.Fields = append(s.Fields[:i], s.Fields[i+1:]...)
			return
		}
	}
}

