/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
// package main
// import "fmt"
// func myfunction(x int,y string) (result int,txt1 string){
//     result =x+x
//     txt1=y+"world"
//     return
// }
// func main(){
//     fmt.Println(myfunction(7,"hello"))
// }

// package main
// import "fmt"
// func myfunction(x int,y string) (result int,txt1 string){
//     result=x+x
//     txt1=y+"world"
//     return
// }
// func main(){
//     a,b:=myfunction(10,"hello")
//     fmt.Println(a,b)
// }

// package main
// import "fmt"
// func myfunction(x int,y string) (result int,txt1 string){
//     result=x+x
//     txt1=y+"world"
//     return
// }
// func main(){
//     a,_:=myfunction(10,"hello")
//     fmt.Println(a)
// }


// package main
// import "fmt"
// func myfunction(x int,y string) (result int,txt1 string){
//     result=x+x
//     txt1=y+"world"
//     return
// }
// func main(){
//     _,b:=myfunction(10,"hello")
//     fmt.Println(b)
// }

// package main
// import "fmt"
// func testcount(x int) int{
//     if x==10{
//         return 0
//     }
//     fmt.Println(x)
//     return testcount(x+1)
    
// }
// func main(){
//     testcount(1)
// }

// package main
// import "fmt"
// func sum(x int) int{
//     if x==0{
//         return 0
//     }else{
//         return x+sum(x-1)
//     }
// }
// func main(){
//     x:=10
//     result:=sum(x)
//     fmt.Println("Total:",result)
// }

// package main
// import "fmt"
// func fact(x int) int{
//     if x<=1{
//         return 1
//     }
//     return x* fact(x-1)
// }
// func main(){
//     x:=5
//     result:=fact(x)
//     fmt.Println("value:",result)
// }

package main
import "fmt"
func odd(x int) int{
    if x<=1{
        return 
    }
    return x*odd(x-1)
}
func main(){
    x:=20
    result:=odd(x)
    fmt.Println(result)
}
