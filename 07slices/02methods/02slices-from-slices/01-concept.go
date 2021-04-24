package main

import "fmt"

func main() {
	bands := []string{"Black Sabbath", "Deep Purple", "Led Zeppelin", "Sepultura", "Kiss", "Metallica"}

	europeanBands := bands[0:3]
	americanBands := bands[3:6]
	from0To3 := bands[:3]
	upTo3 := bands[3:]
	from2To5 := bands[2:5]

	fmt.Println("European Bands = ", europeanBands)
	fmt.Println()
	fmt.Println("American Bands = ", americanBands)
	fmt.Println()
	fmt.Println("From 0 to 3 = ", from0To3)
	fmt.Println()
	fmt.Println("Up to 3 = ", upTo3)
	fmt.Println()
	fmt.Println("From 2 to 5 = ", from2To5)

}
