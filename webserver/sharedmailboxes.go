package webserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koksmat-com/koksmat/model"
)

func addSharedMailboxesRoutes(rg *gin.RouterGroup) {
	sharedMailboxes := rg.Group("/sharedmailboxes")

	sharedMailboxes.GET("/", func(c *gin.Context) {
		sharedMailboxes, err := model.GetSharedMailboxes()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, sharedMailboxes)

	})
	sharedMailboxes.POST("/", func(c *gin.Context) {

		var json model.NewSharedMailbox
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		sharedMailbox, err := model.CreateSharedMailbox(json.DisplayName, json.Alias, json.Name, json.Members, json.Owners, json.Readers)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, sharedMailbox)

	})

}
