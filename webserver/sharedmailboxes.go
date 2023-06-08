package webserver

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/koksmat-com/koksmat/model"
)

func addSharedMailboxesRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/sharedmailboxes")

	users.GET("/", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var sharedMailboxes []model.SharedMailbox
		defer cancel()

		results, err := model.GetSharedMailboxes()
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Error")
			return
		}
		defer results.Close(ctx)
		for results.Next(ctx) {
			var sharedMailbox model.SharedMailbox
			if err = results.Decode(&sharedMailbox); err != nil {
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}

			sharedMailboxes = append(sharedMailboxes, sharedMailbox)
		}
		c.JSON(http.StatusOK, sharedMailboxes)

	})
	users.GET("/comments", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users comments")
	})
	users.GET("/pictures", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users pictures")
	})
}
