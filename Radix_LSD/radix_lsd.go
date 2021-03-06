package radix

/* 
func main() {
	var random = make([]int, arrSize)
	
    seed := rand.NewSource(time.Now().UnixNano())
	seededRand := rand.New(seed)
	
	for i := 0; i < arrSize; i++ {
		random[i] = seededRand.Intn(maxNum)

		tm.Print(random[i], strings.Repeat("-", random[i]))
		tm.Flush()

		time.Sleep(time.Duration(speedMult) * 2 * time.Millisecond)
	}

	radixSort(random)
	
	tm.Println("Done!")
	tm.Flush()
}
 */
func countSort(arr []int, exp1, arrSize int) []int {
	var n = arrSize

	var output = make([]int, n)
	var count [10]int

	for i := 0; i < n; i++ {
		index := (arr[i] / exp1)
		count[(index)%10]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	i := n - 1
	for i >= 0 {
		index := (arr[i] / exp1)
		output[count[(index)%10]-1] = arr[i]
		count[index%10]--
		i--
	}

	i = 0
	for i < n {
		arr[i] = output[i]
		i++
	}

	return arr
}

// Sort is a radix sort algorithm adapted from geeksforgeeks python code
func Sort(arr []int, arrSize, _ int) []int {
	max1 := max(arr)

	var out []int
	var exp = 1

	for max1/exp > 0 {

		// move CLI cursor to top left corner to overwrite previously output data
		//tm.MoveCursor(0, 2)

		out = countSort(arr, exp, arrSize)

/* 		for _, num := range out {
			for i := 0; i < tm.Width(); i++ {
				tm.Print(" ")
			}

			tm.Print("\r", num, strings.Repeat("-", num))
			tm.Flush()

		} */
		exp *= 10
	}

	return out
}

func max(v []int) (m int) {
	for i := 1; i < len(v); i++ {
		if v[i] > m {
			m = v[i]
		}
	}
	return
}
