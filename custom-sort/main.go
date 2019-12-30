package main

import("fmt"; "sort")

type Person struct{
	Name string
	Age int
}

type ByName []Person

func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool{
	return ps[i].Name < ps[j].Name
}

func (ps ByName) Swap(i, j int){
	ps[i], ps[j] = ps[j], ps[i]
}

func main()  {
	// sort package allows you to sort arbitrary data
	// use it to sort your own data i.e Person above

	kids := []Person{
		{"Jill", 9},
		{"Sean", 10},
		{"James", 7},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)
}