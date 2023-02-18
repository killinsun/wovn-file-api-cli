package utils

import "time"

func ClearChannelInterval(concurrency int, ticker *time.Ticker, sem *chan struct{}) {
	for range ticker.C {
		for i:=0; i < concurrency; i++ {
			func(){<- *sem}()
		}
	}
}