package main
import (
    "github.com/oschwald/geoip2-golang"
    "net"
    "log"
)

func GetConnectionType(ipStr string) string {
	ipObj := net.ParseIP(ipStr)
	//mmdb ConnectionType 
	dbConn, err := geoip2.Open("./assets/GeoIP2-Connection-Type.mmdb")
    if err != nil {
        log.Println(err)
    }
    defer dbConn.Close()
    record, err := dbConn.ConnectionType(ipObj)
    if err != nil {
        log.Println(err)
        log.Println("bad IP")
        return "unknown"
    } else {
    	return record.ConnectionType
    }
}

func GetIsp(ipStr string) (string, string) {
    ipObj := net.ParseIP(ipStr)
	//mmdb ISP 
    dbIsp, err := geoip2.Open("./assets/GeoIP2-ISP.mmdb")
    if err != nil {
        log.Println(err)
    }
    defer dbIsp.Close()
    r2, e2 := dbIsp.ISP(ipObj)
    if e2 != nil {
        log.Println(err)
        log.Println("bad IP")
        return "unknown", "unknown"
    } else {
    	return r2.ISP, r2.Organization
    }
}  

func GetGeoInfo(ipStr string) GeoInfo {
    ipObj := net.ParseIP(ipStr)
	//mmdb city and country 
    dbCity, err := geoip2.Open("./assets/GeoIP2-City.mmdb")
    if err != nil {
        log.Println(err)
    }
    defer dbCity.Close()
    record, err := dbCity.City(ipObj)
    if err != nil {
        log.Println(err)
        log.Println("bad IP")
        return GeoInfo{}
    } else {
    	return GeoInfo {
    		record.Country.IsoCode,
			record.City.Names["en"],
			record.Subdivisions[0].Names["en"],
			record.Postal.Code,
			record.Location.Latitude,
			record.Location.Longitude}
	}
}

type GeoInfo struct {
	Country string
	City string
	State string
	PostalCode string
	Latitude float64
	Longitude float64
}