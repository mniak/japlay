package main

import (
	"strings"

	"github.com/arch-mage/mdb"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var cmdRoot = &cobra.Command{
	Use: "louvorja-standalone",
}

func main() {
	cmdRoot.Execute()

	// const mdbfile = "config/BD.mdb"
	// file := lo.Must(os.Open(mdbfile))
	// defer file.Close()

	// tabMusicas := lo.Must(mdb.Rows(file, "MUSICAS"))
	// row := tabMusicas.Next()

	// mdb.Rows()
	// musicas, found := getTable(tables)
	// if !found {
	// 	log.Fatalln("Table MUSICAS not found")
	// }

	// for _, v := range musicas {
	// }
}

func getTable(tables []mdb.Table, name string) (mdb.Table, bool) {
	return lo.Find(tables, func(t mdb.Table) bool {
		return !t.Sys &&
			strings.EqualFold(t.Name, name)
	})
}
