package main

import "CSMSite.Backend/Infrastructure"

func main() {
	db := Infrastructure.NewDb()
	r := Infrastructure.NewRouting(db)
	r.Run()
}
