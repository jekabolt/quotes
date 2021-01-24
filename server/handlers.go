package server

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func (s *Server) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func setCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
}

func handleOptions(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
}

func (s *Server) getQuotes(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) getCategories(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Categories{
		List: []string{
			"фильмы",
			"личности",
			"персонажи",
			"бизнес",
			"отношения",
			"любовь",
		},
	})
}

func (s *Server) postCategories(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func (s *Server) auth(w http.ResponseWriter, r *http.Request) {
	var c Claims
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		log.Error().Err(err).Msg("auth:json.NewDecoder")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	jwt, err := s.jwtWithExpTime(c.UserId, day)
	if err != nil {
		log.Error().Err(err).Msg("auth:json.jwtWithExpTime")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(jwt)
}
