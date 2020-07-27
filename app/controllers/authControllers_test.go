package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetAuthenticationParams(t *testing.T) {
	testCases := []struct {
		Input string
		Want  string
	}{
		{
			Input: "?scope=scope&client_id=client_id&response_type=response_type&redirect_url=redirect_url&state=state",
			Want:  "scope",
		},
		{Input: "", Want: ""},
	}
	for _, testData := range testCases {
		context, _ := gin.CreateTestContext(httptest.NewRecorder())
		req, _ := http.NewRequest("GET", testServer.URL+testData.Input, nil)
		context.Request = req
		result, _ := getAuthenticationParams(context)
		if testData.Want != result.Scope {
			t.Errorf("Want %v but got %v", testData.Want, result.Scope)
		}
	}
}
