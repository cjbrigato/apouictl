package main

import (
	"fmt"
	"github.com/mbndr/figlet4go"
)

func ApouiHeader(s string) {
	ascii := figlet4go.NewAsciiRender()

	// Adding the colors to RenderOptions
	options := figlet4go.NewRenderOptions()
	options.FontName = "larry3d"
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorGreen,
		figlet4go.ColorYellow,
		figlet4go.ColorCyan,
		//figlet4go.NewTrueColorFromHexString("885DBA"),
		//figlet4go.TrueColor{136, 93, 186},
	}

	renderStr, _ := ascii.RenderOpts(s, options)
	fmt.Print(renderStr)
}
