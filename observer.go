package main

import (
	"fmt"
)

type operation interface {
	register(obs observer)
	deregister(obs observer)
	notifyAll(s *score)
}

type publisher struct {
	observers []observer
}

func (p *publisher) register(obs observer) {
	p.observers = append(p.observers, obs)
}

func (p *publisher) deregister(obs observer) {
	var k int
	for i, v := range p.observers {
		if v == obs {
			k = i
		}
	}

	p.observers = append(p.observers[:k], p.observers[k+1:]...)
}

func (p *publisher) notifyAll(s *score) {
	for _, ob := range p.observers {
		ob.update(s)
	}
}

type radio struct {
}

func (r *radio) update(s *score) {
	fmt.Printf("Radio broadcasting score: runs=%d and wickets=%d\n", s.runs, s.wickets)
}

type tv struct {
}

func (t *tv) update(s *score) {
	fmt.Printf("Tv broadcasting score: runs=%d and wickets=%d\n", s.runs, s.wickets)
}

type observer interface {
	update(s *score)
}

type score struct {
	runs    int
	wickets int
}

func main() {
	publisher := &publisher{}
	tv := &tv{}
	radio := &radio{}
	publisher.register(tv)
	publisher.register(radio)
	//publisher.deregister(radio)
	score := &score{0, 0}
	for i := 0; i < 10; i++ {
		score.runs += 1
		if score.runs%5 == 0 {
			score.wickets += 1
		}
		publisher.notifyAll(score)

	}

}
