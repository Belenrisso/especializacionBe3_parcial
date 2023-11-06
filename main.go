package main

import (
	"log"
	"time"
	"fmt"
	"github.com/Belenrisso/especializacionBe3-parcial/internal/tickets"
)

func main() {
	//total, err := tickets.GetTotalTickets("Brazil")
	var archivo string = "tickets.csv"
	listado, e := tickets.ObtenerDatos(archivo)
	defer func(){
		if e != nil {
			log.Fatal(e)
		}
		}()
	

	var destino string
	fmt.Println("Por favor ingrese un destino: ")
	fmt.Scan(&destino)
	var franjaHoraria string
	fmt.Println("\nIngrese una franja horaria (Disponibles: Madrugada, Mañana, Tarde, Noche):")
	fmt.Scan(&franjaHoraria)
	var destinoPorcentaje string
	fmt.Println("\nDigite el destino para saber el porcentaje de pasajeros que viajó en el día")
	fmt.Scan(&destinoPorcentaje)


	canalTicket := make(chan tickets.Ticket)
	defer close(canalTicket)
	canalError := make(chan error)
	defer close(canalError)


	//GO ROUTINE - OBTIENE EL TOTAL DE CUANTAS PERSONAS VIAJAN A UN PAIS DETERMINADO
	go func(chan tickets.Ticket, chan error) {
		total, err := tickets.ObtenerTotalTicketsDestino(destino, &listado)
		if err != nil {
			canalError <- err
			return
		}
		fmt.Printf("\nLa cantidad total de tickets para %s es %d", destino, total)
		
		canalTicket <- tickets.Ticket{}
		}(canalTicket, canalError)//(destino, listado)

	//GO ROUTINE - OBTIENE CUANTAS PERSONAS VIAJAN EN UNA FRANJA HORARIA
		go func(chan tickets.Ticket,chan error) {
		total, err := tickets.ObtenerTicketsFranjaHoraria(franjaHoraria, &listado)
		if err != nil {
			canalError <- err
			return
		}
		fmt.Printf("\nLa cantidad total de tickets para la %s es %d\n", franjaHoraria, total)
		canalTicket <- tickets.Ticket{}
	}(canalTicket, canalError)

	//GO ROUTINE - OBTIENE EL PORCENTAJE DE PERSONAS QUE VIAJAN A UN PAIS DETERMINADO EN UN DIA
	go func(chan tickets.Ticket, chan error) {
		porcentaje, err := tickets.ObtenerPromedioDestinos(destinoPorcentaje, &listado)
		if err != nil {
			canalError <-err
			return
		}
		fmt.Printf("\nEl porcentaje total de tickets para el destino %s es %.2f", destinoPorcentaje, porcentaje)
		canalTicket <- tickets.Ticket{}
	}(canalTicket, canalError)

	time.Sleep(1 * time.Second)


	// Listado completo
	// go func() {
	// 	Listado, e := tickets.ObtenerDatos("tickets.csv")
	// 	if e != nil {
	// 		fmt.Println(e)
	// 	}
	// 	fmt.Println(Listado)
	// }()
	// time.Sleep(1 * time.Second)



}
 
 
	
 



