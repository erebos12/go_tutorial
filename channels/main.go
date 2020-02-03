package main

import "net/http"

import "fmt"

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazone.com",
		"https://redi-be.stg.fs-acc.pwc.delivery/",
	}
	c := make(chan string) // channel of type string
	for _, link := range links {
		// KEYWORD "go" creates a new Go-Routine -> func checkLink() runs in separate routine
		go checkLink(link, c) // use "go" keyword only before function calls
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " might be down!")
		c <- "Might be down"
		return
	}
	fmt.Println(link, "is up!")
	c <- "Its up and running."
}
