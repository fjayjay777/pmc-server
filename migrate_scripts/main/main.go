package main

import migrate "pmc_server/migrate_scripts"

func main() {
	err := migrate.GenerateCourseAssociatedTags()
	if err != nil {
		panic(err)
	}
}
