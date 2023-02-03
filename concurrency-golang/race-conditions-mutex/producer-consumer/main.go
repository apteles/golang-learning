package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizza = 1000

var pizzasMade, pizzasFiled, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

type PizzaOrder struct {
	number  int
	message string
	success bool
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++

	if pizzaNumber <= NumberOfPizza {
		delay := rand.Intn(5) + 1

		fmt.Printf("Received order #%d\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFiled++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("Making pizza #%d. It will take %d seconds...\n", pizzaNumber, delay)

		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)

		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizzaNumber)

		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		}

		p := PizzaOrder{
			number:  pizzaNumber,
			message: msg,
			success: success,
		}

		return &p
	}

	return &PizzaOrder{number: pizzaNumber}
}
func pizzeria(pizzaMaker *Producer) {
	var i = 0

	for {
		currentPizza := makePizza(i)

		if currentPizza != nil {
			i = currentPizza.number

			select {
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}

	}

}

func main() {
	rand.Seed(time.Now().UnixNano())

	color.Cyan("The pizzeria is open for business")
	color.Cyan("=================================")

	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	go pizzeria(pizzaJob)

	for i := range pizzaJob.data {
		if i.number >= NumberOfPizza {
			color.Cyan("Done making pizzas...")
			go func() {
				if err := pizzaJob.Close(); err != nil {
					color.Red("*** Error closing channel!", err)
				}
			}()
		} else {
			if !i.success {
				color.Red(i.message)
				color.Red("The customer is really mad!")
				continue
			}

			color.Green(i.message)
			color.Green("Order #%d is out for delivery!", i.number)
		}
	}
}
