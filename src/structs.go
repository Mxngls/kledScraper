package main

type Result []Ideom

type Ideom struct {
	Id            int
	Word          string
	HomonymNumber int
	Hangul        string
	Type          string
	Senses        []sense
}

type sense struct {
	Translation  string
	Definition   string
	KrDefinition string
	Example      string
	References   []reference
}

type reference struct {
	Type   string
	Values []value
}

type value struct {
	Id     string
	Hangul string
}
