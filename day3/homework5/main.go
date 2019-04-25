package main

import "fmt"

func main() {
	fmt.Println(largeNumberSum("123456789999999", "123"))
}

func largeNumberSum(n, m string) string {
	var (
		nLen          = len(n)
		mLen          = len(m)
		maxLen        = getMax(nLen, mLen)
		flag          = false
		carryBit byte = 0
		result   []byte
	)

	// 翻转字符传
	n = reverse(n)
	m = reverse(m)

	// 较短的字符传补0
	if nLen < mLen {
		for i := nLen; i < mLen; i++ {
			n += "0"
		}
	} else {
		for i := mLen; i < nLen; i++ {
			m += "0"
		}
	}

	for i := 0; i < maxLen; i++ {
		nSum := (n[i] - '0') + (m[i] - '0') + carryBit
		if nSum > 9 {
			// 最后一位超过10，进1
			if i == maxLen-1 {
				flag = true
			}

			carryBit = 1
			result = append(result, nSum-10+'0')
		} else {
			carryBit = 0
			result = append(result, nSum+'0')
		}
	}

	if flag {
		result = append(result, carryBit+'0')
	}

	return reverse(string(result))
}

func reverse(str string) string {
	var result []byte
	tmp := []byte(str)
	temLen := len(tmp)

	for i := 0; i < temLen; i++ {
		result = append(result, tmp[temLen-1-i])
	}

	return string(result)
}

func getMax(n int, m int) int {
	if n > m {
		return n
	}
	return m
}
