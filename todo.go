package main

type Week struct {
	Id       int    `json:"id"`
	FoodTime string `json:"foodtime"`
}

type Weeks []Week

type Lunch struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Semana []Week `json:"semana"`
}

type Lunch []Lunch

type LunchResponse struct {
	Hora string `json:"hora"`
}
