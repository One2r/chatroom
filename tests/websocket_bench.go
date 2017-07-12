package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "websocket_bench"
	app.Usage = "A simple websocket bench tool"
	app.Version = "0.0.1"
	app.Author = "One2r"
	app.Email = "601941036@qq.com"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "ws",
			Value: "",
			Usage: "websocket server url,eg: ws://192.168.1.190",
		},
		cli.IntFlag{
			Name:  "c",
			Value: 10,
			Usage: "establish connections number",
		},
	}
	app.Action = func(c *cli.Context) {
		var wg sync.WaitGroup
		if c.String("ws") == "" {
			fmt.Println("error:Please input the websocket server url ")
			return
		}
		for i := 0; i < c.Int("c"); i++ {
			time.Sleep(10 * time.Millisecond)
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				dialer := websocket.Dialer{ /* set fields as needed */ }
				ws, _, err := dialer.Dial(c.String("ws"), nil)
				if err != nil {
					// handle error
				}
				if err := ws.WriteMessage(websocket.TextMessage, []byte("I am No.:"+strconv.Itoa(i))); err != nil {
					// handle error
				}
				for {
					_, p, err := ws.ReadMessage()
					if err != nil {
						// handle error
					}
					fmt.Println("No.:"+strconv.Itoa(i)+" get message :", string(p))
				}
			}(i)
		}
		wg.Wait()
	}
	app.Run(os.Args)
}
