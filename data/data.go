package data

import "time"

//Global Information
type TotalLst struct {
	Confirmed  int64     `json: "confirmed"`
	Recovered  int64     `json: "recovered"`
	Critical   int64     `json: "critical"`
	Deaths     int64     `json: "deaths"`
	LastChange time.Time `json: "lastChange"`
	LastUpdate time.Time `json: "lastUpdate"`
}

//Country Information
type Country struct {
	Data []struct {
		Location  string `json: "location"`
		Confirmed int64  `json: "confirmed"`
		Deaths    int64  `json: "deaths"`
		Recovered int64  `json: "recovered"`
		Active    int64  `json: "active"`
	}
}

//Data Combined
type Data struct {
	T []TotalLst
	C Country
}
