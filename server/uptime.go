package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type uptimeHandler struct {
	StartTime time.Time
}

func newUptimeHandler() *uptimeHandler {
	now := time.Now()
	log.Print("Creating uptime handler; unix now: ", now.Unix())

	return &uptimeHandler{StartTime: time.Now()}
}

func (h uptimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dt := time.Now().Sub(h.StartTime)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%.4f", dt.Seconds())))
}
