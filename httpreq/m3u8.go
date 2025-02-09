package httpreq

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func Getm3u8(URL string) {
	re, err := regexp.Compile("'hls':\\s*'([^']+)'")
	if err != nil {
		log.Fatal("No mp4 Files, for this episode found: ", err)
	}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:135.0) Gecko/20100101 Firefox/135.0")

	client := &http.Client{}

	resp, err := client.Get(URL)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
	}

	submatch := re.FindStringSubmatch(string(body))

	decoded, err := base64.StdEncoding.DecodeString(submatch[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(decoded))

}
