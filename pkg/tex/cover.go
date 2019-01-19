package tex

// Cover is the struct that represents the document cover
type Cover struct {
	Title       string
	Institution string
	CourseArea  string
	City        string
	Period      string
	Authors     []string
	Orientation string
}

// GenerateCover will generate the cover section from a template to a string
func (cover *Cover) GenerateCover() string {
	return ""
}
