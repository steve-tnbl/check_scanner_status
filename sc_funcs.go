package main

import "encoding/json"
import "fmt"
import "log"
import "github.com/go-resty/resty/v2"

func GetSCScanners (c *resty.Client) (SCScannerList) {
	var result SCScannerList
	resp, err := c.R().Get(fmt.Sprintf("%s/rest/scanner", c.BaseURL))
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode() != 200 {
		log.Fatal("Invalid access and/or secret key")
	}

	json.Unmarshal(resp.Body(), &result)
	return result
}

func GetSCScanZones (c *resty.Client) (SCScanZones) {
	var result SCScanZones

	resp, err := c.R().Get(fmt.Sprintf("%s/rest/zone", c.BaseURL))
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode() != 200 {
		log.Fatal("Invalid access and/or secret key")
	}

	json.Unmarshal(resp.Body(), &result)
	return result
}

func CheckSCScannerStatus (sl SCScannerList) (int, int) {
	var good, bad int

	for _, scanner := range sl.Response {
		if scanner.Status == "1" || scanner.Status == "1025" {
			good++
		} else {
			bad++
		}
	}

	return good, bad
}

func CheckSCOrphanedZones (sz SCScanZones) (int, []string) {
	var count int
	var names []string

	for _, zone := range sz.Response {
		if zone.TotalScanners > 0 && zone.ActiveScanners == 0 {
			count++
			names = append(names, zone.Name)
		}
	}

	return count, names
}
