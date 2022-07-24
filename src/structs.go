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
	Proverbs        []proverb
	Senses          []sense
}

type sense struct {
	Translation  string
	Definition   string
	KrDefinition string
	Examples     []example
	Reference    []ref
}

type proverb struct {
	Hangul string
	Senses []sense
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

func InitProverb() proverb {
	return proverb{}
}

func InitRef() ref {
	return ref{}
}

func InitInflectionLinks() inflectionLink {
	return inflectionLink{}
}
