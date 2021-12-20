package calling

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type config struct {
	accountSid   string
	authToken    string
	twilioNumber string
	mobileNumber string
	callsUrl     string
}

func newConfig() config {
	err := godotenv.Overload()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Let's set some initial default variables
	c := config{
		accountSid:   os.Getenv("TWILIO_ACCOUNT_SID"),
		authToken:    os.Getenv("TWILIO_AUTH_TOKEN"),
		twilioNumber: os.Getenv("TWILIO_NUMBER"),
		mobileNumber: os.Getenv("MOBILE_NUMBER"),
	}
	c.callsUrl = fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Calls.json", c.accountSid)
	return c
}

func Call() {
	cfg := newConfig()

	// fmt.Printf("accountSid: %s\n", cfg.accountSid)
	// fmt.Printf("authToken: %s\n", cfg.authToken)
	// fmt.Printf("twilioNumber: %s\n", cfg.twilioNumber)
	// fmt.Printf("mobileNumber: %s\n", cfg.mobileNumber)
	// fmt.Printf("callsUrl: %s\n", cfg.callsUrl)

	// Build out the data for our message
	v := url.Values{}
	v.Set("To", cfg.mobileNumber)
	v.Set("From", cfg.twilioNumber)
	v.Set("Url", "http://demo.twilio.com/docs/voice.xml")
	//v.Set("Url", "http://twimlets.com/holdmusic?Bucket=com.twilio.music.ambient")
	rb := *strings.NewReader(v.Encode())

	// Create Client
	client := &http.Client{}
	req, _ := http.NewRequest("POST", cfg.callsUrl, &rb)
	req.SetBasicAuth(cfg.accountSid, cfg.authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// make request
	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &data)
		if err != nil {
			fmt.Println("Error unmarshalling Twilio response.")
		}
		fmt.Println(data["sid"])
	} else {
		fmt.Println(resp.Status)
		fmt.Println(resp.Request.GetBody())

		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))
	}
}
