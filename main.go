package main
 import "fmt"

 func main(){
 var num = []int {0,1,2,3,4,5,6,7,8,9,10}
 for _,number := range num{
	 if number%2==0{
		 fmt.Println("%v is even\n")

	 }else{
	 fmt.Println("%v is odd\n")
 }
}
 }