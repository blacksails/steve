package steve

import (
	"net/http"
)

func (s *Server) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var wr webhookRequest
		if err := decode(r, &wr); err != nil {
			s.respondErr(w, r, err, http.StatusBadRequest)
			return
		}

		switch wr.Type {
		case webhookRequestTypePing:
			s.handlePing(w, r)
		}
	}
}

func (s *Server) handlePing(w http.ResponseWriter, r *http.Request) {
	s.respond(w, r, webhookResponse{Type: webhookRequestTypePing}, http.StatusOK)
}
