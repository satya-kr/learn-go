package main

import "fmt"

func main() {
	fmt.Println("Struct in golang")

	akash := User{"Akash", "info@astergo.in", 25, true}
	fmt.Println(akash)

	fmt.Printf("\nName: %v\nEmail: %v\n", akash.Name, akash.Email)
	akash.GetStatus()
	akash.NewEmail()
	fmt.Printf("\nName: %v\nEmail: %v\n", akash.Name, akash.Email)
	akash.UpdateEmail()
	fmt.Printf("\nName: %v\nEmail: %v\n", akash.Name, akash.Email)

}

type User struct {
	Name, Email string
	Age         int
	Status      bool
}

func (u User) GetStatus() {
	fmt.Println("\nIs User active:", u.Status)
}

func (u User) NewEmail() {
	u.Email = "imsatya.kr@gmail.com"
	fmt.Println("New Email is: ", u.Email)
}

func (u *User) UpdateEmail() {
	u.Email = "imsatya.kr@gmail.com"
	fmt.Println("New Email is: ", u.Email)
}
