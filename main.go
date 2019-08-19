package main

import (
	"fmt"
	"github.com/coreos/go-iptables/iptables"
)

func main() {
	it, err := iptables.New()
	if err != nil {
		panic(err)
	}
	l, err := it.Stats("nat", "PREROUTING")
	if err != nil {
		panic(err)
	}
	for _, s := range l {
		p, err := it.ParseStat(s)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", p)
	}
	fmt.Println("")

	rulespec := []string{"-s", "127.0.0.1:9090", "-j", "DNAT", "--to", "127.0.0.1:8080"}
	err = it.Insert("nat", "PREROUTING", 1, rulespec...)
	if err != nil {
		panic(err)
	}
	// err = it.Delete("nat", "PREROUTING", rulespec...)
	// if err != nil {
	// 	panic(err)
	// }
	// l, err = it.Stats("nat", "PREROUTING")
	// if err != nil {
	// 	panic(err)
	// }
	// for _, s := range l {
	// 	p, err := it.ParseStat(s)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Printf("%+v\n", p)
	// }
}
