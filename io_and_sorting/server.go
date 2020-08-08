package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

// func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	router := http.NewServeMux()
// 	router.Handle("/league", http.HandlerFunc(p.leagueHandler))

// 	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
// 	router.ServeHTTP(w, r)
// }

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
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

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(p.store.GetLeague())
	// w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store

	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(p.leagueHandler))

	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router
	return p
}

func (p *PlayerServer) getLeagueTable() []Player {
	return []Player{
		{"Chirs", 20},
	}
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}
