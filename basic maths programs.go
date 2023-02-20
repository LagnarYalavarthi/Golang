/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
/**********PRIME NUMBER*****************/
// package main
// import "fmt"
// func myPrime(n int){
//     for i:=2; i<n; i++ {
//         t := true
//         for j:=2; j<i; j++{
//             if i%j==0 {
//                 t = false
//                 break
//             }
//         }
//         if t{
//             fmt.Println(i)
//         }
        
//     }
// }

// func main(){
//     myPrime(100)
// }

/**********EVEN NUMBER*****************/

// package main
// import "fmt"
// func myEven(n int){
//     for i:=1;i<n;i++{
//         if i%2==0{
//             fmt.Println(i)
//         }
//     }
// }
// func main(){
//     myEven(10)
// }


/**********REVERSE & PALLINDROME OF A STRING*****************/

// package main
// import "fmt"
// func reverseString(str string) (result string){
//     for _,v:=range str{
//         result=string(v)+result
//     }
//     if result==str{
//         fmt.Println(str,"is pallindrome")
//     }else{
//         fmt.Println("its not pallindrome")
//     }
//     return
// }
// func main(){
//     fmt.Println(reverseString("malayalam"))
// }


/**********PALLINDROME OF A NUMBER*****************/

// package main
// import "fmt"
// func pallinNumber(n int) string{
//     temp:=n
//     result:=0
//     var digit int
//     for n>0 {
//         digit=n%10
//         result=(result*10)+digit
//         n=n/10
//     }
//     if(result==temp){
//         return("yes it is pallindrome")
//     }else{
//             return("it is not pallindrome")
//     }
    
// }
// func main(){
//     fmt.Println(pallinNumber(1213))
// }


/**********COUNT OF IN A NUMBER*****************/

package main
import "fmt"
func counting(n int) int{
    count:=0
    for n>0{
        n=n/10
        count=count+1
    }
    return count
}
func main(){
    fmt.Println(counting(1234))
}