package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//define maximum size（unit: mm)
const (
	LENGTH uint64 = 580
	WIDTH  uint64 = 500
	HEIGTH uint64 = 500
)

const (
	LENGTH_IDX uint8 = 0
	WIDTH_IDX  uint8 = 1
	HEIGTH_IDX uint8 = 2
)

type ProductPack struct {
	SolutionType uint64
	ProductCount uint64
	PackageCount uint64

	//Package size
	PackLength uint64
	PackWeigth uint64
	PackHetght uint64

	//Box Size
	BoxLength uint64
	BoxWeight uint64
	BoxHeight uint64
}

func CheckInputSizeValid(l, w, h uint64) bool {

	if l > LENGTH || w > WIDTH || h > HEIGTH {
		return false
	}
	return true
}

func CheckInputSizeBeyondHalfOfMaximum(l, w, h uint64) bool {

	if l > LENGTH/2 || w > WIDTH/2 || h > HEIGTH/2 {
		return true
	}
	return false
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
	PACKCOUNT: {12,8,6,3}

	Algorithm v1.0:
	n: the Maximum product count theoretically.
	1.get the n wich can be div by 3,6,8,12
	2.Find out the longest/shortest side of product and pack them to the shape most familiar to a cube
	3.Find out the longest side of the package, which will match to the longest side of the box(LENGTH)
	4.The other two sides of the box is equal to each other, therefore we don't care the other two side of package

	Algorithm v2.0:
	0.input: in[3]={x1,y1,z1} --> type Sides [3]uint64
	1.get the shortest side index of in[3]: idx_in_short -->  func(pbs *BoxSides) GetShortestSideIndex()
	2.get the max < 580: MaxBoxSide58[3]={X58,Y58,Z58} -->  func(pbs *BoxSides) GetMaxBoxSides()
	3.get the longest side index of MaxBoxSide58: idx58L, and the rest: idx50W,idx50H -->func(pbs *BoxSides) GetMaxBoxSideIndex()
	4.get the max < 500: MaxBoxSide50[3]={X50,Y50,Z50}
	5.get the MaxBoxSide[3]={MaxBoxSide58[idx58L],MaxBoxSide58[idx50W],MaxBoxSide58[idx50H]}
	6.Vol_MaxBoxSide/Vol_in % PACKCOUNT == 0? output : step 7 -->func(pbs *BoxSides) GetMaxBoxVolumn()
	7.MaxBoxSide[?] -= in[idx_in_short], goto step 7
	  switch idx_in_short
		  case idx58L: MaxBoxSide[0] -= in[idx_in_short]
		  case idx50W: MaxBoxSide[1] -= in[idx_in_short]
		  case idx50H: MaxBoxSide[2] -= in[idx_in_short]
*/

type SideLengths [3]uint64

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

func (s *SideLengths) GetMaxSideLengths(L, W, H, maximum uint64) {

	var i uint64 = 1

	for L <= maximum {
		i += 1
		L = L * i
	}
	s[0] = L

	i = 1
	for W <= maximum {
		i += 1
		W = W * i
	}
	s[1] = W

	i = 1
	for H <= maximum {
		i += 1
		H = H * i
	}
	s[2] = H

}

func (s *SideLengths) GetVolume() uint64 {
	return s[0] * s[1] * s[2]
}

func (s *SideLengths) Init(L, W, H uint64) {
	s[0] = L
	s[1] = W
	s[2] = H
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
	x1 := uint64(x0 * 10)
	y1 := uint64(y0 * 10)
	z1 := uint64(z0 * 10)

	fmt.Printf("%dmm %dmm %dmm\n", x1, y1, z1)
	if CheckInputSizeValid(x1, y1, z1) == false {
		fmt.Printf("error: Input size beyond the maximum!\n")
		return
	}

	if CheckInputSizeBeyondHalfOfMaximum(x1, y1, z1) == true {
		fmt.Printf("The size of product beyond the half maximum, you can only pack one product in one box!\n")
		return
	}

	/*----------------------------------------------------------------------*/

	var in SideLengths
	in.Init(x1, y1, z1)
	idx_in_short := in.GetShortestSideLengthIndex()
	var MaxBoxSide58 SideLengths
	MaxBoxSide58.Init(0, 0, 0)
	MaxBoxSide58.GetMaxSideLengths(x1, y1, z1, 580)
	for i, v := range MaxBoxSide58 {
		fmt.Printf("max[%d]: %d\n", i, v)
	}

	var x, y, z uint64 = 0, 0, 0
	var i uint64 = 0

	var maxProductCount uint64 = 0
	var minPackageCount uint64 = 0xffff

	// Get the Maximum product count theoretically.
	n := LENGTH * WIDTH * HEIGTH / (x1 * y1 * z1)

	solutions := make([]*ProductPack, 0)
	var pc3, pc6, pc8, pc12 ProductPack = ProductPack{3, 0, 0, 0, 0, 0, 0, 0, 0},
		ProductPack{6, 0, 0, 0, 0, 0, 0, 0, 0},
		ProductPack{8, 0, 0, 0, 0, 0, 0, 0, 0},
		ProductPack{12, 0, 0, 0, 0, 0, 0, 0, 0}

	//for {

	/*
		if n%12 == 0 {

			switch GetMinimumIndex(x1, y1, z1) {
			case LENGTH_IDX:
				pc12.SelectPackageSide(x1*3, y1*2, z1*2)
			case WIDTH_IDX:
				pc12.SelectPackageSide(x1*2, y1*3, z1*2)
			case HEIGTH_IDX:
				pc12.SelectPackageSide(x1*2, y1*2, z1*3)
			}

			i = 1
			for i*pc12.PackHetght <= HEIGTH {

				x = i * pc12.PackHetght
				i += 1
			}

			i = 1
			for i*pc12.PackWeigth <= WIDTH {

				y = i * pc12.BoxWeight
				i += 1
			}

			i = 1
			for i*pc12.PackLength <= LENGTH {

				z = i * pc12.PackLength
				i += 1
			}

			pc12.ProductCount = x * y * z / (x1 * y1 * z1)
			if pc12.ProductCount != 0 {
				pc12.BoxLength = x / 10
				pc12.BoxWeight = y / 10
				pc12.BoxHeight = z / 10
				pc12.SolutionType = 12
				pc12.PackageCount = pc12.ProductCount / 12
				solutions = append(solutions, &pc12)

				fmt.Println("=============================== Pack 12 in one pokect ===")
				fmt.Printf("Product count  : %d\n", pc12.ProductCount)
				fmt.Printf("Pokects count  : %d\n", pc12.PackageCount)
				fmt.Printf("Box Size(cm)   : %d %d %d\n", x/10, y/10, z/10)
				fmt.Printf("Box Volume(m^2): %.3f\n", float64(x)/1000*float64(y)/1000*float64(z)/1000)

			}
		}

		if n%8 == 0 {
			pc8.SelectPackageSide(x1*2, y1*2, z1*2)
			i = 1
			for i*pc8.PackLength <= LENGTH {
				x = i * pc8.PackLength
				i += 1
			}

			i = 1
			for i*pc8.PackWeigth <= WIDTH {
				y = i * pc8.PackWeigth
				i += 1
			}

			i = 1
			for i*pc8.PackHetght <= HEIGTH {
				z = i * pc8.PackHetght
				i += 1
			}

			pc8.ProductCount = x * y * z / (x1 * y1 * z1)
			if pc8.ProductCount != 0 {
				pc8.SolutionType = 8
				pc8.BoxLength = x / 10
				pc8.BoxWeight = y / 10
				pc8.BoxHeight = z / 10
				pc8.PackageCount = pc8.ProductCount / uint64(8)
				solutions = append(solutions, &pc8)

				fmt.Println("=============================== Pack 8 in one pokect ===")
				fmt.Printf("Product count  : %d\n", pc8.ProductCount)
				fmt.Printf("Pokects count  : %d \n", pc8.PackageCount)
				fmt.Printf("Box Size(cm)   : %d %d %d\n", x/10, y/10, z/10)
				fmt.Printf("Box Volume(m^2): %.3f\n", float64(x)/1000*float64(y)/1000*float64(z)/1000)

			}
		}

		if n%6 == 0 {
			switch GetMinimumIndex(x1, y1, z1) {
			case LENGTH_IDX:
				pc6.SelectPackageSide(x1*3, y1*2, z1)
			case WIDTH_IDX:
				pc6.SelectPackageSide(x1*2, y1*3, z1)
			case HEIGTH_IDX:
				pc6.SelectPackageSide(x1, y1*2, z1*3)
			}

			i = 1
			for i*pc6.PackLength <= LENGTH {

				x = i * pc6.PackLength
				i += 1
			}

			i = 1
			for i*pc6.PackWeigth <= WIDTH {

				y = i * pc6.PackWeigth
				i += 1
			}

			i = 1
			for i*pc6.PackHetght <= HEIGTH {

				z = i * pc6.PackHetght
				i += 1
			}

			pc6.ProductCount = x * y * z / (x1 * y1 * z1)
			if pc6.ProductCount != 0 {
				pc6.SolutionType = 6
				pc6.BoxLength = x / 10
				pc6.BoxWeight = y / 10
				pc6.BoxHeight = z / 10
				pc6.PackageCount = pc6.ProductCount / 6
				solutions = append(solutions, &pc6)

				fmt.Println("=============================== Pack 6 in one pokect ===")
				fmt.Printf("Product count  : %d\n", pc6.ProductCount)
				fmt.Printf("Pokects count  : %d \n", pc6.PackageCount)
				fmt.Printf("Box Size(cm)   : %d %d %d\n", x/10, y/10, z/10)
				fmt.Printf("Box Volume(m^2): %.3f\n", float64(x)/1000*float64(y)/1000*float64(z)/1000)

			}
		}

		if n%3 == 0 {

			switch GetMinimumIndex(x1, y1, z1) {
			case LENGTH_IDX:
				pc3.SelectPackageSide(x1*3, y1, z1)
			case WIDTH_IDX:
				pc3.SelectPackageSide(x1, y1*3, z1)
			case HEIGTH_IDX:
				pc3.SelectPackageSide(x1, y1, z1*3)
			}

			i = 1
			for i*pc3.PackLength <= LENGTH {

				x = i * pc3.PackLength
				i += 1
			}

			i = 1
			for i*pc3.PackWeigth <= WIDTH {

				y = i * pc3.PackWeigth
				i += 1
			}

			i = 1
			for i*pc3.PackHetght <= HEIGTH {

				z = i * pc3.PackHetght
				i += 1
			}

			pc3.ProductCount = x * y * z / (x1 * y1 * z1)
			if pc3.ProductCount != 0 {
				pc3.SolutionType = 3
				pc3.BoxLength = x / 10
				pc3.BoxWeight = y / 10
				pc3.BoxHeight = z / 10
				pc3.PackageCount = pc3.ProductCount / 3
				solutions = append(solutions, &pc3)

				fmt.Println("=============================== Pack 3 in one pokect ===")
				fmt.Printf("Product Count  : %d\n", pc3.ProductCount)
				fmt.Printf("Pokects count  : %d \n", pc3.PackageCount)
				fmt.Printf("Box Size(cm)   : %d %d %d\n", x/10, y/10, z/10)
				fmt.Printf("Box Volume(m^2): %.3f\n", float64(x)/1000*float64(y)/1000*float64(z)/1000)

			}
		}

		if pc3.ProductCount != 0 || pc6.ProductCount != 0 || pc8.ProductCount != 0 || pc12.ProductCount != 0 {

			// 1.find max product count
			for _, s := range solutions {
				if maxProductCount < s.ProductCount {
					maxProductCount = s.ProductCount
				}

			}
			// 2.keep the solutions pack max product
			for _, s := range solutions {
				if s.ProductCount < maxProductCount {
					s.ProductCount = 0 //not the target

				}
			}

			// 3. find the lease packages count from the max solutions
			for _, s := range solutions {
				if s.ProductCount != 0 && minPackageCount > s.PackageCount {
					minPackageCount = s.PackageCount
				}
			}

			// 4. keep the solutions use least packages
			for _, s := range solutions {
				if s.ProductCount != 0 && s.PackageCount > minPackageCount {
					s.ProductCount = 0 //not the target
				}
			}

			// 5. give the solutions
			fmt.Println("=============================== Best Solution ===========>")
			for _, s := range solutions {
				if s.ProductCount != 0 {
					fmt.Printf("Solution Type  : %d products in one pokects\n", s.SolutionType)
					fmt.Printf("Product Count  : %d\n", s.ProductCount)
					fmt.Printf("Pokects count  : %d \n", s.PackageCount)
					fmt.Printf("Box Size(cm)   : %d %d %d\n", s.BoxLength, s.BoxWeight, s.BoxHeight)
					fmt.Printf("Box Volume(m^2): %.3f\n", float64(s.BoxLength)/100*float64(s.BoxWeight)/100*float64(s.BoxHeight)/100)
				}
			}
			fmt.Println("")
			break
		} else if n > 0 {
			n -= 1
		} else {
			fmt.Println("No solution.")
			return
		}

	}*/
}

func GetMinimumIndex(x, y, z uint64) (uint8, uint8, uint8) {
	if x <= y && x <= z {
		if y < z {
			return 1, 2, 3
		} else {
			return 1, 3, 2
		}

	}
	if y <= z && y <= x {
		if z < x {
			return 2, 1, 3
		} else {
			return 2, 3, 1
		}
	}

	if z <= x && z <= y {
		if x < y {
			return 3, 1, 2
		} else {
			return 3, 2, 1
		}
	}
	return 1, 2, 3
}

func GetMaximumIndex(x, y, z uint64) (uint8, uint8, uint8) {
	if x >= y && x >= z {
		if y > z {
			return 1, 2, 3
		} else {
			return 1, 3, 2
		}

	}
	if y >= z && y >= x {
		if z > x {
			return 2, 3, 1
		} else {
			return 2, 1, 3
		}
	}

	if z >= x && z >= y {
		if x > y {
			return 3, 1, 2
		} else {
			return 3, 2, 1
		}
	}
	return 1, 2, 3
}

/*
function SelectPackageSide: Select the longest side as the Length of the package
*/
/*
func (p *ProductPack) SelectPackageSide(l, w, h uint64) {
	switch GetMaximumIndex(l, w, h) {
	case LENGTH_IDX:
		p.PackLength = l
		p.PackWeigth = w
		p.PackHetght = h
	case WIDTH_IDX:
		p.PackLength = w
		p.PackWeigth = l
		p.PackHetght = h
	case HEIGTH_IDX:
		p.PackLength = h
		p.PackWeigth = l
		p.PackHetght = w
	}
}
*/
