package main

import "fmt"

func main() {

	var name string
	var preini int
	var precio int
	var resp string
	fmt.Println("Bienvenido a Subastas JyD. Por favor ingrese el nombre del producto a subastar")
	fmt.Scanf("%s", &name)
	fmt.Println("Por favor ingrese el Valor inicial del producto")
	fmt.Scanf("%d", &preini)
	fmt.Println("Inicia la subasta de", name, "por un valor de", preini)

	fmt.Println("¿Alguien da más?")
	fmt.Println("Digite la palabra si para continuar o la palabra no para terminar la subasta")
	fmt.Scanf("%s", &resp)

	switch resp {
	case "si":
		fmt.Println("Recuerde que el nuevo valor debe ser superior al valor inicial y al valor aceptado anteriormente.\nPor favor ingrese el nuevo valor")
		fmt.Scanf("%d", &precio)
		if precio > preini {
			fmt.Println("Valor aceptado. Ahora", name, "se esta subastando por", precio)
			precio = preini
		} else {
			fmt.Println("Este valor no puede ser aceptado ya que es menor al precio inicial y al valor aceptado anteriormente.")
		}
	case "no":
		fmt.Println("Subasta terminada.", name, "Fue vendido por un valor de", preini)

	default:
		fmt.Println("Lo siento, no reconocimos su respuesta. Por favor intente de nuevo")
	}

}
