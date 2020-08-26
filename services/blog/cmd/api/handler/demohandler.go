package handler

import (
	"net/http"
	"wang/common/logger"
)

type (
	DemoHandler struct {

	}
)

func NewDemoHandler() *DemoHandler {
	return &DemoHandler{}
}

func (dh *DemoHandler) Demo(w http.ResponseWriter, r *http.Request)  {

	//temp, err := template.ParseFiles("template/index.html")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//temp.Execute(w, nil)
	logger.Info("hahaha")
	return
}