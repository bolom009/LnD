package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/bndr/gotabulate"

	"github.com/bolom009/LnD/life"
)

var (
	width, height int
	round         = 1
)

const snapshotsPath = "./snapshots/"

func main() {
	flag.IntVar(&width, "w", 25, "width game field")
	flag.IntVar(&height, "h", 25, "height game field")
	flag.Parse()

	if err := removeContents(snapshotsPath); err != nil {
		_ = fmt.Errorf("%v\n", err)
		os.Exit(1)
		return
	}

	var field = life.GenerateField(width, height)
	if err := makeSnapshot(field, round, snapshotsPath); err != nil {
		_ = fmt.Errorf("%v\n", err)
		os.Exit(1)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Press enter for play game")
		scanner.Scan()

		field = field.NextRound()
		round++
		if err := makeSnapshot(field, round, snapshotsPath); err != nil {
			_ = fmt.Errorf("%v\n", err)
			os.Exit(1)
			return
		}
	}
}

func makeSnapshot(field *life.Field, round int, path string) error {
	t := gotabulate.Create(field.GetCells())
	t.SetAlign("center")

	filename := path + strconv.Itoa(round) + ".txt"
	if err := ioutil.WriteFile(filename, []byte(t.Render("grid")), os.ModePerm); err != nil {
		return err
	}

	return nil
}

func removeContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}

	defer d.Close()

	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}

	for _, name := range names {
		if name == ".gitignore" {
			continue
		}

		if err = os.RemoveAll(filepath.Join(dir, name)); err != nil {
			return err
		}
	}

	return nil
}
