package config

import "os"

var (
	PORT = SetupEnv("PORT")
)

func SetupEnv(name string) func() string {
	return func() string {
		return os.Getenv(name)
	}
}
