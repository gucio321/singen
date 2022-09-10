package main

import (
	"math"
	"time"

	"github.com/AllenDang/giu"
)

var (
	startx, endx int32 = -10, 10
	xvalues            = []float64{float64(startx)}
)

func loop() {
	sinvalues := []float64{}
	cosvalues := []float64{}
	for _, v := range xvalues {
		sinvalues = append(sinvalues, math.Sin(v))
		cosvalues = append(cosvalues, math.Cos(v))
	}

	giu.SingleWindow().Layout(
		giu.Plot("").AxisLimits(float64(startx), float64(endx), -2, 2, giu.ConditionOnce).Plots(
			giu.PlotLineXY("sin", xvalues, sinvalues),
			giu.PlotLineXY("cos", xvalues, cosvalues),
		),
		giu.Row(
			giu.Label("x min: "),
			giu.InputInt(&startx).OnChange(func() {
				xvalues = []float64{0}
			}),
			giu.Label("x max: "),
			giu.InputInt(&endx).OnChange(func() {
				xvalues = []float64{float64(startx)}
			}),
		),
	)
}

func main() {
	wnd := giu.NewMasterWindow("generator", 640, 480, 0)
	go func() {
		for range time.Tick(100 * time.Millisecond) {
			xvalues = append(xvalues, xvalues[len(xvalues)-1]+0.03)
			giu.Update()
		}
	}()
	wnd.Run(loop)
}
