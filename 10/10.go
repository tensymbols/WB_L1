package main

import (
	"fmt"
	"sort"
)

type temperatures []float64

func main() {

	tempsByInterval := map[int64]temperatures{}

	tempData := temperatures{
		-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5,
	}
	sort.Float64s(tempData)

	currInterval := int64(tempData[0]) / 10 * 10
	var currVals temperatures
	for _, v := range tempData {
		thisInterval := int64(v) / 10 * 10
		if thisInterval == currInterval {
			currVals = append(currVals, v)
		} else {
			tempsByInterval[currInterval] = currVals
			currVals = nil
			currVals = append(currVals, v)
			currInterval = thisInterval
		}
	}
	tempsByInterval[currInterval] = currVals

	var allIntervals []int64

	for k := range tempsByInterval {
		allIntervals = append(allIntervals, k)
	}
	sort.Slice(allIntervals, func(i, j int) bool {
		return allIntervals[i] < allIntervals[j]
	})
	for _, v := range allIntervals {
		fmt.Println(v, tempsByInterval[v], "\n------")
	}

}
