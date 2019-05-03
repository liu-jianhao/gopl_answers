package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// return the number of elements
func (s *IntSet) Len() int {
	res := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for i := 0; i < 64; i++ {
			if word&(1<<uint(i)) != 0 {
				res++
			}
		}
	}
	return res
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] &= ^(1 << bit)
}

// remove all elements from the set
func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
	newWords := make([]uint64, len(s.words))
	copy(newWords, s.words)
	return &IntSet{newWords}
}

func (s *IntSet) AddAll(values ...int) {
	for _, v := range values {
		s.Add(v)
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	var min int
	if len(s.words) < len(t.words) {
		min = len(s.words)
	} else {
		min = len(t.words)
	}
	for i := 0; i < min; i++ {
		s.words[i] &= t.words[i]
	}
	s.words = s.words[:min]
}

func (s *IntSet) Elems() []int {
	var res []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				res = append(res, 64*i+j)
			}
		}
	}
	return res
}

func Example_one() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func Example_three() {
	var x IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)

	fmt.Println(x.String()) // "{1 2 3}"

	x.Remove(2)
	fmt.Println(x.String()) // "{1 3}"

	y := x.Copy()
	fmt.Println(y.String()) // "{1 3}"
	fmt.Println(y.Len())    // 2

	y.Clear()
	fmt.Println(y.String()) // "{}"
}

func Example_four() {
	var x IntSet
	x.AddAll(1, 2, 3)
	fmt.Println(x.String()) // "{1 2 3}"
}

func Example_five() {
	var x, y IntSet
	x.AddAll(1, 2, 3)
	y.AddAll(2, 3, 5)
	fmt.Println(x.String()) // "{1 2 3}"
	fmt.Println(y.String()) // "{2 3 5}"
	x.IntersectWith(&y)
	fmt.Println(x.String()) // "{2 3}"
}

func Example_six() {
	var x IntSet
	x.AddAll(1, 2, 3)
	fmt.Println(x.Elems())
}

func main() {
	Example_one()
	Example_two()

	Example_three()
	Example_four()
	Example_five()
	Example_six()
}
