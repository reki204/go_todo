package main

import (
	"fmt"

	"github.com/reki204/go_todo/app/controllers"
	"github.com/reki204/go_todo/app/models"
)

func main() {
	fmt.Println(models.Db)
	
	controllers.StartMainServer()
}
