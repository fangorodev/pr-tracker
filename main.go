package main

import (
	"fmt"
	"errors"
)

func oneRepMax(weight float64, reps int, formula string) (float64, error) {
	switch formula {
	case "ep":
		return maxEpley(weight, reps), nil
	case "br":
		return maxBrzycki(weight, reps), nil
	default:
		return 0, errors.New("not valid formula")
	}
}

func maxEpley(weight float64, reps int) float64 {
	return weight * (1 + float64(reps)/30.0)
}

func maxBrzycki(weight float64, reps int) float64 {
	return weight * (36.0 / float64(37-reps))
}

// Weight for a given rep range given the one rep max
func weightForReps(orm float64, reps int, formula string) (float64, error) {
	switch formula {
	case "ep":
		return reverseEpley(orm, reps), nil
	case "br":
		return reverseBrzycki(orm, reps), nil
	default:
		return 0, errors.New("not valid formula")
	}
}

func reverseEpley(orm float64, reps int) float64 {
	return orm / (1.0 + float64(reps)/30.0)
}

func reverseBrzycki(orm float64, reps int) float64 {
	return orm * float64(37-reps) / 36.0
}

func main() {
	percentages := map[int]float64 {
		1: 1.00,
		2: 0.95,
		4: 0.90,
		6: 0.85,
		8: 0.80,
		10: 0.75,
		12: 0.70,
		16: 0.65,
		20: 0.60,
	}

	repOrder := []int{1, 2, 4, 6, 8, 10, 12, 16, 20}

	var lift float64
	var reps int
	var discipline, formula string
	
	fmt.Println("Please, enter your lift and rep count...")
	fmt.Scan(&lift, &reps)
	fmt.Printf("%g x %d", lift, reps)
	fmt.Println()

	fmt.Println("SBD?")
	fmt.Scan(&discipline)

	if(discipline == "S" || discipline == "D") {
		fmt.Println("Using Epley")
		formula = "ep"
	} else {
		fmt.Println("Using Brzycki")
		formula = "br"
	}

	orm, err := oneRepMax(lift, reps, formula)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("One Rep Max Estimate: %g\n", orm)

	// Conversion Tables
	fmt.Println("Amrap recommendations:")
	// for r, p := range percentages {
	// 	amrap := weightForReps(orm, r, formula)
	// 	fmt.Printf("%d reps (%g) --> %g \n", r, p, amrap)
	// }

	for _, r := range repOrder {
		amrap, err := weightForReps(orm, r, formula)
		if err != nil { fmt.Println(err) }
		fmt.Printf("%d reps (%g) --> %g \n", r, percentages[r], amrap)
	}
}
