package main

import (
	"fmt"

	"github.com/leocaversan/micro-services-goLang/models"
	"github.com/leocaversan/micro-services-goLang/routes"
)

func main() {
	models.Personalities = []models.Personalitie{
		{Id: 1, Name: "name-1", History: "history-1"},
		{Id: 2, Name: "name-2", History: "history-2"},
	}
	fmt.Println("servidor executando")
	routes.HandlerRequest()
}
