package Gravity

import (
    "sync"
    u "../Utils"
)

/* func CallGravity(arrSize, maxNum int) {
    a := utils.MakeArray(arrSize, maxNum)
    GravitySortFastest(a)

    a = utils.MakeArray(arrSize, maxNum)
	GravitySortSlow(a, arrSize, maxNum)
	
	a = utils.MakeArray(arrSize, maxNum)
    GravitySortFast(a, arrSize, maxNum) 
} */

func maxVal(v []int) (m int) {
	for i := 1; i < len(v); i++ {
		if v[i] > m {
			m = v[i]
		}
	}
	return
}

// SortSlow is a naive approach to Gravity Sort, shamefully coded by me
func SortSlow(arr []int, arrSize, maxNum int) []int {
    var twoD = make([][]int, arrSize)
    for i := range twoD {
        twoD[i] = make([]int, maxNum)
    }
	for i := range twoD {
		for j := 0; j < arr[i]; j++ {
			twoD[i][j] = 1
		}
	}

	for outIndex := arrSize-1; outIndex >= 0; outIndex-- {
		for inIndex := maxNum-1; inIndex >= 0; inIndex-- {
			go func(i, j int){
				if twoD[i][j] == 0 {
					for searchIndex := i - 1; searchIndex >= 0; searchIndex-- {
						if twoD[searchIndex][j] == 1 {
							twoD[searchIndex][j] = 0
							twoD[i][j] = 1
							break
						}
					}
				}
			}(outIndex, inIndex)
		}
    }
    
    return func() []int {
        var out []int
        for _, num := range twoD {
            var count int 
            for _, num1 := range num {
                if num1 == 1 {
                    count++
                }
            }
            out = append(out, count)
        }
        return out
    }()
}

// SortFast is the version from rosettacodes
func SortFast(arr []int, arrSize, maxNum int) []int {
    const bead = 'o'
    
    all := make([]byte, maxNum*len(arr))

    abacus := make([][]byte, maxNum)
    for pole, space := 0, all; pole < maxNum; pole++ {
        abacus[pole] = space[:len(arr)]
        space = space[len(arr):]
    }

    var wg sync.WaitGroup
    wg.Add(len(arr))

    for row, n := range arr {
        go func(row, n int) {
            for pole := 0; pole < n; pole++ {
                abacus[pole][row] = bead
            }
            wg.Done()
        }(row, n)
    }
    wg.Wait()

    wg.Add(maxNum)
    for _, pole := range abacus {
        go func(pole []byte) {
            top := 0
            for row, space := range pole {
                if space == bead {
                    pole[row] = 0
                    pole[top] = bead
                    top++
                }
            }
            wg.Done()
        }(pole)
    }
    wg.Wait()

    for row := range arr {
        x := 0
        for pole := 0; pole < maxNum && abacus[pole][row] == bead; pole++ {
            x++
        }
        arr[len(arr)-1-row] = x
    }

    return arr
}

// SortFastest is the method adapted from geeksforgeeks C++ version
func SortFastest(arr []int, _, _ int) []int {
    max := maxVal(arr)
    len := len(arr)
    
    beads := make([]int, max*len)

    for i := 0; i < len; i++ {
        for j := 0; j < arr[i]; j++ {
            beads[i * max + j] = 1
        }
    }
    
    for j := 0; j < max; j++ {
        sum := 0
        
        for i := 0; i < len; i++ {
            sum += beads[i * max + j]
            beads[i*max+j] = 0
        }
        
        for i := len - sum; i < len; i++ {
            beads[i * max + j] = 1
        }
    }
    
    for i := 0; i < len; i++ {
        var j int
        for j = 0; j < max && u.ToBool(beads[i * max + j]); j++ {
            arr[i] = j
        }
    }

    return arr
}
 