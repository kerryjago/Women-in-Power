package main

import "fmt"

func main() {
	// Print message to console
	fmt.Println("This program celebrates women in power.")
	
	// Define women in power
	const (
		EleanorRoosevelt = "Eleanor Roosevelt"
		OprahWinfrey = "Oprah Winfrey"
		AngelaMerkel = "Angela Merkel"
		JaneGoodall = "Jane Goodall"
		MalalaYousafzai = "Malala Yousafzai"
		MichelleObama = "Michelle Obama"
		RosaParks = "Rosa Parks"
		MarieCurie = "Marie Curie"
		CondoleezzaRice = "Condoleezza Rice"
		AungSanSuuKyi = "Aung San Suu Kyi"
	)

	// Create array of women in power
	womenInPower := [10]string{
		EleanorRoosevelt,
		OprahWinfrey,
		AngelaMerkel,
		JaneGoodall,
		MalalaYousafzai,
		MichelleObama,
		RosaParks,
		MarieCurie,
		CondoleezzaRice,
		AungSanSuuKyi,
	}

	// Iterate through array and print message for each woman
	for _, woman := range womenInPower {
		fmt.Printf("Let us always remember and celebrate %s and their many achievements. \n\n", woman)
	}
}