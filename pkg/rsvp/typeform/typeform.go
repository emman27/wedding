package typeform

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/emman27/wedding/pkg/rsvp"
	"github.com/gojektech/heimdall/httpclient"
	"github.com/sirupsen/logrus"
)

// Some configuration for the default form that this was built for
const (
	DefaultFormID                   = "V4SNj7"
	NumberOfQuestionsInForm         = 3
	NumberOfAttendeesQuestionNumber = 2
)

var choiceOptions = map[choiceOption]int{
	"0. I can't come :(": 0,
	"1":                  1,
	"2":                  2,
	"3":                  3,
	"4":                  4,
	"5":                  5,
	"6":                  6,
}

var formID = flag.String("typeform-form-id", DefaultFormID, "Form ID to check. Only applicable if using typeform")

// Timeout for the HTTP API calls to typeform
const Timeout = 2 * time.Second

// BaseURL to typeform API
const BaseURL = "https://api.typeform.com"

// Endpoints used by the counter
const (
	ResponsesAPI = "/forms/%s/responses"
)

// NewCounter initializes a new typeform counter
func NewCounter(apiKey string, logger logrus.StdLogger) *Counter {
	logger.Printf("Initializing new Typeform Counter with API Key %s\n", strings.Repeat("*", len(apiKey)))
	return &Counter{
		apiKey: apiKey,
		logger: logger,
	}
}

// Counter connects to the typeform API to check the number of responses
type Counter struct {
	apiKey string
	logger logrus.StdLogger
}

// GetNumberOfResponses counts the number of responses to the typeform form
func (c *Counter) GetNumberOfResponses() (int, error) {
	resp, err := c.getResponses()
	if err != nil {
		return 0, err
	}
	return resp.TotalItems, nil
}

// GetNumberOfAttendees based on form responses
func (c *Counter) GetNumberOfAttendees() (int, error) {
	resp, err := c.getResponses()
	if err != nil {
		return 0, err
	}
	total := 0
	for _, response := range resp.Items {
		total += response.GetNumberOfAttendees()
	}
	return total, nil
}

var _ rsvp.Counter = (*Counter)(nil)

func (c *Counter) authHeader() *http.Header {
	return &http.Header{"Authorization": []string{fmt.Sprintf("Bearer %s", c.apiKey)}}
}

func (c *Counter) getResponses() (*responsesSchema, error) {
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(Timeout))
	endpoint := fmt.Sprintf("%s%s", BaseURL, fmt.Sprintf(ResponsesAPI, *formID))
	c.logger.Printf("Making API call to %s\n", endpoint)
	res, err := client.Get(endpoint, *c.authHeader())
	if err != nil {
		return nil, err
	}
	c.logger.Printf("API call to %s, response: %v", endpoint, res)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var apiResponse = new(responsesSchema)
	if err = json.Unmarshal(body, &apiResponse); err != nil {
		return nil, err
	}
	return apiResponse, nil
}
