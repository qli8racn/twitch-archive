package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Twitch struct {
		ClientID     string `mapstructure:"client_id"`
		ClientSecret string `mapstructure:"client_secret"`
		RedirectURI  string `mapstructure:"redirect_uri"`
	} `mapstructure:"twitch"`
	AWS struct {
		S3 struct {
			Endpoint string `mapstructure:"endpoint"`
			Buckets struct {
				TwitchArchives struct {
					BucketName string `mapstructure:"bucketName"`
				} `mapstructure:"twitch_archives"`
			} `mapstructure:"buckets"`
		} `mapstructure:"s3"`
	} `mapstructure:"aws"`
}

func New() (cfg *Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("internal/config")

	if err = viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err = viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
