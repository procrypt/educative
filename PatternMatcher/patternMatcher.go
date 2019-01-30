package main

import (
	"fmt"
	"strings"
)

type counts struct {
	x int
	y int
}

func PatterMatcher(pattern string, str string) []string {
	if len(pattern) > len(str) {
		return []string{}
	}
	pattern, switched := getNewPattern(pattern)
	count, firstY := getCountAndFirstYPos(pattern)
	if count.y !=  0 {
		for lenOfX := 1; lenOfX < len(pattern); lenOfX++ {
			totalLenOfY := len(str) - lenOfX*count.x
			if len(str) <= 0 || totalLenOfY%count.y != 0 {
				continue
			}
			lenOfY := totalLenOfY/count.y
			yIdx := firstY * lenOfX
			x, y := str[:lenOfX], str[yIdx:yIdx+lenOfY]
			potentialMatch := doReplace(pattern,x,y,count)
			if str == potentialMatch {
				if !switched {
					return []string{x,y}
				} else {
					return []string{y,x}
				}
			}
		}
	} else {
		if len(str)%count.x == 0 {
			lenOfX := len(str) / count.x
			x := str[:lenOfX]
			potentialMatch := strings.Repeat(x, len(pattern))
			if str == potentialMatch {
				if !switched {
					return []string{x, ""}
				} else {
					return []string{"", x}
				}
			}
		}
	}
	return []string{}
}

func getNewPattern(pattern string) (string, bool)  {
	if pattern[0] == 'x' {
		return pattern, false
	}
	runes := make([]rune, len(pattern))
	for i := range pattern {
		if pattern[i] == 'x' {
			runes[i] = 'y'
		} else {
			runes[i] = 'x'
		}
	}
	return string(runes), true
}

func getCountAndFirstYPos(pattern string) (counts, int)  {
	count := counts{}
	firstY := strings.Index(pattern, "y")
	for _, r := range pattern {
		if r == 'x' {
			count.x++
		} else {
				count.y++
		}
	}
	return count, firstY
}

func doReplace(pattern,x,y string, count counts) string {
	result := make([]byte, 0)
	for _, r := range pattern {
		if r == 'x' {
			result = append(result, []byte(x)...)
		} else {
			result = append(result, []byte(y)...)
		}

	}
	return string(result)
}

func main() {
	fmt.Println(PatterMatcher("xyyx","dogcatcatdog"))
}
