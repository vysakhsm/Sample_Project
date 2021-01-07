package main

import ("fmt"
		"log"
		"net/http")

func MyFunction(w http.ResponseWriter, r *http.Request){
		if r.URL.Path != "/" {
			http.Error(w, "Page Not found", http.StatusNotFound)
			return
		}
		if r.Method != "GET"{
			http.Error(w,"Method Not Support\n", http.StatusNotFound)
			fmt.Fprintf(w, "Second Method Not Found\n")
		} else {
			http.ServeFile(w, r, "AddPage.html")
			fmt.Fprintf(w, "This is my home page\n")	
		}
		if r.Method == "POST"{
			fmt.Fprintf(w, "This is POST METHOD \n")
			// fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)	
			var number1 = r.FormValue("number1")
			var number2 = r.FormValue("number2")
			fmt.Fprintf(w, "Number 1 = %v\n", number1)
			fmt.Fprintf(w, "Number 2 = %v\n", number2)
			var number3 = number1 + number2
			fmt.Fprintf(w, "Sum  = %v \n", number3)
			if err := r.ParseForm(); err != nil{
				fmt.Fprintf(w, "ParseForm() error : %v\n",err)
				return
			}
		}
	}

func main() {
	http.HandleFunc("/", MyFunction)
	if err := http.ListenAndServe(":8080",nil); err != nil{
		log.Fatal(err)
	}	
}