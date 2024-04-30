package ip_mgmt

type IPCounter struct {
	IPCounts map[string]int
}

// initializes a new instance of this struct
func NewIPCounter() *IPCounter {
	return &IPCounter{
		IPCounts: make(map[string]int),
	}
}
