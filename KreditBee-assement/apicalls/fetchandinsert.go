package apicalls

import (
	"KreditBee-assement/data"
	"KreditBee-assement/http_utils"
	"KreditBee-assement/middleware"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func FetchAlbums() ([]data.Album){
	client := http_utils.NewClient()

	req , err := http.NewRequest("GET","http://jsonplaceholder.typicode.com/albums",nil)
	if err != nil {
		log.Fatal("Error in creating New Request")
	}
	res , err := client.Do(req)
	if err != nil {
		log.Fatal("Error in doing Request",err)
	}
	body  , err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error While reading body",err)
	}
	var d []data.Album
	err = json.Unmarshal(body,&d)
	if err != nil {
		log.Fatal("Unmarshell error",err)
	}
	return d
}

func FetchAlbumPhotos(id int)([]data.Photo){
	client := http_utils.NewClient()

	req , err := http.NewRequest("GET","https://jsonplaceholder.typicode.com/photos",nil)
	if err != nil {
		log.Fatal("Error in creating New Request")
	}
	q := req.URL.Query()
	q.Add("albumId",strconv.Itoa(id))
	req.URL.RawQuery = q.Encode()
	res , err := client.Do(req)
	if err != nil {
		log.Fatal("Error in doing Request",err)
	}
	body  , err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error While reading body",err)
	}
	var d []data.Photo
	err = json.Unmarshal(body,&d)
	if err != nil {
		log.Fatal("Unmarshell error",err)
	}
	return d
	
}

func InserttoDatabase(){
	albums := FetchAlbums()
	
	dbclient := middleware.CreateConnection()
	if dbclient == nil {
		log.Fatal("Not able to create a dbclient")
	}
	defer dbclient.Close()
	for _ , alb := range albums{
		middleware.InsertAlbum(dbclient , alb)
		photos := FetchAlbumPhotos(alb.Id)
		for _,ph := range photos{
			middleware.InsertPhoto(dbclient, ph)
		}
	}
	log.Println("insertion is succesfully done")
}