package stack

import "errors"

type Stack struct {
	data []string
}

// Push agrega `x` en el tope de la pila.
func (s *Stack) Push(x string) {
	s.data = append(s.data, x)
}

// Pop remueve y retorna el valor en el tope de la pila.
// Devuelve un `error` si la pila está vacía.
func (s *Stack) Pop() (string, error) {
	if s.Size() == 0 {
		return "", errors.New("empty stack")
	}

	x := s.data[s.Size()-1]
	s.data = s.data[:s.Size()-1]
	return x, nil
}

// Size devuelve el número de elementos en la pila.
func (s *Stack) Size() int {
	return len(s.data)
}
