package main

import "fmt"

func R2I(r string) int {

	m := make(map[rune]int)
	// m['I'] = 1
	// m['V'] = 5
	// m['X'] = 10
	// m['L'] = 50
	// m['C'] = 100
	// m['D'] = 500
	// m['M'] = 1000

	m['И'] = 1
	m['В'] = 5
	m['Х'] = 10
	m['Л'] = 50
	m['С'] = 100
	m['Д'] = 500
	m['М'] = 1000

	rs := []rune(r)

	n := 0
	pn := 0
	fmt.Println(len(r))
	fmt.Println(len(rs))
	for i := len(rs) - 1; i >= 0; i-- {
		//fmt.Printf("%T", r[i])
		cn := m[rs[i]]
		if cn < pn {
			n -= cn
		} else {
			n += cn
		}
		pn = cn
	}
	return n
}

func main() {
	fmt.Println("\nРезультат: ", R2I("МСМХСВИИ"))
}
