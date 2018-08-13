package hlog

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/google/uuid"
)

/*
LogRequest will log an HTTP Request (r) to logrus and return its UUID
*/
func LogRequest(r *http.Request) string {
	/* Generate a Request ID */
	rid := uuid.New().String()

	/* Log the Request */
	logrus.Printf("%s: %s (%s) '%s %s %s' '%s' '%s'", rid, r.RemoteAddr, r.Header.Get("X-Forwarded-For"), r.Method, r.RequestURI, r.Proto, r.Referer(), r.Host)

	/* Return the Request ID, for further logging */
	return rid
}
