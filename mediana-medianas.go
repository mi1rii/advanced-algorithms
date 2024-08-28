package main

import ("fmt")

func particionar(lista []int) [][]int{
	var sublistas [][]int
	n:= len(lista)
	for j:= 0; j < n; j+=5 {
		end:= j + 5
		if end > n {
			end = n
		}
		sublista:= lista[j:end]
		sublistas = append(sublistas, sublista)
	}
	return sublistas
}

func main() {
	lista:= []int{3, 8, 2, 1, 7, 9, 12, 11, 10, 5, 6, 4}
	sublistas:= particionar(lista)
	fmt.Println(sublistas)
}