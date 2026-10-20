package main

import (
	"errors"
	"fmt"
)

var errMissing = errors.New("missing entry")

func lookup(settings map[string]string, key string) (string, error) {
	value, ok := settings[key]
	if !ok {
		return "", fmt.Errorf("setting %q: %w", key, errMissing)
	}
	return value, nil
}
func main() {
	_, err := lookup(map[string]string{"theme": "dark"}, "locale")
	if err != nil {
		fmt.Println(err)
		fmt.Println(errors.Is(err, errMissing))
	}
}
