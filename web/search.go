package web

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

var (
	trailer1 = fakeSearch("trl1.com")
	image1   = fakeSearch("img1.com")
	review1  = fakeSearch("rev1.com")

	trailer2 = fakeSearch("trl2.com")
	image2   = fakeSearch("img2.com")
	review2  = fakeSearch("rev2.com")
)

type SearchResult string
type SearchAlgorithm func(query string) (results []SearchResult)
type SearchFunc func(query string) SearchResult

func fakeSearch(server string) SearchFunc {
	return func(query string) SearchResult {

		start := time.Now()
		time.Sleep(time.Duration(rand.Intn(120)) * time.Millisecond)
		elapsed := time.Since(start)

		results := rand.Intn(20) + 1
		return SearchResult(fmt.Sprintf("%d results on %s in %s", results, server, elapsed))
	}
}

func SerialSearch(query string) (results []SearchResult) {
	results = append(results, trailer1(query))
	results = append(results, image1(query))
	results = append(results, review1(query))
	return
}

func ParallelSearch(query string) (results []SearchResult) {
	c := make(chan SearchResult)
	go func() { c <- trailer1(query) }()
	go func() { c <- image1(query) }()
	go func() { c <- review1(query) }()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}

	return
}

func TimeoutSearch(query string) (results []SearchResult) {
	c := make(chan SearchResult)
	go func() { c <- trailer1(query) }()
	go func() { c <- image1(query) }()
	go func() { c <- review1(query) }()

	timeout := time.After(80 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			log.Println("search timed out")
			return
		}
	}

	return
}

func firstResult(query string, replicas ...SearchFunc) SearchResult {
	c := make(chan SearchResult)
	searchReplicas := func(i int) {
		c <- replicas[i](query)
	}

	for i := range replicas {
		go searchReplicas(i)
	}

	return <-c
}

func FirstSearch(query string) (results []SearchResult) {
	c := make(chan SearchResult)
	go func() { c <- firstResult(query, trailer1, trailer2) }()
	go func() { c <- firstResult(query, image1, image2) }()
	go func() { c <- firstResult(query, review1, review2) }()

	timeout := time.After(80 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			log.Println("search timed out")
			return
		}
	}

	return
}
