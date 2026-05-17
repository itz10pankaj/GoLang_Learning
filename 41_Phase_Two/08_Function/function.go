package main

// Range is basically used for iterating over Data Structure
// func add (a int,b int) int {
//     return a+b
// }

// func getLangs() (string,string,string,int){
//     return "C++","JS","Java",1
// }

func processIt(fn func(a int) int) {
	fn(1)
}
func main() {
	// fmt.Print(add(2,3))
	// fmt.Print(getLangs())
	// lang1,lang2,lang3,value := getLangs()
	// fmt.Println(lang1)
	// fmt.Println(lang2)
	// fmt.Println(lang3)
	// fmt.Println(value)
	fn := func(a int) int {
		return a + 1
	}
	processIt(fn)

}


// Varidic Funtion 
package main
import "fmt"
// Variadic funtions are thoose funtion in which you can pass n number of funtions like fmt.Println

func sum (nums ...int)int{
    total:=0;
    for _,num := range(nums){
        total+=num;
    }
    return total;
}
func main() {
  fmt.Println("Start small. Ship something.")
  nums:=[]int{1,2,3,4}
  fmt.Println(sum(nums...))
}