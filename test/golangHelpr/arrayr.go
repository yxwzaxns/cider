package main

import (
	"github.com/davecgh/go-spew/spew"
)

type P struct {
	Name string
}

func (p *P) Delete() {
	for i, _p := range Ps {
		if _p.Name == p.Name {
			Ps = append(Ps[:i], Ps[i+1:]...)
		}
	}
}

type ps [](*P)

func (p *ps) Get(name string) *P {
	if len(*p) > 0 {
		for _, _p := range *p {
			if _p.Name == name {
				return _p
			}
		}
	}
	return nil
}

func (p *ps) Add(name string) {
	tp := new(P)
	tp.Name = name
	*p = append(*p, tp)
}

var Ps ps

var parr [4]string

func main() {
	parr = [4]string{"aong", "fuck", "tom", "jack"}
	for _, p := range parr {
		Ps.Add(p)
	}
	spew.Dump(Ps)
	Ps.Get("tom").Delete()
	spew.Dump(Ps)
}
