package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"os"

// 	"github.com/mniak/louvorja/pkg/mdblib"

// 	"github.com/samber/lo"
// 	"github.com/spf13/cobra"

// 	_ "github.com/mattn/go-sqlite3"
// )

// func init() {
// 	cmdRoot.AddCommand(cmdMigrate)
// 	cmdMigrate.Flags().BoolP("force", "f", false, "Overwrites target file if already existent")
// }

// var cmdMigrate = &cobra.Command{
// 	Use:   "migrate SOURCE TARGET",
// 	Short: "Migrate data from the input MDB file to a SQLite file",
// 	Args:  cobra.ExactArgs(2),
// 	Run: func(cmd *cobra.Command, args []string) {
// 		sourceFile := args[0]
// 		targetFile := args[1]
// 		force := lo.Must(cmd.Flags().GetBool("force"))

// 		if force {
// 			os.RemoveAll(targetFile)
// 		}

// 		db := lo.Must(sql.Open("sqlite3", targetFile))
// 		defer db.Close()

// 		mdbfile := lo.Must(os.Open(sourceFile))
// 		defer mdbfile.Close()

// 		lo.Must0(ensureTables(db))
// 		lo.Must0(copyRows(mdbfile, db))
// 	},
// }

// func ensureTables(db *sql.DB) error {
// 	// _, err := db.Exec()
// 	return nil
// }

// func copyRows(mdbfile *os.File, db *sql.DB) error {
// 	albumsTable, err := mdblib.Scan(mdbfile, "ALBUM")
// 	fmt.Println(albumsTable...)
// 	return err
// }

// // func copyAll(mdbfile, sqlitefile string) error {
// // 	//
// // 	// list tables
// // 	//
// // 	tables, err := mdb.Tables(file)
// // 	if err != nil {
// // 		return err
// // 	}

// // 	//
// // 	// make transaction
// // 	//
// // 	tx, err := db.Begin()
// // 	if err != nil {
// // 		return err
// // 	}

// // 	for _, table := range tables {
// // 		if table.Sys { // skip system table
// // 			continue
// // 		}
// // 		if err := createTable(tx, file, table); err != nil {
// // 			return err
// // 		}
// // 		if err := copyRows(tx, file, table); err != nil {
// // 			return err
// // 		}
// // 	}

// // 	return tx.Commit()
// // }

// // func createTable(tx *sql.Tx, file io.ReadSeeker, table mdb.Table) error {
// // 	//
// // 	// prepare type conversion
// // 	//
// // 	typeConversion := map[string]string{
// // 		"Bool":       "BOOl",
// // 		"Byte":       "BYTE",
// // 		"Int":        "INTEGER",
// // 		"LongInt":    "INTEGER",
// // 		"Money":      "NUMERIC",
// // 		"Float":      "REAL",
// // 		"Double":     "REAL",
// // 		"DateTime":   "DATETIME",
// // 		"Binary":     "BLOB",
// // 		"Text":       "TEXT",
// // 		"LongBinary": "BLOB",
// // 		"LongText":   "TEXT",
// // 		"GUID":       "TEXT",
// // 		"Numeric":    "NUMERIC",
// // 	}

// // 	//
// // 	// collect columns
// // 	//
// // 	columns := make([]string, len(table.Columns))
// // 	for i, column := range table.Columns {
// // 		columns[i] = fmt.Sprintf("%s %s", column.Name, typeConversion[column.Type])
// // 		columns[i] = column.Name
// // 	}

// // 	//
// // 	// construct query
// // 	//
// // 	query := fmt.Sprintf(
// // 		"CREATE TABLE %s (%s)",
// // 		table.Name,
// // 		strings.Join(columns, ", "),
// // 	)

// // 	//
// // 	// execute
// // 	//
// // 	_, err := tx.Exec(query)
// // 	return err
// // }

// // func copyRows(tx *sql.Tx, file io.ReadSeeker, table mdb.Table) error {
// // 	//
// // 	// collect columns
// // 	//
// // 	columns := make([]string, len(table.Columns))
// // 	bindings := make([]string, len(table.Columns))
// // 	for i, column := range table.Columns {
// // 		columns[i] = column.Name
// // 		bindings[i] = "?"
// // 	}

// // 	//
// // 	// construct query
// // 	//
// // 	query := fmt.Sprintf(
// // 		"INSERT INTO %s (%s) VALUES (%s)",
// // 		table.Name,
// // 		strings.Join(columns, ", "),
// // 		strings.Join(bindings, ", "),
// // 	)

// // 	//
// // 	// prepare it
// // 	//
// // 	stmt, err := tx.Prepare(query)
// // 	if err != nil {
// // 		return err
// // 	}
// // 	defer stmt.Close()

// // 	//
// // 	// iterate rows
// // 	//
// // 	rows, err := mdb.Rows(file, table.Name)
// // 	if err != nil {
// // 		return err
// // 	}
// // 	for {
// // 		fields, err := rows.Next()
// // 		if err == io.EOF {
// // 			break
// // 		}
// // 		if err != nil {
// // 			return err
// // 		}

// // 		//
// // 		// insert
// // 		//
// // 		if _, err = stmt.Exec(fields...); err != nil {
// // 			return err
// // 		}
// // 	}
// // 	return nil
// // }
