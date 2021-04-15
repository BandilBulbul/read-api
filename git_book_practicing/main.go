package main

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}
type PlayerServer struct {
	store PlayerStore
}
/*func PlayerServer(w http.ResponseWriter,r *http.Request){
	player :=strings.TrimPrefix(r.URL.Path,"/players")
	fmt.Fprint(w,GetPlayerScore(player))
}

//server.go
//server.go
//server.go
//server.go
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		p.processWin(w)
	case http.MethodGet:
		p.showScore(w, r)
	}

}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}


func GetPlayerScore(name string) string{
	if player =="Pepper"{
		return "20"
	}
	if player =="Floyd"{
		return "10"
	}
	return  ""
}
type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) RecordWin(name string) {}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

//server.go
func (p *PlayerServer) processWin(w http.ResponseWriter) {
	p.store.RecordWin("Bob")
	w.WriteHeader(http.StatusAccepted)
}
func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}*/