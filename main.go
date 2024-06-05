package main

import "fmt"

func main() {
	fmt.Printf("Problem 1, Match String Weight, hasil pengujian:\n")
	s := "aabbddccc"
	q := []int{2, 4, 8, 9}
	fmt.Printf("Hasil cek `%v` dengan %v adalah %v\n", s, q, matchStringWeight(s, q))
	s = "hello"
	q = []int{8, 5, 20, 15}
	fmt.Printf("Hasil cek `%v` dengan %v adalah %v\n", s, q, matchStringWeight(s, q))
	s = "wooorldd"
	q = []int{23, 15, 12, 15, 8}
	fmt.Printf("Hasil cek `%v` dengan %v adalah %v\n", s, q, matchStringWeight(s, q))
	s = "abddcc"
	q = []int{1, 4, 8, 7}
	fmt.Printf("Hasil cek `%v` dengan %v adalah %v\n", s, q, matchStringWeight(s, q))
	s = "abddcc"
	q = []int{1, 4, 8, 6}
	fmt.Printf("Hasil cek `%v` dengan %v adalah %v\n", s, q, matchStringWeight(s, q))

	fmt.Printf("\nProblem 2, Bracket Balance, hasil pengujian:\n")
	s = "{ [ ( ) ] [ { } ] }"
	fmt.Printf("Hasil cek `%v` adalah %v\n", s, balanceBracket(s))
	s = "{ [ ( { } { } ] ) }"
	fmt.Printf("Hasil cek `%v` adalah %v\n", s, balanceBracket(s))
	s = "{ ( ( [ ] ) [ ] ) [ ] ]"
	fmt.Printf("Hasil cek `%v` adalah %v\n", s, balanceBracket(s))
	s = "{ ( ( [ ] ) [ ] ) [ ] }"
	fmt.Printf("Hasil cek `%v` adalah %v\n", s, balanceBracket(s))
	s = "{ [ [ [ ] [ ] ] { } ] }"
	fmt.Printf("Hasil cek `%v` adalah %v\n", s, balanceBracket(s))

	fmt.Printf("\nProblem 3, Palindrom Search, hasil pengujian:\n")
	s = "3943"
	k := 1
	fmt.Printf("Hasil cek `%v` dengan k=%v adalah %v\n", s, k, palindromSearch(s, k))
	s = "39221243"
	k = 1
	fmt.Printf("Hasil cek `%v` dengan k=%v adalah %v\n", s, k, palindromSearch(s, k))
	k = 2
	fmt.Printf("Hasil cek `%v` dengan k=%v adalah %v\n", s, k, palindromSearch(s, k))
	k = 5
	fmt.Printf("Hasil cek `%v` dengan k=%v adalah %v\n", s, k, palindromSearch(s, k))
	s = "59121243"
	k = 5
	fmt.Printf("Hasil cek `%v` dengan k=%v adalah %v\n", s, k, palindromSearch(s, k))
}
