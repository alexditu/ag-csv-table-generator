package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/alexditu/ag-csv-table-generator/version"
)

type Proprietar struct {
	Apt         string
	Nume        string
	MembruAssoc bool
}

// <nume, Proprietar>
type MembriAssoc map[string]Proprietar

func writeToCsv(data []Proprietar, outFilePath string) error {
	f, err := os.Create(outFilePath)
	if err != nil {
		return err
	}

	defer f.Close()

	f.WriteString("Nr. Crt., Proprietate, Nume Proprietar, Semnătură\n")

	for i, v := range data {
		f.WriteString(strconv.Itoa(i+1) + ", " + v.Apt + ", " + v.Nume + "," + "\n")
	}

	f.Sync()

	return nil
}

func main() {
	fmt.Printf("Running %s version %s\n", version.BinaryName(), version.String())

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <input csv>\n", os.Args[0])
		fmt.Printf("\t<input csv> format: Sc, Ap, Nume Proprietar, Membru in Asociatie\n")
		os.Exit(1)
	}

	fPath := os.Args[1]

	f, err := os.Open(fPath)
	if err != nil {
		fmt.Printf("Error: failed to read file '%s': %s\n", fPath, err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	if scanner.Scan() {
		// skip CSV header
		scanner.Text()
	} else {
		fmt.Printf("Error: file '%s' is empty\n", fPath)
		os.Exit(1)
	}

	members := MembriAssoc{}

	for scanner.Scan() {
		rawData := scanner.Text()
		data := strings.Split(rawData, ",")
		if len(data) < 2 {
			fmt.Printf("Invalid line: '%s', skipping", rawData)
		}

		var p Proprietar

		p.Apt = data[0] + data[1]
		if len(data) >= 3 && data[2] != "" {
			p.Nume = data[2]
		} else {
			fmt.Printf("Skipping '%s' (nu are nume)\n", rawData)
			continue
		}

		if len(data) >= 4 && data[3] == "DA" {
			p.MembruAssoc = true
		} else {
			p.MembruAssoc = false
		}

		oldP, exists := members[p.Nume]
		if exists {
			if oldP.MembruAssoc || p.MembruAssoc {
				oldP.MembruAssoc = true
			} else {
				fmt.Printf("Skipping '%s' (nu e membru în Assoc 2)\n", rawData)
				continue
			}

			oldP.Apt += "; " + p.Apt
			members[p.Nume] = oldP
		} else {
			if !p.MembruAssoc {
				fmt.Printf("Skipping '%s' (nu e membru în Assoc)\n", rawData)
				continue
			}

			members[p.Nume] = p
		}
	}

	// sort by apt
	var agOnlyMembers []Proprietar
	for _, v := range members {
		agOnlyMembers = append(agOnlyMembers, v)
	}

	slices.SortFunc(agOnlyMembers, func(a, b Proprietar) int {
		return cmp.Compare(a.Apt, b.Apt)
	})

	// for _, i := range agOnlyMembers {
	// 	fmt.Printf("%s, %s\n", i.Apt, i.Nume)
	// }

	outFile := "./tabel_prezenta_AG.csv"
	err = writeToCsv(agOnlyMembers, outFile)
	if err != nil {
		fmt.Printf("Error: failed to write to output file '%s': %s\n", outFile, err)
		os.Exit(1)
	}
}
