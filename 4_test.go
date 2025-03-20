package main

import (
    "reflect"
    "testing"
)

func TestCallMethod(t *testing.T) {
    obj := Test{}

    t.Run("ValidMethod", func(t *testing.T) {
        res, err := CallMethod(obj, "Add", 2, 3)
        if err != nil {
            t.Fatalf("Ошибка вызова метода: %v", err)
        }
        if res[0].Int() != 5 {
            t.Errorf("Ожидался результат 5, получено %d", res[0].Int())
        }
    })

    t.Run("InvalidMethod", func(t *testing.T) {
        _, err := CallMethod(obj, "Invalid", 0)
        if err == nil {
            t.Error("Ожидалась ошибка для несуществующего метода")
        }
    })

    t.Run("InvalidArgs", func(t *testing.T) {
        _, err := CallMethod(obj, "Add", "a", "b")
        if err == nil {
            t.Error("Ожидалась ошибка несоответствия типов аргументов")
        }
    })
}
