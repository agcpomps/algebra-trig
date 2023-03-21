package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bloodAlchoolConcentration(n int, w float64) float64 {
	return 600 * float64(n) / (w * (0.6*float64(n) + 169))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number of drinks... ")
	drink, _ := reader.ReadString('\n')
	drink = strings.TrimSpace(drink)
	intDrink, _ := strconv.Atoi(drink)

	fmt.Print("Enter your weight ")
	weight, _ := reader.ReadString('\n')
	weight = strings.TrimSpace(weight)
	floatWeight, _ := strconv.ParseFloat(weight, 64)

	bac := bloodAlchoolConcentration(intDrink, floatWeight)

	fmt.Printf("Your blood alchool concentration is %.4f\n", bac)

	if bac < 0.05 {
		fmt.Println("Feeling of well-being; mild release of inhibitions; absence of observable effects")
	}
	if bac > 0.05 && bac <= 0.08 {
		fmt.Println("Feeling of relaxation; mild sedation; exaggeratio of emotions and behavior slight impairment of motor skills; increase in reaction time")
	}

	if bac > 0.08 && bac <= 0.12 {
		fmt.Println("Muscle control and speech impaired; difficulty performing motor skills; uncoordinate behavior")
	}
	if bac > 0.12 && bac <= 0.15 {
		fmt.Println("Euphoria; major impairment of physical and mental functions; irresponsible behavior; some difficulty standing, walking, and talking")
	}

	if bac > 0.15 && bac <= 0.35 {
		fmt.Println("Surgical anesthesia; lethal dosage for a small percentage of people")
	}

	if bac > 0.35 && bac < 0.40 {
		fmt.Println("Lethal dosage for 50% of people; severe circulatory and respiratory depression; alcohol poisoning/overdose")
	}
}
