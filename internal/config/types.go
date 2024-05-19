package config

type Config struct {
	RequestedNum RequestedNum `json:"requestedNum"`
	Regex        Regex        `json:"regex"`
	Path         Path         `json:"path"`
}

type RequestedNum struct {
	IP  int `json:"ip"`
	URL int `json:"url"`
}

type Regex struct {
	MatchIPsURLsIgnoreQuery string `json:"matchIPsURLsIgnoreQuery"`
}

type Path struct {
	LogPath string `json:"logpath"`
}
