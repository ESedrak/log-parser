package url_mgmt

type URLCounter struct {
	URLCounts map[string]int
}

// initializes a new instance of this struct
func NewUrlCounts() *URLCounter {
	return &URLCounter{
		URLCounts: make(map[string]int),
	}
}
