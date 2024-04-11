package day2

import (
	"os"
	"testing"
)

var answer1 = 2406
var answer2 = 78375

func TestDay2(t *testing.T) {
	file, err := os.ReadFile("../../assets/day2.txt")
	if err != nil {
		t.Fatal(err)
	}

	sum := Day2(string(file))
	if sum != answer1 {
		t.Fatalf("Day 2 sum must be %d, but instead got %d", answer1, sum)
	}

	sum = Day2_2(string(file))
	if sum != answer2 {
		t.Fatalf("Day 2 sum must be %d, but instead got %d", answer2, sum)
	}
}
