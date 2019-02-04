package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    newArray := make([][]uint8, dy)
    for i := range newArray {
        newArray[i] = make([]uint8, dx)
        for j := range newArray[i] {
            newArray[i][j] = uint8(i+j)
        }
    }

    return newArray
}

func main() {
    pic.Show(Pic)
}

