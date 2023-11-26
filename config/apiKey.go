package config

type Api struct {
	ApiId  string `mapstructure:"X-VO-Api-Id" json:"X-VO-Api-Id" yaml:"X-VO-Api-Id"`
	ApiKey string `mapstructure:"X-VO-Api-Key" json:"X-VO-Api-Key" yaml:"X-VO-Api-Key"`
}

type ApiCateGory int

type ApiKeys struct {
	Apis map[ApiCateGory]Api `mapstructure:"apis" json:"apis" yaml:"apis"`
}

const (
	OnlyRead ApiCateGory = iota
	ReadWrite
)
