package config

type Config struct {
	Limit Limit
	Regex Regex
	Path  Path
}

type Limit struct {
	MaxIPs  int
	MaxURLs int
}

type Regex struct {
	MatchIPsURlsIgnoreQuery string
}

type Path struct {
	FilePath string
}
