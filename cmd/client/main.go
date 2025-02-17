package main

import (
	"fmt"
	"log"

	"train-ticket-booking/client"
)

func main() {
	fmt.Println("🚀 Starting Train Ticket Booking Client...")

	// Initialize gRPC client
	ticketClient, err := client.NewTrainTicketClient()
	if err != nil {
		log.Fatalf("Could not create client: %v", err)
	}
	defer ticketClient.CloseConnection()

	// Perform operations
	fmt.Println("\n--- Purchasing Ticket ---")
	ticketClient.PurchaseTicket("CityA", "CityB", "John", "Doe", "john@example.com", 100.0)

	fmt.Println("\n--- Purchasing Ticket ---")
	ticketClient.PurchaseTicket("CityA", "CityB", "John", "Doe", "john1@example.com", 100.0)

	fmt.Println("\n--- Purchasing Ticket ---")
	ticketClient.PurchaseTicket("CityA", "CityB", "John", "Doe", "john2@example.com", 100.0)

	fmt.Println("\n--- Purchasing Ticket ---")
	ticketClient.PurchaseTicket("CityA", "CityB", "John", "Doe", "joh3@example.com", 100.0)

	fmt.Println("\n--- Purchasing Ticket ---")
	ticketClient.PurchaseTicket("CityA", "CityB", "John", "Doe", "joh4@example.com", 100.0)

	fmt.Println("\n--- Purchasing Ticket ---")
	ticketClient.PurchaseTicket("CityA", "CityB", "John", "Doe", "john5@example.com", 100.0)

	fmt.Println("\n--- Getting Receipt ---")
	ticketClient.GetReceipt("john@example.com")

	fmt.Println("\n--- Getting Users in Section A ---")
	ticketClient.GetUsersBySection("A")

	fmt.Println("\n--- Getting Users in Section B ---")
	ticketClient.GetUsersBySection("B")

	fmt.Println("\n--- Modifying Seat ---")
	ticketClient.ModifySeat("john@example.com", "B3", "B")

	fmt.Println("\n--- Removing User ---")
	ticketClient.RemoveUser("john@example.com")

	fmt.Println("\n--- Getting Users in Section A (After Removal) ---")
	ticketClient.GetUsersBySection("A")
}
