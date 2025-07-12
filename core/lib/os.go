package lib

import "os"

type Env struct{}

func (e *Env) GetEnv(k, f string) string {
	if v, ok := os.LookupEnv(k); ok {
		return v
	}
	return f
}
