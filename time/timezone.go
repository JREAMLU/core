package timezone

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	googleURI = "https://maps.googleapis.com/maps/api/timezone/json?location=%f,%f&timestamp=%d&sensor=false"
)

type GoogleTimezone struct {
	DstOffset    float64 `bson:"dstOffset"`
	RawOffset    float64 `bson:"rawOffset"`
	Status       string  `bson:"status"`
	TimezoneID   string  `bson:"timeZoneId"`
	TimezoneName string  `bson:"timeZoneName"`
}

func RetrieveGoogleTimezone(latitude float64, longitude float64) (googleTimezone *GoogleTimezone, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	uri := fmt.Sprintf(googleURI, latitude, longitude, time.Now().UTC().Unix())
	fmt.Println(uri)

	resp, err := http.Get(uri)
	if err != nil {
		return googleTimezone, err
	}

	defer resp.Body.Close()

	// Convert the response to a byte array
	rawDocument, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return googleTimezone, err
	}

	// Unmarshal the response to a GoogleTimezone object
	googleTimezone = new(GoogleTimezone)
	if err = json.Unmarshal(rawDocument, &googleTimezone); err != nil {
		return googleTimezone, err
	}

	if googleTimezone.Status != "OK" {
		err = fmt.Errorf("Error : Google Status : %s", googleTimezone.Status)
		return googleTimezone, err
	}

	if len(googleTimezone.TimezoneID) == 0 {
		err = fmt.Errorf("Error : No Timezone Id Provided")
		return googleTimezone, err
	}

	return googleTimezone, err
}

//e.g
/*
func main() {
	// Call to get the timezone for this lat and lng position
	googleTimezone, err := RetrieveGoogleTimezone(38.85682, -92.991714)
	if err != nil {
		fmt.Printf("ERROR : %s", err)
		return
	}

	// Pretend this is the date and time we extracted
	year := 2013
	month := 1
	day := 1
	hour := 2
	minute := 6

	// Capture the location based on the timezone id from Google
	location, err := time.LoadLocation(googleTimezone.TimezoneID)
	if err != nil {
		fmt.Printf("ERROR : %s", err)
		return
	}

	// Capture the local and UTC time based on timezone
	localTime := time.Date(year, time.Month(month), day, hour, minute, 0, 0, location)
	utcTime := localTime.UTC()

	// Display the results
	fmt.Printf("Timezone:\t%s\n", googleTimezone.TimezoneID)
	fmt.Printf("Local Time: %v\n", localTime)
	fmt.Printf("UTC Time: %v\n", utcTime)
}
*/
