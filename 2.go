package main
 
import (
    "fmt"
    "reflect"
    "strings"
)
// ProcessValue обрабатывает значение в зависимости от его типа:
// - для строк: преобразует в верхний регистр
// - для чисел: умножает на 2
// - для остальных типов: выводит сообщение о неизвестном типе
func ProcessValue(v interface{}) {
    switch val := v.(type) {
    case string:
        fmt.Println(strings.ToUpper(val))
    default:
        rv := reflect.ValueOf(v)
        switch rv.Kind() {
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
            reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
            reflect.Float32, reflect.Float64:
            n := rv.Convert(reflect.TypeOf(float64(0))).Float()
            fmt.Println(n*2)
        default:
            fmt.Println("Неизвестный тип")
        }
    }
}
