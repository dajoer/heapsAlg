package heapsAlg

import (
	"fmt"
	"flag"
	"strings"
)

func generate(n int, a []string, o *[][]string) {
	if n == 1 {
		// Rückgabe
		*o = append(*o, a)
	} else {
		for i := 0; i < n-1; i++ {
			generate(n-1, a, o)
			if n%2 == 0 {
				a[i], a[n-1] = a[n-1], a[i]
			} else {
				a[0], a[n-1] = a[n-1], a[0]
			}
		}
		generate(n-1, a, o)
	}
}

func HeapsAlg(words []string) ([][]string) {
	var tmp [][]string
	generate(len(words), words, &tmp)
	return tmp
}

func main() {
	flag.Parse()
	for _,l := range HeapsAlg(flag.Args()) {
		fmt.Println(strings.Join(l, " "))
	}
}
