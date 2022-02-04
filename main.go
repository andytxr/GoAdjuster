package GoAdjuster

import (
	"fmt"
)

func main (){

	fmt.Println("How many employeers? ")
	var qtdE int
	fmt.Scanln(&qtdE)

	if qtdE < 1 || qtdE > 100{
		for{
			fmt.Println("The quantity of employeers needs to be bigger than 1 and minor than 100")
			fmt.Println("How many employeers? ")
			fmt.Scanln(&qtdE)
			if qtdE >= 1 || qtdE < 100{
				break
			}
		}
	}

	for i:=0; i<qtdE;i++{
		fmt.Println("Enter with the value per hour: ")
		var valueHour float64
		fmt.Scanln(&valueHour)

		if valueHour < 2 || valueHour > 100{
			for{
				fmt.Println("Error! The value per hour needs to be bigger than 2 and minor than 100")
				fmt.Println("Enter with the value per hour: ")
				fmt.Scanln(&valueHour)
				if valueHour >= 2 || valueHour <= 100{
					break
				}
			}
		}

		fmt.Println("Enter with the days quantity: ")
		var qtdD int
		fmt.Scanln(&qtdD)
		if qtdD < 1 || qtdD > 31{
			for{
				fmt.Println("Error! The quantity of days needs to be bigger than 1 and minor than 31")
				fmt.Println("Enter with the days quantity: ")
				fmt.Scanln(&qtdD)
				if qtdD >= 1 || qtdD <= 31{
					break
				}
			}
		}

		fmt.Println("Enter with the readjust percentage in the wage: ")
		var perc float64
		fmt.Scanln(&perc)
		if perc < 10 || perc > 50{
			fmt.Println("Error! The percentage of readjust needs to be bigger than 10 and minor than 50")
			fmt.Println("Enter with the readjust percentage in the wage: ")
			fmt.Scanln(&perc)
			if perc >= 10 || perc <= 50{
				break
			}
		}

		qtdDF:=float64(qtdD)

		valueHour = valueHour+valueHour*(perc/100)
		var wage = valueHour*qtdDF

		fmt.Println("Your wage is: %wage", wage)

	}

}


