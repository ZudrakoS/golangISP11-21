package main

import (
 "fmt"
 "sort"
)

func main() {
 months := []string{"март", "февраль", "январь", "апрель", "май"}

 order := map[string]int{
  "январь":   1,
  "февраль":  2,
  "март":     3,
  "апрель":   4,
  "май":      5,
  "июнь":     6,
  "июль":     7,
  "август":   8,
  "сентябрь": 9,
  "октябрь":  10,
  "ноябрь":   11,
  "декабрь":  12,
 }

 sort.Slice(months, func(i, j int) bool {
  return order[months[i]] < order[months[j]]
 })

 fmt.Println("Отсортированные месяцы:")
 for _, month := range months {
  fmt.Println(month)
 }
}
