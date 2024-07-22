package main

import (
	"log/slog"
	"mytechblog/test"
)
func main() {
	slog.Info("The Main Is Deubugging...")
	test.Debug()	

	slog.Info("The Server Is Lauching...")
	test.Server()
}