package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	fl "randomtify_app/flags"
	"randomtify_app/services/commands"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
		return
	}
	command, flags, err := fl.ParseCommandLineArgs(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	commands.GetProcessor().Process(command, flags)
}
