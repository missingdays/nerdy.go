/*
    type: algorithm
    theme: arrays
    sub-theme: sorting
    name: bubble sorting
    author of code: Evgeny @missingdays Bovykin

*/

package sorting

import "fmt"

func BubbleSort(a []int) {

    for itemCount := len(a) - 1; itemCount > 0 ; itemCount-- {

        hasChanged := false

        for index := 0; index < itemCount; index++ {
            if a[index] > a[index+1] {
                a[index], a[index+1] = a[index+1], a[index]

                hasChanged = true
            }
        }

        if hasChanged == false {
            break
        }
    }
}
