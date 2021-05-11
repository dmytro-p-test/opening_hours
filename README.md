## Description
opening_hours is a test application for
- structure of opening_hours [OpenStreetMap tag](https://wiki.openstreetmap.org/wiki/Key:opening_hours)
- parser for the opening_hours
- simple REST-API that consumes a string, parses it and returns a JSON formatted representation 

for test purposes only

## Data structure
defined in parser/parser.go

Type _OpenHours_ is an array of _Rulesets_ separated by semicolon and space
_Rulesets_ is a compound type which consists of _Day_ and _Time_ string types
 - OpenHours 
    - Ruleset 1
      - Day
      - Time
   - Ruleset 2
      - Day
      - Time   
    - ...

Text string "Mo-Fr 10:00-20:00; Sa,Su 09:00-12:00; PH off" is represented as:

```
OpenHours{
	Ruleset{Day: "Mo-Fr", Time: "10:00-20:00"},
	Ruleset{Day: "Sa,Su", Time: "09:00-12:00"},
	Ruleset{Day: "PH", Time: "off"},
	}
```

## Usage
Type _OpenHours_ is available in parser/parser.go
The type has two methods
 - _Parse(str string) error_ - to parse a string into internal structure 
 - _ToJSON() ([]byte, error)_ - to marshal a variable into JSON byte slice

```
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
```

## Run
To start a webserver on localhost:8080
```
go run main.go 
```
## Examples
Input string: `Mo-Fr 10:00-20:00; Sa,Su 09:00-12:00; PH off`
```
curl --location --request POST 'localhost:8080' \
--header 'Content-Type: text/plain' \
--data-raw 'Mo-Fr 10:00-20:00; Sa,Su 09:00-12:00; PH off'
```

Console output:
```
[
 {
  "day": "Mo-Fr",
  "time": "10:00-20:00"
 },
 {
  "day": "Sa,Su",
  "time": "09:00-12:00"
 },
 {
  "day": "PH",
  "time": "off"
 }
]
```

## Tests
Parser
```
go test -v parser/parser.go parser/parser_test.go
```
Web Service
```
 go test -v cmd/webservice_test.go cmd/webservice.go
```
## Build
```
 cd opening_hours
 go build
```
Executable starts the web service

Dockerfile is also available you want to run web service in a container.