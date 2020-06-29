package main

import (
	"image/color"
	"log"
	"os"
	"strconv"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/dolanor/katas/roll"
)

type d20App struct {
	lineEditor *widget.Editor
	rollClic   *widget.Clickable
	errLabel   *widget.Label
}

func main() {
	go func() {
		defer os.Exit(0)
		w := app.NewWindow(app.Size(unit.Dp(800), unit.Dp(600)))
		da := d20App{
			lineEditor: &widget.Editor{
				SingleLine: true,
			},
			rollClic: &widget.Clickable{},
			errLabel: &widget.Label{},
		}

		err := da.loop(w)
		if err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func (da *d20App) loop(w *app.Window) error {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops
	var err error
	var result int

	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)

				for da.rollClic.Clicked() {
					var n int
					n, err = strconv.Atoi(da.lineEditor.Text())
					if err == nil {
						result = roll.D20(n)
					} else {
						result = 0
					}
				}

				da.layoutWidgets(gtx, th, result, err)

				e.Frame(gtx.Ops)
			}
		}
	}
	return nil
}
func (da *d20App) layoutWidgets(gtx layout.Context, th *material.Theme, result int, err error) layout.Dimensions {
	fl := layout.Flex{Axis: layout.Vertical}
	//topOSInset := layout.Inset{}
	label := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		body := material.Body1(th, "Number of dice to roll")
		return body.Layout(gtx)
	})
	input := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		e := material.Editor(th, da.lineEditor, "number of dice")
		return e.Layout(gtx)
	})
	btn := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		e := material.Button(th, da.rollClic, "ROLL!")
		return e.Layout(gtx)
	})
	res := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		body := material.Body1(th, strconv.Itoa(result))
		body.Color = color.RGBA{R: 255, B: 255, A: 255}
		return body.Layout(gtx)
	})
	widgets := []layout.FlexChild{label, input, btn, res}

	if err != nil {
		errLbl := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			l := material.Body1(th, err.Error())
			l.Color = color.RGBA{R: 255, A: 255}
			return l.Layout(gtx)
		})
		widgets = append(widgets, errLbl)
	}

	return fl.Layout(gtx, widgets...)
}
