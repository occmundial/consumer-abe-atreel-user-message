package models

// Configuration :
type Configuration struct {
	KafkaBrokers      []string
	GroupID           string
	Topics            []string
	QueueRequestDelay int64
	QueueTimeout      int
	ArtifactVersion   string `env:"VERSION"`

	// Dependencies
	TxtSalary     string
	Occ           string
	AbeUtmEmpresa string
	ABEUTMJob     string

	// APIs
	RetryWaitMin int
	RetryWaitMax int
	RetryMax     int
	APITimeout   int
	APIAtreel    string `env:"API_ATREEL"`
	APITlaloc    string `env:"API_TLALOC"`
	TokenTlaloc  string `env:"TLALOC_TOKEN"`
}
