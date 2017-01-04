package main
import (
    "fmt"
   	"encoding/json"
   	"time"
    "xojoc.pw/useragent"
    "net/http"
)

type Impression struct {
	Ip string
	Country string
	Timestamp string
	Source string
	Os string
	Browser string
	Isp string
	IspOrg string
	ConnectionType string
	City string
	State string
	PostalCode string
	Latitude float64
	Longitude float64
}

func (i *Impression) SetGeoInfo(gi GeoInfo) {
    i.Country = gi.Country
    i.City = gi.City
    i.State = gi.State
    i.PostalCode = gi.PostalCode
    i.Latitude = gi.Latitude
    i.Longitude = gi.Longitude
}

func (i *Impression) SetUaInfo(r *http.Request) {
    uaStr := r.UserAgent()
    ua :=  useragent.Parse(uaStr)  
    i.Os = ua.OS
    i.Browser = ua.Name
}
func (i *Impression) SetTimestamp() {
    i.Timestamp = time.Now().Format(time.RFC850)
}

func (i *Impression) ToJson() string {
 	b, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b) + "\n"
}

