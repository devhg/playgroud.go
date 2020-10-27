package Test

import (
	"fmt"
	merror "github.com/QXQZX/LearnGo/pkg/error"
	"log"
	"net/http"
	"strconv"
	"testing"
)

type result struct {
}

func getFromRepository(id int) (result, error) {
	err := fmt.Errorf("no found %d", id)
	//msg := fmt.Sprintf("error getting the  result with id %d", id)
	//err = merror.Wrap(err, msg)
	return result{}, err
}

func getFromService(idString string) (result, error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		err = merror.BadRequest.Wrapf(err, "Service converting id to int")
		err = merror.AddErrorContext(err, "id", "wrong id format, should be an integer")
		return result{}, err
	}
	return getFromRepository(id)
}

func handleError(err error) {
	Type := merror.GetType(err)
	var status int

	switch Type {
	case merror.BadRequest:
		status = http.StatusBadRequest
	case merror.NotFound:
		status = http.StatusNotFound
	default:
		status = http.StatusBadRequest
	}
	fmt.Println(status)
	if Type == merror.NoType {
		log.Printf(err.Error())
	}
	// 写回ResponseWriter

	context := merror.GetErrorContext(err)
	if context != nil {
		log.Println(context)
	}
}

func Test(t *testing.T) {
	_, err := getFromService("1s")
	if err != nil {
		handleError(err)
	}
}
