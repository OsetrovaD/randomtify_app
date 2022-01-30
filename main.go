package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path"
	fl "randomtify_app/flags"
	"randomtify_app/services/commands"
)

func main() {
	e, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = godotenv.Load(path.Dir(e) + "/.env"); err != nil {
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
