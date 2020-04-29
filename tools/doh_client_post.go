package main

import (
	"fmt"

	"github.com/babolivier/go-doh-client"
)

// DNSType implements DNS values.
type DNSType uint16

const (
	// A implements the DNS A type.
	A DNSType = 1
	// NS implements the DNS NS type.
	NS = 2
	// CNAME implements the DNS CNAME type.
	CNAME = 5
	// SOA implements the DNS SOA type.
	SOA = 6
	// PTR implements the DNS PTR type.
	PTR = 12
	// MX implements the DNS MX type.
	MX = 15
	// TXT implements the DNS TXT type.
	TXT = 16
	// AAAA implements the DNS AAAA type.
	AAAA = 28
	// SRV implements the DNS SRV type.
	SRV = 33
)

// DNSClass implements DNS classes.
type DNSClass uint16

const (
	// IN implement the DNS Internet class.
	IN DNSClass = 1
	// CS implements the DNS CSNET class.
	CS = 2
	// CH implements the DNS CH class.
	CH = 3
	// HS implements the DNS Hesiod class.
	HS = 4
	// ANYCLASS implements the DNS * QCLASS.
	ANYCLASS = 255
)

func main() {
	resolver := doh.Resolver{
		Host:  "157.245.57.91", // Change this with your favourite DoH-compliant resolver.
		Class: doh.IN,
	}

	// Perform a A lookup on example.com
	a, _, err := resolver.LookupA("www.example.com")
	if err != nil {
		panic(err)
	}

	if a[0].IP4 == "93.184.216.34" {
		fmt.Println("Method:POST. OK!")
	}

	b, _, err := resolver.LookupTXT("_esni.157.245.57.91")
	if err != nil {
		panic(err)
	}

	fmt.Println(b)

	// client := &http.Client{}
	// req, _ := http.NewRequest("GET", "https://157.245.57.91/dns-query?dns=q80BAAABAAAAAAAAA3d3dwdleGFtcGxlA2NvbQAAAQAB", nil)

	// req.Header.Set("accept", "application/dns-message")
	// res, _ := client.Do(req)
	// resData, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(resData)
}
