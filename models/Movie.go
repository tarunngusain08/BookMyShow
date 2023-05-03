package models

type Movie struct {
	Id       int
	Name     string
	Duration int
	Language string
	Cast     []*Actor
}
