package main

import (
	"imports/dependencies/app/controller"
	"imports/dependencies/app/model"
)

func main() {
	model.Init()
	controller.Start()
}
