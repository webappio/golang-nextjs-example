package main

import (
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"log"
	"net/http"
	"time"
)

func main() {
	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var resp []byte
		if req.URL.Path == "/handler-initial-data" {
			resp = []byte(`{"text": "initial"}`)
		} else if req.URL.Path == "/handler" {
			conn, _, _, err := ws.UpgradeHTTP(req, rw)
			if err != nil {
				log.Println("Error with WebSocket: ", err)
				rw.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			go func() {
				defer conn.Close()

				time.Sleep(time.Second) //TODO HACK: sleep a second to check everything is working properly
				err = wsutil.WriteServerMessage(conn, ws.OpText, []byte(`{"text": "from-websocket"}`))
				if err != nil {
					log.Println("Error writing WebSocket data: ", err)
					return
				}
			}()
			return
		} else {
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.Header().Set("Content-Length", fmt.Sprint(len(resp)))
		rw.Write(resp)
	})

	log.Println("Server is available at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
