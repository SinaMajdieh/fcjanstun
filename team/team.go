package team

import (
	"fmt"

	"github.com/pattack/rcss"
)

type janstun struct {
	name  string
	side  rcss.Side
	match rcss.Match
	world World
}

type World struct {
	// LeftLine struct {
	// 	Dis  float64
	// 	Dir  float64
	// 	Time int
	// }
	// RightLine struct {
	// 	Dis  float64
	// 	Dir  float64
	// 	Time int
	// }
	// TopLine struct {
	// 	Dis  float64
	// 	Dir  float64
	// 	Time int
	// }
	// BottomLine struct {
	// 	Dis  float64
	// 	Dir  float64
	// 	Time int
	// }

	Ball struct {
		Dis     float64
		Dir     float64
		DisChng float64
		DirChng float64
		Time    float64
	}

	Flags map[string]rcss.Flag
	Goals map[string]rcss.Goal
	Lines map[string]rcss.Line
}

func (j janstun) Name() string {
	return j.name
}
func (j *janstun) SetSide(side rcss.Side) {
	// j.side = side
}
func (j *janstun) Invite(m rcss.Match, unum rcss.UniformNumber) {
	j.match = m
	fmt.Println("Joined")
	j.match.Move(0, 0)

}
func (j *janstun) SetPlayMode(mode rcss.PlayMode) {

	// fmt.Printf("Side: %c, Uniform Number: %d, Mode: %s\n", side, unum, mode)
}

func (j *janstun) ServerParam(sp rcss.ServerParameters) {
	//fmt.Printf("Server Parameters: %#v\n", sp)
}

func (j *janstun) PlayerParam(pp rcss.PlayerParameters) {
	//fmt.Printf("Player Parameters: %#v\n", pp)
}

func (j *janstun) PlayerType(pt rcss.PlayerType) {
	//fmt.Printf("Player Type: %#v\n", pt)
}

func (j *janstun) Hear() {

}

func (j *janstun) SenseBody() {

}

func (j *janstun) Score() {

}

func (j *janstun) Kickoff() {
	j.match.Kick(100, 0)
	//j.match.See()
	//fmt.Println(j.world.Flags[Center].Distance())
	// for j.world.Flags[Center].Distance() != 0 {
	// 	//fmt.Println("", j.world.Flags[Center].Distance(), j.world.Flags[Center].DataArriveTime())
	// 	j.match.Dash(100)
	// }

}

func NewTeam(name string) rcss.Team {
	var j janstun
	j.name = name
	j.world.Lines = make(map[string]rcss.Line)
	j.world.Goals = make(map[string]rcss.Goal)
	j.world.Flags = make(map[string]rcss.Flag)
	return &j
}
