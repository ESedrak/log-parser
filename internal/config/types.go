package config

type Config struct {
	Limit Limit `json:"limit"`
	Regex Regex `json:"regex"`
	Path  Path  `json:"path"`
}

type Limit struct {
	MaxIPs  int `json:"maxIPs"`
	MaxURLs int `json:"maxURLs"`
}

type Regex struct {
	MatchIPsURlsIgnoreQuery string `json:"matchIPsURLsIgnoreQuery"`
}

type Path struct {
	LogPath string `json:"logpath"`
}
