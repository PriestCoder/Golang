package GolangBasics

import ("fmt"
   "net/http")

func main(){

	fmt.Println("Hello , Maddy ")

	resp , err := http.Get("https://pkg.go.dev/fyne.io/fyne/v2#readme-shipping-the-fyne-toolkit")
	if err!=nil {
		fmt.Println("Error :",err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("HTTP Rsponse Status :",resp.Status)




}