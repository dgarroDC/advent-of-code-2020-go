package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func reverse(s string) string {
	reversed := ""
	for _, c := range s {
		reversed = string(c) + reversed
	}
	return reversed
}

func normalize(s string) string {
	reversed := reverse(s)
	if strings.Compare(s, reversed) < 0 {
		return s
	} else {
		return reversed
	}
}

func transform(image []string, rot int, flipV bool) []string {
	for rot != 0 {
		image = rotate(image)
		rot--
	}
	if flipV {
		dim := len(image)
		for i, j := 0, dim-1; i < j; i, j = i+1, j-1 {
			image[i], image[j] = image[j], image[i]
		}
	}
	return image
}

func rotate(image []string) []string {
	dim := len(image)
	newImage := make([]string, dim)
	for i := 0; i < dim; i++ {
		newImage[i] = ""
		for j := 0; j < dim; j++ {
			newImage[i] += string(image[dim-1-j][i])
		}
	}
	return newImage
}

func printImage(image []string) {
	for _, row := range image {
		fmt.Println(row)
	}
}

func rotateBorders(borders *[]string) {
	*borders = []string{reverse((*borders)[3]), (*borders)[0], reverse((*borders)[1]), (*borders)[2]}
}

func main() {
	fileBytes, _ := ioutil.ReadFile("20.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines) - 1]

	borderTiles := make(map[string][]int)
	tileBorders := make(map[int][]string)
	tileImages := make(map[int][]string)
	id := -1
	borders := make([]string, 4)
	image := make([]string, 0)
	for i, line := range lines {
		if id == -1 {
			id, _ = strconv.Atoi(strings.Split(strings.Split(line, "Tile ")[1], ":")[0])
		} else if line == "" {
			borders[2] = lines[i-1]
			image = image[:len(image)-1]
			for _, border := range borders {
				borderTiles[normalize(border)] = append(borderTiles[normalize(border)], id)
			}
			tileBorders[id] = borders
			tileImages[id] = image
			id = -1
			borders = make([]string, 4)
			image = make([]string, 0)
		} else {
			if borders[0] == "" {
				borders[0] = line
			} else {
				image = append(image, line[1:len(line)-1])
			}
			borders[3] += line[:1]
			borders[1] += line[len(line)-1:]
		}
	}

	var corner int
	for id, borders := range tileBorders {
		matchCount := 0
		for _, border := range borders {
			if len(borderTiles[normalize(border)]) == 2 {
				matchCount++
			}
		}
		if matchCount == 2 {
			corner = id
			break
		}
	}

	cornerRot := 0
	cornerBorders := tileBorders[corner]
	for len(borderTiles[normalize(cornerBorders[0])]) != 1 || len(borderTiles[normalize(cornerBorders[3])]) != 1 {
		rotateBorders(&cornerBorders)
		cornerRot++
	}
	tileImages[corner] = transform(tileImages[corner], cornerRot, false)

	targetLeftBorder := cornerBorders[1]
	idOnLeft := corner
	targetUpBorder := cornerBorders[2]
	idOnUp := corner
	side := int(math.Sqrt(float64(len(tileImages))))
	fullImageTiles := make([][]int, side)
	for i := 0; i < side; i++ {
		rowTiles := make([]int, side)
		if i != 0 {
			var idOnDown int
			idsWithTargetBorder := borderTiles[normalize(targetUpBorder)]
			if idsWithTargetBorder[0] != idOnUp {
				idOnDown = idsWithTargetBorder[0]
			} else {
				idOnDown = idsWithTargetBorder[1]
			}
			rot := 0
			borders := tileBorders[idOnDown]
			for normalize(borders[0]) != normalize(targetUpBorder) {
				rotateBorders(&borders)
				rot++
			}
			flipH := borders[0] != targetUpBorder
			if flipH {
				targetUpBorder = reverse(borders[2])
				targetLeftBorder = borders[3]
			} else {
				targetUpBorder = borders[2]
				targetLeftBorder = borders[1]
			}
			idOnUp = idOnDown
			idOnLeft = idOnUp
			rowTiles[0] = idOnUp
			if flipH {
				rot += 2
				rot %= 4
			}
			tileImages[idOnUp] = transform(tileImages[idOnUp], rot, flipH)
		}
		rowTiles[0] = idOnUp
		for j := 1; j < side; j++ {
			var idOnRight int
			idsWithTargetBorder := borderTiles[normalize(targetLeftBorder)]
			if idsWithTargetBorder[0] != idOnLeft {
				idOnRight = idsWithTargetBorder[0]
			} else {
				idOnRight = idsWithTargetBorder[1]
			}
			rot := 0
			borders := tileBorders[idOnRight]
			for normalize(borders[3]) != normalize(targetLeftBorder) {
				rotateBorders(&borders)
				rot++
			}
			flipV := borders[3] != targetLeftBorder
			if flipV {
				targetLeftBorder = reverse(borders[1])
			} else {
				targetLeftBorder = borders[1]
			}
			idOnLeft = idOnRight
			rowTiles[j] = idOnLeft
			tileImages[idOnLeft] = transform(tileImages[idOnLeft], rot, flipV)
		}
		fullImageTiles[i] = rowTiles
	}

	dim := len(tileImages[corner])
	fullDim := side * dim
	fullImage := make([]string, fullDim)

	for i := 0; i < fullDim; i++ {
		fullImage[i] = ""
		iTile := i / dim
		iOff := i % dim
		for j := 0; j < fullDim; j++ {
			jTile := j / dim
			jOff := j % dim
			fullImage[i] += string(tileImages[fullImageTiles[iTile][jTile]][iOff][jOff])
		}
	}

	monster := [3]string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}

	monsterHeight := len(monster)
	monsterWidth := len(monster[0])
	for rot := 0; rot < 4; rot++ {
		for _, flipV := range [2]bool{false, true} {
			fullImage = transform(fullImage, rot, flipV)
			for i := 0; i < fullDim - monsterHeight; i++ {
				for j := 0; j < fullDim - monsterWidth; j++ {
					found := true
					for iM := 0; iM < monsterHeight; iM++ {
						for jM := 0; jM < monsterWidth; jM++ {
							if monster[iM][jM] == '#' {
								if fullImage[i+iM][j+jM] == '.' {
									found = false
									break
								}
							}
						}
						if !found {
							break
						}
					}
					if found {
						for iM := 0; iM < monsterHeight; iM++ {
							for jM := 0; jM < monsterWidth; jM++ {
								if monster[iM][jM] == '#' {
									fullImage[i+iM] = fullImage[i+iM][:j+jM] + "O" + fullImage[i+iM][j+jM+1:]
								}
							}
						}
					}
				}
			}
			fullImage = transform(fullImage, (4-rot)%4, flipV)
		}
	}

	roughness := 0
	for i := 0; i < fullDim; i++ {
		for j:= 0; j < fullDim; j++ {
			if fullImage[i][j] == '#' {
				roughness++
			}
		}
	}
	fmt.Println(roughness)
}
