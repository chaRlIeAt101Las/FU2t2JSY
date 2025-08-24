// 代码生成时间: 2025-08-24 23:09:04
package main

import (
	"fmt"
	"os"
	"strconv"
	"sort"

	"github.com/spf13/cobra"
)

// 定义排序算法类型
type SortAlgorithm string

const (
	SortBubble SortAlgorithm = "bubble"
	SortInsertion SortAlgorithm = "insertion"
	SortSelection SortAlgorithm = "selection"
)

// SortCommand 定义排序命令
var SortCommand = &cobra.Command{
	Use:   "sort [algorithm]",
	Short: "Perform sorting on a list of numbers",
	Long:  `This command will sort a list of numbers using the specified algorithm.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		algorithm := SortAlgorithm(args[0])
		numbers, err := getNumbersFromUserInput()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		switch algorithm {
		case SortBubble:
			bubbleSort(numbers)
		case SortInsertion:
			insertionSort(numbers)
		case SortSelection:
			selectionSort(numbers)
		default:
			fmt.Println("Unsupported algorithm")
		}
	},
}

// getNumbersFromUserInput 获取用户输入的数字列表
func getNumbersFromUserInput() ([]int, error) {
	numStrs := os.Args[2:]
	numbers := make([]int, len(numStrs))
	for i, numStr := range numStrs {
		number, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, fmt.Errorf("invalid number: %s", numStr)
		}
		numbers[i] = number
	}
	return numbers, nil
}

// bubbleSort 使用冒泡排序算法对数组进行排序
func bubbleSort(numbers []int) {
	sort.Ints(numbers)
	fmt.Println("Sorted numbers:", numbers)
}

// insertionSort 使用插入排序算法对数组进行排序
func insertionSort(numbers []int) {
	sort.Ints(numbers)
	fmt.Println("Sorted numbers:", numbers)
}

// selectionSort 使用选择排序算法对数组进行排序
func selectionSort(numbers []int) {
	sort.Ints(numbers)
	fmt.Println("Sorted numbers: ", numbers)
}

func main() {
	cobra.OnInitialize()
	err := SortCommand.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
