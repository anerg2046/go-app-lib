package str

import (
	"math"
	"strings"
)

// 计算两个字符串的余弦近似度
func CosineSimilarity(a, b string) float64 {
	var (
		aWords = strings.Split(a, "")
		bWords = strings.Split(b, "")
		words  = append(aWords, bWords...)
		dict   = make(map[string]int)
		aVec   = make([]int, len(words))
		bVec   = make([]int, len(words))
	)

	for i, w := range words {
		if _, ok := dict[w]; !ok {
			dict[w] = len(dict)
		}

		if i < len(aWords) {
			aVec[dict[w]]++
		} else {
			bVec[dict[w]]++
		}
	}

	var (
		dotProduct = 0.0
		aNorm      = 0.0
		bNorm      = 0.0
	)

	for i := range words {
		dotProduct += float64(aVec[i] * bVec[i])
		aNorm += math.Pow(float64(aVec[i]), 2)
		bNorm += math.Pow(float64(bVec[i]), 2)
	}

	return dotProduct / (math.Sqrt(aNorm) * math.Sqrt(bNorm))
}
