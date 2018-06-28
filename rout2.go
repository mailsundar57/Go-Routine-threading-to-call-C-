package main

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
	"strconv"
)

var wg sync.WaitGroup

func callCpp(s string){

    supplyArg := "./a.out" 
	
	fmt.Println(supplyArg)
	
    out, err := exec.Command("./a.out", s).Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The out of Cpp is %s\n", out)
	wg.Done()
}

func callCppint(v int){

    supplyArg :=  strconv.Itoa(v) 
	
	// fmt.Println(supplyArg)
	
    out, err := exec.Command("./a1.out",supplyArg).Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The out of Cpp is %s\n", out)
	wg.Done()
}

func main() {

	wg.Add(1)
	go callCpp("heythis")
	wg.Add(1)
	go callCpp("heythat")
	wg.Wait()

	for i := 1; i< 5;i++{
		// fmt.Println(i)
    	wg.Add(1)
		go callCppint(i)
		wg.Wait()
	}
	       
    wg.Add(1)   
	go callCppint(82)
	wg.Wait()
}

