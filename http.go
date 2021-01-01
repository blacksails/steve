package steve

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// decode decodes the request body into v and calls
// its OK() function to validate the object.
func decode(r *http.Request, v interface{}) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return err
	}
	if valid, ok := v.(interface {
		OK() error
	}); ok {
		err = valid.OK()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, v interface{}, code int) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(v)
	if err != nil {
		s.respondErr(w, r, err, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if _, err := buf.WriteTo(w); err != nil {
		s.log.WithValues(
			"ctx", r.Context(),
		).Error(err, "respond")
	}
}

func (s *Server) respondErr(w http.ResponseWriter, r *http.Request, err error, code int) {
	errObj := struct {
		Error string `json:"error"`
	}{Error: err.Error()}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(errObj); err != nil {
		s.log.WithValues(
			"ctx", r.Context(),
		).Error(err, "respondErr")
	}
}
