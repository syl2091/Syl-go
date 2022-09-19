package main

//排序

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
