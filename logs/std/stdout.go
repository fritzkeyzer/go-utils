package std

import "fmt"

type stdErr struct {
}

func Out() *stdErr {
	return &stdErr{}
}

func (s *stdErr) Print(args ...any) {
	fmt.Print(args...)
	fmt.Println()
}
