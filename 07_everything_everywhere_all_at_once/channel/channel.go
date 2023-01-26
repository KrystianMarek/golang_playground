package channel

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Pulsar struct {
	Prefix string
	Hz     int
	Pipe   chan string
	index  int
	Wg     sync.WaitGroup
	Ctx    context.Context
}

func (p Pulsar) Pulse() {
	go func() {
		for {
			select {
			case <-p.Ctx.Done():
				p.Wg.Done()
				return
			default:
				{
					p.index += 1
					p.Pipe <- fmt.Sprintf("%s|%v", p.Prefix, p.index)
					second := time.Millisecond * 1000
					sleepTime := int(second) / p.Hz
					time.Sleep(time.Duration(sleepTime))
				}
			}
		}
	}()
}

func NewPulsar(wg sync.WaitGroup, ctx context.Context, prefix string, pipe chan string) *Pulsar {
	wg.Add(1)
	p := Pulsar{Prefix: prefix, Hz: 2, index: 0, Wg: wg, Ctx: ctx, Pipe: pipe}
	return &p
}

func pulsarDemo() {
	var wg sync.WaitGroup
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(60))
	pipe := make(chan string, 1)

	p1 := NewPulsar(wg, ctx, "P1", pipe)
	p1.Pulse()
	go func() {
		for msg := range pipe {
			fmt.Println(msg)
		}
	}()

}

func simpleChannel() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)
}

func withWaitGroups() {
	var wg sync.WaitGroup //  initialise a counter
	wg.Add(2)             // two go routines to wait for

	messages := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			messages <- i
			time.Sleep(time.Millisecond * 500)
		}

		close(messages) // close channel when done
		wg.Done()
	}()

	go func() {
		for {
			msg, open := <-messages

			if !open {
				break
			}
			fmt.Println(msg)
		}

		wg.Done()
	}()

	wg.Wait() // block until all routines are done executing
}

func withCapacity() {
	messages := make(chan string, 2)
	messages <- "Yer"
	messages <- "a"

	msg := <-messages
	fmt.Println(msg)

	msg = <-messages
	fmt.Println(msg)

	messages <- "Wizard"

	msg = <-messages
	fmt.Println(msg)
}

type SimpleChannel struct{}

func (s SimpleChannel) Demo() {
	fmt.Println("Simple Channel:")
	simpleChannel()
	fmt.Println("withWaitGroups:")
	withWaitGroups()
	fmt.Println("withCapacity:")
	withCapacity()
	fmt.Println("pulsarDemo:")
	pulsarDemo()
}
