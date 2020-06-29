package main
//test
import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("测试Push，增加一行")
	fmt.Println("哈哈哈....")
	fmt.Println("测试pull")


	fmt.Println("aaaaa")
	fmt.Println("bbbbb")
	fmt.Println("ccccc")
	fmt.Println("ddddd")

	now :=time.Now()
	//123
	nostr :=now.Format("2006-01")
	fmt.Println(nostr)

	fmt.Println(time.Unix(1576674119,0))
    //fmt.Println(time.Now() )
	//
	////var test float64
	////
	////test =3/2
	//var  a int64
	//a =20
	//
	//
	//fmt.Println(float64(a))

	//hh :=int64(0)
	//fmt.Println(hh)

   maps :=make(map[string]interface{},0)
   maps["1"]=1
   maps["2"]=2
   maps["3"]=3
   maps["4"]=4
   fmt.Println(maps)


	fmt.Println("test fetch")

}


