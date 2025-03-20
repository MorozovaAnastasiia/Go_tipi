// Package sorting демонстрирует сортировку пользователей по разным критериям.
package main

import "sort"

// User представляет информацию о пользователе
type User struct {
    Name string
    Age  int
}

// ByName реализует sort.Interface для сортировки пользователей по имени
type ByName []User

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

// ByAge реализует sort.Interface для сортировки пользователей по возрасту
type ByAge []User

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
