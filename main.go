package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"io/ioutil"
	"strings"
)
type coctail struct {
	name  string
	items map[string]float64
	lemon float64
}

func newCoctail(name string) coctail {
	c := coctail{
		name:  name,
		items: map[string]float64{},
		lemon: 0,
	}
	return c
}

func (c coctail) format() string {
	fs := "*** Coctail contains: *** \n"
	var total float64 = 0

	for k, v := range c.items {
		fs += fmt.Sprintf("%-10v %v ml\n", k+":", v)
		total += v
	}
	fs += fmt.Sprintf("%-10v %0.2f ml\n", "lemon juice:", c.lemon)

	fs += fmt.Sprintf("%-10v %0.2f ml", "total:", total+c.lemon)
	return fs

}
func (c *coctail) checkTotal() bool {
	var total float64 = 0
	for _, v := range c.items {
		total += v
	}
	return total+c.lemon > 200
}
func (c *coctail) printTotal() float64 {
	var total float64 = 0
	for _, v := range c.items {
		total += v
	}
	return total + c.lemon
}

func (c *coctail) updateLemon(lemon float64) {
	c.lemon = lemon
}

func (c *coctail) addItem(name string, amount float64) {
	c.items[name] = amount
}

func (c *coctail) save() {
	data := []byte(c.format())

	err := ioutil.WriteFile("coctails/"+c.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("*** Coctail was saved to file ***")
}

//geting user input
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err

}

func createCoctail() coctail {
	//reads from terminal -os.
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create new coctail name :", reader)

	c := newCoctail(name)
	fmt.Println("Created coctail name -", c.name)
	return c
}

func promptOptions(c coctail) {

	if c.checkTotal() {
		fmt.Println("!!! Reached the maximum amount (amount must be less than 200ml in total) !!!")
		c.save()
		fmt.Println("*** You have saved your coctail", c.name)
		os.Exit(1)

	}
	reader := bufio.NewReader(os.Stdin)
	option, _ := getInput("\n*** Choose option:*** \n j - add juice, a - add alcohol,l - add squeezed lemon,\n s - save coctail, t - total in glass right now, e - exit\n ", reader)
	switch option {
	case "j":
		name, _ := getInput("---> Juice name: ", reader)
		amount, _ := getInput("---> Amount (in ml):", reader)

		p, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			fmt.Println("*** The amuont must be a number ***")
			promptOptions(c)
		}
		if c.checkTotal() {
			fmt.Println("!!! Reached the maximum amount  !!!")
			promptOptions(c)

		}
		if p > 200 {
			fmt.Println("!!! The amuont must be less than 200ml  !!!")
			promptOptions(c)
		}
		c.addItem(name, p)
		fmt.Println("*** Item added -", name, amount)
		promptOptions(c)

	case "a":
		name, _ := getInput("---> Alcohol name: ", reader)
		amount, _ := getInput("---> Amount (in ml):", reader)

		p, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			fmt.Println("!!! The amuont must be a number !!!")
			promptOptions(c)
		}
		if c.checkTotal() {
			fmt.Println("!!! Reached the maximum amount  !!!")
			promptOptions(c)

		}
		if p > 200 {
			fmt.Println("!!! The amuont must be less than 200ml  !!!")
			promptOptions(c)
		}
		c.addItem(name, p)
		fmt.Println("*** Item added -", name, amount)
		promptOptions(c)

	case "s":
		c.save()
		fmt.Println("*** You have saved your coctail", c.name)

	case "l":
		lemon, _ := getInput("--> Enter lemon juice amount (ml): ", reader)

		l, err := strconv.ParseFloat(lemon, 64)
		if err != nil {
			fmt.Println("!!! The lemon juice amount must be a number !!!")
			promptOptions(c)
		}
		if c.checkTotal() {
			fmt.Println("!!! Reached the maximum amount (amount must be less than 200ml in total) !!!")
			promptOptions(c)

		}
		if l > 200 {
			fmt.Println("!!! The amuont must be less than 200ml  !!!")
			promptOptions(c)
		}
		c.updateLemon(l)
		fmt.Println("*** Lemon juice added -", lemon)
		promptOptions(c)

	case "e":
		fmt.Println("!!! Your data wasn't saved !!!")
		os.Exit(1)
	case "t":
		fmt.Printf("--- Amount in the glass right now %v ml \n", c.printTotal())
		promptOptions(c)

	default:
		fmt.Println("!!! That was not a valid option !!!")
		promptOptions(c)
	}

}

func main() {
	mycoctail := createCoctail()
	promptOptions(mycoctail)

}
