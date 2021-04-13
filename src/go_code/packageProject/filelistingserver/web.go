package main

import (
	"fmt"
	"go_code/packageProject/filelistingserver/filelisting"
	"net/http"
	"os"
)

type appHandler func(http.ResponseWriter, *http.Request) error

type userError interface {
	error
	Message() string
}

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		//panic
		defer func() {
			if r := recover(); r != nil{
				fmt.Println("Panic", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {
			fmt.Println("Error handling request", err.Error())

			//user err
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
			}

			// system err
			switch {
			case os.IsNotExist(err):
				http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			default:
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}
	}
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe("localhost:8888", nil)
	if err != nil {
		panic(err)
	}
}
