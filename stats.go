package main

func Mean(data []float64) float64 {
	var sum float64 = 0

	for _, num := range data {
		sum += num
	}
	return sum / float64(len(data))
}

func BubbleSort(data []float64) []float64 {
	// sort the numbers using BubbleSort
	for i := 0; i < (len(data))-1; i++ {
		for j := 0; j < (len(data))-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}

func Median(data []float64) float64 {
	data = BubbleSort(data)
	n := len(data)

	if n%2 == 0 {
		middle1 := n/2 - 1
		middle2 := n / 2
		return (data[middle1] + data[middle2]) / 2
	} else {
		middle := n / 2
		return data[middle]
	}
}

func Variance(data []float64) float64 {
	miu := Mean(data)
	var sqrs float64 = 0
	for _, value := range data {
		sqrs += (value - miu) * (value - miu)
	}
	return sqrs / float64(len(data)-1)
}

func StDev(variance float64) float64 {
	var wholePart, decimalPart float64

	for i := 0; variance >= 0; i++ {
		wholePart = float64(i)
		variance -= 2*wholePart + 1
	}

	decimalPart = (variance + 2*wholePart + 1) / (2*wholePart + 1)

	exactSqrt := wholePart + decimalPart

	return exactSqrt
}
