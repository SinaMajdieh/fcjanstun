package team

import (
	"fmt"

	"github.com/pattack/rcss"
)

func (j *janstun) See(obj rcss.Object) {
	j.world.ProcessSee(obj)
}
func (w *World) ProcessSee(obj rcss.Object) {
	switch obj.Head() {
	case "l":
		l, _ := obj.(rcss.Line)
		var Type string = ""
		if l.IsLeft() == true {
			Type = "left"
		} else if l.IsRight() == true {
			Type = "right"
		} else if l.IsTop() == true {
			Type = "top"
		} else if l.IsBottom() == true {
			Type = "bottom"
		}
		w.Lines[Type] = l
		//fmt.Println("Line : ", Type, l)
	case "f":
		f, _ := obj.(rcss.Flag)
		w.Flags[f.Flag()] = f
		//fmt.Println(f.Flag(), ":", f.Distance())
		if f.Flag() == "(c)" {
			fmt.Println("Distance : ", f.Distance(), "Time : ", f.DataArriveTime())
		}
	case "g":
		g, _ := obj.(rcss.Goal)
		var Type string = ""
		if g.IsLeft() {
			Type = "left"
		} else if g.IsRight() {
			Type = "right"
		}
		w.Goals[Type] = g
		//fmt.Println("Goal : ", Type, g)
	case "b":
		b, _ := obj.(rcss.Ball)
		w.Ball.Dir = b.Dir
		//fmt.Println("Ball : ", b)
	}
}
