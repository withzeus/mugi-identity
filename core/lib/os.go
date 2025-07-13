package lib

import "os"

type OS struct{}

func (e *OS) GetEnv(k, f string) string {
	if v, ok := os.LookupEnv(k); ok {
		return v
	}
	return f
}
