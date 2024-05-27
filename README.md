**THIS CODE IS NOT COVERED BY TENABLE SUPPORT. IT IS PROVIDED ONLY AS A PROOF OF CONCEPT.**

## Compilation

To build a binary from the source code, just run the `make` command.
This will, by default, create binaries for multiple Operating Systems and architectures.

If you need binaries for more/other OS/archs, use the environment variables `GOOS` and `GOARCH` as describe in
https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63

## Usage

This code supports two Tenable products, Security Center (on-prem) and Vulnerability Management (SaaS).
The default mode is on-prem, use the paramter `--cloud` to query scanners connect to the cloud product

### Required Parameters

`-a | --accesskey` -> To provide your access key for the API (SC or VM)

`-s | --secretkey` -> To provide your secret key for the API (SC or VM)

### Optional Parameters

`-h | --server` -> Set the hostname of your Security Center console. Does not apply to VM.

`--insecure` -> Use this if your Security Center certificate is untrusted. Does not apply to VM.

`-p | --pedantic` -> Generate alarms for offline scanners

`-w | --warn` -> Generate a warning instead of a criticial alarm (default)



