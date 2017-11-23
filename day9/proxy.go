package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/mssola/user_agent"
	"gopkg.in/elazarl/goproxy.v1"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.OnRequest(goproxy.ReqHostMatches(regexp.MustCompile("^.*$"))).DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			//parse user agent string
			ua := user_agent.New(r.UserAgent())
			fmt.Printf("Is mobile: %v\n", ua.Mobile()) // => false
			fmt.Printf("Is bot: %v\n", ua.Bot())       // => false
			fmt.Printf("Mozilla: %v\n", ua.Mozilla())  // => "5.0"

			fmt.Printf("Platform: %v\n", ua.Platform()) // => "X11"
			fmt.Printf("OS: %v\n", ua.OS())             // => "Linux x86_64"

			nameE, versionE := ua.Engine()
			fmt.Printf("Engine: %v\n", nameE)            // => "AppleWebKit"
			fmt.Printf("Engine version: %v\n", versionE) // => "537.11"

			nameB, versionB := ua.Browser()
			fmt.Printf("Browser: %v\n", nameB)            // => "Chrome"
			fmt.Printf("Browser version: %v\n", versionB) // => "23.0.1271.97"

			if ua.Bot() || nameB == "curl" {
				return r, goproxy.NewResponse(r,
					goproxy.ContentTypeText, http.StatusForbidden,
					"Don't waste your time!")
			}
			return r, nil
		})
	log.Fatal(http.ListenAndServe(":8080", proxy))
}
