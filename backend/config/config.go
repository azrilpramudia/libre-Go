package config

import (
	"errors"
	"net"
	"net/url"
	"os"
	"strconv"
)

func AppAddress() (string, error) {
	port := os.Getenv("APP_PORT")
	portNumber, err := strconv.Atoi(port)
	if err != nil || portNumber < 1 || portNumber > 65535 {
		return "", errors.New("APP_PORT must be a valid port number between 1 and 65535")
	}

	return ":" + port, nil
}

func DatabaseURL() (string, error) {
	if databaseURL := os.Getenv("DATABASE_URL"); databaseURL != "" {
		return databaseURL, nil
	}

	user := os.Getenv("DB_USER")
	database := os.Getenv("DB_NAME")
	if user == "" || database == "" {
		return "", errors.New("set DATABASE_URL or both DB_USER and DB_NAME")
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}
	if _, err := strconv.Atoi(port); err != nil {
		return "", errors.New("DB_PORT must be a valid port number")
	}

	sslMode := os.Getenv("DB_SSLMODE")
	if sslMode == "" {
		sslMode = "disable"
	}

	connectionURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(user, os.Getenv("DB_PASSWORD")),
		Host:   net.JoinHostPort(host, port),
		Path:   database,
	}
	query := connectionURL.Query()
	query.Set("sslmode", sslMode)
	connectionURL.RawQuery = query.Encode()

	return connectionURL.String(), nil
}
