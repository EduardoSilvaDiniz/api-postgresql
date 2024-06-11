package main

import (
	"fmt"
	"github.com/EduardoSilvaDiniz/api-postgresql/src/configs"
)

func main() {
	fmt.Println("Run")
	configs.Load().Error()
}
