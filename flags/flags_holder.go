package flags

import "randomtify_app/commands"

const (
	AllFullFlag          = "--all"
	AllShortFlag         = "-a"
	QueryFullFlag        = "--query"
	QueryShortFlag       = "-q"
	AlphabetFullFlag     = "--alphabet"
	AlphabetShortFlag    = "-al"
	CharsAmountFullFlag  = "--chars-amount"
	CharsAmountShortFlag = "-ca"
	NameFullFlag         = "--name"
	NameShortFlag        = "-n"
)

var (
	CommandsFlags = map[string]map[string]struct{}{
		commands.HelpCommand: {},
		commands.RandomtifyCommand: {
			QueryShortFlag:       {},
			QueryFullFlag:        {},
			AlphabetFullFlag:     {},
			AlphabetShortFlag:    {},
			CharsAmountFullFlag:  {},
			CharsAmountShortFlag: {},
		},
		commands.ShowArtistCommand: {
			AllFullFlag:   {},
			AllShortFlag:  {},
			NameFullFlag:  {},
			NameShortFlag: {},
		},
		commands.AvailableAlphabetsCommand: {},
	}
	FlagsAliases = map[string]string{
		AllFullFlag:          AllShortFlag,
		AllShortFlag:         AllFullFlag,
		NameShortFlag:        NameFullFlag,
		NameFullFlag:         NameShortFlag,
		QueryFullFlag:        QueryShortFlag,
		QueryShortFlag:       QueryFullFlag,
		AlphabetShortFlag:    AlphabetFullFlag,
		AlphabetFullFlag:     AlphabetShortFlag,
		CharsAmountShortFlag: CharsAmountFullFlag,
		CharsAmountFullFlag:  CharsAmountShortFlag,
	}
	FlagsWithoutValue = map[string]struct{}{
		AllFullFlag:  {},
		AllShortFlag: {},
	}
)
