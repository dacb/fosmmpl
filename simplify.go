package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Mmpl struct {
	Atoms        []Atom        `json:"atoms"`
	Groups       []Group       `json:"groups"`
	Chains       []Chain       `json:"chains"`
	Molecules    []Molecule    `json:"molecules"`
	BondTypes    []BondType    `json:"bonds"`
	AngleTypes   []AngleType   `json:"angles"`
	TorsionTypes []TorsionType `json:"torsions"`
}

type Atom struct {
	Mass    float64  `json:"mass"`
	R       float64  `json:"r"`
	Epsilon float64  `json:"epsilon"`
	Names   []string `json:"names"`
	Comment string   `json:"comment"`
}

type Group struct {
	Name    string      `json:"name"`
	Atoms   []GroupAtom `json:"atoms"`
	Bonds   []Bond      `json:"bonds"`
	Comment string      `json:"comment"`
}

type GroupAtom struct {
	Name    string `json:"name"`
	Q       string `json:"q"`
	Index   int    `json:"index"`
	Comment string `json:"comment"`
}

type Bond struct {
	ChainIndexA int    `json:"chain_index_A"`
	GroupIndexA int    `json:"group_index_A"`
	AtomIndexA  int    `json:"atom_index_A"`
	ChainIndexB int    `json:"chain_index_B"`
	GroupIndexB int    `json:"group_index_B"`
	AtomIndexB  int    `json:"atom_index_B"`
	Comment     string `json:"comment"`
}

type Chains struct {
	Chains []Chain `json:"chains"`
}

type Chain struct {
	Name    string            `json:"name"`
	Groups  []ChainGroup      `json:"groups"`
	Names   []ChainGroupNames `json:"names"`
	Bonds   []Bond            `json:"bonds"`
	Comment string            `json:"comment"`
}

type ChainGroup struct {
	Name    string `json:"name"`
	Index   int    `json:"index"`
	Comment string `json:"comment"`
}

type ChainGroupNames struct {
	Name       string `json:"name"`
	GroupIndex int    `json:"group_index"`
	AtomIndex  int    `json:"atom_index"`
	Comment    string `json:"comment"`
}

type Molecules struct {
	Molecules []Molecules `json:"molecules"`
}

type Molecule struct {
	Name    string  `json:"name"`
	Comment string  `json:"comment"`
	Chains  []Chain `json:"chains"`
	Bonds   []Bond  `json:"bonds"`
}

type BondType struct {
	NameA   string  `json:"name_A"`
	NameB   string  `json:"name_B"`
	L       float64 `json:"l"`
	K       float64 `json:"k"`
	Comment string  `json:"comment"`
}

type AngleType struct {
	NameA   string  `json:"name_A"`
	NameB   string  `json:"name_B"`
	NameC   string  `json:"name_C"`
	Theta   float64 `json:"theta"`
	K       float64 `json:"k"`
	Comment string  `json:"comment"`
}

type TorsionType struct {
	NameA   string  `json:"name_A"`
	NameB   string  `json:"name_B"`
	NameC   string  `json:"name_C"`
	NameD   string  `json:"name_D"`
	Type    int     `json:"type"`
	Phi     float64 `json:"phi"`
	N       float64 `json:"n"`
	K       float64 `json:"k"`
	Comment string  `json:"comment"`
}

func main() {
	jsonFile, err := os.Open("fosmmpl_complex.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var dat map[string]interface{}
	if err := json.Unmarshal(byteValue, &dat); err != nil {
		panic(err)
	}

	var mmpl Mmpl
	if err := json.Unmarshal(byteValue, &mmpl); err != nil {
		panic(err)
	}
	//mmpl.TorsionTypes = []TorsionType{}
	//mmpl.Atoms = []Atom{}
	//mmpl.Groups = []Group{}
	//mmpl.Chains = []Chain{}
	//mmpl.Molecules = []Molecule{}

	//fmt.Println(dat["mmpl"].(map[string]interface{})["atom"])
	//for key, _ := range dat["mmpl"].(map[string]interface{}) {
	//fmt.Println(key)
	//}
	//fmt.Println(dat["mmpl"].(map[string]interface{})["ap"])
	//for i, value := range dat["mmpl"].(map[string]interface{})["torsions"].([]interface{}) {
	//fmt.Println(key, value)
	//fmt.Println(value)
	//fmt.Println(value.(map[string]interface{})["A"])

	//torsionType := TorsionType{"A", "B", "C", "D", i, 0, 0, 1., "nothing here"}
	//fmt.Println(torsionType)
	//mmpl.TorsionTypes = append(mmpl.TorsionTypes, torsionType)
	//}

	fmt.Println(mmpl)

}
