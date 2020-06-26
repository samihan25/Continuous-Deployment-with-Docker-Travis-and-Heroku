package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"runtime"
	"time"
)

//Prints Hello World!!
func hello() string {
	return "Hello World!!"
}

//Print Indian Standard Time (IST) which equals to UTC + 5:30 hr
func ist() (string, string) {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	nowTime := time.Now().In(loc).Format("3:04 PM")
	nowDate := time.Now().In(loc).Format("02-01-2006 Monday")
	return nowTime, nowDate
}

func helloWorld(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "\n\n")
	fmt.Fprintf(w, hello())
	fmt.Fprintf(w, "\n\n")
	//fmt.Fprintf(w, "And I am ..... Ironman\n")

	//Print IP address
	ipAddress, _, _ := net.SplitHostPort(r.RemoteAddr)
	fmt.Fprintf(w, "Your IP address : %s\n\n", ipAddress)

	//Print Indian standard time
	nowTime, nowDate := ist()
	fmt.Fprintf(w, "Indian Standard Time :-\n")
	fmt.Fprintf(w, "Time : %s\n", nowTime)
	fmt.Fprintf(w, "Date : %s\n\n", nowDate)

	//Print the Operating System
	fmt.Fprintf(w, "Operating System of server is : %s\n\n", runtime.GOOS)
}

func main() {
	fmt.Println("Server is running ....")
	fmt.Println("Press Ctrl+C to close the server.")
	http.HandleFunc("/", helloWorld)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
	//srv := &http.Server{Addr: ":8080"}
	//go func() {
	//	srv.ListenAndServe()
	//}()
	//time.Sleep(10 * time.Second)
	//srv.Shutdown(context.TODO())
	//s.Shutdown(context.Background())
}
