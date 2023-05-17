package models

type Theatre struct {
    ID           int
    Name         string
    Address      string
    City         *City
    Auditoriums  []*Auditorium
}
