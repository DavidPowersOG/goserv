package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
)

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func clickpost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    i := new(Impression)
    r.ParseForm()
    i.Source = r.Form["source"][0]

    i.SetTimestamp()
    i.SetUaInfo(r)
    //get geo info
    ipStr := FindIp(r)
    i.Isp, i.IspOrg = GetIsp(ipStr)
    i.ConnectionType = GetConnectionType(ipStr)
    i.SetGeoInfo( GetGeoInfo(ipStr) )

    j := i.ToJson()
    //output to:
    //1 - console, 
	fmt.Println(j)   
    //2 - file, 
	WriteLineToFile("./assets/log.txt",j)	
    //3 - http response
    fmt.Fprintf(w, j) 
}

func main() {
    router := httprouter.New()
	router.GET("/", home)
    router.POST("/clickpost", clickpost)
    log.Fatal(http.ListenAndServe(GetPort(), router))
}


/*
https://pacific-wave-83896.herokuapp.com/
*/