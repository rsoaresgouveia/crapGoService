package main

import (
	"net/http"
	"player-leaderboard/core/rest"
	"player-leaderboard/database"
)

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/echo", rest.JsonBodytoJsonResponseexampleHandler) 	//Aceita json no body e retorna o json do body parseado
	mux.HandleFunc("/truck", rest.ReturnJsonExampleHandler) 				//Requisição vazia que retorna um json hardcoded
	mux.HandleFunc("/", rest.ParamsFromUrlExampleHandler)				//Requisição com parametros, retorno dos parametros impressos na pagina
	mux.HandleFunc("/insert", rest.InsertIntodatabaseCrapDataExample)	//Requisição Post com body json para inserção no banco de dados

	http.ListenAndServe(":8080", mux)

	database.ConnectoDatabaseExample()

}
