package day1

import (
	"os"
	"testing"
)

const answer1 = 54239
const answer2 = 55343

func TestDay1(t *testing.T) {
	file, err := os.ReadFile("../../assets/day1.txt")
	if err != nil {
		t.Fatal(err)
	}

	sum := Day1(string(file))
	if sum != answer1 {
		t.Fatalf("Day 1 sum must be %d, but instead got %d", answer1, sum)
	}

	sum = Day1_2(string(file))
	if sum != answer2 {
		t.Fatalf("Day 1_2 sum must be %d, but instead got %d", answer2, sum)
	}
}
