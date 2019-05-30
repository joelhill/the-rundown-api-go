package therundown

import (
	"context"
	"errors"
	"fmt"
	"time"

	blaster "github.com/joelhill/go-rest-http-blaster"
	logrus "github.com/sirupsen/logrus"
)

// EventBySportByDateOptions - Are the options to hit the event by sport by date endpoint
type EventBySportByDateOptions struct {
	// URL Parts
	URL   string
	Sport string
	Date  string

	// Optional URL Params
	AllPeriods bool
	Scores     bool
	Teams      bool
	Offset     string
}

// NewEventBySportByDateOptions - Returns the options with most url parts already set to hit the event by sport endpoint
func (s *Service) NewEventBySportByDateOptions() *EventBySportByDateOptions {
	return &EventBySportByDateOptions{
		URL:   s.Config.BaseURL,
		Sport: s.Config.Sport,
	}
}

// EventBySportByDate - hits the https://therundown-therundown-v1.p.rapidapi.com/sports/3/events endpoint
func (s *Service) EventBySportByDate(options *EventBySportByDateOptions) (TheRunDownAPILines, int, error) {
	errorPayload := make(map[string]interface{})
	mapping := TheRunDownAPILines{}

	// make sure we have all the required elements to build the full required url string.
	err := validateEventBySportByDateURI(options)
	if err != nil {
		return mapping, 0, err
	}

	t := time.Now()
	cacheBuster := t.Format("20060102150405")

	uri := fmt.Sprintf("%s/sports/%s/events/%s?cachebuster=%s", options.URL, options.Sport, options.Date, cacheBuster)

	if options.AllPeriods {
		uri = fmt.Sprintf("%s&include=%s", uri, "all_periods")
	}

	if options.Scores {
		uri = fmt.Sprintf("%s&include=%s", uri, "scores")
	}

	if options.Teams {
		uri = fmt.Sprintf("%s&include=%s", uri, "teams")
	}

	if len(options.Offset) > 0 {
		uri = fmt.Sprintf("%s&offset=%s", uri, options.Offset)
	}

	s.Logger = s.Logger.WithFields(logrus.Fields{
		"URI": uri,
	})
	s.Logger.Debug("EventBySportByDate API Call")

	// make you a client
	client, err := blaster.NewClient(uri)
	if err != nil {
		s.Logger.Errorf("failed to create a http client: %s", err.Error())
		return mapping, 0, err
	}

	client.SetHeader("Accept", "application/json")
	client.SetHeader("X-RapidAPI-Key", s.Config.XRapidAPIKey)
	client.SetHeader("X-RapidAPI-Host", s.Config.XRapidAPIHost)
	client.WillSaturateOnError(&errorPayload)
	client.WillSaturate(&mapping)

	statusCode, err := client.Get(context.Background())
	if err != nil {
		s.Logger.Errorf("something went wrong making the get request for EventBySportByDate: %s", err.Error())
		return mapping, statusCode, err
	}

	s.Logger.Infof("EventBySportByDate Status Code: %d", statusCode)

	if client.StatusCodeIsError() {
		s.Logger.Errorf("EventBySportByDate retuned an unsuccessful status code. Error: %+v", errorPayload)
	}

	return mapping, statusCode, nil
}

func validateEventBySportByDateURI(options *EventBySportByDateOptions) error {
	if len(options.URL) == 0 {
		return errors.New("missing required option to build the url: URL")
	}
	if len(options.Sport) == 0 {
		return errors.New("missing required option to build the url: Sport")
	}
	if len(options.Date) == 0 {
		return errors.New("missing required option to build the url: Date")
	}

	return nil
}
