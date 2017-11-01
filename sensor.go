package main

import (
	//"fmt"
	//"encoding/json"
	"fmt"
	"github.com/briandowns/spinner"
	//"github.com/gin-gonic/gin"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
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
	return "http://" + s.Masterip + "/" + s.Localname
}

func (ss Sensors) rmax() int {
	return len(ss)
}

func (ss Sensors) rmin() int {
	return ss[0].ID
}

func (ss Sensors) Command(sensorid int, command string) {
	fmt.Println("Command:", command)
	fmt.Println("Sensor:", sensorid)

	//GetRid
	var rid int
	for rid = range ss {
		if ss[rid].ID == sensorid {
			break
		}
	}

	fmt.Println("rid:", rid)
	var getUrl string
	getUrl = ss[rid].Url() + "/" + command
	fmt.Println(getUrl)
	GetUrl(getUrl)
	fmt.Printf("D> Done.\n")
}

func (ss Sensors) list(update bool) {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)

	if update {
		s.Prefix = "  "
		s.Color("red")
		s.Suffix = "  Fetching Statuses..."
		s.Writer = os.Stderr
		s.FinalMSG = "  [Statuses up-to-date] \n\n"
		s.Start()

	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetColWidth(160)
	table.SetColMinWidth(0, 3)
	table.SetColMinWidth(1, 20)
	table.SetColMinWidth(2, 30)
	table.SetColMinWidth(3, 10)
	table.SetColMinWidth(4, 3)

	table.SetHeader([]string{"ID", "NAME", "DESCRIPTION", "MASTERIP", "[~]"})
	table.SetCenterSeparator(" ")
	table.SetColumnSeparator(" ")
	table.SetRowSeparator("-")
	table.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})

	for _, relay := range ss {
		if update {
			status := relay.Url() + "/" + "status"
			table.Append([]string{strconv.Itoa(relay.ID), relay.Name, relay.Description, relay.Masterip, GetUrlQuiet(status)})
		} else {
			table.Append([]string{strconv.Itoa(relay.ID), relay.Name, relay.Description, relay.Masterip, relay.Status})
		}
	}
	//table.SetFooter([]string{"", "TOTAL", strconv.Itoa(ss.rmax())}) // Add Footer
	table.SetColumnAlignment([]int{tablewriter.ALIGN_CENTER})
	table.SetAutoFormatHeaders(false)
	//table.SetFooterAlignment(tablewriter.ALIGN_RIGHT)
	table.SetRowLine(false)
	table.SetHeaderLine(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT) // Set Alignment
	if update {
		s.Stop()
	}
	table.Render() // Send output

}

func (ss Sensors) updateStatuses() {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Prefix = "  "
	s.Color("red")
	s.Suffix = "  Fetching Statuses..."
	s.Writer = os.Stderr
	s.FinalMSG = "  [Statuses up-to-date] \n\n"
	s.Start()

	for i := range ss {
		status := ss[i].Url() + "/" + "status"
		ss[i].Status = GetUrlQuiet(status)
	}
	s.Stop()
}

//////////////////////////////////////////// OLD
/* func GetSensors() []byte {

  relaysBody := []byte(`
   [{
     "id" : 1,
     "name" :"salon.canap.relay1" ,
     "localname" :"relay1" ,
     "place" :"Salon - Canap" ,
     "masterip" :"192.168.1.42" ,
     "description" :"Salon - GAUCHE Canap" ,
     "status" :"OFF"
      },
   {
     "id" : 2,
     "name" :"salon.canap.relay2" ,
     "localname" :"relay2" ,
     "place" :"Salon - Canap" ,
     "masterip" :"192.168.1.42" ,
     "description" :"Salon - DROITE Canap" ,
     "status" :"OFF"
      },
   {
     "id" : 3,
     "name" :"chambre.relay1" ,
     "localname" :"relay1" ,
     "place" :"Chambre - Fond gauche" ,
     "masterip" :"192.168.1.28" ,
     "description" :"Chambre - Grande Lumiere" ,
     "status" :"OFF"
      },
   {
     "id" : 4,
     "name" :"chambre.relay2" ,
     "localname" :"relay2" ,
     "place" :"Chambre - arriere-lit" ,
     "masterip" :"192.168.1.28" ,
     "description" :"Chambre - Veilleuse Anaïs <3" ,
     "status" :"OFF"
      },
   {
     "id" : 5,
     "name" :"salon.meubletv.relay1" ,
     "localname" :"relay1" ,
     "place" :"Salon - Meuble TV" ,
     "masterip" :"192.168.1.6" ,
     "description" :"TV / Wii / FREEBOX (Combo)" ,
     "status" :"ON"
      },
   {
     "id" : 6,
     "name" :"pitiburo.porte.relay1" ,
     "localname" :"relay1" ,
     "place" :"Pitiburo - Cote Porte" ,
     "masterip" :"192.168.1.25" ,
     "description" :"Pitiburo - Cote Porte" ,
     "status" :"OFF"
      },
   {
     "id" : 7,
     "name" :"chambre.relay3" ,
     "localname" :"relay1" ,
     "place" :"Chambre - arriere-lit" ,
     "masterip" :"192.168.1.39" ,
     "description" :"Chambre - Veilleuse Colin <3" ,
     "status" :"OFF"
      }]
      `)

  return relaysBody
}*/
