package main

import (
	"fmt"
	"gig-echo-template/database/seeder"
)

func main() {
	seeder.Execute()
	fmt.Println("Database seeded!")
}
