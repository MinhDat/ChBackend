package main

import (
	helper "ChGo/helpers"
	seed "ChGo/seeds"
)

func main() {
	helper.Init()
	seed.User()
	seed.Category()
	helper.CloseDB()
}
