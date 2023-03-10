package config

import (
	"fmt"
	"os"
	"strconv"
)

const (
	DefaultLimit    int    = 50
	DefaultOffset   int    = 0
	DefaultSortType string = "DESC"
	DefaultSortBy   string = "created_at"
)

var (
	Locale       string
	JwtSecretKey string
	JwtTTL       int
)

func getConfig() {
	Locale = os.Getenv("PJ_LOCALE")
	JwtSecretKey = os.Getenv("JWT_SECRET_KEY")

	// Get int
	var err error

	if JwtTTL, err = strconv.Atoi(os.Getenv("JWT_TTL")); err != nil {
		panic(fmt.Errorf("Get JWT_TTL error: %w", err))
	}
}
