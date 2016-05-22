/*
    type: algorithm
    theme: math
    sub-theme: calculate
    name: factorial
    author of code: Evgeny @missingdays Bovykin

*/

package math

func Factorial(n int) int{

  if n == 0 || n == 1{

    return 1

  }

  a := 1

  for i := 1; i < n + 1; i++ {
    a = a * i
  }

  return a

}
