package main

import (
   "fmt"
   "os"
   "bytes"
)

func main() {

    phonetic_dict:= map[string]string{
      
            "0": "Zero",
            "1": "One",
            "2": "Two",
            "3": "Three",
            "4": "Four",
	    "5": "Five",
	    "6": "Six",
  	    "7": "Seven",
	    "8": "Eight",
	    "9": "Nine",
    }
    
   
    input_digits:= os.Args[1:]
    var result bytes.Buffer
    
    for index,digit:= range input_digits {
        
        var digitString string
        for i := 0; i < len(digit); i++ {
            digitString = phonetic_dict[string(digit[i])]
            result.WriteString(digitString)
        }

	if index < len(input_digits)-1 {
		result.WriteString(",")
	} 
    }
    fmt.Print(result.String())
}