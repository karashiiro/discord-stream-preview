package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var streamURL = "https://discord.com/api/v8/streams/guild:%s:%s:%s/preview?version=1612055983157"

type streamPreviewRequestParameters struct {
	GuildID string `uri:"guildID" binding:"required"`
	ChannelID string `uri:"channelID" binding:"required"`
	UserID string `uri:"userID" binding:"required"`
}

func main() {
	token := os.Getenv("DISCORD_STREAM_PREVIEW_API_TOKEN")
	port := os.Getenv("DISCORD_STREAM_PREVIEW_API_PORT")

	router := gin.Default()

	router.GET("/streams/:guildID/:channelID/:userID", func(ctx *gin.Context) {
		params := streamPreviewRequestParameters{}

		if err := ctx.ShouldBindUri(&params); err != nil {
			ctx.JSON(400, gin.H{"message": err})
			return
		}

		reqURL := fmt.Sprintf(streamURL, params.GuildID, params.ChannelID, params.UserID)
		req, err := http.NewRequest("GET", reqURL, nil)
		if err != nil {
			ctx.JSON(400, gin.H{"message": err})
			return
		}
		req.Header.Set("authorization", token)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			ctx.JSON(400, gin.H{"message": err})
			return
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			ctx.JSON(500, gin.H{"message": err})
			return
		}

		ctx.Header("content-type", "application/json")
		ctx.String(res.StatusCode, string(body))
	})

	if err := router.Run(":" + port); err != nil {
		log.Println(err)
	}
}
