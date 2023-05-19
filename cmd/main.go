package main

import (
	"gig-echo-template/pkg/database"
	"gig-echo-template/pkg/router"
)

func main() {
	router.Run(database.Connect())
}
