package main

import (
	"RedDocMD/decision_tree/ida3"
	"RedDocMD/decision_tree/parser"
	"fmt"
)

func main() {
	filename := "/home/deep/work/go/decision-tree/data/data1_19.csv"
	inputData, err := parser.ParseFile(filename)
	if err != nil {
		fmt.Println("Wrong filename")
	} else {
		decisionTree := ida3.IDA3(inputData)
		dotFileName := "./dot/graph1.svg"
		decisionTree.ToGraphvizSVG(dotFileName)
		fmt.Println(decisionTree)
	}
}
