package main

import (
	"crypto/tls"
	"fmt"
	"github.com/jackdanger/collectlinks"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)
var visited = make(map[string]bool)
var parent = make(map[string]string)
func intCrawl(path string) {
	/*flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Please provide a uri to crawl")
		os.Exit(1)
	}*/

	root := path
	queue := make(chan string)
	f, err := os.OpenFile("result.csv", 	os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	go func() {queue <- root}()

	for uri := range queue {
		enqueue(uri, queue, f)
	}

}

func enqueue(uri string, queue chan string, f *os.File) {
	visited[uri] = true

	if strings.Contains(uri, "/sv/sv") || strings.Contains(uri, "/en/en") || strings.Contains(uri, "/en/sv") || strings.Contains(uri, "/sv/en") {
		path := fmt.Sprintf("%s%s%s\n", parent[uri]," => ", uri)
		if _, err := f.WriteString(path); err != nil {
			log.Println(err)
		}
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	client := http.Client{Transport: transport}
	resp, err := client.Get(uri)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	links := collectlinks.All(resp.Body)

	for _, link := range links {
		absolute := normalizeUrl(link, uri)

		if absolute != "" && !visited[absolute] {
			parent[absolute] = uri
			go func() { queue <- absolute }()
		}
	}
}

func normalizeUrl(href, base string) string {
	/*if href == "sv/what-is-heated-tobacco" {
		fmt.Println(href)
		//href = "what-is-heated-tobacco"
	}*/
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return ""
	}

	uri = baseUrl.ResolveReference(uri)
	/*if strings.Contains(uri.String(), "/sv/sv") || strings.Contains(uri.String(), "/en/en") {
		fmt.Println(uri)
	}*/
	return uri.String()
}
