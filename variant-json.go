package csljson

import "strconv"

// Variant parses any primitive json type (string, number, boolean, null) as a string
type Variant string

func (v *Variant) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		s = string(b)
	}
	*v = Variant(s)
	return nil
}
