package backend

import "fmt"

func init() {
	bindHandlers = append(bindHandlers, &GreetHandler{})
}

type GreetHandler struct {
}

// Greet returns a greeting for the given name
func (g *GreetHandler) Greet(name string) string {
	return fmt.Sprintf("Hello 1 %s, It's show time!", name)
}
