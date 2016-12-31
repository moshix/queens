package main

/* n queens solver by Moshe Bar Dec 31 2016
   orignally programmed in PL/1 back in 1982
   on an IDF mainframe. Ported to Golang in Dec 2016
   Copyright 2016 by Moshe Bar

   Program source is under Apache license             */

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

/* global variables */
var a [28]int64
var count int64

/* the below validate the command line flags passed
   for presence, and wether it's a numerical value */
type CustomStruct struct {
	FlagValue int
}

func (cF *CustomStruct) String() string {
	return strconv.Itoa(cF.FlagValue)
}

func (cF *CustomStruct) Set(s string) error {
	cF.FlagValue, _ = strconv.Atoi(s)

	return nil
}

func (cF *CustomStruct) Get() int {
	return cF.FlagValue
}

func (cF *CustomStruct) IsSet() bool {
	if cF.FlagValue == -1 {
		return false
	}

	return true
}

func main() {
	var n int64 = 8 // default size of chess board
	var loopvar int64

	limitFlag := CustomStruct{-1}

	flag.Var(&limitFlag, "n", "chessboard size")
	boolPtr := flag.Bool("s", false, "single chessboard size")

	helpPtr := flag.Bool("h", false, "show help dialogue")
	flag.Parse()
	if *helpPtr {
		fmt.Println("Usage:   ")
		fmt.Println("queens -h -s -n=14")
		fmt.Println(" ")
		fmt.Println("-h         show this help         ")
		fmt.Println("-s         single n value calculation switch")
		fmt.Println("-n=14      calculate all values from 4 to 14 in list form")
		fmt.Println("Smallest n value is 4. Highest practical value as of 2016 is 27 for supercomputers")
		fmt.Println("World record as of 2016 is n=26. More info http://penguin.ewu.edu/~trolfe/SCCS-95/SCCS-95.html")
		fmt.Println(" ")
		fmt.Println("queens is written by Moshe Bar in Go, derived from his original 1982 work in MVS PL/I")
		fmt.Println(" ")
		return
	}
	if limitFlag.IsSet() {
		fmt.Printf("\nn= %d\n\n", limitFlag.Get())
		n = int64(limitFlag.Get())
	} else {
		fmt.Printf("\nn command line argument is not included.\n")
		fmt.Println("Switching to interacive session....")
		fmt.Println("Please enter chessboard size: ")
		fmt.Scan(&n)

	}

	if n < 4 {
		fmt.Println("There are no values for n < 4. Please retry")
		return
	}

	if n > 27 {
		fmt.Println("n > 27 requires supercomputers as of 2016. Pls retry")
		return
	}
	if !*boolPtr {
		for loopvar = 4; loopvar <= n; loopvar++ {
			start := time.Now()
			queen(loopvar)               // we start computing solutions
			elapsed := time.Since(start) //elapsed time for this calculation

			/* we just returned from computation */
			fmt.Println(" n= ", loopvar, "solutions:  ", count, "-- exec time:", elapsed)
			count = 0
		}
	} else {
		start := time.Now()
		count = 0 // reset count
		queen(n)
		elapsed := time.Since(start) //elapsed time for this calculation

		/* we just returned from computation */
		fmt.Println(" n= ", n, "solutions:  ", count, "-- exec time:", elapsed)
	}

	return
}

func place(ps int64) int64 {
	/* here we check if a position is valid or not */
	/* 0 is invalid queen position                 */
	/* 1 is a valid position with no threats       */
	var i int64
	var absholder int64

	for i = 1; i < (ps); i++ {
		if a[i] == a[ps] {
			return 0
		}
		absholder = a[i] - a[ps]
		if absholder < 0 { // abs(absholder)
			absholder = -absholder //negative number multiplied by -1
		} // now it's a positive number

		if absholder == (ps - i) {
			return 0
		}
	}
	return 1
}

func queen(n int64) {
	/* in this function we backtrack the queens on the chess board, while checking for
	   a valid queen positioning

           This is the CPU intensive loop                                                */
	var k int64 = 1
	a[k] = 0
	for k > 0 {
		a[k] = a[k] + 1
		for (a[k] <= n) && (place(k) == 0) {
			a[k]++
		}
		if a[k] <= n {
			if k == n {
				count = count + 1
			} else {
				k++
				a[k] = 0
			}

		} else {
			k--
		}
	}
}
