/* Instead of timeout, make the genNos() stop when the user hits ENTER key
DO NOT handle user interaction (accepting ENTER key) in the genNos() function
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	stopCh := make(chan struct{})
	fmt.Println("Hit ENTER to stop!")
	go func() {
		fmt.Scanln()
		close(stopCh)
	}()
	wg.Add(1)
	go genNos(wg, stopCh)
	wg.Wait()
	fmt.Println("Done!")
}

// maintainable
func genNos(wg *sync.WaitGroup, stopCh <-chan struct{}) {
	defer wg.Done()
LOOP:
	for i := 1; ; i++ {
		select {
		case <-stopCh:
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println(i)
		}
	}
	fmt.Println("[genNos] stopped")
}
