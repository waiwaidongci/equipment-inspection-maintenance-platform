package taskpipeline

import "sync"

func Launch(wg *sync.WaitGroup, beforeAdd <-chan struct{}, fn func()) {
	go func() { <-beforeAdd; wg.Add(1); defer wg.Done(); fn() }()
}
