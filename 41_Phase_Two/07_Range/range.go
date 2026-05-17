package main

import "fmt"

// Range is basically used for iterating over Data Structure
func main() {

// nums := []int{6, 7, 8}
   
// for i := 0; i < len(nums); i++ {
// fmt.Println(nums[i])
// }
    // var sum int = 0
    // for i,num := range(nums){
    //     fmt.Println(num,"index=",i)
    //     sum+=num  
    // }
    // fmt.Println(sum)
   
    var name string = "Pankaj"
    for i,c := range(name) {
        fmt.Println(i,c,string(c))
    }
   
}