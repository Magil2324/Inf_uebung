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
// Die Funktion soll rekursiv geschrieben werden
func FlattenRekursiv(nested [][]int) []int {
	var list []int

	if len(nested) == 0 {
		return list
	}

	list = append(nested[0], FlattenRekursiv(nested[1:])...)

	for _, i := range list {
		if slices.Contains(list, list[i]) {
			list = slices.Delete(list, i-1, i)
		}
	}

	return list
}
