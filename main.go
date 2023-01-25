package main

import (
	"fmt"

	"catboostcapi/internal"
)

func main() {
	var (
		numericFeatures = [][]float32{
			{72, 41, 70, 4, 45, 95, 13, 34},
			{93, 80, 82, 43, 4, 25, 6, 75},
			{83, 50, 44, 15, 78, 31, 27, 78},
			{79, 72, 47, 82, 16, 63, 34, 29},
			{48, 83, 67, 99, 1, 28, 69, 93},
			{0, 15, 24, 4, 75, 87, 69, 82},
			{86, 51, 89, 20, 38, 5, 59, 78},
			{0, 1, 52, 0, 94, 89, 91, 37},
			{82, 81, 39, 0, 58, 96, 97, 78},
			{4, 14, 54, 48, 19, 90, 29, 92},
		}
		categoryFeatures = [][]string{
			{"улица", "foo"},
			{"аптека", "foo"},
			{"фонарь", "bar"},
			{"улица", "foo"},

			{"фонарь", "bar"},
			{"улица", "foo"},
			{"улица", "bar"},
			{"аптека", "bar"},
			{"фонарь", "foo"},
			{"фонарь", "foo"},
		}
	)

	classifier, err := internal.LoadClassifierFromFile("catboost_model")
	if err != nil {
		fmt.Print("Load error: ", err)
	}

	results, err := classifier.Predict(numericFeatures, len(numericFeatures[0]), categoryFeatures, len(categoryFeatures[0]))
	if err != nil {
		fmt.Print("Predict error: ", err)
	}
	fmt.Print("Result: ", results)
}
