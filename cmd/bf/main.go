package main

import (
	"fmt"

	bloomfilter "jrc356.com"
)

func main() {
	bf := bloomfilter.NewSimpleBloomfilter(1000)

	strings_to_insert := []string{
		"Lorem",
		"ipsum",
		"dolor",
		"sit",
		"amet",
		"consectetur",
		"adipiscing",
		"elit",
		"Integer",
		"nec",
		"odio",
		"Praesent",
		"cursus",
		"ante",
		"dapibus",
		"diam",
		"Sed",
		"nisi",
		"Nulla",
		"quis",
		"sem",
		"at",
		"nibh",
		"elementum",
		"imperdiet",
		"Duis",
		"sagittis",
		"libero",
	}

	strings_to_check := []string{
		"Praesent",
		"hello",
		"mom",
		"and",
		"imperdiet",
		"dad",
		"Integer",
	}

	for _, s := range strings_to_insert {
		bf.Insert(s)
	}

	for _, s := range strings_to_check {
		fmt.Printf("%s in filter? %v\n", s, bf.Check(s))
	}
}
