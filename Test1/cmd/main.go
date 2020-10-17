package main

import (
	"github/M1KUS3Q/test1/package1"
)

func main() {
	//w visual studio pan zaznaczy wszystko i ctrl + k + u to panu odkomentuje
	//to wszystko są testy czy wszystko ładnie działa, pusty wiersz pomiedzy mandatory a optional
	//spoiler : wszystko raczej ładnie działa
	//i do zadania 3 w optional dodałem możliwość sprawdzania słów po spacji, żebym mógł w plik txt wkleić skrypt shreka i żeby zadziałało

	// fmt.Println(package1.Suma(1, 2))
	// s1 := []int{2, 3, 4}
	// s2 := []int{5, 6, 7}
	// fmt.Println(package1.Concatenate(s1, s2))
	// fmt.Println(package1.Is18(16))
	// fmt.Println(package1.PrintDividedBy4(43))
	// s3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	// fmt.Println(package1.FilterEven(s3))
	// ussr := package1.NewUser("lol", "lol2", 34)
	// fmt.Println(ussr.ToString())
	// ussr.ChangeAge(3)
	// fmt.Println(ussr.ToString())

	//fmt.Println(package1.StringLength(str))
	// lol := package1.StringLength("")
	// fmt.Println(lol)
	m, _ := package1.CheckForRepeats("cmd/lol.txt")
	// fmt.Println(m)
	package1.SaveMapToFile("lol2.txt", m)
	package1.CreateServer()

	// fmt.Print(package1.ReadHttp("https://github.com"))

}
