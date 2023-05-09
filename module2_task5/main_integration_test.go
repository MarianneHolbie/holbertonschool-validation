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
      responseCode: 200,
      body:         "<!DOCTYPE html>\n<html lang=\"en-us\">\n  <head>\n    <meta charset=\"utf-8\">\n    <meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge,chrome=1\">\n    \n    <title>Awesome Inc</title>\n    <meta name=\"viewport\" content=\"width=device-width,minimum-scale=1\">\n    <meta name=\"description\" content=\"\">\n    <meta name=\"generator\" content=\"Hugo 0.84.0\" />\n    \n    \n    \n    \n      <meta name=\"robots\" content=\"noindex, nofollow\">\n    \n\n    \n<link rel=\"stylesheet\" href=\"/ananke/css/main.min.css\" >\n\n\n\n    \n    \n    \n      \n\n    \n\n    \n    \n      <link href=\"/index.xml\" rel=\"alternate\" type=\"application/rss+xml\" title=\"Awesome Inc\" />\n      <link href=\"/index.xml\" rel=\"feed\" type=\"application/rss+xml\" title=\"Awesome Inc\" />\n      \n    \n    \n    <meta property=\"og:title\" content=\"Awesome Inc\" />\n<meta property=\"og:description\" content=\"\" />\n<meta property=\"og:type\" content=\"website\" />\n<meta property=\"og:url\" content=\"http://example.org/\" />\n\n<meta itemprop=\"name\" content=\"Awesome Inc\">\n<meta itemprop=\"description\" content=\"\"><meta name=\"twitter:card\" content=\"summary\"/>\n<meta name=\"twitter:title\" content=\"Awesome Inc\"/>\n<meta name=\"twitter:description\" content=\"\"/>\n\n\t\n  </head>\n\n  <body class=\"ma0 avenir bg-near-white\">\n\n    \n\n  <header>\n    <div class=\"pb3-m pb6-l bg-black\">\n      <nav class=\"pv3 ph3 ph4-ns\" role=\"navigation\">\n  <div class=\"flex-l justify-between items-center center\">\n    <a href=\"/\" class=\"f3 fw2 hover-white no-underline white-90 dib\">\n      \n        Awesome Inc\n      \n    </a>\n    <div class=\"flex-l items-center\">\n      \n\n      \n      \n<div class=\"ananke-socials\">\n  \n</div>\n\n    </div>\n  </div>\n</nav>\n\n      <div class=\"tc-l pv3 ph3 ph4-ns\">\n        <h1 class=\"f2 f-subheadline-l fw2 light-silver mb0 lh-title\">\n          Awesome Inc\n        </h1>\n        \n      </div>\n    </div>\n  </header>\n\n\n    <main class=\"pb7\" role=\"main\">\n      \n <article class=\"cf ph3 ph5-l pv3 pv4-l f4 tc-l center measure-wide lh-copy mid-gray\">\n    \n  </article>\n  \n  \n  \n  \n  \n  \n  \n\n    </main>\n    <footer class=\"bg-black bottom-0 w-100 pa3\" role=\"contentinfo\">\n  <div class=\"flex justify-between\">\n  <a class=\"f4 fw4 hover-white no-underline white-70 dn dib-ns pv2 ph3\" href=\"http://example.org/\" >\n    &copy;  Awesome Inc 2023 \n  </a>\n    <div>\n<div class=\"ananke-socials\">\n  \n</div>\n</div>\n  </div>\n</footer>\n\n  </body>\n</html>\n",
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
	  URI:			"/hello",
	  responseCode:	200,
	  body:			"Hello there!",
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