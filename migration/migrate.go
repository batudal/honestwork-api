package main

import "sync"

func main() {
	wg := sync.WaitGroup{}
	wg.Add(4)
	go MigrateJobs(&wg)
	go MigrateSkills(&wg)
	go MigrateUsers(&wg)
	go MigrateTxs(&wg)
	wg.Wait()
}
