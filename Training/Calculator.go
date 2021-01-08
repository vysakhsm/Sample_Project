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
	var string1 string
	var string2 string
	//var number3 int
	//var number4 int
	if r.URL.Path != "/"{
		http.Error(w, "Page Not found", http.StatusNotFound)
		return
	}
	switch r.Method{
	case "GET": 
		http.ServeFile(w, r, "AddPage.html")
		fmt.Fprintf(w,"This is String")
	case "POST":
		fmt.Fprintf(w,"This is after POST\n ")
		string1 = r.FormValue("string1")
		string2 = r.FormValue("string2")
		fmt.Fprintf(w, "String 1 = %v\n", string1)
		fmt.Fprintf(w, "String 2 = %v\n", string2)
		//if number3, err := strconv.ParseInt(number1,6,12); err != nil{
		//	fmt.Fprintf(w,"Error occured ")
		//}
		//number4 = number3
		fmt.Fprintf(w, "Appended String  = %v \n", string1 + string2)
	}
}
