package worker

import (
	"time"
)

type DealWatcher struct {
}

func NewDealWatcher() *DealWatcher {
	return &DealWatcher{}
}

func (r *DealWatcher) WatchDeals() {
	for {
		time.Sleep(time.Duration(30) * time.Minute)
	}
}

// checks instead

// get network_id, recruiter_addr + job_id (job:recruiter_addr:job_id)
// check if job has already been updated
// update job on database with the deal_network_id and deal_id
