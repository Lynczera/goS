package handlers

import (
	"errors"
	"fmt"

	"github.com/Lynczera/goS/managers/tvmanager"
	"github.com/gin-gonic/gin"
)

func GetListing(c *gin.Context) {
	tvmanager.GetListings()

	c.JSON(200, gin.H{
		"msg": "test Listing ok",
	})
}

func Testpost(c *gin.Context) {

	message := c.PostForm("message")
	fmt.Printf("message received: %v\n", message)

	c.JSON(200, gin.H{
		"msg": "testpost ok",
	})

}

func getLineup(zip string) (string, error) {

	lineup, err := tvmanager.GetLineup(zip)

	if zip != "" {
		if err != nil {
			return "", errors.New("invalid zip code")
		}

		return lineup, nil
	}

	return "", errors.New("zip code not provided")

}
