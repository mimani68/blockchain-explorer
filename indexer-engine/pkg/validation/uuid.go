package validation

import (
	"regexp"

	"app.io/pkg/logHandler"
)

func UidValidation(value string) bool {
	regexUuid, err := regexp.Compile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$")
	if err != nil {
		logHandler.Log(logHandler.ERROR, "Unable to use regex for UUID")
	}
	return regexUuid.MatchString(value)
}
