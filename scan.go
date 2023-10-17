package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func monitorPort(host string, port int) (int, string) {
	target := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", target)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	conn.Close()
	return 200, fmt.Sprintf("200 Everything is OK")
}

func portHandler(w http.ResponseWriter, r *http.Request) {
	host := "HOST"
	port := 1433

	_, statusMessage := monitorPort(host, port)
	fmt.Fprintf(w, "%s\n", statusMessage)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Read the HTML content from the file
	html, err := os.ReadFile("./html/index.html")
	if err != nil {
		fmt.Println("Error reading HTML file:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(html)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/port", portHandler)
	fmt.Println("Port Monitor server is online at port 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
