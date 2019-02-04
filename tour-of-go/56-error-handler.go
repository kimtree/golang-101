package main

import (
    "fmt"
)

type ErrorNegativeSqrt float64

func (e ErrorNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %g", e)
}

func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, ErrorNegativeSqrt(f)
    }
    return 0, nil
}

func main() {
    fmt.Println(Sqrt(-2))
}

