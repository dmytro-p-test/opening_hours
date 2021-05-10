// Defines structure of the open_hours tag
// Parse a string into the structure

package opening_hours

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

type Ruleset struct {
	// common English day abbreviations Mo, Tu, We, Th, Fr, Sa, Su
	// and intervals e.g Mo-Fr or list Mo,Fr
	// Also PH for public holidays
	Days 	string
	// Times are divided by comma without space 08:00-12:00,13:00-16:00
	// Time format is HH24:MM
	// Time is either an interval 08:00-12:00 or open-ended interval 08:00+ or off if facility is closed
	Times   string
}

func (r Ruleset) Parse(str string)(rs Ruleset, err error){
	rs = Ruleset{}
	var splittedString = strings.Split(str, " ")
	if len(splittedString)!=2{
		//todo: return empty ruleset and err here
		log.Fatalf(`input string doesn't have 2 parts "days times"" separated by space: %v`, str)
	}
	//todo: add check for valid days
	rs.Days = splittedString[0]
	rs.Times = splittedString[1]
	return rs, nil
}


func ParseOpenHours(str string){
	rulesets := []Ruleset{}
	for i, rsString := range strings.Split(str, ";"){
		fmt.Printf("ruleset %d: %v\n", i, rsString)
		var rs Ruleset
		rs, err := rs.Parse(rsString)
		if err!= nil{
			log.Fatalf("Error in parsing a ruleset %v", rsString)
		}
		rulesets = append(rulesets, rs)
	}
}

func main() {
	/*
	// Parse open_hours.json into Rulesets type
	file, err := os.Open("open_hours.json")
	if err!=nil{
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	rulesets := Rulesets{}
	err = decoder.Decode(&rulesets)
	if err!= nil{
		log.Fatal(err)
	}

	for i, _ := range rulesets{
		fmt.Println("ruleset " + string(i))
	}
	*/

	var openHoursString = flag.String("open_hours", "", `open_hours string enclosed in ""`)
	flag.Parse()
	fmt.Printf("string: %v\n", *openHoursString)

	// Test input string
	testOpenHoursString := "Mo-Fr 08:00-12:00,13:00-17:30; Sa 08:00-12:00; PH off"

	ParseOpenHours(testOpenHoursString)

}