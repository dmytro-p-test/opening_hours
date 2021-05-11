package parser

import (
	"encoding/json"
	"errors"
	"strings"

	log "github.com/sirupsen/logrus"
)

// example strings
// Mo-Fr 10:00-20:00; Sa,Su 09:00-12:00; PH off
// Mo-Fr 09:00-12:00,13:00-19:00; PH off
// Mo 09:00+; PH off
// Mo-Fr 09:00-12:00

type OpenHours []Ruleset

type Ruleset struct {
	Day  string `json:"day"`
	Time string `json:"time,omitempty"`
}

// Parse OpenHours.Parse converts a string into a slice of OpenHours type
func (oh *OpenHours) Parse(str string) error {
	for _, rs := range strings.Split(str, "; ") { //caution: rulesets separated by semicolon and space!
		ruleset := Ruleset{}

		rulesetParts := strings.Count(rs, " ")
		if rulesetParts != 1 && rulesetParts != 2 {
			dataError := errors.New("invalid ruleset")
			log.Warningf("%v, must contain only one or zero spaces %v",
				dataError, rs)
			return dataError
		}
		ruleset.Day = strings.Split(rs, " ")[0]
		ruleset.Time = strings.Split(rs, " ")[1]

		*oh = append(*oh, ruleset)

	}
	return nil
}

// ToJSON serialize an OpenHours variable into a JSON-byte slice
func (oh *OpenHours) ToJSON() ([]byte, error) {
	data, err := json.MarshalIndent(*oh, "", " ")
	if err != nil {
		log.Warnf("marshaling error %v", err)
		return nil, err

	}
	return data, nil
}

// Contains returns true if a ruleset r exists in the rs slice
// otherwise returns false
func Contains(rs []Ruleset, r Ruleset) bool {
	for _, ruleset := range rs {
		if ruleset.Day == r.Day && ruleset.Time == r.Time {
			return true
		}
	}
	return false
}
