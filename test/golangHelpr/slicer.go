package main

import (
	"fmt"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

type Player struct {
	Name string
}
type Players []*Player

func (p *Player) Delete() {
	for i, _p := range players {
		if _p.Name == p.Name {
			players = append(players[:i], players[i+1:]...)
		}
	}
}

var players Players

func main() {
	for i := 0; i < 1; i++ {
		players = append(players, &Player{Name: strconv.Itoa(i)})
	}
	// fmt.Printf("%v: %T|cap: %d,len: %d\n", players, players, cap(players), len(players))
	// spew.Dump(players[:0])
	fmt.Printf("----------%d-----------\n", len(players))
	// players[0].Delete()
	// players[0].Delete()
	// players[0].Delete()
	for _, v := range players {
		if v.Name == "0" {
			v.Delete()
		}
	}
	spew.Dump(players)
	// fmt.Printf("start delete demo\n")
	// for i, v := range players {
	// 	// if v.Name == "0" {
	// 	v.Delete()
	// 	// }
	// 	spew.Dump(players)
	// 	fmt.Printf("----------%d-----------\n", i)
	// }

	// delete(demo, 4)
	// spew.Dump(players)
	// fmt.Printf("%v: %T|cap: %d,len: %d\n", players, players, cap(players), len(players))
}
