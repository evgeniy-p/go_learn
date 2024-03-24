package main

import "fmt"

func appendSliceByPointer(sl *[]string) {
	fmt.Printf("\t PP %v [%d/%d]%p %p\n", sl, len(*sl), cap(*sl), *sl, sl)
	*sl = append(*sl, "four")

	fmt.Printf("\t PP %v [%d/%d]%p %p\n", sl, len(*sl), cap(*sl), *sl, sl)
	*sl = append(*sl, "five")

	fmt.Printf("\t PP %v [%d/%d]%p %p\n", sl, len(*sl), cap(*sl), *sl, sl)
	*sl = append(*sl, "six")

}

func appendSliceByVal(sl []string) {
	fmt.Printf("\t VV %v [%d/%d]%p %p\n", sl, len(sl), cap(sl), sl, &sl)
	sl = append(sl, "four")

	fmt.Printf("\t VV %v [%d/%d]%p %p\n", sl, len(sl), cap(sl), sl, &sl)
	sl = append(sl, "five")

	fmt.Printf("\t VV %v [%d/%d]%p %p\n", sl, len(sl), cap(sl), sl, &sl)
	sl = append(sl, "six")
}

// -----

func main() {
	sl := make([]string, 0, 4)
	sl = append(sl, "one", "two", "three")

	fmt.Printf("IN  ByPointer %v [%d/%d]%p %p\n", sl, len(sl), cap(sl), sl, &sl)
	appendSliceByPointer(&sl)
	fmt.Printf("OUT ByPointer %v [%d/%d]%p %p\n", sl, len(sl), cap(sl), sl, &sl)

	// ------
	fmt.Println()

	sl = make([]string, 0, 4)
	sl = append(sl, "one", "two", "three")

	fmt.Printf("IN ByVal %v [%d/%d]%p %p\n", sl, len(sl), cap(sl), sl, &sl)
	appendSliceByVal(sl)
	fmt.Printf("OUT ByVal %v [%d/%d]%p %p\n", sl, len(sl), cap(sl), sl, &sl)
}
