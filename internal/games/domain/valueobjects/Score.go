package valueobjects

import "fmt"

type Score struct {
	current     int
	maxPossible int
}

func NewScore(current, maxPossible int) (Score, error) {
	if maxPossible <= 0 {
		return Score{}, fmt.Errorf("puntación maxima debe ser mayor a 0")
	}

	if current <= 0{
		current = 0
	}

	if current > maxPossible {
		current = maxPossible
	}

	return Score{
		current: current,
		maxPossible: maxPossible,
	}, nil
}

func (s Score) Current() int {
    return s.current
}

func (s Score) MaxPossible() int {
    return s.maxPossible
}