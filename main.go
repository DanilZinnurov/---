package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	flag.Parse()
	log.SetFlags(0)
	serv := server{}
	serv.addr = flag.String("addr", "localhost:8080", "http service address")
	serv.serve()
	/*http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*serv.addr, nil))*/
	//unit тесты
	//Тест должен поднять сервер подключиться по вебсокету посылать команды sleep(10 мс) в гитхаб залить 126-128 строчки в метод server
}
