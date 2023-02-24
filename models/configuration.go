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
	TxtSalary       string
	Occ             string
	ABE_UTM_EMPRESA string

	//APIs
	RetryWaitMin int
	RetryWaitMax int
	RetryMax     int
	APITimeout   int
	APIAtreel    string `env:"API_ATREEL"`

	// Database
	DBTimeout         int
	DBUser            string `env:"DB_USER"`
	DBPassword        string `env:"DB_PASSWORD"`
	DBName            string `env:"DB_NAME"`
	DBPort            int    `env:"DB_PORT"`
	DBHost            string `env:"DB_HOST"`
	DBMaxOpenConns    int    `env:"DB_MaxOpenConns"`
	DBMaxIdleConns    int    `env:"DB_MaxIdlesConns"`
	DBConnMaxLifetime int    `env:"DB_ConnMaxLifetime"`
}
