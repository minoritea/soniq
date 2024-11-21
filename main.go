package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wk8/go-ordered-map/v2"
)

func main() {
	o := orderedmap.New[string, struct{}]()
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		o.Set(s.Text(), struct{}{})
	}
	for p := o.Oldest(); p != nil; p = p.Next() {
		fmt.Fprintln(os.Stdout, p.Key)
	}
}
