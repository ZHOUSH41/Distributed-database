package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// slice
func Pic(dx, dy int) [][]uint8 {
	ans := make([][]uint8, dy)
	for i:= 0; i < dy; i++ {
		ans[i] = make([]uint8, dx)
		for j:=0;j < dx; j++ {
			ans[i][j] = uint8(i*j)
		}
	}
	return ans
}

// map
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		m[word] += 1
	}
	return m
}

// Stringer
type IPAddr [4]byte

func (v IPAddr) String() string {
	s := make([]string, len(v))
	for i,val := range v {
		s[i] = strconv.Itoa(int(val))
	}
	return fmt.Sprintf(strings.Join(s, "."))
}

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("My favorite number is", rand.Intn(100))
}