package main

import (
	"fmt"
	"log"
	"os"

	"./app/setup/prompts"
)

func main() {

	port := prompts.Port()
	password := prompts.Password()
	prompts.ConfirmPassword()

	fmt.Printf("The setup has been saved to .env")

}

func writeConfig(port string, path string, services string, iplookup string) {

	f, err := os.OpenFile(".env", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}

	defer f.Close()
	f.WriteString(fmt.Sprintf("SWERVER_PORT=%v\n", port))
	f.WriteString(fmt.Sprintf("SWERVER_PATH=%v\n", path))
	f.WriteString(fmt.Sprintf("SWERVER_SERVICES=%v\n", services))
	f.WriteString(fmt.Sprintf("SWERVER_IPLOOKUP_URL=%v\n", iplookup))

}
