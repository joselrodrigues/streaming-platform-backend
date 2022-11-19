package routes_test

import (
	"fmt"
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	// . "github.com/onsi/gomega"

	"stream/routes"
)

var _ = Describe("Stream", func() {
	It("", func() {

		app := routes.Setup()

		tests := []struct {
			description   string
			route         string
			expectedError bool
			expectedCode  int
			expectedBody  string
		}{
			{
				description:   "index route",
				route:         "/series/one-piece/1",
				expectedError: false,
				expectedCode:  200,
				expectedBody:  "OK",
			},
		}
		for _, test := range tests {
			req, _ := http.NewRequest(
				"GET",
				test.route,
				nil,
			)

			res, err := app.Test(req, -1)

			fmt.Printf("%v %v", res, err)
		}

	})
})
