package migration

import "embed"

//go:embed *.sql
var MigrationFs embed.FS
