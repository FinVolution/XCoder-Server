package xmongo

type Config struct {
	c    *dbConfig
}

type dbConfig struct {
	Hosts      string `yaml:"hosts"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	AuthSource string `yaml:"authsource"`
	ReplicaSet string `yaml:"replicaSet"`
}
