/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/

/******************VOWELS IN STRING***************************/
// package main
// import "fmt"
// func vowels(s string) (v string){
//     for _,v:=range s{
//         if v=='a'|| v=='e'|| v=='i'||v=='o'|| v=='u'{
//             fmt.Println(string(v))
//         }
//         if v=='A'|| v=='E'|| v=='I'||v=='O'|| v=='U'{
//             fmt.Println(string(v))
        
//         }
//     }
//     return string(v)
// }
// func main(){
//     fmt.Println(vowels("lAgner"))
// }

/*************ARMSTRONG NUMBER********/


// package main
// import "fmt"
// func armstrong(n int){
//     sum:=0
//     temp:=n
//     for temp>0{
//         digit:=temp%10
//         sum+=digit*digit*digit
//         temp=temp/10
//     }
//     if n==sum{
//         fmt.Println("number is armstrong")
//     }else{
//         fmt.Println("number is not armstrong")
//     }
//     return
// }
// func main(){
//     armstrong(156)
//     armstrong(153)
// }


/*************** FIBANOCCI SERIES USING RECURSION*********/

// package main
// import "fmt"
// func fabnocciseries(n int) int{
//     if n<=1{
//         return n
//     }
//     return fabnocciseries(n-1)+fabnocciseries(n-2)
// }
// func main(){
//     for i:=1;i<=5;i++{
//         fmt.Println(fabnocciseries(i))
//     }
    
// }



/************ RECURSION**********/

// package main
// import "fmt"
// func countdown(n int) int {
//     if n>30{
//         fmt.Println(n)
//         return countdown(n-1)
//     }else{
//         return 0
//     }
    
// }
// func main(){
//     fmt.Println(countdown(100))
// }

// package main
// import "fmt"
// func countdown(n int) int{
//     if n<30{
//         return 0
//     }else{
//         fmt.Println(n)
//         return countdown(n-1)
//     }
// }
// func main(){
//     fmt.Println(countdown(100))
// }


/*********** FACTORIAL USING RECURSION**********/

package main
import "fmt"
func factorial(n int) int{
    if n==0||n==1{
        return 1
    }else{
        return n*factorial(n-1)
    }
}
func main(){
    fmt.Println(factorial(5))
}

