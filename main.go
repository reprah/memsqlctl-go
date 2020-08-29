package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/olekukonko/tablewriter"
	"os"
)

type MemSql struct {
	conn *sql.DB
}

func (m *MemSql) showLeaves() error {
	rows, err := m.conn.Query(`SHOW LEAVES`)

	if err != nil {
		return err
	}

	defer rows.Close()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	columnNames, err := rows.Columns()

	if err != nil {
		return err
	}

	table.SetHeader(columnNames)

	for rows.Next() {
		var (
			host                         string
			port                         string
			availability_group           string
			pair_host                    sql.NullString
			pair_port                    sql.NullString
			state                        string
			opened_connections           string
			average_roundtrip_latency_ms string
			node_id                      string
			grace_period                 sql.NullString
		)

		err := rows.Scan(
			&host,
			&port,
			&availability_group,
			&pair_host,
			&pair_port,
			&state,
			&opened_connections,
			&average_roundtrip_latency_ms,
			&node_id,
			&grace_period,
		)

		if err != nil {
			return err
		}

		row := []string{
			host,
			port,
			availability_group,
			pair_host.String,
			pair_port.String,
			state,
			opened_connections,
			average_roundtrip_latency_ms,
			node_id,
			grace_period.String,
		}

		table.Append(row)
	}

	table.Render()

	return nil
}

func (m *MemSql) showLicense() error {
	rows, err := m.conn.Query(`SHOW STATUS LIKE "%License%"`)

	if err != nil {
		return err
	}

	defer rows.Close()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for rows.Next() {
		var (
			name  string
			value string
		)

		if err := rows.Scan(&name, &value); err != nil {
			return err
		}

		table.Append([]string{name, value})
	}

	table.Render()

	return nil
}

func (m *MemSql) setLicense(license string) error {
	_, err := m.conn.Exec(`SET LICENSE = ?`, license)

	if err != nil {
		return err
	} else {
		fmt.Printf("Set license to %s\n", license)
	}

	return nil
}

func NewMemSql() (*MemSql, error) {
	dbConfig := mysql.Config{
		User:                 "root",
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", "127.0.0.1", "3306"),
		AllowNativePasswords: true,
	}

	conn, err := sql.Open("mysql", dbConfig.FormatDSN())

	if err != nil {
		return &MemSql{}, err
	}

	if err = conn.Ping(); err != nil {
		return &MemSql{}, err
	}

	return &MemSql{conn: conn}, nil
}

func main() {
	commands := []string{"show-leaves", "set-license", "show-license"}

	// first element is the program name, second is the subcommand
	if len(os.Args) < 2 {
		fmt.Printf("Available subcommands: %v\n", commands)
		return
	}

	memSql, err := NewMemSql()

	if err != nil {
		fmt.Printf("Error connecting to MemSQL: %v\n", err)
		return
	}

	// ensure we close the connection
	defer memSql.conn.Close()

	switch command := os.Args[1]; command {
	case "set-license":
		fs := flag.NewFlagSet("set-license", flag.ContinueOnError)
		license := fs.String("license", "", "A base-64 encoded MemSQL license")
		fs.Parse(os.Args[2:])
		err = memSql.setLicense(*license)
	case "show-license":
		err = memSql.showLicense()
		return
	case "show-leaves":
		err = memSql.showLeaves()
	default:
		fmt.Printf("Available subcommands: %v\n", commands)
		return
	}

	if err != nil {
		fmt.Println(err)
	}
}
