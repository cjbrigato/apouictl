package main

import (
	//"fmt"
	//"encoding/json"
	//"fmt"
	"github.com/briandowns/spinner"
	//"github.com/gin-gonic/gin"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"strings"
	"time"
)

type Sensors []Sensor

type Sensor struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Localname   string `json:"localname"`
	Place       string `json:"place"`
	Type        string `json:"type"`
	Masterip    string `json:"masterip"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (s Sensor) Url() string {
	if s.Type == "doorlock" {
		return "http://" + s.Masterip + "/status"
	}
	return "http://" + s.Masterip + "/" + s.Type
}

func (ss Sensors) rmax() int {
	return len(ss)
}

func (ss Sensors) rmin() int {
	return ss[0].ID
}

// Command for Sensor is nonsense
/*
func (ss Sensors) Command(sensorid int, command string) {
	fmt.Println("Command:", command)
	fmt.Println("Sensor:", sensorid)

	//GetRid
	var sid int
	for sid = range ss {
		if ss[sid].ID == sensorid {
			break
		}
	}

	fmt.Println("sid:", sid)
	var getUrl string
	getUrl = ss[sid].Url() + "/" + command
	fmt.Println(getUrl)
	GetUrl(getUrl)
	fmt.Printf("D> Done.\n")
}
*/

func (ss Sensors) list(update bool, quiet bool) {
	if update {
		ss.updateStatuses(quiet)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetColWidth(160)
	table.SetColMinWidth(0, 3)
	table.SetColMinWidth(1, 20)
	table.SetColMinWidth(2, 30)
	table.SetColMinWidth(3, 10)
	table.SetColMinWidth(4, 10)
	table.SetColMinWidth(5, 3)
	table.SetHeader([]string{"ID", "NAME", "DESCRIPTION", "TYPE", "MASTERIP", "[~]"})
	table.SetCenterSeparator(" ")
	table.SetColumnSeparator(" ")
	table.SetRowSeparator("-")
	table.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
	for _, sensor := range ss {
		table.Append([]string{strconv.Itoa(sensor.ID), sensor.Name, sensor.Description, sensor.Type, sensor.Masterip, sensor.Status})
	}
	//table.SetFooter([]string{"", "TOTAL", strconv.Itoa(ss.rmax())}) // Add Footer
	table.SetColumnAlignment([]int{tablewriter.ALIGN_CENTER})
	table.SetAutoFormatHeaders(false)
	//table.SetFooterAlignment(tablewriter.ALIGN_RIGHT)
	table.SetRowLine(false)
	table.SetHeaderLine(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT) // Set Alignment
	table.Render()                             // Send output
}

func (ss Sensors) updateStatuses(quiet bool) {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	if !quiet {
		s.Prefix = "  "
		s.Color("red")
		s.Suffix = "  Fetching Statuses..."
		s.Writer = os.Stderr
		s.FinalMSG = "  [Statuses up-to-date] \n\n"
		s.Start()
	}
	for i := range ss {
		status := ss[i].Url() // + "/" + "status"
		prestatus := GetUrlQuiet(status)
		if prestatus == "ERR" {
			//ss[i].Status = prestatus
			ss[i].Status = "OFFLINE"
			continue
		}
		if ss[i].Type == "doorlock" {
			ss[i].Status = strings.Split(prestatus, ";")[1]
			continue
		}
		ss[i].Status = strings.Split(prestatus, ";")[1] + " (" + strings.Split(prestatus, ";")[2] + ")"

	}
	if !quiet {
		s.Stop()
	}
}

//////////////////////////////////////////// OLD
/* func GetSensors() []byte {

  sensorsBody := []byte(`
      [
    {
      "id" : 1,
      "name" : "salon.temp.sensor1",
      "localname" : "temp.sensor1",
      "place" : "Salon - Meuble Tv",
      "type"  : "temperature",
      "masterip" : "192.168.1.18",
      "description" : "Temperature Salon",
      "status" : "27.50"
      },
    {
      "id" : 2,
      "name" : "chambre.temp.sensor1",
      "localname" : "sensor1",
      "place" : "Chambre - Table de nuit Colin",
      "type"  : "temperature",
      "masterip" : "192.168.1.16",
      "description" : "Temperature Chambre",
      "status" : ""
      },
    {
      "id" : 3,
      "name" : "salon.humid.sensor1",
      "localname" : "humid.sensor1",
      "place" : "Salon - Meuble Tv",
      "type"  : "humidity",
      "masterip" : "192.168.1.18",
      "description" : "Humidite Salon",
      "status" : "16.90"
      },
    {
      "id" : 4,
      "name" : "chambre.humid.sensor1",
      "localname" : "sensor1",
      "place" : "Chambre - Table de nuit Colin",
      "type"  : "humidity",
      "masterip" : "192.168.1.16",
      "description" : "Humidite Chambre",
      "status" : ""
      },
    {
      "id" : 5,
      "name" : "door.lock.magnet.sensor",
      "localname" : "sensor",
      "place" : "Porte d'entree",
      "type"  : "doorlock",
      "masterip" : "192.168.1.24",
      "description" : "Door Lock Status",
      "status" : "CLOSE"
      }
      ]
      `)

  return sensorsBody
}*/
