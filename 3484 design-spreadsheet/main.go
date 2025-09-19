package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Spreadsheet struct {
	cell map[string]int
}

func Constructor(rows int) Spreadsheet {
	return Spreadsheet{cell: make(map[string]int)}
}

func (this *Spreadsheet) SetCell(cell string, value int) {
	this.cell[cell] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	this.cell[cell] = 0
}

func (this *Spreadsheet) GetValue(formula string) int {
	var num2 int
	str, _ := strings.CutPrefix(formula, "=")
	newStr := strings.Split(str, "+")
	num1, err := strconv.Atoi(newStr[0])
	if err != nil {
		num1 = this.cell[newStr[0]]
	}
	num2, err = strconv.Atoi(newStr[1])
	if err != nil {
		num2 = this.cell[newStr[1]]
	}
	return num1 + num2
}

func main() {
	obj := Constructor(0)
	obj.SetCell("A1", 10)
	fmt.Println(obj.GetValue("=A1+6"))
}

/**
 * Your Spreadsheet object will be instantiated and called as such:
 * obj := Constructor(rows);
 * obj.SetCell(cell,value);
 * obj.ResetCell(cell);
 * param_3 := obj.GetValue(formula);
 */
