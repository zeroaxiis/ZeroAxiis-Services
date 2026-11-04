package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                  string
	MongoURI              string
	RedisURI              string
	JWTSecret             string
	CloudinaryCloudName   string
	CloudinaryAPIKey      string
	CloudinaryAPISecret   string
	AppEnv                string
	PublicFrontend        string
	AdminFrontend         string
	YouTubeAPIKey         string
	HalomailWebhookSecret string
}

func Load() (*Config, error) {
	//config laod and it does not exist it just ignores it
	_ = godotenv.Load()

	//creating a pointer to Config
	cfg := &Config{
		Port:                  os.Getenv("PORT"),
		MongoURI:              os.Getenv("MONGO_URI"),
		RedisURI:              os.Getenv("REDIS_URI"),
		JWTSecret:             os.Getenv("JWT_SECRET"),
		CloudinaryCloudName:   os.Getenv("CLOUDINARY_CLOUD_NAME"),
		CloudinaryAPIKey:      os.Getenv("CLOUDINARY_API_KEY"),
		CloudinaryAPISecret:   os.Getenv("CLOUDINARY_API_SECRET"),
		AppEnv:                os.Getenv("APP_ENV"),
		PublicFrontend:        os.Getenv("PUBLIC_FRONTEND"),
		AdminFrontend:         os.Getenv("ADMIN_FRONTEND"),
		YouTubeAPIKey:         os.Getenv("YouTubeAPIKey"),
		HalomailWebhookSecret: os.Getenv("HALOMAIL_WEBHOOK_SECRET"), 
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

var appConfig *Config

func MustLoad() *Config {

	if appConfig != nil {
		return appConfig
	}

	cfg, err := Load()
	if err != nil {
		panic(err)
	}

	appConfig = cfg

	return appConfig
}

// here func (c *Config) --> this means attach me to Config struct
func (c *Config) validate() error {
	switch {
	case c.Port == "":
		return errors.New("PORT is Required")
	case c.MongoURI == "":
		return errors.New("MONGODB_URI is Required")
	case c.RedisURI == "":
		return errors.New("REDIS_URI is Required")
	case c.JWTSecret == "":
		return errors.New("JWT_SECRET is Required")
	case c.CloudinaryCloudName == "":
		return errors.New("CLOUDINARY_CLOUD_NAME is Required")
	case c.CloudinaryAPIKey == "":
		return errors.New("CLOUDINARY_API_KEY is Required")
	case c.CloudinaryAPISecret == "":
		return errors.New("CLOUDINARY_API_SECRET is Required")
	case c.AppEnv == "":
		return errors.New("Logger Env is required")
	case c.PublicFrontend == "":
		return errors.New("Public Frontend URL is needed for CORS")
	case c.AdminFrontend == "":
		return errors.New("Admin Frontend url is needed for Data Management")
	case c.YouTubeAPIKey == "":
		return errors.New("YouTube API Key is Required")
	case c.HalomailWebhookSecret == "":
		return errors.New("Halomail webhook secert is missing")
	}
	
	return nil
}
