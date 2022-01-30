package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"randomtify_app/clients"
	"randomtify_app/commands"
	"randomtify_app/entities"
	fl "randomtify_app/flags"
)

const (
	helpInfo = `Usage: {command} [{flag "value"} {flag "value"} ...]
 Commands with its flags:
  help - shows help info
      available flags: none
  available-alphabets - shows all supported alphabets
      available flags: none
  search - searches random artist
      available flags:
          //if no flags are used, all values will be random
          //all these flags can be used together at the same time
       -q, --query - specifies a query for search
       -al, --alphabet - specifies the alphabet for search
       -ca, --chars-amount - specifies the amount of characters in a random search query
  show-artist - shows artist's info
      available flags:
       -a, --all - shows all artists that were found (no value required)
       -n, --name - shows artist's info with its songs that were added`
)

type requestSender func(flags map[string]string, client clients.RandomtifyClient) (err error, res interface{})

var funcs = map[string]requestSender{
	commands.AvailableAlphabetsCommand: getAlphabets,
	commands.HelpCommand:               showHelpInfo,
	commands.ShowArtistCommand:         getArtistsInfo,
	commands.SearchCommand:             getRandomArtist,
}

type Processor interface {
	Process(command string, flags map[string]string)
}

type processor struct {
	client clients.RandomtifyClient
}

func GetProcessor() Processor {
	return &processor{clients.GetRandomtifyClient()}
}

func (p *processor) Process(command string, flags map[string]string) {
	f := funcs[command]
	err, result := f(flags, p.client)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}

func getAlphabets(_ map[string]string, client clients.RandomtifyClient) (err error, res interface{}) {
	resp, err := client.GetAlphabets()
	if err != nil {
		return
	}

	body, err := getBodyAsBytes(resp)
	if err != nil {
		return
	}
	res = new(entities.Alphabets)
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}
	return
}

func showHelpInfo(_ map[string]string, _ clients.RandomtifyClient) (err error, res interface{}) {
	return nil, helpInfo
}

func getArtistsInfo(flags map[string]string, client clients.RandomtifyClient) (err error, res interface{}) {
	name := getFlagValue(fl.NameFullFlag, fl.NameShortFlag, &flags)
	if name != "" {
		res = new(entities.ArtistExtendedInfo)
	} else {
		res = new(entities.AllArtistsInfo)
	}
	resp, err := client.GetArtist(name)
	if err != nil {
		return
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, "Nothing is found"
	}

	body, err := getBodyAsBytes(resp)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}
	return
}

func getRandomArtist(flags map[string]string, rc clients.RandomtifyClient) (err error, res interface{}) {
	query := getFlagValue(fl.QueryFullFlag, fl.QueryShortFlag, &flags)
	alphabet := getFlagValue(fl.AlphabetFullFlag, fl.AlphabetShortFlag, &flags)
	charsAmount := getFlagValue(fl.CharsAmountFullFlag, fl.CharsAmountShortFlag, &flags)

	resp, err := rc.GetRandomArtist(query, alphabet, charsAmount)
	if err != nil {
		return
	}
	body, err := getBodyAsBytes(resp)
	if err != nil {
		return
	}
	res = new(entities.RandomArtistInfo)
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}
	return
}

func getBodyAsBytes(resp *http.Response) ([]byte, error) {
	defer func(body io.ReadCloser) {
		e := body.Close()
		if e != nil {
			fmt.Println(e.Error())
		}
	}(resp.Body)

	return ioutil.ReadAll(resp.Body)
}

func getFlagValue(firstFlagName, secondFlagName string, flags *map[string]string) string {
	name := (*flags)[firstFlagName]
	if name == "" {
		name = (*flags)[secondFlagName]
	}
	return name
}
