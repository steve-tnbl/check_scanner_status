package main

import "encoding/json"
import "fmt"
import "log"
import "github.com/go-resty/resty/v2"

func GetIOScanners (c *resty.Client) (IOScannerList) {
        var result IOScannerList
        resp, err := c.R().Get(fmt.Sprintf("%s/scanners", c.BaseURL))
        if err != nil {
                log.Fatal(err)
        }
        if resp.StatusCode() != 200 {
                log.Fatal("Invalid access and/or secret key")
        }

        json.Unmarshal(resp.Body(), &result)
        return result
}

func GetIOScannerGroups (c *resty.Client) (IOScannerGroups) {
        var result IOScannerGroups
        resp, err := c.R().Get(fmt.Sprintf("%s/scanner-groups", c.BaseURL))
        if err != nil {
                log.Fatal(err)
        }
        if resp.StatusCode() != 200 {
                log.Fatal("Invalid access and/or secret key")
        }

        json.Unmarshal(resp.Body(), &result)
        return result
}

func CheckIOScannerStatus (sl IOScannerList) (int, int) {
        var good, bad int

        for _, scanner := range sl.Scanners {
		if scanner.Type != "managed" { continue }
                if scanner.Status == "on" {
                        good++
                } else {
                        bad++
                }
        }

        return good, bad
}

func CheckIOOrphanedGroups (c *resty.Client, sg IOScannerGroups) (int, []string) {
	var ogcount int
	var ognames []string
	groupid2name := make(map[int]string)

	// First we need the Group IDs
	var groups []int
	for _, group := range sg.ScannerPools {
		groupid2name[group.ID] = group.Name
		groups = append(groups, group.ID)
	}

	for _, groupid := range groups {
		scanners := GetIOScannersInGroup(c, groupid)
		if scanners.Count() > 0 && !scanners.HasOnlineScanner() {
			ogcount++
			ognames = append(ognames, groupid2name[groupid])
		}
	}

	return ogcount, ognames
}

func GetIOScannersInGroup (c *resty.Client, id int) (IOScannersInGroup) {
        var result IOScannersInGroup
        resp, err := c.R().Get(fmt.Sprintf("%s/scanner-groups/%d/scanners", c.BaseURL, id))
        if err != nil {
                log.Fatal(err)
        }
        if resp.StatusCode() != 200 {
                log.Fatal("Invalid access and/or secret key")
        }

        json.Unmarshal(resp.Body(), &result)
        return result
}

func (sig IOScannersInGroup) HasOnlineScanner() bool {
	for _, scanner := range sig.Scanners {
		if scanner.Status == "on" {
			return true
		}
	}

	return false
}

func (sig IOScannersInGroup) Count() int {
	return len(sig.Scanners)
}
