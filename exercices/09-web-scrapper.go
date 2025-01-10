package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func fetchURL(cont context.Context, url string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	request, err := http.NewRequestWithContext(cont, "GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		select {
		case <- cont.Done():
			fmt.Println("Timed out.")
			return
		default:
			fmt.Println("Error fetching URL")
			return
		}
	}
	defer response.Body.Close()

	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Fetched URL", url)
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.linkedin.com",
		"https://www.linkedin.com",
		"https://github.com/m-elhamlaoui/refactoring-24-younessothmane/commits/main/",
		"https://www.linkedin.com",
	}

	cont, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() 

	var waitGroup sync.WaitGroup

	for _, url := range urls {
		waitGroup.Add(1) 
		go fetchURL(cont, url, &waitGroup)
	}

	waitGroup.Wait()
	fmt.Println("done")
}
