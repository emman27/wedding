package api

import (
	"github.com/emman27/wedding/pkg/rsvp-saver"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// RSVPRequest http structure
type RSVPRequest struct {
	Name    string `json:"name" binding:"required"`
	Count   int    `json:"number_of_attendees" binding:"required"`
	Contact string `json:"contact" binding:"required"`
}

// NewServer to serve HTTP requests
func NewServer(db saver.Database) *gin.Engine {
	r := gin.Default()
	s := saver.NewSaver(db)
	r.POST("/api/rsvp", func(c *gin.Context) {
		r := new(RSVPRequest)
		if err := c.ShouldBind(r); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error(), "request": r})
			return
		}
		if err := s.Save(rsvp{r}); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error(), "request": r})
			return
		}
		c.JSON(http.StatusCreated, map[string]interface{}{"request": r})
	})
	return r
}

// Implements saver.RSVP
type rsvp struct{ *RSVPRequest }

func (r rsvp) Name() string {
	return r.RSVPRequest.Name
}

func (r rsvp) Contact() string {
	return r.RSVPRequest.Contact
}

func (r rsvp) Date() time.Time {
	return time.Now()
}

func (r rsvp) NumberOfPeople() int {
	return r.RSVPRequest.Count
}
