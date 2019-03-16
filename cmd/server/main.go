package main

import (
	"github.com/emman27/wedding/internal/database"
	"github.com/emman27/wedding/pkg/api"
)

func main() {
	s := api.NewServer(database.NewInMemoryDB())
	s.Run()
}
