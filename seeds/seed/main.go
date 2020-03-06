package main

import (
	helper "ChGo/helpers"
	seed "ChGo/seeds"
)

func main() {
	helper.Init()
	seed.Auth()
	seed.User()
	seed.Category()
	helper.CloseDB()
}
