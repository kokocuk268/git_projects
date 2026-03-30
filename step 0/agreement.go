package main

import (
	"fmt"
	"time"
)

func main() {
	var dateOfContract string
	fmt.Scanln(&dateOfContract)

	contractDate, err := time.Parse("02.01.2006", dateOfContract)
	if err != nil{
		fmt.Println("error of date format!")
		return
	}
	newDateOfContract := contractDate.AddDate(0,0,15)
	resultDate := newDateOfContract.Format("02.01.2006")
	
	var name,surname,middlename string
	fmt.Scanln(&name); fmt.Scanln(&surname); fmt.Scanln(&middlename)

	var salaryForFirstWork,salaryForSecondWork,salaryForThirdWork float64
	fmt.Scanln(&salaryForFirstWork); fmt.Scanln(&salaryForSecondWork); fmt.Scanln(&salaryForThirdWork)

	sumFullSalary := salaryForFirstWork + salaryForSecondWork + salaryForThirdWork
	rub := int(sumFullSalary)
	kop := int((sumFullSalary - float64(rub)) * 100)

	fmt.Printf("Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.\n" +
	"Дата подписания договора: %s. Просим вас подойти в офис в любое удобное для вас время в этот день.\n" +
	"Общая сумма выплат составит %d руб. %d коп.\n" +
	"\n" +
	"С уважением,\n" +
	"Гл. бух. Иванов А.Е." ,	
	surname,
	name,
	middlename,
	resultDate,
	rub,
	kop,
	)
}