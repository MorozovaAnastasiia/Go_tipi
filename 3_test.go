package main

import (
    "sort"
    "testing"
)

func TestUserSorting(t *testing.T) {
    users := []User{
        {"Nastya", 31},
        {"Andrey", 25},
        {"Natasha", 28},
    }

    t.Run("SortByName", func(t *testing.T) {
        sort.Sort(ByName(users))
        expected := []string{"Andrey", "Natasha", "Nastya"}
        for i, name := range expected {
            if users[i].Name != name {
                t.Errorf("Позиция %d: ожидалось %s, получено %s",
                    i, name, users[i].Name)
            }
        }
    })

    t.Run("SortByAge", func(t *testing.T) {
        sort.Sort(ByAge(users))
        expected := []int{25, 28, 31}
        for i, age := range expected {
            if users[i].Age != age {
                t.Errorf("Позиция %d: ожидалось %d, получено %d",
                    i, age, users[i].Age)
            }
        }
    })
}
