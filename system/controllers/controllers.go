package controllers

import (

	//"gin-stegosaurus/models"

	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/miekg/dns"
	"golang.org/x/net/idna"
)

// XinController hi.
type XinController struct{}

// DNSResponse is dns over https packet for json
type DNSResponse struct {
	Status    int        `json:"Status"`
	TC        bool       `json:"TC"`
	RD        bool       `json:"RD"`
	RA        bool       `json:"RA"`
	AD        bool       `json:"AD"`
	CD        bool       `json:"CD"`
	Questions []Question `json:"Question"`
	Answer    []Answer   `json:"Answer"`
}

type Question struct {
	Name string `json:"name"`
	Type int    `json:"type"`
}

type Answer struct {
	Name string `json:"name"`
	Type int    `json:"type"`
	TTL  int    `json:"TTL"`
	Data string `json:"data"`
}

// DNSQuery hi.
func (t XinController) DNSQuery(c *gin.Context) {
	// h := core.NewH()
	var dnsQuery string
	var dnsPacket *dns.Msg

	checkDNS := c.Request.Header.Get("accept")

	if checkDNS == "" {
		c.String(http.StatusBadRequest, "Invalid argument value: \"dns\"")
		fmt.Println(http.StatusBadRequest, "Invalid argument value: \"dns\"")
		return
	}

	fmt.Println("\ncheckDNS", checkDNS)

	if checkDNS == "application/dns-message" {
		fmt.Println("Method", c.Request.Method)
		switch c.Request.Method {
		case "GET":
			fmt.Println("------------------I am GET------------------")
			dnsQuery = c.Query("dns")

			requestBinary, err := base64.RawURLEncoding.DecodeString(dnsQuery)
			if err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("Invalid argument value: \"dns\" = %q", dnsQuery))
				fmt.Printf("Invalid argument value: \"dns\" = %q\n", dnsQuery)
				return
			}

			if len(requestBinary) == 0 && (checkDNS == "application/dns-message" || checkDNS == "application/dns-udpwireformat") {
				c.String(http.StatusBadRequest, fmt.Sprintf("Invalid argument value: \"dns\" = %q", dnsQuery))
				fmt.Printf("Invalid argument value: \"dns\" = %q\n", dnsQuery)
				return
			}

			if len(requestBinary) == 0 {
				c.String(http.StatusBadRequest, "Invalid argument value: \"dns\"")
				fmt.Printf("Invalid argument value: \"dns\"\n")
				return
			}

			dnsPacket = new(dns.Msg)
			err = dnsPacket.Unpack(requestBinary)
			if err != nil {
				c.String(http.StatusBadRequest, "DNS packet parse failure (%s)", err.Error())
				fmt.Printf("DNS packet parse failure (%s)\n", err.Error())
			}

		case "POST":
			fmt.Println("------------------I am POST------------------")
			body, _ := ioutil.ReadAll(c.Request.Body)
			dnsPacket = new(dns.Msg)
			err := dnsPacket.Unpack(body)
			if err != nil {
				c.String(http.StatusBadRequest, "DNS packet parse failure (%s)", err.Error())
				fmt.Printf("DNS packet parse failure (%s)\n", err.Error())
			}
		}

		fmt.Println("------------------Start Query------------------")
		replay, _, err := new(dns.Client).Exchange(dnsPacket, "8.8.8.8:53")
		if err != nil {
			fmt.Println("Error Exchange", err.Error())
		}

		fmt.Println("Client IP:", c.Request.RemoteAddr)
		fmt.Println("DNS Query:", dnsPacket.Question[0].String())

		fmt.Println("------------------Set Replay------------------")
		replay.SetReply(dnsPacket)

		now := time.Now().UTC().Format(http.TimeFormat)
		c.Writer.Header().Set("Date", now)
		c.Writer.Header().Set("Last-Modified", now)
		c.Writer.Header().Set("Content-Type", "application/dns-message")
		c.Writer.Header().Set("Vary", "Accept")
		c.Writer.Header().Set("Access-control-allow-origin", "*")
		fmt.Println("------------------Set Header------------------")

		fmt.Println("DNS Answer:", replay.Answer)

		respBytes, err := replay.Pack()

		c.String(http.StatusOK, string(respBytes))
	}

	if checkDNS == "application/dns-json" {
		dnsQuery, _ := idna.ToASCII(c.Query("name"))
		fmt.Println("Query:", dnsQuery)
		queryType := c.Query("type")
		fmt.Println("queryType:", queryType)

		m := new(dns.Msg)
		m.SetQuestion(dnsQuery+".", dns.StringToType[queryType])
		replay, _, err := new(dns.Client).Exchange(m, "8.8.8.8:53")
		if err != nil {
			fmt.Println("Error Exchange", err.Error())
		}

		fmt.Println(replay.Answer[0].String())

		var response DNSResponse
		response = DNSResponse{
			Status: 0, TC: false, RD: true, RA: true, AD: true, CD: false,
			Questions: []Question{
				Question{
					Name: replay.Question[0].Name,
					Type: int(replay.Question[0].Qtype),
				},
			},
			Answer: []Answer{
				Answer{
					Name: replay.Answer[0].Header().Name,
					Type: int(replay.Answer[0].Header().Rrtype),
					TTL:  int(replay.Answer[0].Header().Ttl),
					Data: strings.Split(replay.Answer[0].String(), "\t")[4],
				},
			},
		}

		resp, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
			return
		}

		c.String(http.StatusOK, string(resp))
	}

	return
}

// NotFound hi.
func (t XinController) NotFound(c *gin.Context) {

	c.String(http.StatusOK, "This is the 404 page.\nI watch you watch me.")
	return
}
