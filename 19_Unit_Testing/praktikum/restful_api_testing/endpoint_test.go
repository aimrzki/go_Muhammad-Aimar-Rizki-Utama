package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"restful_api_testing/controllers"
)

func TestGetUsersController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTUyODg2ODUsInVzZXJfaWQiOjJ9.-icUzvDuDtwRFM4LejYXV1Nhks_C0JUmtcZcM5NUEuw"
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.GetUsersController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetUsersControllerWithoutAuthorizationCode(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token := ""
	req.Header.Set("Authorization", token)

	// Assertions
	if assert.NoError(t, controllers.GetUsersController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

}

func TestGetUsersControllerWithWrongtAuthorizationCode(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token := "wrong"
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.GetUsersController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetUserController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTUyODg2ODUsInVzZXJfaWQiOjJ9.-icUzvDuDtwRFM4LejYXV1Nhks_C0JUmtcZcM5NUEuw"
	req.Header.Set("Authorization", token)

	// Assertions
	if assert.NoError(t, controllers.GetUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetUserControllerWithWrongId(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/100", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("100")

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTUyODg2ODUsInVzZXJfaWQiOjJ9.-icUzvDuDtwRFM4LejYXV1Nhks_C0JUmtcZcM5NUEuw"
	req.Header.Set("Authorization", token)

	// Assertions
	if assert.NoError(t, controllers.GetUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetUserControllerWithWrongToken(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	token := "eyJleHAiOjE2OTUyODg2ODUsInVzZXJfaWQiOjJ9.-icUzvDuDtwRFM4LejYXV1Nhks_C0JUmtcZcM5NUEuw"
	req.Header.Set("Authorization", token)

	// Assertions
	if assert.NoError(t, controllers.GetUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetUserControllerWithoutToken(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTUyODg2ODUsInVzZXJfaWQiOjJ9.-icUzvDuDtwRFM4LejYXV1Nhks_C0JUmtcZcM5NUEuw"
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.GetUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCreateUserController(t *testing.T) {
	e := echo.New()

	reqBody := `{
		"name": "ayam",
		"email": "ayam@gmail.com",
		"password": "admin123"
	}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.CreateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestLoginUserController(t *testing.T) {
	e := echo.New()
	reqBody := `{
		"email": "aimrzki@gmail.com",
		"password": "admin123"
	}`
	req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.LoginUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestLoginUserControllerWithWrongEmail(t *testing.T) {
	e := echo.New()

	reqBody := `{
		"email": "aimrzki123@gmail.com",
		"password": "admin1234"
	}`
	req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.LoginUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestLoginUserControllerWithWrongPassword(t *testing.T) {
	e := echo.New()

	reqBody := `{
		"email": "aimrzki@gmail.com",
		"password": "admin1234"
	}`
	req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.LoginUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteUserController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users/5", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("5")

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTUyOTY4MzUsInVzZXJfaWQiOjV9.noVcAZj7DaRf5198zmha8MaaSbCRVmA-mY9GiBxLbEA"
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.DeleteUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteUserControllerWithWrongId(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users/100", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("100")

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTUyOTMzMTEsInVzZXJfaWQiOjd9.4DvT823Fl7WniAOS6kQ9c-_DK2kJyONSNv8vi8D1H8k"
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.DeleteUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteUserControllerWithWrongToken(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users/7", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("7")

	token := "wrong"
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.DeleteUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteUserControllerWithoutToken(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users/7", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("7")

	token := ""
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.DeleteUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateUserController(t *testing.T) {
	e := echo.New()
	reqBody := `{
		"name": "Aimar Rizki",
		"email": "aimrzki@gmail.com",
		"password": "admin123"
	}`
	req := httptest.NewRequest(http.MethodPut, "/users/2", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTUyODg2ODUsInVzZXJfaWQiOjJ9.-icUzvDuDtwRFM4LejYXV1Nhks_C0JUmtcZcM5NUEuw"
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.UpdateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateUserControllerWithWrongInput(t *testing.T) {
	e := echo.New()
	reqBody := `{
		"name": 1
		"email": 1
		"password": 1
	}`
	req := httptest.NewRequest(http.MethodPut, "/users/2", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTUyODg2ODUsInVzZXJfaWQiOjJ9.-icUzvDuDtwRFM4LejYXV1Nhks_C0JUmtcZcM5NUEuw"
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.UpdateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateUserControllerWithWrongId(t *testing.T) {
	e := echo.New()
	reqBody := `{
		"name": "Aimar Rizki",
		"email": "aimrzki@gmail.com",
		"password": "admin123"
	}`
	req := httptest.NewRequest(http.MethodPut, "/users/100", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("100")
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTUyODg2ODUsInVzZXJfaWQiOjJ9.-icUzvDuDtwRFM4LejYXV1Nhks_C0JUmtcZcM5NUEuw"
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.UpdateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

}

func TestUpdateUserControllerWithWrongAuthorizaionCode(t *testing.T) {
	e := echo.New()
	reqBody := `{
		"name": "Aimar Rizki",
		"email": "aimrzki@gmail.com",
		"password": "admin123"
	}`
	req := httptest.NewRequest(http.MethodPut, "/users/2", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	token := "wrong"
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.UpdateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateUserControllerWithoutAuthorizationCode(t *testing.T) {
	e := echo.New()
	reqBody := `{
		"name": "Aimar Rizki",
		"email": "aimrzki@gmail.com",
		"password": "admin123"
	}`
	req := httptest.NewRequest(http.MethodPut, "/users/2", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	token := ""
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.UpdateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetAllBooksController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTUyOTM4NzAsInVzZXJfaWQiOjV9.3Yq-TmH0ZI1jah4AErWu1EKVcT6UQx7X-0NvTEfXkow"
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.GetAllBooksController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetAllBooksControllerWithWrongToken(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token := "wrong"
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.GetAllBooksController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetAllBooksControllerWithoOutoken(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token := ""
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.GetAllBooksController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetBookController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books/3", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("3")

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTUyOTM4NzAsInVzZXJfaWQiOjV9.3Yq-TmH0ZI1jah4AErWu1EKVcT6UQx7X-0NvTEfXkow"
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.GetBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetBookControllerWithWrongId(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books/1000", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1000")

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTUyOTM4NzAsInVzZXJfaWQiOjV9.3Yq-TmH0ZI1jah4AErWu1EKVcT6UQx7X-0NvTEfXkow"
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.GetBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetBookControllerWrongToken(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books/1000", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1000")

	token := "wrong"
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.GetBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetBookControllerWithoutToken(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books/1000", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1000")

	token := ""
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.GetBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCreateBookController(t *testing.T) {
	e := echo.New()
	reqBody := `{
		"title": "john wick",
		"author": "Francine Jay",
		"publisher": "Gramedia"
	}`
	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTUyOTM4NzAsInVzZXJfaWQiOjV9.3Yq-TmH0ZI1jah4AErWu1EKVcT6UQx7X-0NvTEfXkow"
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.CreateBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCreateBookControllerWithToken(t *testing.T) {
	e := echo.New()
	reqBody := `{
		"title": "john wick",
		"author": "Francine Jay",
		"publisher": "Gramedia"
	}`
	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token := "wrong"
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.CreateBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCreateBookControllerWithWrongInput(t *testing.T) {
	e := echo.New()
	reqBody := `{
		"title": 1,
		"author": 1,
		"publisher": 1
	}`
	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token := "wrong"
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.CreateBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCreateBookControllerWithOutToken(t *testing.T) {
	e := echo.New()
	reqBody := `{
		"title": "john wick",
		"author": "Francine Jay",
		"publisher": "Gramedia"
	}`
	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token := ""
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.CreateBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateBookController(t *testing.T) {
	e := echo.New()
	reqBody := `{
		"title": "5 cm - English Sub",
		"author": "Donny Dhirgantoro",
		"publisher": "Kompas Gramedia"
	}`
	req := httptest.NewRequest(http.MethodPut, "/books/6", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("6")

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTUyOTM4NzAsInVzZXJfaWQiOjV9.3Yq-TmH0ZI1jah4AErWu1EKVcT6UQx7X-0NvTEfXkow"
	req.Header.Set("Authorization", token)
	// Assertions
	if assert.NoError(t, controllers.UpdateBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateBookControllerWithWrongData(t *testing.T) {
	e := echo.New()
	reqBody := `{
		"title": 2,
		"author": 2,
		"publisher": 2
	}`
	req := httptest.NewRequest(http.MethodPut, "/books/6", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("6")

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTUyOTM4NzAsInVzZXJfaWQiOjV9.3Yq-TmH0ZI1jah4AErWu1EKVcT6UQx7X-0NvTEfXkow"
	req.Header.Set("Authorization", token)
	if assert.NoError(t, controllers.UpdateBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateBookControllerWithWrongId(t *testing.T) {
	e := echo.New()
	reqBody := `{
		"title": "5 cm - English Sub",
		"author": "Donny Dhirgantoro",
		"publisher": "Kompas Gramedia"
	}`
	req := httptest.NewRequest(http.MethodPut, "/books/1000", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1000")

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTUyOTM4NzAsInVzZXJfaWQiOjV9.3Yq-TmH0ZI1jah4AErWu1EKVcT6UQx7X-0NvTEfXkow"
	req.Header.Set("Authorization", token)
	// Assertions
	if assert.NoError(t, controllers.UpdateBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateBookControllerWithWrongToken(t *testing.T) {
	e := echo.New()
	reqBody := `{
		"title": "5 cm - English Sub",
		"author": "Donny Dhirgantoro",
		"publisher": "Kompas Gramedia"
	}`
	req := httptest.NewRequest(http.MethodPut, "/books/6", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("6")

	token := "wrong"
	req.Header.Set("Authorization", token)
	// Assertions
	if assert.NoError(t, controllers.UpdateBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateBookControllerWithoutToken(t *testing.T) {
	e := echo.New()
	reqBody := `{
		"title": "5 cm - English Sub",
		"author": "Donny Dhirgantoro",
		"publisher": "Kompas Gramedia"
	}`
	req := httptest.NewRequest(http.MethodPut, "/books/6", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("6")

	token := ""
	req.Header.Set("Authorization", token)
	// Assertions
	if assert.NoError(t, controllers.UpdateBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteBookController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/books/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("33")

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTUyOTM4NzAsInVzZXJfaWQiOjV9.3Yq-TmH0ZI1jah4AErWu1EKVcT6UQx7X-0NvTEfXkow"
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.DeleteBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteBookControllerWithWrongId(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/books/1000", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1000")

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTUyOTM4NzAsInVzZXJfaWQiOjV9.3Yq-TmH0ZI1jah4AErWu1EKVcT6UQx7X-0NvTEfXkow"
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.DeleteBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteBookControllerWithWrongToken(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/books/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	token := "wrong"
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.DeleteBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

	}
}

func TestDeleteBookControllerWithoutToken(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/books/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	token := ""
	req.Header.Set("Authorization", token)

	if assert.NoError(t, controllers.DeleteBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
