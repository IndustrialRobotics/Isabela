package main

import (
	"os"
	"fmt"
)

//24.justforfunc #24- what's the most common identifier in the Go stdlib-

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage:\n\t%s [files]\n", os.Args[0])
		os.Exit(1)
	}

	//fs := token.NewFileSet()

	//for _, arg := range os.Args[1:] {
	//	b, err := ioutil.ReadFile(arg)
	//	if err != nil {
	//		log.Fatal(err)
	//	}

		//f := fs.AddFile(arg, fs.Base(), len(b))
		//var s scanner.Scanner
		//s.Init(f, b, nil, scanner.ScanComments)
		//
		//for {
		//	_, tok, lit := s.Scan()
		//	if tok == token.EOF {
		//		break
		//	}
		//
		//	if tok == token.IDENT {
		//		fmt.Println(lit)
		//	}
		//}


}
