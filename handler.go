package steve

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"net/http"
)

func (s *Server) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			s.respondErr(w, r, err, http.StatusInternalServerError)
			return
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		if !s.verifySignature(r, body) {
			s.respondErr(w, r, errors.New("unauthorized"), http.StatusUnauthorized)
			return
		}

		var i interaction
		if err := decode(r, &i); err != nil {
			s.respondErr(w, r, err, http.StatusBadRequest)
			return
		}

		switch i.Type {
		case interactionTypePing:
			s.handlePing(w, r)
		default:
			s.log.Info("got command", "cmd", i)
		}
	}
}

func (s *Server) handlePing(w http.ResponseWriter, r *http.Request) {
	s.respond(w, r, interaction{Type: interactionTypePing}, http.StatusOK)
}

func (s *Server) verifySignature(r *http.Request, body []byte) bool {
	sig := r.Header.Get("X-Signature-Ed25519")
	sigDecoded, err := hex.DecodeString(sig)
	if err != nil {
		s.log.Error(err, "could not decode signature")
	}

	pkDecoded, err := hex.DecodeString(s.appPubKey)
	if err != nil {
		s.log.Error(err, "could not decode pubkey")
	}

	t := r.Header.Get("X-Signature-Timestamp")
	var b bytes.Buffer
	if _, err := b.Write([]byte(t)); err != nil {
		s.log.Error(err, "could not write ts")
		return false
	}
	if _, err := b.Write(body); err != nil {
		s.log.Error(err, "could not write body")
		return false
	}

	return ed25519.Verify(pkDecoded, b.Bytes(), sigDecoded)
}
