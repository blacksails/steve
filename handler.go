package steve

import "net/http"

func (s *Server) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("hello world")); err != nil {
			s.log.Error(err, "could not respond")
		}
	}
}
