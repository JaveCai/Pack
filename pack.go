package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	LevelError = iota
	LevelWarning
	LevelInformational
	LevelDebug
)

//define maximum size（unit: mm)
const (
	LENGTH int64 = 580
	WIDTH  int64 = 500
	HEIGTH int64 = 500
)

const (
	LENGTH_IDX uint8 = 0
	WIDTH_IDX  uint8 = 1
	HEIGTH_IDX uint8 = 2
)

type ProductPack struct {
	SolutionType uint8
	ProductCount int64
	PackageCount int64

	//Package size
	PackLength int64
	PackWeigth int64
	PackHetght int64

	//Box Size
	BoxSides SideLengths
}

func main() {

	for {
		fmt.Println("Please enter the size of your product:")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		args := strings.Split(input.Text(), " ")
		GetPackSolution(args)
	}

}

/*
	Algorithm v1.0:
	PACKCOUNT: {12,8,6,3}
	n: the Maximum product count theoretically.
	1.get the n wich can be div by PACKCOUNT
	2.Find out the longest/shortest side of product and pack them to the shape most familiar to a cube
	3.Find out the longest side of the package, which will match to the longest side of the box(LENGTH)
	4.The other two sides of the box is equal to each other, therefore we don't care the other two side of package

	Algorithm v2.0:
	1.input: in[3]={x1,y1,z1}
	2.get the max box sides < 580: MaxBoxSide580[3]={X58,Y58,Z58}
	3.get the max box sides < 500: MaxBoxSide500[3]={X50,Y50,Z50}
	4.get 3 group of MaxBoxSides test[3][3]={
			{MaxBoxSides580[0], MaxBoxSides500[1], MaxBoxSides500[2]},
			{MaxBoxSides580[1], MaxBoxSides500[1], MaxBoxSides500[2]},
			{MaxBoxSides580[2], MaxBoxSides500[1], MaxBoxSides500[2]},}
	5.check every group
		CheckNumberCanBePacked() == true? if yes then save the best one and return
	  									  if no then goto next step
	6.recursive call GetSolution enter next level which cut one side a bit to get solution and check as previous step
				GetSolution(ll-in[maxIdx[i][0]], ww, hh, i)
				GetSolution(ll, ww-in[maxIdx[i][1]], hh, i)
				GetSolution(ll, ww, hh-in[maxIdx[i][2]], i)
*/

func CheckInputSizeValid(l, w, h int64) bool {
	if l > LENGTH || w > WIDTH || h > HEIGTH {
		return false
	}
	return true
}

func CheckInputSizeBeyondHalfOfMaximum(l, w, h int64) bool {
	if l > LENGTH/2 && w > WIDTH/2 && h > HEIGTH/2 {
		return true
	}
	return false
}

type SideLengths [3]int64

func (s *SideLengths) Init(L, W, H int64) {
	s[0] = L
	s[1] = W
	s[2] = H
}

func (s *SideLengths) GetMaxSideLengths(l, w, h, maximum int64) {
	s[0] = GetMaxSideLength(l, maximum)
	s[1] = GetMaxSideLength(w, maximum)
	s[2] = GetMaxSideLength(h, maximum)
}

func (s *SideLengths) GetVolume() int64 {
	return s[0] * s[1] * s[2]
}

func GetMaxSideLength(l, max int64) int64 {
	var i int64 = 1
	var L int64 = 0
	for l*i <= max {
		L = l * i
		i += 1
	}
	return L
}

func GetVolume(l, w, h int64) int64 {
	return l * w * h
}

func GetPackSolution(args []string) {

	if len(args[:]) != 3 {
		fmt.Printf("Please enter 3 parameters.\n")
		return
	}

	//convert data type from string to float64
	x0, _ := strconv.ParseFloat(args[0], 64)
	y0, _ := strconv.ParseFloat(args[1], 64)
	z0, _ := strconv.ParseFloat(args[2], 64)

	if x0 == 0 || y0 == 0 || z0 == 0 {
		fmt.Printf("error :Input parameters should be 3 and greater than 0 : %f, %f, %f\n", x0, y0, z0)
		return
	}
	//convert cm to mm
	x1 := int64(x0 * 10)
	y1 := int64(y0 * 10)
	z1 := int64(z0 * 10)

	fmt.Printf("%dmm %dmm %dmm\n", x1, y1, z1)
	if CheckInputSizeValid(x1, y1, z1) == false {
		fmt.Printf("error: Input size beyond the maximum!\n")
		return
	}

	if CheckInputSizeBeyondHalfOfMaximum(x1, y1, z1) == true {
		fmt.Printf("The size of product beyond the half maximum, you can only pack one product in one box!\n")
		return
	}

	solution := GetPackSolutionImp(x1, y1, z1)
	l, w, h := solution.BoxSides[0], solution.BoxSides[1], solution.BoxSides[2]

	fmt.Printf("SolutionType\t: %d \n", solution.SolutionType)
	fmt.Printf("ProductCount\t: %d \n", solution.ProductCount)
	fmt.Printf("Box Size\t: %.1f  %.1f  %.1f\n", float64(l)/10, float64(w)/10, float64(h)/10)
	fmt.Printf("Box Volume\t: %.3f\n", float64(l)/1000*float64(w)/1000*float64(h)/1000)
}

func CheckNumberCanBePacked(n int64) (bool, uint8) {
	if n%12 == 0 {
		return true, 12
	} else if n%8 == 0 {
		return true, 8
	} else if n%6 == 0 {
		return true, 6
	} else if n%3 == 0 {
		return true, 3
	} else {
		return false, 0
	}
}

//for testing
func GetPackSolutionImp(l, w, h int64) (solution ProductPack) {

	/*----------------------------------------------------------------------*/

	fileName := "pack.log"
	logFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Fail to create the log file!")
	}
	defer logFile.Close()
	debugLog := log.New(logFile, "[D]", log.Lshortfile /*log.LstdFlags*/)
	/*----------------------------------------------------------------------*/

	var in SideLengths
	in.Init(l, w, h)

	var MaxBoxSides580 SideLengths
	MaxBoxSides580.GetMaxSideLengths(l, w, h, 580)

	var MaxBoxSides500 SideLengths
	MaxBoxSides500.GetMaxSideLengths(l, w, h, 500)

	maxIdx := [3][3]int{
		{0, 1, 2},
		{1, 2, 0},
		{2, 0, 1},
	}

	maxInput := [3][3]int64{
		{MaxBoxSides580[0], MaxBoxSides500[1], MaxBoxSides500[2]},
		{MaxBoxSides580[1], MaxBoxSides500[2], MaxBoxSides500[0]},
		{MaxBoxSides580[2], MaxBoxSides500[0], MaxBoxSides500[1]},
	}

	var bHaveSolution bool = false

	/*----------- start define function GetSolution ---------------------*/
	var depth int = 0
	var GetSolution func(ll, ww, hh int64, i int)

	GetSolution = func(ll, ww, hh int64, i int) {
		if ll <= 0 || ww <= 0 || hh <= 0 {
			return
		}
		depth += 1
		debugLog.Printf("[GetSolution] depth: %d\n", depth)

		n := ll * ww * hh / in.GetVolume()
		can, SolutionType := CheckNumberCanBePacked(n)
		if (can == true && n > solution.ProductCount) || (can == true && n == solution.ProductCount && SolutionType > solution.SolutionType) {
			solution.SolutionType = SolutionType
			solution.ProductCount = n
			solution.BoxSides.Init(ll, ww, hh)
			bHaveSolution = true
			debugLog.Printf("[GetSolution] 2:   %d %d %d | sol: %d\n", ll, ww, hh, SolutionType)
		} else if bHaveSolution == false { /*recursive call: cut one side and continue to next level */

			debugLog.Printf("[GetSolution] 3.1: %d - %d, %d, %d\n", ll, in[maxIdx[i][0]], ww, hh)
			GetSolution(ll-in[maxIdx[i][0]], ww, hh, i)
			debugLog.Printf("[GetSolution] 3.2: %d, %d - %d, %d\n", ll, ww, in[maxIdx[i][1]], hh)
			GetSolution(ll, ww-in[maxIdx[i][1]], hh, i)
			debugLog.Printf("[GetSolution] 3.3: %d, %d, %d - %d\n", ll, ww, hh, in[maxIdx[i][2]])
			GetSolution(ll, ww, hh-in[maxIdx[i][2]], i)
		} else {

			//1.no solution
			//2.have solution but not the best
			debugLog.Printf("[GetSolution] 4: nothing to do\n")

		}
		depth -= 1
	}

	/*----------- end define function GetSolution ---------------------*/

	for i, mi := range maxInput {
		debugLog.Printf("[GetSolution] data group input: %d, %d, %d\n", mi[0], mi[1], mi[2])
		GetSolution(mi[0], mi[1], mi[2], i)
		//bHaveSolution = false
	}

	return solution
}

/*

func (s *SideLengths) GetShortestSideLengthIndex() uint8 {
	if s[0] <= s[1] && s[0] <= s[2] {
		return 0
	}

	if s[1] <= s[2] && s[1] <= s[0] {
		return 1
	}

	if s[2] <= s[0] && s[2] <= s[1] {
		return 2
	}
	return 0
}

func (s *SideLengths) GetLongestSideLengthIndex() uint8 {
	if s[0] >= s[1] && s[0] >= s[2] {
		return 0
	}

	if s[1] >= s[2] && s[1] >= s[0] {
		return 1
	}

	if s[2] >= s[0] && s[2] >= s[1] {
		return 2
	}
	return 0

}

func GetMinSortIndex(x, y, z int64) (uint8, uint8, uint8) {
	if x <= y && x <= z {
		if y < z {
			return 0, 1, 2
		} else {
			return 0, 2, 1
		}

	}
	if y <= z && y <= x {
		if z < x {
			return 1, 0, 2
		} else {
			return 1, 2, 0
		}
	}

	if z <= x && z <= y {
		if x < y {
			return 2, 0, 1
		} else {
			return 2, 1, 0
		}
	}
	return 0, 1, 2
}

func (s *SideLengths) GetMaxSortIndex() (uint8, uint8, uint8) {
	if s[0] >= s[1] && s[0] >= s[2] {
		if s[1] > s[2] {
			return 0, 1, 2
		} else {
			return 0, 2, 1
		}

	}
	if s[1] >= s[2] && s[1] >= s[0] {
		if s[2] > s[0] {
			return 1, 2, 0
		} else {
			return 1, 0, 2
		}
	}

	if s[2] >= s[0] && s[2] >= s[1] {
		if s[0] > s[1] {
			return 2, 0, 1
		} else {
			return 2, 1, 0
		}
	}
	return 0, 1, 2
}
*/
