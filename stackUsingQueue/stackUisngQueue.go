package main

import "fmt"

type Stack struct {
	container []int
}

func (s *Stack) Push(value int)  {
	s.container = append(s.container, value)
}

func(s *Stack) Pop() int {
	if !s.isEmpty() {
		pop := s.container[len(s.container)-1]
		s.container = s.container[:len(s.container)-1]
		return pop
	}
	return -1
}

func(s *Stack) Peek() int {
	return s.container[len(s.container)-1]
}

func(s *Stack) isEmpty() bool {
	if len(s.container) == 0 {
		return true
	}
	return false
}

func(s *Stack) enque(value int) {
	s.Push(value)
}

func(s *Stack) deque() int {
	if !s.isEmpty() {
		pop := s.container[0]
		s.container = s.container[1:]
		return pop
	}
	return -1
}

func main() {
	s := &Stack{}
	values := []int{1,2,3,4,5}
	for _, value := range values {
		s.enque(value)
	}
	fmt.Println(s.container)
	fmt.Println(s.Pop())
	fmt.Println(s.container)
	fmt.Println(s.deque())
	fmt.Println(s.container)
	fmt.Println(s.isEmpty())
}
