package main

import (
	"flag"

	"github.com/puppetlabs/relay-sdk-go/pkg/taskutil"

	"github.com/relay-integrations/relay-twilio/actions/steps/send-sms/pkg/logs"
	"github.com/relay-integrations/relay-twilio/actions/steps/send-sms/pkg/twilio"
)

var logger = logs.WithPackage("main")

type TwilioConnectionSpec struct {
	AccountSID string `spec:"accountSID"`
	AuthToken  string `spec:"authToken"`
}

type Spec struct {
	// Connection contains the authentication context for Twilio.
	Connection TwilioConnectionSpec `spec:"twilio"`

	// From is the phone number to send the request from.
	From string `spec:"from"`

	// To is the phone number to deliver the relevant message to.
	To string `spec:"to"`

	// This is the actual message that we should send.
	Body string `spec:"body"`
}

func main() {
	defaultMetadataSpecURL, err := taskutil.MetadataSpecURL()

	if err != nil {
		logger.WithError(err).Fatal("failed to get default Metadata Spec URL")
	}

	var (
		specURL = flag.String("spec-url", defaultMetadataSpecURL, "url to fetch the spec from")
	)

	flag.Parse()

	planOpts := taskutil.DefaultPlanOptions{SpecURL: *specURL}

	var spec Spec

	if err := taskutil.PopulateSpecFromDefaultPlan(&spec, planOpts); err != nil {
		logger.WithError(err).Fatal("failed to populate spec from plan")
	}

	// TODO: Validate that all the appropriate parameters are here.

	sender := twilio.NewSender(spec.Connection.AccountSID, spec.Connection.AuthToken, spec.From)

	if err := sender.Send(spec.To, spec.Body); err != nil {
		logger.WithError(err).Fatal("failed to deliver SMS using Twilio.")
	}
}
