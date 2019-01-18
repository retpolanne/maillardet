package tex

// Cover is the struct that represents the document cover
type Cover struct {
	Title       string
	Institution []string
	City        string
	Period      Period
	Authors     []string
	Orientation string
	Area        string
}

type Period struct {
}

// GenerateCover will generate the cover section from a template to a string
func (cover *Cover) GenerateCover() string {
	return ""
}

// GenerateSubcover will generate the subcover from a template to a string
func (cover *Cover) GenerateSubcover() string {
	return ""
}
