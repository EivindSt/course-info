package main

import (
	"fmt"
	"sync"
	"time"
)

type syncmap struct {
	ma map[string]string
	mu *sync.Mutex
}

func main() {
	smap := &syncmap{make(map[string]string), &sync.Mutex{}}

	for i := 0; i < 1000; i++ {
		go func() {
			for {
				x := smap.lookup("UiS")
				fmt.Println(x)
				time.Sleep(10 * time.Millisecond)
			}
		}()
	}

	for i := 0; i < 1000; i++ {
		go func(j int) {
			val := fmt.Sprintf("University of Stavanger %d", j)
			for {
				smap.insert("UiS", val)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	ch := make(chan bool)
	<-ch
}

func (m *syncmap) lookup(key string) string {

}

func (m *syncmap) insert(key, value string) {

}
