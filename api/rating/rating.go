package rating

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/takez0o/honestwork-api/utils/schema"
	"github.com/takez0o/honestwork-api/utils/web3"
)

func WatchRatings(rs_job *redisearch.Client, rs_user *redisearch.Client) {
	for {
		fetchAllRatings(rs_job, rs_user)
		time.Sleep(time.Duration(30) * time.Minute)
	}
}

func fetchAllRatings(rs_job *redisearch.Client, rs_user *redisearch.Client) {
	listers := fetchAllListers(rs_job)
	members := fetchAllMembers(rs_user)
	fmt.Println("Members -> ", members)
	for _, lister := range listers {
		fmt.Println("Fetching lister rating...")
		updateRating(lister)
	}
	for _, member := range members {
		fmt.Println("Fetching member rating...")
		updateRating(member)
	}
}

func search(length int, f func(index int) bool) int {
	for index := 0; index < length; index++ {
		if f(index) {
			return index
		}
	}
	return -1
}

func fetchAllListers(redis *redisearch.Client) []string {
	data, _, err := redis.Search(redisearch.NewQuery("*"))
	if err != nil {
		return []string{}
	}
	var jobs []schema.Job
	for _, d := range data {
		translationKeys := make([]string, 0, len(d.Properties))
		for key := range d.Properties {
			translationKeys = append(translationKeys, key)
		}
		var job schema.Job
		err = json.Unmarshal([]byte(fmt.Sprint(d.Properties[translationKeys[0]])), &job)
		if err != nil {
		}
		jobs = append(jobs, job)
	}
	listers := make([]string, 0)
	for _, j := range jobs {
		idx := search(len(listers), func(index int) bool {
			return listers[index] == j.UserAddress
		})
		if idx == -1 {
			listers = append(listers, j.UserAddress)
		}
	}
	return listers
}

func fetchAllMembers(redis *redisearch.Client) []string {
	data, _, err := redis.Search(redisearch.NewQuery("*"))
	if err != nil {
		return []string{}
	}
	var members []string
	for _, d := range data {
		arr := strings.Split(d.Id, ":")
		members = append(members, arr[1])
	}
	return members
}

func updateRating(address string) float64 {
	rating := web3.FetchAggregatedRating(address)
	fmt.Println("Rating for " + address + " is " + fmt.Sprint(rating))
	return rating
}
