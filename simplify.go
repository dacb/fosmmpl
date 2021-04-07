// Package simplify takes a wordy abstract representation of the MMPL
// and converts it to a more compact set of structures that map onto
// a json structure for easy i/o in other settings.
// This tool will only ever be run once to generate the compact representation.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type MMPL struct {
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
	Name    string  `json:"name"`
	Q       float64 `json:"q"`
	Index   int     `json:"index"`
	Comment string  `json:"comment"`
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
	N       int     `json:"n"`
	K       float64 `json:"k"`
	Comment string  `json:"comment"`
}

func unpackAtom(m map[string]interface{}) Atom {
	mass, err := strconv.ParseFloat(m["mass"].(string), 64)
	if err != nil {
		log.Fatal(err)
	}
	epsilon, err := strconv.ParseFloat(m["epsilon"].(string), 64)
	if err != nil {
		log.Fatal(err)
	}
	r, err := strconv.ParseFloat(m["r"].(string), 64)
	if err != nil {
		log.Fatal(err)
	}

	cmt, ok := m["comment"].(string)
	if !ok {
		cmt = ""
	}

	names := []string{}

	nlist, ok := m["names"].(map[string]interface{})["n"]
	for _, v := range nlist.([]interface{}) {
		if mv, ok := v.(map[string]interface{}); ok {
			names = append(names, mv["name"].(string))
		}
	}
	atom := Atom{
		Epsilon: epsilon, Mass: mass, R: r,
		Names:   names,
		Comment: cmt}
	return atom
}

func unpackGroup(m map[string]interface{}) Group {
	name, ok := m["name"].(string)
	if !ok {
		log.Fatal("unable to find name for group")
	}
	fmt.Println(name)
	cmt, ok := m["comment"].(string)
	if !ok {
		cmt = ""
	}
	atoms := []GroupAtom{}
	alist, ok := m["atoms"].(map[string]interface{})["a"]
	fmt.Printf("%T\n", alist)
	fmt.Println(alist)
	for _, v := range alist.([]interface{}) {
		if mv, ok := v.(map[string]interface{}); ok {
			mvname, ok := mv["name"].(string)
			if !ok {
				log.Fatal("in group, unable to find name of atom")
			}
			q, err := strconv.ParseFloat(mv["q"].(string), 64)
			if err != nil {
				log.Fatal(err)
			}
			idx, err := strconv.Atoi(mv["idx"].(string))
			if err != nil {
				log.Fatal(err)
			}
			agcmt, ok := mv["comment"].(string)
			if !ok {
				agcmt = ""
			}
			atoms = append(atoms, GroupAtom{
				Name:    mvname,
				Q:       q,
				Index:   idx,
				Comment: agcmt,
			})
		}
	}
	group := Group{Name: name,
		Atoms:   atoms,
		Comment: cmt}
	return group
}

func unpackBond(m map[string]interface{}) BondType {
	l, err := strconv.ParseFloat(m["l"].(string), 64)
	if err != nil {
		log.Fatal(err)
	}
	k, err := strconv.ParseFloat(m["k"].(string), 64)
	if err != nil {
		log.Fatal(err)
	}

	cmt, ok := m["comment"].(string)
	if !ok {
		cmt = ""
	}

	bondType := BondType{
		NameA: m["A"].(string), NameB: m["B"].(string),
		L: l, K: k,
		Comment: cmt}
	return bondType
}

func unpackAngle(m map[string]interface{}) AngleType {
	theta, err := strconv.ParseFloat(m["theta"].(string), 64)
	if err != nil {
		log.Fatal(err)
	}
	k, err := strconv.ParseFloat(m["k"].(string), 64)
	if err != nil {
		log.Fatal(err)
	}

	cmt, ok := m["comment"].(string)
	if !ok {
		cmt = ""
	}

	angleType := AngleType{
		NameA: m["A"].(string), NameB: m["B"].(string),
		NameC: m["C"].(string),
		Theta: theta, K: k,
		Comment: cmt}
	return angleType
}

func unpackTorsion(m map[string]interface{}) TorsionType {
	ttype, err := strconv.Atoi(m["type"].(string))
	if err != nil {
		log.Fatal(err)
	}
	n, err := strconv.Atoi(m["n"].(string))
	if err != nil {
		log.Fatal(err)
	}
	phi, err := strconv.ParseFloat(m["phi"].(string), 64)
	if err != nil {
		log.Fatal(err)
	}
	k, err := strconv.ParseFloat(m["k"].(string), 64)
	if err != nil {
		log.Fatal(err)
	}

	cmt, ok := m["comment"].(string)
	if !ok {
		cmt = ""
	}

	torsionType := TorsionType{
		NameA: m["A"].(string), NameB: m["B"].(string),
		NameC: m["C"].(string), NameD: m["D"].(string),
		Type: ttype, Phi: phi, N: n, K: k,
		Comment: cmt}
	return torsionType
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

	var mmpl MMPL
	mmpl.Atoms = []Atom{}
	mmpl.Groups = []Group{}
	mmpl.Chains = []Chain{}
	mmpl.Molecules = []Molecule{}
	mmpl.BondTypes = []BondType{}
	mmpl.AngleTypes = []AngleType{}
	mmpl.TorsionTypes = []TorsionType{}
	// not yet
	//if err := json.Unmarshal(byteValue, &mmpl); err != nil {
	//panic(err)
	//}

	// extract out the mmpl object and prepare to loop over the lists
	// within it
	mmplx := dat["mmpl"].(map[string]interface{})

	// atoms
	atoms := mmplx["atoms"]
	for k, v := range atoms.([]interface{}) {
		if mv, ok := v.(map[string]interface{}); ok {
			atom := unpackAtom(mv)
			mmpl.Atoms = append(mmpl.Atoms, atom)
		} else {
			log.Fatal("should not happen", k, v)
		}
	}
	// groups
	groups := mmplx["groups"]
	for k, v := range groups.([]interface{}) {
		if mv, ok := v.(map[string]interface{}); ok {
			group := unpackGroup(mv)
			mmpl.Groups = append(mmpl.Groups, group)
		} else {
			log.Fatal("should not happen", k, v)
		}
	}

	// bond types
	bonds := mmplx["bonds"]
	for k, v := range bonds.([]interface{}) {
		if mv, ok := v.(map[string]interface{}); ok {
			bondType := unpackBond(mv)
			mmpl.BondTypes = append(mmpl.BondTypes, bondType)
		} else {
			log.Fatal("should not happen", k, v)
		}
	}
	// bond angle types
	angles := mmplx["angles"]
	for k, v := range angles.([]interface{}) {
		if mv, ok := v.(map[string]interface{}); ok {
			angleType := unpackAngle(mv)
			mmpl.AngleTypes = append(mmpl.AngleTypes, angleType)
		} else {
			log.Fatal("should not happen", k, v)
		}
	}
	// torsion angle types
	torsions := mmplx["torsions"]
	for k, v := range torsions.([]interface{}) {
		if mv, ok := v.(map[string]interface{}); ok {
			torsionType := unpackTorsion(mv)
			mmpl.TorsionTypes = append(mmpl.TorsionTypes, torsionType)
		} else {
			log.Fatal("should not happen", k, v)
		}
	}

	//log.Fatal("nuff")

	json, err := json.MarshalIndent(mmpl.Groups, " ", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))

}
