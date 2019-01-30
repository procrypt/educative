package main

import "strings"

type counts struct {
	a int
	b int
}

func wordPattern(pattern string, str string) bool {
	if len(pattern) > len(str) {
		return false
	}
	pattern = getNewPattern(pattern)
	count, firstB := getCountAndFirstB(pattern)
	if count.b != 0 {
		for lenOfA:=1; lenOfA < len(str); lenOfA++ {
			totalLenOfB := len(str)- lenOfA*count.a
			if len(str) <= 0 || totalLenOfB%count.b != 0 {
				continue
			}
			lenOfB := totalLenOfB / count.b
			bIdx := firstB * lenOfA
			a,b := str[:lenOfA], str[bIdx:bIdx+lenOfB]
			potentialMatch := doReplace(pattern, a,b, count)
			if str == potentialMatch {
				return true
			} else {
				return false
			}
		}
	} else {
		if len(str)%count.a == 0 {
			lenOfA := len(str)/count.a
			a := str[:lenOfA]
			potentialMatch := strings.Repeat(a, len(pattern))
			if potentialMatch == str {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

func getNewPattern(pattern string) (string) {
	if pattern[0] == 'a' {
		return pattern
	}
	runes := make([]rune, len(pattern))
	for i,char := range pattern {
		if char == 'b' {
			runes[i] = 'a'
		} else {
			runes[i] = 'b'
		}
	}
	return string(runes)
}

func getCountAndFirstB(pattern string) (counts, int) {
	count := counts{}
	firstB := strings.Index(pattern, "b")
	for _, char := range pattern {
		if char == 'x' {
			count.a++
		} else {
			count.b++
		}
	}
	return count, firstB
}

func doReplace(pattern, a, b string, count counts) string {
	result := make([]byte, 0)
	for _, r := range pattern {
		if r == 'a' {
			result = append(result, []byte(a)...)
		} else {
			result = append(result, []byte(b)...)
		}
	}
	return string(result)
}