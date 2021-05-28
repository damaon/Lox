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

func run(source string){
    tokens := scanTokens(source);

    // For now, just print the tokens.
    for token := range tokens {
      fmt.Println(token);
    }
}

func scanTokens(source string) []string{
	// TODO: move to scanner modules
	tokens := []string{"a", "b", "c", "d"}
	return tokens;
}

func error(line int, message string){
	report(line, "", message)
}

var hadError = false; // global for now
func report(line int, where, message string){
	fmt.Println("[line %v] Error %v:%v", line, where, message)
	hadError = true;
}


