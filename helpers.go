package main
import (
	"os"
	"strings"
	"fmt"
	"net/http"
)
func WriteLineToFile(path string, line string) bool {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
		return false
	}
	defer f.Close()

	if _, err = f.WriteString(line); err != nil {
		panic(err)
		return false
	}
	return true
}
//set the ip if on local machine
//otherwise bypass Heroku proxy to get ip
func FindIp(r *http.Request) string {
	ip := r.RemoteAddr
	if(strings.Contains(ip, "[::1]")) {
		ip = "2603:300a:1b00:7700:f1cf:3f49:5fef:88cf"
	} else {
		xf := r.Header.Get("X-Forwarded-For")
		ips := strings.Split(xf, ", ")
		ip = ips[len(ips)-1]
	}
	fmt.Println(ip)
	return ip    
}

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "8080"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
