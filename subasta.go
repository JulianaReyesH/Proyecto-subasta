package main

import "fmt"

func main() {

	var name string
	var preini int
	var precio_oferta int
	var precio_actual int
	var resp string
	fmt.Println("Bienvenido a Subastas JyD. Por favor ingrese el nombre del producto a subastar")
	fmt.Scanf("%s", &name)
	fmt.Println("Por favor ingrese el Valor inicial del producto")
	fmt.Scanf("%d", &preini)

	precio_actual = preini
	fmt.Println("Inicia la subasta de", name, "por un valor de", precio_actual)

	for {
		fmt.Println("El valor actual es:", precio_actual)
		fmt.Println("¿Alguien da más?")

		fmt.Println("Digite la palabra si para continuar o la palabra no para terminar la subasta")
		fmt.Scanf("%s", &resp)

		if resp == "si" {

			fmt.Println("Recuerde que el nuevo valor debe ser superior al valor inicial y al valor aceptado anteriormente.\nPor favor ingrese el nuevo valor")
			fmt.Scanf("%d", &precio_oferta)
			if precio_oferta > precio_actual {
				precio_actual = precio_oferta
				fmt.Println("Valor aceptado. Ahora", name, "se esta subastando por", precio_actual)
			} else {
				fmt.Println("Este valor no puede ser aceptado ya que es menor al precio inicial y al valor aceptado anteriormente.")
			}
		} else {
			fmt.Println("Subasta terminada.", name, "Fue vendido por un valor de", precio_actual)
			break
		}
	}
}
