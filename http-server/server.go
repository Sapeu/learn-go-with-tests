package main

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

type StubPlayerStore struct {
	scores map[string]int
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

// func GetPlayerScore(name string) int {
// 	if name == "Peper" {
// 		return 20
// 	}

// 	if name == "Floyd" {
// 		return 10
// 	}
// 	return 0
// }

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}
