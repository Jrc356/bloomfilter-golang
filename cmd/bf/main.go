package main

import (
	bloomfilter "jrc356.com"
)

func main() {
	bf := bloomfilter.Bloomfilter{}.New(100)
	println(bf.Hash("hello"))
	println(bf.Hash("world"))
	println(bf.Hash("from"))
	println(bf.Hash("bloomfilter"))
}
