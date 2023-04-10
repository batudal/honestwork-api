package main

import (
	"fmt"
	"github.com/takez0o/honestwork-api/api/worker"
)

func startWorkers() {
	fmt.Println("Starting workers...")
	rating_watcher := worker.NewRatingWatcher()
	go rating_watcher.WatchRatings()
	revenue_watcher := worker.NewRevenueWatcher()
	go revenue_watcher.WatchRevenues()
	event_subscriber := worker.NewEventSubscriber()
	go event_subscriber.Subscribe()
	deal_watcher := worker.NewDealWatcher()
	go deal_watcher.WatchDeals()
	fmt.Println("Workers online.")
}
