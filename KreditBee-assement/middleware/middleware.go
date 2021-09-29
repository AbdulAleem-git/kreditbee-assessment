package middleware

import (
	"KreditBee-assement/data"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq" // postgres golang driver
)
const(
		host = "localhost"
		port = 5432
		user = "postgres"
		password = "postgres"
		dbname = "test"
)
func CreateConnection() *sql.DB {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host,port,user,password,dbname)
	 // Open the connection
	 db, err := sql.Open("postgres", psqlinfo)
 
	 if err != nil {
		 panic(err)
	 }
 
	 // check the connection
	 err = db.Ping()
 
	 if err != nil {
		 panic(err)
	 }
 
	 fmt.Println("Successfully connected!")
	 // return the connection
	 return db
 }

 func InsertAlbum(db *sql.DB, album data.Album) {

       // create the insert sql query
    // returning userid will return the id of the inserted user
    sqlStatement := `INSERT INTO album (id ,userid,title) VALUES ($1, $2, $3) RETURNING id`

    // the inserted id will store in this id
    var id int64

    // execute the sql statement
    // Scan function will save the insert id in the id
    err := db.QueryRow(sqlStatement,album.Id,album.Userid,album.Title).Scan(&id)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    fmt.Printf("%v\n", id)

}
func InsertPhoto(db *sql.DB, photo data.Photo) {

	// create the insert sql query
 // returning userid will return the id of the inserted user
 sqlStatement := `INSERT INTO photo (id ,albumid,photoid,title,url,thumbnailurl) VALUES ($1, $2, $3,$4,$5,$6) RETURNING id`

 // the inserted id will store in this id
 var id int64

 // execute the sql statement
 // Scan function will save the insert id in the id
 err := db.QueryRow(sqlStatement,photo.Id, photo.Albumid,photo.Id,photo.Title,photo.Url,photo.ThumbnailUrl).Scan(&id)

 if err != nil {
	 log.Fatalf("Unable to execute the query. %v", err)
 }

 fmt.Printf("Inserted a single record %v", id)

}

func Search(w http.ResponseWriter, r *http.Request){

	t := r.URL.Query().Get("type")
	id ,err:= strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		    log.Printf("Unable to convert the string into int.  %v", err)
			return
	}
    // convert the id type from string to int
    // typ := params["type"]
	
    // id, err := strconv.Atoi(params["id"])
	// if err != nil {
    //     log.Printf("Unable to convert the string into int.  %v", err)
	// 	return
    // }
    // fmt.Println("id",id)
    // fmt.Println("type",typ)

    // call the getUser function with user id to retrieve a single user
    if t == "album"{
		albums, err := getAllAlbums(id)

    	if err != nil {
        	log.Printf("Unable to get all user. %v", err)
   		 }
    // send all the users as response
    	json.NewEncoder(w).Encode(albums)
	}else{
		photos, err := getAllphotos(id)

    	if err != nil {
        	log.Printf("Unable to get all user. %v", err)
   		 }
    // send all the users as response
    	json.NewEncoder(w).Encode(photos)
	}
}
// get one user from the DB by its userid
func getAllAlbums(id int) ([]data.Album, error) {
    // create the postgres db connection
    db := CreateConnection()

    // close the db connection
    defer db.Close()

    var albums []data.Album

    // create the select sql query
    sqlStatement := `SELECT * FROM users where id=$1`

    // execute the sql statement
    rows, err := db.Query(sqlStatement, id)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    // close the statement
    defer rows.Close()

    // iterate over the rows
    for rows.Next() {
        var album data.Album

        // unmarshal the row object to user
        err = rows.Scan(&album.Id, &album.Userid, &album.Title)

        if err != nil {
            log.Fatalf("Unable to scan the row. %v", err)
        }

        // append the user in the users slice
        albums = append(albums, album)

    }

    // return empty user on error
    return albums, err
}

// get one user from the DB by its userid
func getAllphotos(id int) ([]data.Photo, error) {
    // create the postgres db connection
    db := CreateConnection()

    // close the db connection
    defer db.Close()

    var photos []data.Photo

    // create the select sql query
    sqlStatement := `SELECT * FROM photo where id=$1`

    // execute the sql statement
    rows, err := db.Query(sqlStatement, id)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    // close the statement
    defer rows.Close()

    // iterate over the rows
    for rows.Next() {
        var photo data.Photo

        // unmarshal the row object to user
        err = rows.Scan(&photo.Albumid, &photo.Id, &photo.Title,&photo.Url,&photo.ThumbnailUrl)

        if err != nil {
            log.Fatalf("Unable to scan the row. %v", err)
        }

        // append the user in the users slice
        photos = append(photos, photo)

    }

    // return empty user on error
    return photos, err
}