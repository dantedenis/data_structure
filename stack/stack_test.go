package stack

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestStack_Push(t *testing.T) {
	stack := NewStack()
	var expected []int
	rand.Seed(time.Now().Unix())

	for i := 0; i < (rand.Int()%100)+10; i++ {
		temp := rand.Int() % 1000
		expected = append(expected, temp)
		stack.Push(temp)
	}

	for i := 0; i < len(expected)/2; i++ {
		j := len(expected) - i - 1
		expected[i], expected[j] = expected[j], expected[i]
	}

	if stack.String() != strings.Trim(fmt.Sprint(expected), "[]") {
		t.Error("Invalid test: expected - ", expected, "actual -", stack)
	}
	log.Printf("Test for size: %d", stack.Len())
}
