package main
import "fmt"
const NMAX int = 1024
type arrInteger[NMAX]int
func main(){
	var A arrInteger
	var size, i int
	fmt.Scan(&size)
	for i = 0;i <size; i++{
		fmt.Scan(&A[i])
	}
	fmt.Print(cariNilaiMin(A, size))
}
func cariNilaiMin(A arrInteger, size int)int{
	var min, i int
	min = A[0]
	for i = 0; i < size; i++{
		if A[i] < min{
			min = A[i]
		}
	}
	return min
}