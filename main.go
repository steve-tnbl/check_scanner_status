package main

import "crypto/tls"
import "fmt"
import "os"
import "strings"
import "time"
//import "github.com/davecgh/go-spew/spew"
import "github.com/DavidGamba/go-getoptions"
import "github.com/go-resty/resty/v2"
import "github.com/olorin/nagiosplugin"

func main() {
	// Parse options
	opt := getoptions.New()
	var cloudvm bool
	var server string
	var insecure bool
	var pedantic bool
	var accesskey string
	var secretkey string
	var warn bool
	// Vulnerability Management / Tenable.io Settings
	opt.BoolVar(&cloudvm, "cloud", false)
	// Security Center / Tenable.sc Settings
	opt.StringVar(&server, "server", "", opt.Alias("h"))
	opt.BoolVar(&insecure, "insecure", false, opt.Alias("k"))
	// Settings for both
	opt.StringVar(&accesskey, "accesskey", "", opt.Alias("a"))
	opt.StringVar(&secretkey, "secretkey", "", opt.Alias("s"))
	opt.BoolVar(&pedantic, "pedantic", false, opt.Alias("P"))
	opt.BoolVar(&warn, "warn", false, opt.Alias("W"))
	_, opterr := opt.Parse(os.Args[1:])
        
	if opterr != nil || opt.Called("help") {
		fmt.Fprint(os.Stderr, opt.Help())
		os.Exit(1)
	}

	// Initialize Nagios
	check := nagiosplugin.NewCheck()
	defer check.Finish()

	// Initialize HTTP REST Client
	client := resty.New()
	client.SetHeader("Content-Type", "application/json")
	client.SetHeader("Accept", "application/json")
	client.SetTimeout(5 * time.Second)

	// Construct the Base URL
	if cloudvm {
		client.BaseURL = "https://cloud.tenable.com"
	} else {
		client.BaseURL = "https://"+server
	}

	// Construct authentication header
	if cloudvm {
		client.SetHeader("X-ApiKeys", fmt.Sprintf("accessKey=%s; secretKey=%s", accesskey, secretkey))
	} else {
		client.SetHeader("x-apikey", fmt.Sprintf("accesskey=%s; secretkey=%s;", accesskey, secretkey))
	}
		
	// Allow untrusted cert if the user specified --insecure/-k
	client.SetTLSClientConfig(&tls.Config{ InsecureSkipVerify: insecure })

	// Vulnerability Management logic
	if cloudvm {
		if len(accesskey) != 64 ||
		   len(secretkey) != 64 {
			check.AddResult(nagiosplugin.UNKNOWN, "Access key and/or secret key not set")
			check.Finish()
		}

		scanners := GetIOScanners(client)
		ok, notok := CheckIOScannerStatus(scanners)
		if notok == 0 {
			// All good!
			check.AddResult(nagiosplugin.OK, fmt.Sprintf("%d scanners online", ok))
			check.Finish()
		}
		if notok > 0 && pedantic {
			check.AddResult(DefaultStatus(warn), fmt.Sprintf("%d scanners are offline", notok))
			check.Finish()
		}

		scannergroups := GetIOScannerGroups(client)
		ogcount, ognames := CheckIOOrphanedGroups(client, scannergroups)
		
		if ogcount > 0 {
			check.AddResult(DefaultStatus(warn), fmt.Sprintf("These groups have no active scanners: %s", strings.Join(ognames, ",")))
			check.Finish()
		}
	}

	// Security Center logic (default)
	if !cloudvm {
		if len(server) == 0 {
			check.AddResult(nagiosplugin.UNKNOWN, "No server specified")
			check.Finish()
		}

		if len(accesskey) != 32 ||
		   len(secretkey) != 32 {
			check.AddResult(nagiosplugin.UNKNOWN, "Access key and/or secret key not set")
			check.Finish()
		}

		scanners := GetSCScanners(client)
		ok, notok := CheckSCScannerStatus(scanners)
		if notok == 0 {
			// All good!
			check.AddResult(nagiosplugin.OK, fmt.Sprintf("%d scanners online", ok))
			check.Finish()
		}
		if notok > 0 && pedantic {
			check.AddResult(DefaultStatus(warn), fmt.Sprintf("%d scanners are offline", notok))
			check.Finish()
		}

		scanzones := GetSCScanZones(client)
		ozcount, oznames := CheckSCOrphanedZones(scanzones)

		if ozcount > 0 {
			check.AddResult(DefaultStatus(warn), fmt.Sprintf("These zones have no working scanners: %s", strings.Join(oznames, ",")))
			check.Finish()
		}
	}
}

func DefaultStatus (warn bool) (nagiosplugin.Status) {
	if warn { return nagiosplugin.WARNING }
	return nagiosplugin.CRITICAL
}
