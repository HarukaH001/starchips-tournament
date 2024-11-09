package config

type Config struct {
	Debug bool `default:"false"`

	ClientServer ServerConfig `split_words:"true" required:"true"`
	AdminServer  ServerConfig `split_words:"true" required:"true"`

	Postgres PostgresConfig `split_words:"true" required:"true"`
}
