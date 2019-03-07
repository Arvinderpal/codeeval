package codeeval

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const maxSends = 20

// MReceiversOneSender creates m receiver go routines and a single sender which
// writes int values to an output chan processed by the receivers. The sender
// closes the output chan after a certain number of sends.
func MReceiversOneSender(m int) {

	sendCh := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(m)

	// sender routine
	go func() {
		for i := 0; i < maxSends; i++ {
			sendCh <- i
		}
		close(sendCh)
	}()

	// receivers
	for i := 0; i < m; i++ {
		go func(id int) {
			defer wg.Done()
			for v := range sendCh {
				fmt.Printf("(%v): %v\n", id, v)
			}
		}(i)
	}

	wg.Wait()
}

const MaxRandomNumber = 50

// MReceiversNSenders has M receivers, N senders, any one of them says
// "let's end the game" by notifying a moderator to close an additional
// signal channel.
func MReceiversNSenders(m, n int) {
	var stoppedBy int
	rand.Seed(time.Now().UnixNano())
	dataCh := make(chan int)
	trigStopCh := make(chan int)
	stopCh := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(m)

	// moderator
	go func() {
		stoppedBy = <-trigStopCh
		close(stopCh)
		// WHY DOES THE BELOW CODE NOT WORK???
		// for {
		// select {
		// case v := <-trigStopCh:
		// 	fmt.Printf("stop triggered by %v", v)
		// 	close(stopCh)
		// 	return
		// default:
		// }
		// }
	}()

	// senders
	for i := 0; i < n; i++ {
		go func(id int) {
			for {
				select {
				case <-stopCh:
					return
				default:
					value := rand.Intn(MaxRandomNumber)
					if value == 42 {
						// this select is necessary -- two concurrent writes may cause one of them to block
						select {
						case trigStopCh <- id:
						default:
						}
					} else {
						dataCh <- value
					}
				}
			}
		}(i)
	}

	// receivers
	for i := 0; i < m; i++ {
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-stopCh:
					return
				case v := <-dataCh:
					if v == MaxRandomNumber-1 {
						select {
						case trigStopCh <- id:
						default:
						}
					}

				}
			}
		}(i)
	}
	wg.Wait()
	fmt.Printf("stop triggered by %v", stoppedBy)
}
