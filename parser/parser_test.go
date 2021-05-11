package parser

import (
	"testing"
)

func TestOpenHours_Parse(t *testing.T) {
	tests := []struct {
		str           string
		wantOpenHours OpenHours
	}{
		{"Mo-Fr 10:00-20:00; Sa,Su 09:00-12:00; PH off",
			OpenHours{
				Ruleset{Day: "Mo-Fr", Time: "10:00-20:00"},
				Ruleset{Day: "Sa,Su", Time: "09:00-12:00"},
				Ruleset{Day: "PH", Time: "off"},
			},
		},
		{"Mo-Fr 09:00-12:00,13:00-19:00; PH off",
			OpenHours{
				Ruleset{Day: "Mo-Fr", Time: "09:00-12:00,13:00-19:00"},
				Ruleset{Day: "PH", Time: "off"},
			},
		},
		{"Mo 09:00+; PH off",
			OpenHours{
				Ruleset{Day: "Mo", Time: "09:00+"},
				Ruleset{Day: "PH", Time: "off"},
			},
		},
		{"Mo-Fr 09:00-12:00",
			OpenHours{
				Ruleset{Day: "Mo-Fr", Time: "09:00-12:00"},
			},
		},
	}

	for _, test := range tests {
		var openHour OpenHours
		err := (&openHour).Parse(test.str)
		if err != nil {
			t.Errorf("Cannot parse test string %v: %v", test.str, err)
		}
		if len(openHour) != len(test.wantOpenHours) {
			t.Errorf("%v: want %v rulesets, got %v", test.str, len(openHour), len(test.wantOpenHours))
		}

		for _, ruleset := range test.wantOpenHours {
			if !Contains(openHour, ruleset) {
				t.Errorf("cannot find %v in %v", ruleset, openHour)
			}
		}
	}
}
