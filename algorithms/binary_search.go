/*
    type: algorithm
    theme: arrays
    sub-theme: search
    name: binary-search
    author of code: Evgeny @missingdays Bovykin

*/


package algorithms

func BinarySearch(a []float64, value float64) int {

    low := 0

    high := len(a) - 1

    for low <= high {

        //Select middle point
        mid := (low + high) / 2

        if value < a[mid] {
            high = mid - 1
		} else if a[mid] < value {
            low = mid + 1
		} else {
            return mid
        }
    }

    return -1
}
