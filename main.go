package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
)

var (
	args  []string
	flags map[string]string
)

func main() {
	if err := parseCommandLineArgs(os.Args); err != nil {
		panic(err)
	}
	fmt.Println(args)
	fmt.Println(flags)
}
func parseCommandLineArgs(osArgs []string) error {
	args = []string{}
	flags = map[string]string{}
	for i := 1; i < len(osArgs); i++ {
		arg := osArgs[i]
		switch {
		case strings.HasPrefix(arg, "--"):
			if err := addFlag(&osArgs, &i, &arg, 2); err != nil {
				return err
			}
		case strings.HasPrefix(arg, "-"):
			if err := addFlag(&osArgs, &i, &arg, 1); err != nil {
				return err
			}
		default:
			args = append(args, arg)
		}
	}
	return nil
}

func addFlag(osArgs *[]string, index *int, flagName *string, firstSymbolIndex int) error {
	if len(*osArgs) <= *index+1 {
		return errors.Errorf("'%s' flag doesn't have value", *flagName)
	}
	flags[(*flagName)[firstSymbolIndex:]] = (*osArgs)[*index+1]
	*index++
	return nil
}
