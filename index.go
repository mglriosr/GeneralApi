package main

import ("net/http")

func main(){
	// Rutas
	http.HandleFunc("/", homeHandler)

	//Iniciar servidor
	println("Inicio del servidor")
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hola mundo"))
}