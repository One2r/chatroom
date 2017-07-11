package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		time.Sleep(10 * time.Millisecond)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			dialer := websocket.Dialer{ /* set fields as needed */ }
			ws, _, err := dialer.Dial("ws://192.168.190.131:8088?room=1&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTk4MjI4MzMsImlzcyI6ImNoYXRyb29tIiwibmJmIjoxNDk5NzM2NDMzLCJzdWIiOnsiSUQiOi0xLCJUeXBlIjoiYWRtaW4iLCJVc2VybmFtZSI6ImFkbWluIn19.4WCrjL6rlhP8Xih88KXSmYAJMk0r11cpd59gn3Rfobw", nil)
			if err != nil {
				// handle error
			}
			if err := ws.WriteMessage(websocket.TextMessage, []byte("i am : "+strconv.Itoa(i))); err != nil {
				// handle error
			}
			for {
				_, p, err := ws.ReadMessage()
				if err != nil {
					// handle error
				}
				fmt.Println(strconv.Itoa(i)+" get message :", string(p))
			}
		}(i)
	}
	wg.Wait()
}
