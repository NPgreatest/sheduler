package config

type Server struct {
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	ApiKeys ApiKeys `mapstructure:"api_keys" json:"api_keys" yaml:"api_keys"`
}
