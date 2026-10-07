package main

import (
	"fmt"
	keyvalues "server/keyValues"
	"strings"
)

func dispatch(parts []string, kv *keyvalues.KeyValues) string {
	switch strings.ToUpper(parts[0]) {
	case "SET":
		kv.Set(parts[1], parts[2])
		return "+OK\n"
	case "GET":
		value, exists := kv.Get(parts[1])
		if !exists {
			return "$-1\n"
		}
		return fmt.Sprintf("OK %d %s\n", len(value), value)
	case "DEL":
		deleted := kv.Del(parts[1])
		if deleted {
			return ":1\n"
		}
		return ":0\n"
	case "EXISTS":
		exists := kv.Exists(parts[1])

		if exists {
			return ":1\n"
		}
		return "0\n"

	}
	return "-ERR unknown command\n"
}
