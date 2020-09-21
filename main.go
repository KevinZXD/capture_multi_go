package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type NewMap struct {
	lock *sync.RWMutex
	sm   map[interface{}]interface{}
}

func (m *NewMap) Set(k interface{}, v interface{}) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.sm[k] = v
	return true
}
func request(url string) string {
	//resp, _ := http.Get(url)
	//defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	return string(rand.Intn(100))
}
func capture_multi(n int, result NewMap) {
	//url := "https://suggest.taobao.com/sug?q=VR&code=utf-8"
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		fmt.Printf(string(i))
		go func() {
			res :=  string(rand.Intn(100))
			result.Set(i, res)
			wg.Done()
		}()
	}
	wg.Wait()

}
func main() {
	//channel:=make(chan string)
	result := NewMap{
		lock: new(sync.RWMutex),
		sm:   make(map[interface{}]interface{}),
	}

	var stime = time.Now().Unix()
	capture_multi(10, result)
	var ctime = time.Now().Unix() - stime
	result_map := result.sm
	for k := range result_map {
		fmt.Println(k)
		//fmt.Println(result_map[k])
	}

	fmt.Println(ctime)

}
