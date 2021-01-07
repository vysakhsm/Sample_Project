package main

import ("fmt"
		"net/http"
		//"strconv"
		)

func main(){
	http.HandleFunc("/", Calculate)
	http.ListenAndServe(":8080",nil)
}
func Calculate(w http.ResponseWriter, r *http.Request){
	var number1 string
	var number2 string
	//var number3 int
	//var number4 int
	if r.URL.Path != "/"{
		http.Error(w, "Page Not found", http.StatusNotFound)
		return
	}
	switch r.Method{
	case "GET": 
		http.ServeFile(w, r, "AddPage.html")
		fmt.Fprintf(w,"This is calculator")
	case "POST":
		fmt.Fprintf(w,"This is calculator after POST\n ")
		number1 = r.FormValue("number1")
		number2 = r.FormValue("number2")
		fmt.Fprintf(w, "Number 1 = %v\n", number1)
		fmt.Fprintf(w, "Number 2 = %v\n", number2)
		//if number3, err := strconv.ParseInt(number1,6,12); err != nil{
		//	fmt.Fprintf(w,"Error occured ")
		//}
		//number4 = number3
		fmt.Fprintf(w, "Sum  = %v \n", number1 + number2)
	}
}