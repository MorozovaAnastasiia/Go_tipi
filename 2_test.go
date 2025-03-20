package main

import (
    "testing"
)

func TestProcessValue(t *testing.T) {
    tests := []struct {
        input  interface{}
        expect string
    }{
        {"test", "TEST"},
        {42, "84"},
        {3.14, "6.28"},
        {struct{}{}, "Неизвестный тип"},
    }

    for _, tt := range tests {
        t.Run(fmt.Sprintf("%T", tt.input), func(t *testing.T) {
            // Перехватываем вывод
            rescue := captureOutput()
            ProcessValue(tt.input)
            output := rescue()

            if !strings.Contains(output, tt.expect) {
                t.Errorf("Ожидалось %q, получено %q", tt.expect, output)
            }
        })
    }
}

// Вфункция для перехвата вывода
func captureOutput() func() string {
    rescue := fmt.Println
    var buf []byte
    fmt.Println = func(a ...interface{}) (int, error) {
        buf = append(buf, fmt.Sprintln(a...)...)
        return len(buf), nil
    }
    return func() string {
        fmt.Println = rescue
        return string(buf)
    }
}
