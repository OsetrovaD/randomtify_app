package services

import "os"

type Config struct {
	RandomtifyAppUrl string
	AlphabetsPath    string
	SearchPath       string
	ArtistsPath      string
}

func GetConfig() *Config {
	return &Config{
		RandomtifyAppUrl: getEnv("RANDOMTIFY_APP_URL", ""),
		AlphabetsPath:    getEnv("ALPHABETS_PATH", ""),
		SearchPath:       getEnv("SEARCH_PATH", ""),
		ArtistsPath:      getEnv("ARTISTS_PATH", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
