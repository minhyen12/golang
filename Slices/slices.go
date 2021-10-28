package main

import "golang.org/x/tour/pic" //ctrl + .

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
    for y := range p {
        p[y] = make([]uint8, dx)
        for x := range p[y] {
        // p[y][x] = uint8(x^y)
        // p[y][x] = uint8(x*y)
        p[y][x] = uint8((x+y)/2)
        }
    }
    return p
}

func main() {
	pic.Show(Pic)
}
