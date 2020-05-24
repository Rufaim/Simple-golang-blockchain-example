package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/Rufaim/Simple-golang-blockchain-example/blockchain"
)

type Data struct {
	Data string `json:"data"`
}

func (data *Data) ToString() string {
	return data.Data
}

type Server struct {
	blockchain   *blockchain.Blockchain
	config       Config
	rootFileName string
}

func NewServer(config Config) *Server {
	return &Server{
		blockchain:   blockchain.NewBlockchain(),
		config:       config,
		rootFileName: "index.html",
	}
}

func (s *Server) getRootHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(s.config.StaticPath, s.rootFileName)
		dat, err := ioutil.ReadFile(path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(dat)
	}
}

func (s *Server) getBCHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			writeBlock(w, r, s.blockchain)
		case http.MethodGet:
			returnBlockchain(w, s.blockchain)
		default:
			return
		}
	}
}

func writeBlock(w http.ResponseWriter, r *http.Request, bc *blockchain.Blockchain) {
	blockData := &Data{}
	err := json.NewDecoder(r.Body).Decode(blockData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not write block: " + err.Error()))
		return
	}
	bc.AddBlock(blockData)
	w.WriteHeader(http.StatusOK)
}

func returnBlockchain(w http.ResponseWriter, bc *blockchain.Blockchain) {
	jbytes, err := json.MarshalIndent(bc.GetWritiableBlocks(), "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	// write JSON string
	w.WriteHeader(http.StatusOK)
	w.Write(jbytes)
}
