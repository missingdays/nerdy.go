/*
    type: algorithm
    theme: arrays
    sub-theme: sorting
    name: insertion sorting
    author of code: Evgeny @missingdays Bovykin

*/

package sorting

import "fmt"

func InsertionSort(a []int) {

    for i := 1; i < len(a); i++ {
        value := a[i]

        j := i - 1

        for j >= 0 && a[j] > value {
            a[j+1] = a[j]

            j = j - 1
        }

        a[j+1] = value
    }
}
