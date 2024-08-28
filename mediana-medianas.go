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

func buscarMediana(lst []int) int {
	tamañoLista:= len(lst)
	for j:=0; j < tamañoLista; j++ {
		for k:= j + 1; k < tamañoLista; k++ {
			if lst[j] > lst[k] {
				lst[j], lst[k] = lst[k], lst[j]
			}
		}
	}
	return lst[tamañoLista / 2]
}

func main() {
	lista:= []int{3, 8, 2, 1, 7, 9, 12, 11, 10, 5, 6, 4}
	sublistas:= particionar(lista)
	fmt.Println(sublistas)

	for _, sublista:= range sublistas {
		mediana:= buscarMediana(sublista)
		fmt.Println(mediana)
	}
}