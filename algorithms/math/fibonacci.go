/*
    type: algorithm
    theme: math
    sub-theme: calculate
    name: fibonacci
    author of code: example is taken from golang.org (official golang site)

*/

package math

func Fib(n int) int {

  a, b := 0, 1

  for i:=0; i < n; i++ {
	a, b = b, a+b
  }

  return a

}
