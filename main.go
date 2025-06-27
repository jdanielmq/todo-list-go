package main

import (
	"bufio"
	"fmt"
	"os"
)

type Persona struct {
	nombre string
	edad   int
	correo string
}

func (p *Persona) diHola() {
	fmt.Println("Hola, mi nombre es : ", p.nombre)
}

type Tarea struct {
	nombre     string
	desc       string
	completado bool
}

type ListaTareas struct {
	tareas []Tarea
}

func main() {

	//declaracion de una matriz
	var matriz [5]int
	fmt.Println(matriz)
	fmt.Println("###############################")

	//declaracion de una matriz inicializada
	var matriz_uno = [5]int{10, 20, 30, 40, 50}
	fmt.Println(matriz_uno)
	fmt.Println("###############################")

	//declaracion de una matriz inicializada con una cantidad elemento no conicida
	var matriz_dos = [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	matriz_dos[1] = 100
	matriz_dos[2] = 200
	for i := 0; i < len(matriz_dos); i++ {
		fmt.Println(matriz_dos[i])
	}
	fmt.Println("###############################")

	//recorrer la matriz y retornar 2 valores (index y valor)
	for index, value := range matriz_dos {
		fmt.Printf("index: %d y valor: %d \n", index, value)
	}
	fmt.Println("###############################")

	//declaracion de matriz bidimensional
	var matriz_tres = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matriz_tres)
	fmt.Println("###############################")

	// vamos a declarar y ver las rebanadas (slice)
	var matriz_cuatro []int
	fmt.Println(matriz_cuatro)
	fmt.Println("###############################")

	// inicializamos una rebanada
	matriz_cinco := []string{"Domingo", "Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado"}
	fmt.Println(matriz_cinco)
	fmt.Println("###############################")

	//vamos a sacar una rebanada de los dias de la semana
	diaRebanada := matriz_cinco[0:5]
	fmt.Println(diaRebanada)
	fmt.Println("###############################")

	//mostrar cantidad de la rebanada (len()) y capacidad de la rebanada (cap())
	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada))
	fmt.Println("###############################")

	//vamos agregar un elemento a la rebanada y vamos imprimir el punto anterior
	diaRebanada = append(diaRebanada, "Viernes")
	fmt.Println(diaRebanada)
	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada))
	fmt.Println("###############################")

	//se puede agregar mas elemento que la capacidad dice, si se pasa,
	//la rebanada dianmicamente se duplica la capacidad.
	diaRebanada = append(diaRebanada, "Viernes", "Sabado", "Otro Dia")
	fmt.Println(diaRebanada)
	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada))
	fmt.Println("###############################")

	//cuando queremos eliminar un elemnto de la rebanada se hace de la siguiente manera
	diaRebanada = append(diaRebanada[:2], diaRebanada[3:]...) // diaRebanada[:2]= recoje desde el inicio hasta el numero puesto menos 1
	fmt.Println(diaRebanada)
	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada))
	fmt.Println("###############################")

	//crear una rebanada con tipo de dato, largo y capacidad
	nombres := make([]string, 5, 10)
	nombres[0] = "Juan Daniel"
	fmt.Println(nombres)
	fmt.Println(len(nombres))
	fmt.Println(cap(nombres))
	fmt.Println("###############################")

	//vamos a ver como se utiliza la funcionalidad copy, en este caso veremos con las rebanadas
	rebanada1 := []int{1, 2, 3, 4, 5}
	rebanada2 := make([]int, 5)
	fmt.Println(rebanada1)
	fmt.Println(rebanada2)
	fmt.Println("###############################")
	//rebanada2 recibira las copia de la rebanada1. y la funcion de vuelve la cantidad de elemetos copiados
	copy(rebanada2, rebanada1)
	fmt.Println(rebanada1)
	fmt.Println(rebanada2)
	fmt.Println("###############################")

	//ahora vamos aprender como utlizar las Map.
	colors := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}
	fmt.Println(colors)
	fmt.Println(colors["rojo"]) // obtener el dato por su clave (key)

	// agregando un registro al mapa.
	colors["negro"] = "#000000"
	fmt.Println(colors)
	fmt.Println("###############################")

	valor := colors["rojo"]
	fmt.Println(valor)

	// verificar si la variable es vacia
	valor, ok := colors["rojo"]
	fmt.Println(valor, ok)

	valor1, ok1 := colors["amarillo"]
	fmt.Println(valor1, ok1)
	fmt.Println("###############################")

	// vamos utilizar la verificacion con el ok1
	if ok {
		fmt.Println("si existe la clave OK")
	} else {
		fmt.Println("no existe este clave OK")
	}
	if ok1 {
		fmt.Println("si existe la clave OK1")
	} else {
		fmt.Println("no existe este clave OK1")
	}

	// vamos a simplicar la evalucion del if con el OK3, tomamos el ejemplo de OK1
	if valor3, ok3 := colors["amarillo"]; ok3 {
		fmt.Println(valor3)
	} else {
		fmt.Println("no existe este clave OK3")
	}
	fmt.Println("###############################")
	// vamos eliminar un elemento de uin map
	delete(colors, "azul") //enviamos el map y la clave a eliminar
	fmt.Println(colors)
	fmt.Println("###############################")

	// vamos a recorrer el map y vamos a capturar la clave y valor
	for clave, valorMap := range colors {
		fmt.Printf("clave: %s y valor: %s \n", clave, valorMap)
	}
	fmt.Println("###############################")

	//ahora vamos a trabajar con estruturas, esto se crea arriba del metodo principal main()
	var p Persona
	p.nombre = "Juan Daniel"
	p.edad = 47
	p.correo = "jdannymq@gmail.com"
	fmt.Println(p)
	fmt.Println("###############################")

	// otra manera de instanciar la clase Persona
	pp := Persona{"Juan Daniel", 44, "jdannymq@gmail.com"}
	fmt.Println(pp)
	fmt.Println("###############################")

	// declarando variable y colocando punteros
	var x int = 10
	var pu *int = &x
	fmt.Println(&x)
	fmt.Println(pu)
	fmt.Println("###############################")

	//vamos a crear un editar para modificar la variable.
	editar(&x)
	fmt.Println(x)
	fmt.Println("###############################")

	// vamos s trabajar con punteros y una estructura
	ppp := Persona{"Juan daniel", 44, "sdfghjkhgf@gmasil.com"}
	ppp.diHola()
	fmt.Println("###############################")

	//aca vamos a trabajar con el ejercicio de session.
	lista := ListaTareas{}
	/*creamos una instancia de bufio para la entrada de datos*/
	leer := bufio.NewReader(os.Stdin)

	for {
		var opcion int
		fmt.Println("Seleccione una opcion : \n",
			"1.- Agregar Tarea \n",
			"2.- Marcar tarea como completada \n",
			"3.- Editar tarea \n",
			"4.- Eliminar tarea \n",
			"5.- Salir")

		fmt.Print("Ingrese la oipcion: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Print("Ingrese nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la descripcion de la tarea : ")
			t.desc, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Tarea agregada correctamente")
		case 2:
			var index int
			fmt.Print("Ingrese el index de la tarea que desea marcar como completada :")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println("Tarea marcada como completada correctamente.")
		case 3:
			var index int
			var t Tarea
			fmt.Print("Ingrese el indice de la tarea que sea actualizar: ")
			fmt.Scanln(&index)
			fmt.Print("Ingrese nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la descripcion de la tarea : ")
			t.desc, _ = leer.ReadString('\n')
			lista.editarTarea(index, t)
			fmt.Println("Tarea actualizada correctamente")
		case 4:
			var index int
			fmt.Print("Ingrese el indice de la tarea que desea eliminar: ")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Println("Tarea eliminada correctamente")
		case 5:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opcion no valida")
		}
		//listar todas las tareas
		fmt.Println("Lista de Tareas:")
		fmt.Println("###############################")
		for i, t := range lista.tareas {
			fmt.Printf("%d. %s - %s - completado: %t \n", i, t.nombre, t.desc, t.completado)
		}
		fmt.Println("###############################")
	}
}
func editar(x *int) {
	*x = 20
}

// metodo para agregar una tarea
func (l *ListaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

// metodo para marcar una tarea como completada
func (l *ListaTareas) marcarCompletado(index int) {
	l.tareas[index].completado = true
}

// vamos editar una tarea
func (l *ListaTareas) editarTarea(index int, t Tarea) {
	l.tareas[index] = t
}

// vamos a eliminar una tarea
func (l *ListaTareas) eliminarTarea(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}
