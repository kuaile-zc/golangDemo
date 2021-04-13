package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)


/**
 * @author  CoreyChen Zhang
 * @version  2021/4/12 17:04
 * @modified
 */

const Prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, Prefix) == -1 {
		return userError("Path must start with" + Prefix)
	}
	path := request.URL.Path[len(Prefix):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
