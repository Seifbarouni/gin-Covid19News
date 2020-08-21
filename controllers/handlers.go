package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"projects/GinFramework/gin-Covid/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

var t []models.TotalLst
var c models.Country
var resData models.Data
 
func handleErrPanic(err error) {
	if err != nil {
		panic(err)
	}
}
func handleErrLog(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func getGlobInfo() {
		url := "https://covid-19-data.p.rapidapi.com/totals?format=json"

		req, err := http.NewRequest("GET", url, nil)
		handleErrPanic(err)

		req.Header.Add("x-rapidapi-host", "covid-19-data.p.rapidapi.com")
		req.Header.Add("x-rapidapi-key", "23ca074be8msh1d676ab48fd8954p1ad701jsn4cc1f45d865d")

		res, err := http.DefaultClient.Do(req)
		handleErrPanic(err)

		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		handleErrLog(err)
		json.Unmarshal(body, &t)

}
func getCountriesInfo() {
		url := "https://covid2019-api.herokuapp.com/v2/current"
		resp, err := http.Get(url)
		handleErrLog(err)

		body, err := ioutil.ReadAll(resp.Body)
		handleErrLog(err)

		json.Unmarshal(body, &c)

}

var cacheVar = cache.New(5*time.Minute,10*time.Minute)

func getAllData() {
	getGlobInfo()
	getCountriesInfo()
	resData.T = t
	resData.C = c

	cacheVar.Set("data",resData,cache.DefaultExpiration)
}

func Redirect(c *gin.Context) {
	c.Redirect(http.StatusFound,"/corona")
}

func RenderHTML(c *gin.Context){
	data,found:=cacheVar.Get("data")
	if found{
		c.HTML(200,"glob.html",data)
	}else{
		getAllData()
		c.HTML(200,"glob.html",resData)
	}
}