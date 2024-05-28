package config

type Config struct {
	RequestedNum RequestedNum `mapstructure:"requestedNum"`
	Regex        Regex        `mapstructure:"regex"`
	Path         Path         `mapstructure:"path"`
}

type RequestedNum struct {
	IP  int `mapstructure:"ip"`
	URL int `mapstructure:"url"`
}

type Regex struct {
	MatchIPsURLsIgnoreQuery string `mapstructure:"matchIPsURLsIgnoreQuery"`
}

type Path struct {
	LogPath string `mapstructure:"logpath"`
}
