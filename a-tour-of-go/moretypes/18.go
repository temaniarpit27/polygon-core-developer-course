package moretypes

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	resp := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		resp[i] = make([]uint8, dx)

		for j := 0; j < dx; j++ {
			resp[i][j] = uint8((i + j) / 2)
		}
	}

	return resp
}

func MoreTypes18() {
	pic.Show(Pic)
}
