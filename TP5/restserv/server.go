package restserv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest_tp5/comsoc"
	"sync"
	"time"
	"math/rand"
)

type PollingStation struct {
	sync.Mutex // pas bien, il faudrait dans le meilleur des cas crééer un microservice (et donc un channel) ou au pire juste mettre le mutex pour le reqCount
	// la on fait un lock sur l'entiereté de l'application à achque fois qu'on fait une requete/reponse
	id       string
	reqCount comsoc.Count
	addr     string
}

func NewPollingStation(addr string) *PollingStation {
	m := make(comsoc.Count)
	return &PollingStation{id: addr, reqCount: m, addr: addr}
}

func (rsa *PollingStation) checkMethod(method string, w http.ResponseWriter, r *http.Request) bool {
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "method %q not allowed", r.Method)
		return false
	}
	return true
}

func (*PollingStation) decodeRequest(r *http.Request) (req Request, err error) { //transforme la requete à l'origine en bytes en json
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (rsa *PollingStation) doVote(w http.ResponseWriter, r *http.Request) {
	rsa.Lock() // pas vraiment du rest à cause de ça
	defer rsa.Unlock()

	// vérification de la méthode de la requête
	if !rsa.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := rsa.decodeRequest(r)
	if err != nil { 
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error()) //pas bien de propager directement l'erreur, pourrait donner des infos sur notre api
		return
	}

	// traitement de la requête
	var resp Response // on crée un objet réponse
	rsa.reqCount[req.Alt]++
	resp.List = rsa.reqCount
	w.WriteHeader(http.StatusOK)
	serial, _ := json.Marshal(resp)
	w.Write(serial)
}

func (rsa *PollingStation) getWinner(w http.ResponseWriter, r *http.Request) {
	if !rsa.checkMethod("GET", w, r) {
		return
	}

	rsa.Lock()
	defer rsa.Unlock()

	var resp Response // on crée un objet réponse
	winner := comsoc.MaxCount(rsa.reqCount)
	if len(winner) > 1 {
		rnd := rand.Intn(len(winner))
		resp.Winner = winner[rnd]
	} else {
		resp.Winner = winner[0]
	}
	resp.List = rsa.reqCount
	w.WriteHeader(http.StatusOK) //il faut write le header avant le corps (sinon c'est trop tard)
	serial, _ := json.Marshal(resp)
	w.Write(serial)
}

func (rsa *PollingStation) Start() {
	// création du multiplexer
	mux := http.NewServeMux()
	mux.HandleFunc("/vote", rsa.doVote)
	mux.HandleFunc("/winner", rsa.getWinner)

	// création du serveur http
	s := &http.Server{
		Addr:           rsa.addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20}

	// lancement du serveur
	log.Println("Listening on", rsa.addr)
	go log.Fatal(s.ListenAndServe())
}






