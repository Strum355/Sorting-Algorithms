package utils

import (
	"math/rand"
	"time"
)

func MakeArray(arrSize, maxNum int) []int {
    var a = make([]int, arrSize)
	
    seed := rand.NewSource(time.Now().UnixNano())
	seededRand := rand.New(seed)
	
	for i := 0; i < arrSize; i++ {
		a[i] = seededRand.Intn(maxNum)
	}

	return a
}

func ToBool(b int) bool {
    if b == 1 {
        return true
    }
    
    return false
 }