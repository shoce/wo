/*
history:
2020/04/20 v1

GoFmt GoBuildNull GoBuild GoRelease GoRun
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var err error
	var n int
	var l string
	var ww []string

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: wo n\n	n = position of the word to select, negative to select from the end of the line\n")
		os.Exit(1)
	}

	n, err = strconv.Atoi(os.Args[1])
	if err != nil || n == 0 {
		fmt.Fprintf(os.Stderr, "invalid n = '%s'", os.Args[1])
		os.Exit(1)
	}

	sp := regexp.MustCompile(`[ \t]+`)
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		l = sc.Text()
		l = strings.TrimSpace(l)

		ww = sp.Split(l, -1)

		if n > 0 && len(ww) < n {
			continue
		}
		if n < 0 && len(ww) < -n {
			continue
		}

		if n > 0 {
			fmt.Println(ww[n-1])
		} else {
			fmt.Println(ww[len(ww)+n])
		}
	}
	if err = sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "read stdin: %v", err)
	}
}
