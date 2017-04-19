package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
Solving Air Cargo Problem 3 using breadth_first_search...

Expansions   Goal Tests   New Nodes
  14663       18098       129631

Plan length: 12  Time elapsed in seconds: 103.0314349620021
Load(C1, P1, SFO)
Load(C2, P2, JFK)
Fly(P2, JFK, ORD)
Load(C4, P2, ORD)
Fly(P1, SFO, ATL)
Load(C3, P1, ATL)
Fly(P1, ATL, JFK)
Unload(C1, P1, JFK)
Unload(C3, P1, JFK)
Fly(P2, ORD, SFO)
Unload(C2, P2, SFO)
Unload(C4, P2, SFO)
*/

type DataSet struct {
	Name       string
	Expansions string
	GoalTests  string
	NewNodes   string
	PathLength string
}

func main() {
	// arg: 1-problem-1.log
	body, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	dataSets := make([]DataSet, 0)

	nextIs := ""
	lines := strings.Split(string(body), "\n")
	for _, line := range lines {
		words := strings.Split(line, " ")
		if nextIs == "data" {
			index := 0
			for _, w := range words {
				if len(strings.TrimSpace(w)) > 0 {
					switch index {
					case 0:
						dataSets[len(dataSets)-1].Expansions = w
					case 1:
						dataSets[len(dataSets)-1].GoalTests = w
					case 2:
						dataSets[len(dataSets)-1].NewNodes = w
					}
					index++
				}
			}
			nextIs = ""
			continue
		}
		for wnm, word := range words {
			if len(word) > 3 && word[len(word)-3:] == "..." {
				ds := DataSet{}
				ds.Name = word[0 : len(word)-3]
				if ds.Name == "h_1" || ds.Name == "h_ignore_preconditions" {
					ds.Name = words[wnm-2] + " " + words[wnm][0:len(words[wnm])-3]
				}
				dataSets = append(dataSets, ds)
			}
			if word == "Expansions" {
				nextIs = "data"
			}
			if word == "Plan" {
				dataSets[len(dataSets)-1].PathLength = words[wnm+2]
			}
		}
	}
	for _, ds := range dataSets {
		fmt.Printf("%+v\n", ds)
	}

	fmt.Println("Expansions")
	for _, ds := range dataSets {
		name := strings.Replace(strings.Title(strings.Replace(ds.Name, "_", " ", -1)), " ", "", -1)
		fmt.Printf("%s %s\n", name, ds.Expansions)
	}
	fmt.Println("Goal Tests")
	for _, ds := range dataSets {
		name := strings.Replace(strings.Title(strings.Replace(ds.Name, "_", " ", -1)), " ", "", -1)
		fmt.Printf("%s %s\n", name, ds.GoalTests)
	}
	fmt.Println("New Nodes")
	for _, ds := range dataSets {
		name := strings.Replace(strings.Title(strings.Replace(ds.Name, "_", " ", -1)), " ", "", -1)
		fmt.Printf("%s %s\n", name, ds.NewNodes)
	}
	fmt.Println("Path Length")
	for _, ds := range dataSets {
		name := strings.Replace(strings.Title(strings.Replace(ds.Name, "_", " ", -1)), " ", "", -1)
		fmt.Printf("%s %s\n", name, ds.PathLength)
	}
}
