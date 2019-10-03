package main

import (
	"fmt"
	"strconv"
	"time"
)

var currentId int

var lunches Lunch

// Give us some seed data
func init() {
	//semena tier 1
	var semanaT1 []Week
	semanaT1 = append(semanaT1, Week{Id: 1, FoodTime: "T2"})
	semanaT1 = append(semanaT1, Week{Id: 2, FoodTime: "T1"})
	semanaT1 = append(semanaT1, Week{Id: 3, FoodTime: "T2"})
	semanaT1 = append(semanaT1, Week{Id: 4, FoodTime: "T1"})
	//semena tier 2
	var semanaT2 []Week
	semanaT2 = append(semanaT2, Week{Id: 1, FoodTime: "T1"})
	semanaT2 = append(semanaT2, Week{Id: 2, FoodTime: "T2"})
	semanaT2 = append(semanaT2, Week{Id: 3, FoodTime: "T1"})
	semanaT2 = append(semanaT2, Week{Id: 4, FoodTime: "T2"})
	//semena tier 3
	var semanaT3 []Week
	semanaT3 = append(semanaT3, Week{Id: 1, FoodTime: "T1"})
	semanaT3 = append(semanaT3, Week{Id: 2, FoodTime: "T1"})
	semanaT3 = append(semanaT3, Week{Id: 3, FoodTime: "T2"})
	semanaT3 = append(semanaT3, Week{Id: 4, FoodTime: "T2"})

	RepoCreateLunch(Lunch{Name: "Alfredo", Semana: semanaT1})
	RepoCreateLunch(Lunch{Name: "Juan", Semana: semanaT1})
	RepoCreateLunch(Lunch{Name: "AlvaroC", Semana: semanaT2})
	RepoCreateLunch(Lunch{Name: "Roberto", Semana: semanaT2})
	RepoCreateLunch(Lunch{Name: "Cesar", Semana: semanaT2})
	RepoCreateLunch(Lunch{Name: "Jonay", Semana: semanaT3})
	RepoCreateLunch(Lunch{Name: "Amisadai", Semana: semanaT1})
	RepoCreateLunch(Lunch{Name: "Jorge", Semana: semanaT2})
	RepoCreateLunch(Lunch{Name: "LuisC", Semana: semanaT2})
	//RepoCreateLunch(Lunch{Name: "Marino", Semana: semana})
	//RepoCreateLunch(Lunch{Name: "Pedro", Semana: "1", FoodTime: "T1"})

}

func RepoFindLunch(idS string) Lunch {
	id, err := strconv.Atoi(idS)
	if err == nil {
		fmt.Println("Error")
	}
	for _, t := range lunches {
		if t.Id == id {
			return t
		}
	}
	// return empty Lunch if not found
	return Lunch{}
}

func RepoFindEmployeeByName(name string) LunchResponse {
	dt := time.Now()
	var idSemana int
	var lunchT LunchResponse

	//semana 1
	if dt.Day() < 30 {
		idSemana = 4
	}
	if dt.Day() < 22 {
		idSemana = 3
	}
	if dt.Day() < 15 {
		idSemana = 2
	}
	if dt.Day() < 8 {
		idSemana = 1
	}

	for _, t := range lunches {
		if t.Name == name {
			for _, s := range t.Semana {
				if s.Id == idSemana {
					lunchT = LunchResponse{Hora: s.FoodTime}

					return lunchT
				}
			}
		}
	}
	// return empty Lunch if not found

	return LunchResponse{}
}

func RepoCreateLunch(t Lunch) Lunch {
	currentId += 1
	t.Id = currentId
	lunches = append(lunches, t)
	return t
}

func RepoDestroyLunch(id int) error {
	for i, t := range lunches {
		if t.Id == id {
			lunches = append(lunches[:i], lunches[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Lunch with id of %d to delete", id)
}
