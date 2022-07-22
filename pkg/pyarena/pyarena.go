package pyarena

import (
	_ "embed"
)

//go:embed check_roms.py
var CheckRoms string

//go:embed list_roms.py
var ListRoms string
