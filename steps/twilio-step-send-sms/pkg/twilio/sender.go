package twilio

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/relay-integrations/relay-twilio/actions/steps/send-sms/pkg/logs"
)

var logger = logs.WithPackage("twilio")

type Sender struct {
	accountSID string
	authToken  string
	fromNumber string
}

type twilioResponse struct {
	SID string `json:"sid"`
}

type twilioErrorResponse struct {
	Message string `json:"message"`
}

func (s Sender) Send(to, body string) error {
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + s.accountSID + "/Messages.json"

	msgData := url.Values{}
	msgData.Set("To", to)
	msgData.Set("From", s.fromNumber)
	msgData.Set("Body", body)

	client := &http.Client{}
	// Timeout in the HTTP Client includes the entire request liftime, including connection + response body read.
	// If somehow Twilio starts behaving maliciously and returning a slow request, this will handle that case
	// The time is extremely long, but it really just shouldn't be infinite
	client.Timeout = 15 * time.Minute
	req, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(msgData.Encode()))
	req.SetBasicAuth(s.accountSID, s.authToken)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	if err != nil {
		logger.WithError(err).Error("failed to send request to Twilio")
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data twilioResponse
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			logger.WithError(err).Error("failed to parse Twilio response")
		} else {
			logger.Debugf("message `%s` delivered", data.SID)
		}
	} else {
		var errResp twilioErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			logger.WithError(err).Error("failed to parse Twilio error response")
			return err
		}
		logger.Errorf("invalid twilio response code: %s - %s", resp.Status, errResp.Message)
		// TODO: We should actually return an error here...
	}

	return nil
}

func NewSender(accountSID, authToken, fromNumber string) Sender {
	return Sender{
		accountSID,
		authToken,
		fromNumber,
	}
}
