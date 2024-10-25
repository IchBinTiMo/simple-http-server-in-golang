package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func get_ip(r *http.Request) (string, error) {
	ip := r.Header.Get("X-Real-Ip")

	if net.ParseIP(ip) != nil {
		fmt.Print("from X-Real-Ip")
		return ip, nil
	}

	ip = r.Header.Get("X-Forwarded-For")

	for _, v := range strings.Split(ip, ",") {
		if net.ParseIP(v) != nil {
			fmt.Print("from X-Forwarded-For")
			return v, nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	return "", fmt.Errorf("no valid ip found")

}

func main() {
	// srv := api.NewServer()
	// http.ListenAndServe(":8080", srv)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip, err := get_ip(r)

		if err == nil {
			fmt.Println(ip)
		}

		w.Write([]byte("Hello World"))
	})
	http.ListenAndServe(":8080", nil)
}
