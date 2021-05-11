package main

import (
	log "github.com/sirupsen/logrus"
	"opening_hours/cmd"
)

func main() {

	log.Infoln("staring  webservice...")
	cmd.Serve()

	/*
		log.Infoln("this is main!")
		args := os.Args[1:]
		if len(args) != 1 {
			log.Fatalf(`one arguments open hours optionally enclosed in "" is expected, got %v`, args)
		}


		str := os.Args[1]
		log.Infof("running parser for string %v\n", str)

		var openHours parser.OpenHours
		err := (&openHours).Parse(str)
		if err!=nil{
			log.Fatalf("parsing error for string %v: %v", str, err)
		}

		jsonBytes, err := (&openHours).ToJSON()
		if err!=nil{
			log.Fatal(err)
		}
		fmt.Printf("%s\n", jsonBytes)

		log.Infof("open_hours: %v", openHours)*/

}
