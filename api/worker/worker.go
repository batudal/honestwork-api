package worker

import (
	"fmt"
	"time"
)

func Start() {
	rating_watcher := NewRatingWatcher()
	go rating_watcher.WatchRatings()
	revenue_watcher := NewRevenueWatcher()
	go revenue_watcher.WatchRevenues()
	event_subscriber := NewEventSubscriber()
	go event_subscriber.Subscribe()
	deal_watcher := NewDealWatcher()
	go deal_watcher.WatchDeals()
	fmt.Println("Workers online @", time.Now().Format("2006-01-02 15:04:05"))
}
