package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "\n\nHello World\n\n")
	//fmt.Fprintf(w, "And I am ..... Ironman\n")
	ipAddress, _, _ := net.SplitHostPort(r.RemoteAddr)
	dt := time.Now()
	fmt.Fprintf(w, "Your IP address : %s\n\n", ipAddress)
	fmt.Fprintf(w, "Date and Time : %s\n\n", dt)
}

func main() {
	fmt.Println("Server is running ....")
	fmt.Println("Press Ctrl+C to close the server. ")
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
	//srv := &http.Server{Addr: ":8080"}
	//go func() {
	//	srv.ListenAndServe()
	//}()
	//time.Sleep(10 * time.Second)
	//srv.Shutdown(context.TODO())
	//s.Shutdown(context.Background())
}
