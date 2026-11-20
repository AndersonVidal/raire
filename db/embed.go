// Package db embeds the forward-only SQL migrations and idempotent seeds.
package db

import "embed"

//go:embed migrations/*.sql
var Migrations embed.FS

//go:embed seeds/*.sql
var Seeds embed.FS
