package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Person struct{
	First string
	Last string
	Age int
}

func main(){
	ej5()
}
func ej1(){
	//TODO://Marhsall a structure to a Json object
	p1:= Person{
		First: "Johan",
		Last: "Chacon",
		Age: 23,
	}
	p2:= Person{
		First: "Hannia",
		Last: "Chacon",
		Age: 50,
	}
	p3:= Person{
		First: "Fernando",
		Last: "Aguilar",
		Age: 30,
	}
	//This is an array of structs
	group:=[]Person{
		p1,p2,p3,
	}
	//You need to turn this into []byte to use the marshall func

	bs, err := json.Marshal(group)
	if err!= nil{
		fmt.Println("An error ocured in the mutuacion of the slice:",err)
	}else{
		fmt.Println("This is the slice encode in JSON: ",bs)
		fmt.Println("This is the JSON conver to string:",string(bs))
	}

}


func ej2(){
	//TODO: Unmarshal the object in to JSON in to a Go Struct 

	//This is in JSON Format
	s:=`[{"First":"Johan","Last":"Chacon","Age":32},{"First":"Fle","Last":"Bass", "Age":51},{"First":"Francisco","Last":"Aguilar", "Age":30}]`
	//This is in JSON Format
	fmt.Println("This is the Json Object ", s)
	//Need to turn it in to  []Bit
	bs := []byte(s)
	//You need a struct in go To save the Json
	p:=[]Person{}
	//Using UnMarshall to Mutuaite with the pointer to the structure 
	err := json.Unmarshal(bs,&p)
	if err != nil{
		fmt.Println("Has been an error in the Mutuacion of the Object ",err )
	}else{
		fmt.Println("The object has been Mutuated to the struct: ",p)
	}

}
func ej3(){
	//So in summary, this line of code is encoding the "group" slice as JSON and writing it to standard output, and checking for any errors that occur during the encoding process.
	p1:= Person{
		First: "Johan",
		Last: "Chacon",
		Age: 23,
	}
	p2:= Person{
		First: "Hannia",
		Last: "Chacon",
		Age: 50,
	}
	p3:= Person{
		First: "Fernando",
		Last: "Aguilar",
		Age: 30,
	}
	//This is an array of structs
	group:=[]Person{
		p1,p2,p3,
	}
	err:= json.NewEncoder(os.Stdout).Encode(group)
	
	if err != nil{
		fmt.Println("This is the error:",err)
	}

}

func ej4(){
	//Sort throught this
	xi := []int{6,6,72,12,123,123213,4311,2,3,4,5,24,3412312,123340912,4432,}
	sort.Ints(xi)
	fmt.Println(xi)
	xx:= []string{"Johan","Josue","Fernando","Jefferson","Hannia"}
	sort.Strings(xx)
	fmt.Println(xx)
}
//TODO:////////////////////////TODO:////////////////////TODO:////////////
type Pperson struct {
	Name string
	Age  int
}
//This one does by Age (Int)
type byAge []Pperson

func (a byAge) Len() int           {return len(a) }
func (a byAge) Swap(i, j int)      {a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool {return a[i].Age < a[j].Age}

//This one does by Name (string)
type byLast []Pperson

func (a byLast) Len() int           { return len(a) }
func (a byLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byLast) Less(i, j int) bool { return a[i].Name < a[j].Name }

func ej5(){
	p1:= Pperson{
		Name: "Josue",
		Age: 23,
	}
	p2:= Pperson{
		Name: "Pedro",
		Age: 44,
	}
	p3:= Pperson{
		Name: "Adrian",
		Age: 55,
	}
	group :=[]Pperson{
		p1,p2,p3,
	}
	fmt.Println(group)
	//Call the Method sort.Sort([]slice of the struct)
	sort.Sort(byLast(group))
	fmt.Println("Order by Name.")
	fmt.Println(group)
	//Call the Method sort.Sort([]slice of the struct)
	sort.Sort(byAge(group))
	fmt.Println("Order By Age")
	fmt.Println(group)
	
}
