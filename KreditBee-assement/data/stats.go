package data

type Album struct {
	Userid int    `json:"userid"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
}
type Photo struct {
	Albumid      int    `json:"albumid"`
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	ThumbnailUrl string `json:"thumbnailUrl"`
}
