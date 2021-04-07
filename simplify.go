package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Atoms struct {
	Atoms []Atom `json:"atoms"`
}

type Atom struct {
	Mass    float64  `json:"mass"`
	R       float64  `json:"r"`
	Epsilon float64  `json:"epsilon"`
	Names   []string `json:"names"`
}

type Groups struct {
	Groups []Group `json:"groups"`
}

type Group struct {
	Name  string      `json:"name"`
	Atoms []GroupAtom `json:"atoms"`
	Bonds []GroupBond `json:"bonds"`
}

type GroupAtom struct {
	Name  string `json:"name"`
	Q     string `json:"q"`
	Index int    `json:"index"`
}

type GroupBond struct {
	AtomIndexA int `json:"atom_index_A"`
	AtomIndexB int `json:"atom_index_B"`
}

type Chains struct {
	Chains []Chain `json:"chains"`
}

type Chain struct {
	Name   string            `json:"name"`
	Groups []ChainGroup      `json:"groups"`
	Names  []ChainGroupNames `json:"names"`
	Bonds  []ChainGroupBonds `json:"bonds"`
}

type ChainGroup struct {
	Name  string `json:"name"`
	Index int    `json:"index"`
}

type ChainGroupNames struct {
	Name       string `json:"name"`
	GroupIndex int    `json:"group_index"`
	AtomIndex  int    `json:"atom_index"`
}

type ChainGroupBonds struct {
	GroupIndexA int `json:"group_index_A"`
	AtomIndexA  int `json:"atom_index_A"`
	GroupIndexB int `json:"group_index_B"`
	AtomIndexB  int `json:"atom_index_B"`
}

type Molecules struct {
	Molecules []Molecules `json:"molecules"`
}

type Molecule struct {
	Name []string `json:"name"`
}

func main() {
	jsonFile, err := os.Open("fosmmpl.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var dat map[string]interface{}
	if err := json.Unmarshal(byteValue, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat["mmpl"].(map[string]interface{})["atom"])

}
