package config

import "os"

var (
	PORT = prepare("PORT")
)

func prepare(name string) func() string {
	return func() string {
		return os.Getenv(name)
	}
}
