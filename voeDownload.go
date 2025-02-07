package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

func voeDownload() {

	re, err := regexp.Compile("'mp4':\\s*'([^']+)'")
	if err != nil {
		log.Fatal("No mp4 Files, for this episode found: ", err)
	}

	req, err := http.NewRequest("GET", "https://maxfinishseveral.com/e/kotd0uw6kf4l", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:135.0) Gecko/20100101 Firefox/135.0")

	client := &http.Client{}

	resp, err := client.Do(req)
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

	out, err := os.Create("./file.mp4")
	if err != nil {
		log.Fatal(err)
	}

	downResp, err := http.Get(string(decoded))
	if err != nil {
		log.Fatal(err)
	}

	_, err3 := io.Copy(out, downResp.Body)
	if err3 != nil {
		log.Fatal(err3)
	}

}
