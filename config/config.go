package config

import "github.com/spf13/viper"

const envFileName = ".env"

type Config struct{}

func NewConfig() *Config {
	viper.SetConfigFile(envFileName)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	return &Config{}
}

func (c *Config) GetGrpcPort() string {
	return viper.GetString("SELF_GRPC_PORT")
}

func (c *Config) GetMandrillHost() string {
	return viper.GetString("MANDRILL_HOST")
}

func (c *Config) GetMandrillPort() string {
	return viper.GetString("MANDRILL_PORT")
}

func (c *Config) GetMandrillApiKey() string {
	return viper.GetString("MANDRILL_API_KEY")
}

func (c *Config) GetTemplatePath() string {
	return viper.GetString("TEMPLATE_PATH")
}

func (c *Config) GetSmtpHost() string {
	return viper.GetString("SMTP_HOST")
}

func (c *Config) GetSmtpPort() string {
	return viper.GetString("SMTP_PORT")
}

func (c *Config) GetSmtpPassword() string {
	return viper.GetString("SMTP_PASSWORD")
}

func (c *Config) GetFrom() string {
	return viper.GetString("FROM")
}

func (c *Config) GetKafkaBroker1() string {
	return viper.GetString("KAFKA_BROKER_1")
}

func (c *Config) GetKafkaBroker2() string {
	return viper.GetString("KAFKA_BROKER_2")
}

func (c *Config) GetTopicUserRegister() string {
	return viper.GetString("TOPIC_USER_REGISTER")
}
