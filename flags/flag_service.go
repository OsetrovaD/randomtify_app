package flags

import (
	"github.com/pkg/errors"
	"randomtify_app/commands"
	"strings"
)

const unknownCommandInfoMessage = "Unknown command '%s'. Please, use command 'help' to show all available commands"

func ParseCommandLineArgs(osArgs []string) (command string, flags map[string]string, err error) {
	if len(osArgs) == 1 {
		err = errors.Errorf("No command used. Please, use command 'help' to show all available commands")
		return
	}
	if _, ok := CommandsFlags[osArgs[1]]; !ok {
		err = errors.Errorf(unknownCommandInfoMessage, osArgs[1])
		return
	}

	command = osArgs[1]
	if command == commands.HelpCommand {
		return
	}

	flags = map[string]string{}
	if len(osArgs) > 2 {
		for i := 2; i < len(osArgs); i++ {
			arg := osArgs[i]
			if _, ok := CommandsFlags[command][arg]; !ok {
				err = errors.Errorf("Flag '%s' is not available for command '%s'. Please, use command 'help' to show commands info", arg, command)
				return
			}
			if _, ok := flags[FlagsAliases[arg]]; ok {
				err = errors.Errorf("It's forbidden to use the same flag twice")
				return
			}
			if _, ok := FlagsWithoutValue[arg]; ok {
				if err = addFlagWithoutValue(&osArgs, &i, &arg, &flags); err != nil {
					return
				}
				continue
			}
			if err = addFlag(&osArgs, &i, &arg, &flags); err != nil {
				return
			}
		}
	}
	return
}

func addFlag(osArgs *[]string, index *int, flagName *string, flags *map[string]string) error {
	if len(*osArgs) <= *index+1 {
		return errors.Errorf("'%s' flag doesn't have value", *flagName)
	}
	(*flags)[*flagName] = (*osArgs)[*index+1]
	*index++
	return nil
}

func addFlagWithoutValue(osArgs *[]string, index *int, flagName *string, flags *map[string]string) error {
	if len(*osArgs) > *index+1 && !strings.HasPrefix((*osArgs)[*index+1], "--") && !strings.HasPrefix((*osArgs)[*index+1], "-") {
		return errors.Errorf("Flag '%s' doesn't require value", *flagName)
	}
	(*flags)[*flagName] = ""
	return nil
}
