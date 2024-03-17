package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Jack Sparrow in Love", Artist: "Luis Smith", Price: 49.99},
	{ID: "2", Title: "Mia Kalifa New Story", Artist: "Marck Lust", Price: 29.99},
	{ID: "3", Title: "Jenefer Anistar", Artist: "Roger Jason", Price: 79.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	r := gin.Default()
	r.GET("/albums", getAlbums)
	r.GET("/albums/:id", getAlbumsByID)
	r.POST("/albums", postAlbums)

	r.Run("localhost:8080")
}
