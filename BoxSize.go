package main

import (
	"fmt"
	"os"
	"strconv"
)

//define maximum size
const (
	LENGTH uint64 = 500
	WIDTH  uint64 = 500
	HEIGTH uint64 = 580
)

type ProductPack struct {
	SolutionType uint64
	ProductCount uint64
	PackageCount uint64
	//Box Size
	Length uint64
	Weight uint64
	Height uint64
}

func CheckInputSizeValid(a, b, c uint64) bool {
	x, y, z := MaximumFirst(a, b, c)
	if x > HEIGTH || y > LENGTH || z > LENGTH {
		return false
	}
	return true
}

func CheckInputSizeBeyondHalfOfMaximum(a, b, c uint64) bool {
	x, y, z := MaximumFirst(a, b, c)
	if x > HEIGTH/2 && y > LENGTH/2 && z > LENGTH/2 {
		return true
	}
	return false
}

func main() {
	if len(os.Args[1:]) < 3 {
		fmt.Printf("error :need 3 parameters\n")
		return
	}
	x0, _ := strconv.ParseFloat(os.Args[1], 64)
	y0, _ := strconv.ParseFloat(os.Args[2], 64)
	z0, _ := strconv.ParseFloat(os.Args[3], 64)

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

	var x, y, z uint64 = 0, 0, 0
	var i uint64 = 0

	var maxProductCount uint64 = 0
	var minPackageCount uint64 = 0xffff

	n := LENGTH * WIDTH * HEIGTH / (x1 * y1 * z1)

	solutions := make([]*ProductPack, 0)
	var pc3, pc6, pc8, pc12 ProductPack = ProductPack{}, ProductPack{}, ProductPack{}, ProductPack{}

	for {

		if n%12 == 0 {

			//pc12 = n / 12
			zz, xx, yy := MinimumFirst(x1, y1, z1)

			i = 1
			for i*xx*2 <= LENGTH {

				x = i * xx * 2
				i += 1
			}

			i = 1
			for i*yy*2 <= WIDTH {

				y = i * yy * 2
				i += 1
			}

			i = 1
			for i*zz*3 <= HEIGTH {

				z = i * zz * 3
				i += 1
			}

			pc12.ProductCount = x * y * z / (x1 * y1 * z1)
			if pc12.ProductCount != 0 {
				pc12.Length = x / 10
				pc12.Weight = y / 10
				pc12.Height = z / 10
				pc12.SolutionType = 12
				pc12.PackageCount = pc12.ProductCount / 12
				solutions = append(solutions, &pc12)
				fmt.Println("=============================== Pack 12 in one ===")
				fmt.Printf("Product count  : %d\n", pc12.ProductCount)
				fmt.Printf("Package count  : %d\n", pc12.PackageCount)
				fmt.Printf("Box Size(cm)   : %d %d %d\n", x/10, y/10, z/10)
				fmt.Printf("Box Volume(m^2): %f\n", float64(x)/1000*float64(y)/1000*float64(z)/1000)
			}
		}

		if n%8 == 0 {
			zz, xx, yy := MaximumFirst(x1, y1, z1)

			i = 1
			for i*xx*2 <= LENGTH {

				x = i * xx * 2
				i += 1
			}

			i = 1
			for i*yy*2 <= WIDTH {

				y = i * yy * 2
				i += 1
			}

			i = 1
			for i*zz*2 <= HEIGTH {

				z = i * zz * 2
				i += 1
			}

			pc8.ProductCount = x * y * z / (x1 * y1 * z1)
			if pc8.ProductCount != 0 {
				pc8.SolutionType = 8
				pc8.Length = x / 10
				pc8.Weight = y / 10
				pc8.Height = z / 10
				pc8.PackageCount = pc8.ProductCount / uint64(8)
				solutions = append(solutions, &pc8)
				fmt.Println("=============================== Pack 8 in one ===")
				fmt.Printf("Product count  : %d\n", pc8.ProductCount)
				fmt.Printf("Package count  : %d \n", pc8.PackageCount)
				fmt.Printf("Box Size(cm)   : %d %d %d\n", x/10, y/10, z/10)
				fmt.Printf("Box Volume(m^2): %f\n", float64(x)/1000*float64(y)/1000*float64(z)/1000)
			}
		}

		if n%6 == 0 {
			zz, xx, yy := MinimumFirst(x1, y1, z1)

			i = 1
			for i*xx <= LENGTH {

				x = i * xx
				i += 1
			}

			i = 1
			for i*yy*2 <= WIDTH {

				y = i * yy * 2
				i += 1
			}

			i = 1
			for i*zz*3 <= HEIGTH {

				z = i * zz * 3
				i += 1
			}

			pc6.ProductCount = x * y * z / (x1 * y1 * z1)
			if pc6.ProductCount != 0 {
				pc6.SolutionType = 6
				pc6.Length = x / 10
				pc6.Weight = y / 10
				pc6.Height = z / 10
				pc6.PackageCount = pc6.ProductCount / 6
				solutions = append(solutions, &pc6)
				fmt.Println("=============================== Pack 6 in one ===")
				fmt.Printf("Product count  : %d\n", pc6.ProductCount)
				fmt.Printf("Package count  : %d \n", pc6.PackageCount)
				fmt.Printf("Box Size(cm)   : %d %d %d\n", x/10, y/10, z/10)
				fmt.Printf("Box Volume(m^2): %f\n", float64(x)/1000*float64(y)/1000*float64(z)/1000)
			}
		}
		//fmt.Println("111", n)
		if n%3 == 0 {
			//fmt.Println("222")
			zz, xx, yy := MinimumFirst(x1, y1, z1)
			//fmt.Println("333", zz)
			i = 1
			for i*xx <= LENGTH {

				x = i * xx
				i += 1
			}

			i = 1
			for i*yy <= WIDTH {

				y = i * yy
				i += 1
			}

			i = 1
			for i*zz*3 <= HEIGTH {

				z = i * zz * 3
				i += 1
			}
			//fmt.Println("444", zz)

			pc3.ProductCount = x * y * z / (x1 * y1 * z1)
			//fmt.Println("555", x * y * z)
			//fmt.Println("666",(x1 * y1 * z1))
			if pc3.ProductCount != 0 {
				pc3.SolutionType = 3
				pc3.Length = x / 10
				pc3.Weight = y / 10
				pc3.Height = z / 10
				pc3.PackageCount = pc3.ProductCount / 3
				solutions = append(solutions, &pc3)
				fmt.Println("=============================== Pack 3 in one ===")
				fmt.Printf("Product Count  : %d\n", pc3.ProductCount)
				fmt.Printf("Packages count : %d \n", pc3.PackageCount)
				fmt.Printf("Box Size(cm)   : %d %d %d\n", x/10, y/10, z/10)
				fmt.Printf("Box Volume(m^2): %f\n", float64(x)/1000*float64(y)/1000*float64(z)/1000)
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
			fmt.Println("=============================== Best Solution ===>")
			for _, s := range solutions {
				if s.ProductCount != 0 {
					fmt.Printf("Solution Type  : %d products in one package\n", s.SolutionType)
					fmt.Printf("Product Count  : %d\n", s.ProductCount)
					fmt.Printf("Packages count : %d \n", s.PackageCount)
					fmt.Printf("Box Size(cm)   : %d %d %d\n", s.Length, s.Weight, s.Height)
					fmt.Printf("Box Volume(m^2): %f\n", float64(s.Length)/100*float64(s.Weight)/100*float64(s.Height)/100)
				}
			}
			break
		} else if n > 0 {
			n -= 1
		} else {
			fmt.Println("no solution.")
			return
		}

	}

}

func MinimumFirst(x, y, z uint64) (uint64, uint64, uint64) {
	if x <= y && x <= z {
		return x, y, z

	}

	if y <= z && y <= x {
		return y, x, z
	}

	if z <= x && z <= y {
		return z, x, y
	}
	return x, y, z
}

func MaximumFirst(x, y, z uint64) (uint64, uint64, uint64) {
	if x >= y && x >= z {
		return x, y, z

	}

	if y >= z && y >= x {
		return y, x, z
	}

	if z >= x && z >= y {
		return z, x, y
	}
	return x, y, z
}
