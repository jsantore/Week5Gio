package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
	"golang.org/x/image/colornames"
	"log"
)

func eventLoop() {
	mainWindow := app.NewWindow()
	var operationsQ op.Ops
	demoTheme := material.NewTheme(gofont.Collection())
	for {
		event := <-mainWindow.Events() //get the events - read from channel
		switch eventType:= event.(type) {
		case system.DestroyEvent:
			log.Fatal(eventType.Err)
		case system.FrameEvent:
			graphicsContext := layout.NewContext(&operationsQ, eventType)
			drawText("Hello Comp510", graphicsContext, demoTheme)
			eventType.Frame(graphicsContext.Ops)
		}
	}
}

func drawText(textToDisplay string, graphicsContext layout.Context, theme *material.Theme){
	label := material.H1(theme, textToDisplay)
	label.Color = colornames.Blue
	label.Alignment = text.Middle
	label.Layout(graphicsContext)
}

func main(){
	go eventLoop()
	app.Main()
}