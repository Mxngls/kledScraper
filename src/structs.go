package main

type Result struct {
	Alpha           int
	Frequency       int
	Id              int
	Hangul          string
	HomonymNumber   int
	Hanja           string
	TypeKr          string
	TypeEng         string
	Pronounciation  string
	Audio           string
	Level           int
	Inflections     string
	InflectionLinks []inflectionLink
	Senses          []sense
}

type sense struct {
	Translation  string
	Definition   string
	KrDefinition string
	Examples     []example
	Reference    []ref
}

type example struct {
	Type  string
	Value string
}

type ref struct {
	Type  string
	Value string
	Id    int
}

type inflectionLink struct {
	Id     int
	Hangul string
}

func InitSense() sense {
	return sense{}
}

func InitInflectionLinks() inflectionLink {
	return inflectionLink{}
}
