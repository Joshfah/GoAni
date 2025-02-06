package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/hashicorp/go-getter"
)

func main() {

	re, err := regexp.Compile("'mp4':\\s*'([^']+)'")
	if err != nil {
		log.Fatal("No mp4 Files, for this episode found: ", err)
	}

	req, err := http.NewRequest("GET", "https://maxfinishseveral.com/e/m63gnnnjygor", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:131.0) Gecko/20100101 Firefox/131.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

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
		return
	}

	submatch := re.FindStringSubmatch(string(body))

	decoded, err := base64.StdEncoding.DecodeString(submatch[1])
	if err != nil {
		panic(err)
	}

	err2 := getter.Get("./file.mp4", string(decoded))
	if err2 != nil {
		fmt.Println("Error downloading File: ", err2)
	}

	fmt.Println(string(decoded))
}
