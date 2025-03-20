// Package reflection содержит утилиты для работы с рефлексией.
package main

import (
    "fmt"
    "reflect"
)

// CallMethod динамически вызывает метод с заданным именем для объекта.
// Возвращает результаты вызова метода или ошибку, если метод не найден
// или аргументы не соответствуют ожидаемым типам.
func CallMethod(obj interface{}, methodName string, args ...interface{}) ([]reflect.Value, error) {
    val := reflect.ValueOf(obj)
    method := val.MethodByName(methodName)
    if !method.IsValid() {
        return nil, fmt.Errorf("метод %s не найден", methodName)
    }
 
    methodType := method.Type()
    in := make([]reflect.Value, len(args))
    for i, arg := range args {
        argType := methodType.In(i)
        argVal := reflect.ValueOf(arg)
        if !argVal.Type().ConvertibleTo(argType) {
            return nil, fmt.Errorf("неверный тип аргумента %d", i)
        }
        in[i] = argVal.Convert(argType)
    }
 
    return method.Call(in), nil
}

// Test структура для демонстрации работы CallMethod
type Test struct{}

// Add возвращает сумму двух целых чисел
func (t Test) Add(a, b int) int { return a + b }

// Greet возвращает приветствие с заданным именем
func (t Test) Greet(name string) string { return "Hello " + name }
