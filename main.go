package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexflint/go-arg"
	"os"
)

const version = "0.1"

type args struct {
	Command string `arg:"positional,help:<on|off|status|rlist|slist|rulist|sulist|apiserver|help>"`
	Relayid int    `arg:"positional,help:<id>"`
	Verbose bool   `arg:"-v,help:verbosity lego get vel"`
	Address string `arg:"-a,help:Address for api binding (default 127.0.0.1)"`
	Port    string `arg:"-p,help:port for api binding (default 9191)"`
}

func (args) Description() string {
	return "Apoui Home-Control CLI and APIServer <3"
}

var relays Relays
var sensors Sensors

func main() {
	ApouiHeader("Apoui")
	ApouiHeader("Control")
	fmt.Println("\t\t\t\t\t\t[ Version", version, "]")

	program := os.Args[0] + ":"
	var cli args
	cli.Relayid = -1
	cli.Command = "help"
	cli.Address = "127.0.0.1"
	cli.Port = "9191"
	arg.MustParse(&cli)

	commands := []string{"on", "off", "status", "rlist", "slist", "rulist", "sulist", "apiserver", "help"}

	cok := contains(commands, cli.Command)
	if true != cok {
		fmt.Println(program, "'"+cli.Command+"' is not a "+program+" command.")
		fmt.Println("See '" + program + " help")
		return
	}
	if cli.Command == "help" {
		arg.MustParse(&cli).WriteHelp(os.Stderr)
		return
	}

	datar, err := Asset("data/relays.json")
	if err != nil {
		fmt.Println("FATAL::NO_BASE_DB_JSON")
		return
	}
	json.Unmarshal(datar, &relays)

	datas, err := Asset("data/sensors.json")
	if err != nil {
		fmt.Println("FATAL::NO_BASE_DB_JSON")
		return
	}
	json.Unmarshal(datas, &sensors)

	if cli.Command == "rlist" {
		relays.list(false, false)
		return
	}
	if cli.Command == "slist" {
		sensors.list(false, false)
		return
	}

	if cli.Command == "rulist" {
		relays.list(true, false)
		return
	}
	if cli.Command == "sulist" {
		sensors.list(true, false)
		return
	}

	if cli.Command == "apiserver" {
		ApiServer(cli.Address, cli.Port, relays, sensors)
		return
	}

	if cli.Relayid == -1 {
		arg.MustParse(&cli).WriteHelp(os.Stderr)
		relays.list(false, false)
		return
	}

	if cli.Relayid > relays.rmax() {
		arg.MustParse(&cli).Fail("Wrong RELAYID: Over MAX")
	}
	if cli.Relayid < relays.rmin() {
		arg.MustParse(&cli).Fail("Wrong RELAYID: Under MIN")
	}

	relays.Command(cli.Relayid, cli.Command)

}
