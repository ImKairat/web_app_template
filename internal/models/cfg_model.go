package models

type AppConfig struct {
	Psql struct {
		Host      string `yaml:"Host" default:"localhost"`
		Port      string `yaml:"Port" default:"5432"`
		User      string `yaml:"User" default:"postgres"`
		Password  string `yaml:"Password" default:"postgres"`
		DbName    string `yaml:"DbName" default:"postgres"`
		TableName string `yaml:"TableName" default:"users"`
	}
}
