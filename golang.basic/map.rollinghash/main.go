package main

import (
	"fmt"

	dPkg "lec.youtube.tucker/map.rollinghash/dataStruct"
)

func main() {
	fmt.Println("abcde = ", dPkg.Hash("abcde"))
	fmt.Println("abcde = ", dPkg.Hash("abcde"))
	fmt.Println("abcdf = ", dPkg.Hash("abcdf"))
	fmt.Println("tbcde = ", dPkg.Hash("tbcde"))
	fmt.Println("tbcdeasdkjhgaskjhgas = ", dPkg.Hash("tbcdeasdkjhgaskjhgas"))
	fmt.Println()

	m := dPkg.CreateMap()
	m.Add("AAA", "01011111111")
	m.Add("BBB", "01022222222")
	m.Add("CDEFGHIJKLMN", "01011111111")
	m.Add("CCC", "01044444444")

	fmt.Println("AAA = ", m.Get("AAA"))
	fmt.Println("CCC = ", m.Get("CCC"))
	fmt.Println("DDD = ", m.Get("DDD"))
	fmt.Println("CDEFGHIJKLMN = ", m.Get("CDEFGHIJKLMN"))

}
