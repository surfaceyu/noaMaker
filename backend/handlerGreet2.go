package backend

import "fmt"

func init() {
	bindHandlers = append(bindHandlers, &GreetHandler2{})
}

type Person struct {
	Name    string   `json:"name"`
	Age     uint8    `json:"age"`
	Address *Address `json:"address"`
}

type Address struct {
	Street   string `json:"street"`
	Postcode string `json:"postcode"`
}

type GreetHandler2 struct {
}

// Greet returns a greeting for the given name
func (g *GreetHandler2) Greet2(p Person) string {
	return fmt.Sprintf("Hello 2 %s-%d, It's show time!", p.Name, p.Age)
}
