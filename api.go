package main

import (
	//"fmt"
	//"encoding/json"
	"fmt"
	//"github.com/briandowns/spinner"
	"github.com/gin-gonic/gin"
	//"github.com/olekukonko/tablewriter"
	//"os"
	"strconv"
	//"time"
)

type Apouiapi interface {
	ApiRelayCommand() string
}

func (ss Sensors) ApiRelayCommand(sensorid int, command string) string {

	//GetRid
	var sid int
	for sid = range ss {
		if ss[sid].ID == sensorid {
			break
		}
	}

	commands := []string{"status", "updateall", "update", "help"}

	cok := contains(commands, command)
	if true != cok {
		resp := "Command " + command + " is not an Sensor command. try 'help'"
		return resp
	}

	if command == "help" {
		resp := "Commands: status / update / updateall / help"
		return resp
	}

	if command == "updateall" {
		go func() { ss.updateStatuses(true) }()
		resp := "All Relays statuses update Scheduled !"
		return resp
	}

	if command == "update" {
		status := ss[sid].Url() //+ "/" + "status"
		ss[sid].Status = GetUrlQuiet(status)
		resp := "Relay status up to date !"
		return resp
	}

	var getUrl string
	getUrl = ss[sid].Url() + "/" + command
	resp := GetUrlQuiet(getUrl)
	go func() {
		status := ss[sid].Url() + "/" + "status"
		ss[sid].Status = GetUrlQuiet(status)
	}()
	return resp
}

func (rs Relays) ApiRelayCommand(relayid int, command string) string {

	//GetRid
	var rid int
	for rid = range rs {
		if rs[rid].ID == relayid {
			break
		}
	}

	commands := []string{"on", "off", "status", "updateall", "update", "help"}

	cok := contains(commands, command)
	if true != cok {
		resp := "Command " + command + " is not an Relay command. try 'help'"
		return resp
	}

	if command == "help" {
		resp := "Commands: on / off / status / update / updateall / help"
		return resp
	}

	if command == "updateall" {
		go func() { rs.updateStatuses(true) }()
		resp := "All Relays statuses update Scheduled !"
		return resp
	}

	if command == "update" {
		status := rs[rid].Url() + "/" + "status"
		rs[rid].Status = GetUrlQuiet(status)
		resp := "Relay status up to date !"
		return resp
	}

	var getUrl string
	getUrl = rs[rid].Url() + "/" + command
	resp := GetUrlQuiet(getUrl)
	go func() {
		status := rs[rid].Url() + "/" + "status"
		rs[rid].Status = GetUrlQuiet(status)
	}()
	return resp
}

func ApiServer(address string, port string, rs Relays, ss Sensors) {

	// Disable Console Color
	// gin.DisableConsoleColor()
	//⬡‣⬢
	fmt.Println("‣ ⬢ Launching APIServer...")
	fmt.Println("  ⬡ Updating Relays...")
	rs.updateStatuses(true)
	fmt.Println("  ⬡ Updating Sensors...")
	ss.updateStatuses(true)
	fmt.Println("‣ ⬢ Endpoints :")
	fmt.Println("  ⬡ /ping")
	fmt.Println("  ⬢ /api/v1")
	fmt.Println("      ⬢ /sensor")
	fmt.Println("      ⬡ /sensor/updateall")
	fmt.Println("      ⬢ /relay")
	fmt.Println("      ⬡ /relay/:id")
	fmt.Println("      ⬡ /relay/:id/:command")
	fmt.Println("‣ ⬢ Listing Relays :")
	rs.list(false, false)
	fmt.Println("‣ ⬢ Listing Sensors :")
	ss.list(false, false)
	fmt.Printf("\n -- [ LISTENING ON %v:%v ] --\n", address, port)
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v1 := r.Group("/api/v1")
	{

		v1.GET("/sensor", func(c *gin.Context) {
			c.IndentedJSON(200, ss)
		})
		v1.GET("/sensor/updateall", func(c *gin.Context) {
			go func() { ss.updateStatuses(true) }()
			c.IndentedJSON(200, gin.H{"sensors": "all", "command": "updateall", "response": "Scheduled!"})
		})
		v1.GET("/relay", func(c *gin.Context) {
			c.IndentedJSON(200, rs)
		})

		v1.GET("/relay/:id/:command", func(c *gin.Context) {
			id := c.Param("id")
			rid, _ := strconv.Atoi(id)
			rmax := rs.rmax()
			rmin := rs.rmin()
			if rid > rmax {
				c.IndentedJSON(200, gin.H{"err": "No such relay"})
				return
			}
			if rid < rmin {
				c.IndentedJSON(200, gin.H{"err": "No such relay"})
				return
			}
			command := c.Param("command")
			resp := rs.ApiRelayCommand(rid, command)
			c.IndentedJSON(200, gin.H{"relay(id)": rid, "command": command, "response": resp})
		})

		v1.GET("/relay/:id", func(c *gin.Context) {
			id := c.Param("id")
			rid, _ := strconv.Atoi(id)
			rmax := rs.rmax()
			rmin := rs.rmin()
			if rid > rmax || rid < rmin {
				c.IndentedJSON(200, gin.H{"err": "No such relay"})
				return
			}
			relay := rs[rid-1]

			c.IndentedJSON(200, relay)
		})
	}

	routerconfig := address + ":" + port
	//knownRoutes := r.Routes()
	//hiprintApiSpecs(knownRoutes)
	r.Run(routerconfig)
}

/*func printApiSpecs(rti gin.RoutesInfo){
	rti..
	var ri gin.RouteInfo
	for _,ri = range rti {
		ri.nuHandlers = len(handlers)
		ri.handlerName := nameOfFunction(handlers.Last())
		log.Printf("[GIN-debug] "+"%-6s %-25s --> %s (%d handlers)\n", ri.httpMethod, ri.absolutePath, ri.handlerName, ri.nuHandlers)

		}
}*/
