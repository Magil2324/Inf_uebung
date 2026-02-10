package rekursion

import "slices"

// Erwartet eine verschachtelte Liste von Ganzzahlen (int).
// Jede innere Liste kann beliebig viele Zahlen enthalten.
// Die Funktion soll alle Zahlen zu einem einzigen Slice zusammenführen („flatten“).
//
// Beispiel:
// Flatten([][]int{
//     {1, 2},
//     {},
//     {3, 4, 5},
//     {6},
// }) → []int{1, 2, 3, 4, 5, 6}
//
// Die Funktion soll iterativ geschrieben werden
func FlattenIterativ(nested [][]int) []int {

	var list []int

	if len(nested) == 0 {
		return list
	}

	for _, innerlist := range nested {
		for _, i := range innerlist {
			if !slices.Contains(list, i) {
				list = append(list, i)
			}
		}
	}

	return list
}
