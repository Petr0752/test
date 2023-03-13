package main

import (
	"fmt"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	tempGroups := make(map[int][]float64)

	for _, temp := range temperatures {
		group := int(temp / 10)
		tempGroups[group] = append(tempGroups[group], temp)
	}

	for group, temps := range tempGroups {
		fmt.Printf("%d: %v\n", group*10, temps)
	}
}

// Этот код создает массив температурных колебаний и проходит по нему, вычисляет группу (разделяет температуру на 10 и округляет до целого),
//  и добавляет температуру в соответствующую группу в tempGroups.
//  Затем код проходит по tempGroups и выводит группу и температуры в этой группе в stdout.
