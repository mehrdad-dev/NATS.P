package main

import (
	"github.com/mbndr/figlet4go"
	"log"
	"fmt"
)


func ShowFiglet(text string) {
	ascii := figlet4go.NewAsciiRender()

	// Adding the colors to RenderOptions
	options := figlet4go.NewRenderOptions()
	blue,_ :=  figlet4go.NewTrueColorFromHexString("1089ff")
	green,_ :=  figlet4go.NewTrueColorFromHexString("009975")
	blue2,_ :=  figlet4go.NewTrueColorFromHexString("3e64ff") 
	green2,_ :=  figlet4go.NewTrueColorFromHexString("85ef47")

	color :=  []figlet4go.Color{blue,green,blue2,green2}
	options.FontColor = color
	options.FontName = "standard"

	renderStr, err := ascii.RenderOpts(text, options)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(renderStr)
}