/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
/***********************CREATING MAPS****************************/

// package main
// import "fmt"
// func main(){
//     var subjectMarks=map[string]float32{"English":80,"Telugu":85,"Maths":90}
//     fmt.Println(subjectMarks)
// }

// package main
// import "fmt"
// func main(){
//     marks:=map[string]float32{"Maths":70,"Physics":55,"Chemistry":50}
//     fmt.Println(marks)
// }

// package main
// import "fmt"
// func main(){
//     cars:=make(map[string]int)
//     cars["bmw"]=1
//     cars["benz"]=2
//     cars["ford"]=3
//     fmt.Println(cars)
// }


/***********************ACCESSING VALUES IN MAPS***************/

// package main
// import "fmt"
// func main(){
//     flowerColor:=map[string]string{"rose":"red","jasmine":"white","hibiscus":"red"}
//     fmt.Println(flowerColor)
//     //fmt.Println(flowerColor["rose"])
//     //fmt.Println(flowerColor["jasmine"])
//     flowerColor["hibiscus"]="orange"
//     fmt.Println(flowerColor)
//     flowerColor["sunflower"]="yellow"
//     fmt.Println(flowerColor)
//     delete(flowerColor,"hibiscus")
//     fmt.Println(flowerColor)
// }

// package main
// import "fmt"
// func main(){
//     squaredNumber:=map[int]int{2:4,3:9,4:16,5:25,6:36}
//     for number,square:= range squaredNumber{
//         fmt.Printf("square of %d is %d\n", number,square)
//     }
// }

/******************CHECKING KEY IN MAP*************/
// package main
// import "fmt"
// func main(){
//     fruits:=map[string]int{"apples":1,"mangoes":2,"banana":5}
//     value,ok:=fruits["apples"]
//     fmt.Println(value,ok)
// }

/******************COUNT OF ITEMS IN MAP***********/
// package main
// import "fmt"
// func main(){
//     details:=map[string]int{"a":1,"b":2,"c":3,"d":4,"e":78}
//     fmt.Println(len(details))
// }

/************ITERATION OVER MAP****************/

// package main
// import "fmt"
// func main(){
//     employee:=map[string]int{"a":101,"b":102,"c":103,"d":104}
//     for key,element:=range employee{
//         fmt.Println("key:",key,"=>", "Element:",element)
//     }
// }

/************SORTING MAP KEYS***********/
package main
import "fmt"
import "sort"
func main(){
    unSortedMap := map[string]int{"India": 1, "Canada": 2, "Germany": 3}
    keys:=make([]string,0,len(unSortedMap))
    for k:=range unSortedMap{
        keys=append(keys,k)
    }
    sort.Strings(keys)
    for _,k:=range keys{
        fmt.Println(k,unSortedMap[k])
    }
}
