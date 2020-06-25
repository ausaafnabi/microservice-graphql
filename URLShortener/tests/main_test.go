package main_test

import (
  "testing"
  "net/http"
)

func TestGetOriginalURL(t *testing.T){
  response,err :=http.Get("http://localhost:8000/v1/short/1")

  if http.StatusOK != response.StatusCode {
    t.Error("Expected respose code %d. | Got %d\n",http.StatusOK,response.StatusCode)
  }

  if err != nil{
    t.Error("Encountered an error:",err)
  }
}
