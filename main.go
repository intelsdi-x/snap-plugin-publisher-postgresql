package main

import (
	"os"

	"github.com/intelsdi-x/pulse-plugin-publisher-postgresql/postgresql"
	"github.com/intelsdi-x/pulse/control/plugin"
)

func main() {
	meta := postgresql.Meta()
	plugin.Start(meta, postgresql.NewPostgreSQLPublisher(), os.Args[1])
}
