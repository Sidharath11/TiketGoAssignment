package main

import (
	"university/app"
	"university/domain"
)

func main() {
	app.StartApp()
	defer domain.ClosePgSQL()
}
