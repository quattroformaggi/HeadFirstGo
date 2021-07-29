package main

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
	"fmt"
)

type Guestbook struct{ //для view.html
	SignatureCount int
	Signatures []string
}

func getStrings(fileName string)[]string{
	var lines []string
	file,err := os.Open(fileName)
	if os.IsNotExist(err){
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		lines = append(lines,scanner.Text())
	}
	check(scanner.Err())
	return lines
	//возвращает сегмент считанных из fileName строк 1:1 элементов сегмента
}

func check(err error) { //проверка ошибок
	if err != nil {
		log.Fatal(err)
	}
}

func newHandler(writer http.ResponseWriter, request *http.Request){
	html,err := template.ParseFiles("new.html")
	check(err)
	err = html.Execute(writer,nil)
	check(err)
	//форма для ввода
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt")
	html,err:=template.ParseFiles("view.html")
	check(err)
	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures: signatures,
	}
	err = html.Execute(writer,guestbook)
	check(err)
	//читает записи с ГК и выводит их + записей всего
}

func createHandler(writer http.ResponseWriter, request *http.Request){
	signature := request.FormValue("signature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file,err:= os.OpenFile("signatures.txt",options,os.FileMode(0600))
	check(err)
	_,err = fmt.Fprintln(file, signature)
	check(err)
	err = file.Close()
	check(err)
	//открывает файл signatures.txt / создаёт его
	//добавляет строку текста в конец файла
	//закрывает файл
	http.Redirect(writer,request,"/guestbook",http.StatusFound)
	//POST /gb/create -> redir /gb -> GET /gb
	
	//получает запрос POST с записью и присоединяет ее к signatures.txt
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
