package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	ls := []struct {
		item        string
		csvDuration int
		zipDuration int
		uplDuration int
	}{
		{item: "item: 1", csvDuration: 2, zipDuration: 3, uplDuration: 5},
		{item: "item: 2", csvDuration: 5, zipDuration: 6, uplDuration: 9},
		{item: "item: 3", csvDuration: 6, zipDuration: 3, uplDuration: 1},
		{item: "item: 4", csvDuration: 9, zipDuration: 1, uplDuration: 9},
		{item: "item: 5", csvDuration: 12, zipDuration: 4, uplDuration: 4},
	}

	for _, l := range ls {
		wg.Add(1)
		go func(str string) {
			defer wg.Done()
			processItem(l)
		}(l.item)
	}

	wg.Wait()
	log.Println("Completed")
}

func processItem(item struct {
	item        string
	csvDuration int
	zipDuration int
	uplDuration int
}) {
	// creating csv
	time.Sleep(time.Second * time.Duration(item.csvDuration))
	log.Printf("CSV created for: %v ", item)

	// creating zip
	time.Sleep(time.Second * time.Duration(item.zipDuration))
	log.Printf("creating zip: %v ", item)
	// sending zip
	time.Sleep(time.Second * time.Duration(item.uplDuration))
	log.Printf("sending zip: %v ", item)
}
