package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main(){
	if len(os.Args) > 2 {
		fmt.Println("Usage: golox [script]");
		os.Exit(64) 
	  } else if len(os.Args) == 2 {
		runFile(os.Args[1]);
	  } else {
		runPrompt();
	  }
}

func check(e error) {
	// Good enough for now
    if e != nil {
        panic(e)
    }
}

func runFile(fileName string){
	bytes, err := ioutil.ReadFile(fileName)
	check(err)
	run(string(bytes));
}

func runPrompt(){
	reader := bufio.NewReader(os.Stdin)
    for { 
		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
      	if line == "" { break; } 
      	run(line);
    }
}

func run(code string){

}

