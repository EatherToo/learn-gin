package router

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestRouter(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin get method", w.Body.String())
}

func TestUserSave(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/astaxie", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户astaxie已保存", w.Body.String())
}

func TestUserRegister(t *testing.T) {
	values := url.Values{}
	values.Add("email", "rrr@gmail.com")
	values.Add("username", "rrr")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/register", bytes.NewBufferString(values.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	router := SetupRouter()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}
