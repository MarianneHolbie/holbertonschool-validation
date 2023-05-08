package main

import (
  //nolint:staticcheck
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "testing"
)

func Test_server(t *testing.T) {
  if testing.Short() {
    t.Skip("Flag `-short` provided: skipping Integration Tests.")
  }

  tests := []struct {
    name         string
    URI          string
    responseCode int
    body         string
  }{
    {
      name:         "Home page",
      URI:          "",
      responseCode: 404,
      body:         "404 page not found\n",
    },
    {
      name:         "Hello page",
      URI:          "/hello?name=Holberton",
      responseCode: 200,
      body:         "Hello Holberton!",
    },
	{
	  name:			"Space in name, bad url",
	  URI: 			"/hello?name=Grace Hoper",
	  responseCode:	400,
	  body:			"400 Bad Request",
	},
	{
	  name: 		"Encode space in URI with %20",
	  URI:			"/hello?name=Rosalind%20Franklin",
	  responseCode:	200,
	  body:			"Hello Rosalind Franklin!",
	},
	{
	  name:			"if two name",
	  URI:			"/hello?name=Betty&name=Holberton",
	  responseCode:	200,
	  body:			"Hello Betty!",
	},
	{
	  name:			"Empty name",
	  URI:			"/hello?name=",
	  responseCode:	400,
	},
	{
	  name:			"without value name",
	  URI:			"/hello?name",
	  responseCode:	400,
	  },
	{
	  name:			"name is B212#",
	  URI:			"/hello?name=B212%23",
	  responseCode:	200,
	  body:			"Hello B212#!",
	},
	{
	  name: 		"There",
	  URI:			"/hello?name=there",
	  responseCode:	200,
	  body:			"Hello there!",
	},
	{
	  name:			"health server",
	  URI:			"/health",
	  responseCode:	200,
	  body:			"ALIVE",
	},
	{
	  name:			"home path",
	  URI:			"/",
	  responseCode:	404,
	  body:			"404 page not found\n",
	},
	{
	  name: 		"invalid path",
	  URI:			"/invalid",
	  responseCode:	404,
	  body:			"404 page not found\n",
	},
	{
		name:         "No name parameter",
		URI:          "/hello",
		responseCode: 200,
		body:         "Hello there!",
	},

  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      ts := httptest.NewServer(setupRouter())
      defer ts.Close()

      res, err := http.Get(ts.URL + tt.URI)
      if err != nil {
        t.Fatal(err)
      }

      // Check that the status code is what you expect.
      expectedCode := tt.responseCode
      gotCode := res.StatusCode
      if gotCode != expectedCode {
        t.Errorf("handler returned wrong status code: got %q want %q", gotCode, expectedCode)
      }

      // Check that the response body is what you expect.
      expectedBody := tt.body
      bodyBytes, err := ioutil.ReadAll(res.Body)
      res.Body.Close()
      if err != nil {
        t.Fatal(err)
      }
      gotBody := string(bodyBytes)
      if gotBody != expectedBody {
        t.Errorf("handler returned unexpected body: got %q want %q", gotBody, expectedBody)
      }
    })
  }
}