package pushkin

type ProducerConfig struct {
	Name    string
	Cluster []string
}

type ConnConfig struct {
	Host string
}