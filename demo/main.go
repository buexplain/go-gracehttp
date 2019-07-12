package main

import (
	"fmt"
	"github.com/buexplain/go-gracehttp"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/slow", func(writer http.ResponseWriter, request *http.Request) {
		_ = request.ParseForm()
		t, _ := strconv.Atoi(request.Form.Get("ms"))
		if t > 0 {
			<-time.After(time.Duration(t) * time.Millisecond)
		}
		writer.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(writer, "kill -SIGUSR2 %d\n", os.Getpid())
	})

	server := gracehttp.NewServer(":1991", mux, 60*time.Second, 60*time.Second)

	log.Println(server.ListenAndServe())
}
