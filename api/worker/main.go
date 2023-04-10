package worker

import (
	"fmt"
)

func main() {
	fmt.Println("Starting workers...")
	rating_watcher := NewRatingWatcher()
	go rating_watcher.WatchRatings()
	revenue_watcher := NewRevenueWatcher()
	go revenue_watcher.WatchRevenues()
	event_subscriber := NewEventSubscriber()
	go event_subscriber.Subscribe()
	deal_watcher := NewDealWatcher()
	go deal_watcher.WatchDeals()
	fmt.Println("Workers online.")
}
