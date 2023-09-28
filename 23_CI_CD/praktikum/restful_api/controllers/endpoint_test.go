package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func generateAuthToken() (string, error) {
	// Buat request untuk login
	loginRequest := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    "aimrzki@gmail.com",
		Password: "admin123",
	}

	// Marshal request ke JSON
	requestBody, err := json.Marshal(loginRequest)
	if err != nil {
		return "", err
	}

	// Buat HTTP request
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// Buat konteks Echo
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Panggil handler LoginUserController untuk mendapatkan token
	if err := LoginUserController(c); err != nil {
		return "", err
	}

	// Periksa status kode HTTP
	if rec.Code != http.StatusOK {
		return "", fmt.Errorf("Login failed with status code %d", rec.Code)
	}

	// Ambil token dari respons
	var response map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		return "", err
	}

	token, ok := response["token"].(string)
	if !ok {
		return "", errors.New("Token not found in response")
	}

	return token, nil
}

func TestCreateUserController(t *testing.T) {
	e := echo.New()

	reqBody := `{
		"name": "Aimar Rizki",
		"email": "aimrzki@gmail.com",
		"password": "admin123"
	}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserController(c)) {
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

	if assert.NoError(t, LoginUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetUsersController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token, err := generateAuthToken()
	if err != nil {
		t.Fatalf("Failed to generate auth token: %v", err)
	}

	// Kasus 1: Token valid
	req.Header.Set("Authorization", "Bearer "+token)
	if assert.NoError(t, GetUsersController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

	}

	// Kasus 2: Token tidak ada (Unauthorized)
	req.Header.Del("Authorization")
	if assert.NoError(t, GetUsersController(c)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	}

	// Kasus 3: Token tidak valid (Unauthorized)
	req.Header.Set("Authorization", "InvalidTokenFormat")
	if assert.NoError(t, GetUsersController(c)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	}

	// Kasus 4: Token tidak valid (Unauthorized)
	req.Header.Set("Authorization", "Bearer InvalidToken")
	if assert.NoError(t, GetUsersController(c)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	}
}

func TestGetUserController(t *testing.T) {
	e := echo.New()

	// Kasus 1: Pengguna yang sah mendapatkan data pengguna mereka sendiri
	t.Run("ValidUserOwnData", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		token, err := generateAuthToken()
		if err != nil {
			t.Fatalf("Failed to generate auth token: %v", err)
		}

		req.Header.Set("Authorization", "Bearer "+token)

		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// Kasus 2: Pengguna yang sah mencoba mendapatkan data pengguna lain (Forbidden)
	t.Run("ValidUserOtherUserData", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/2", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("2")

		token, err := generateAuthToken()
		if err != nil {
			t.Fatalf("Failed to generate auth token: %v", err)
		}

		req.Header.Set("Authorization", "Bearer "+token)

		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, http.StatusForbidden, rec.Code)
		}
	})

	// Kasus 3: Token tidak valid (Unauthorized)
	t.Run("InvalidToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

	// Kasus 4: Pengguna tidak memiliki akses token (Unauthorized)
	t.Run("NoToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

	// Kasus 5: ID yang salah (Bad Request)
	t.Run("InvalidID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/invalid_id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("invalid_id")

		token, err := generateAuthToken()
		if err != nil {
			t.Fatalf("Failed to generate auth token: %v", err)
		}

		req.Header.Set("Authorization", "Bearer "+token)

		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}

func TestUpdateUserController(t *testing.T) {
	e := echo.New()

	// Kasus 1: Pengguna yang sah berhasil memperbarui data pengguna mereka sendiri
	t.Run("ValidUserUpdateOwnData", func(t *testing.T) {
		reqBody := `{
			"name": "Muhammad Aimar Rizki Utama",
			"email": "aimrzki@gmail.com",
			"password": "admin123"
		}`

		req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		token, err := generateAuthToken()
		if err != nil {
			t.Fatalf("Failed to generate auth token: %v", err)
		}

		req.Header.Set("Authorization", "Bearer "+token)

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// Kasus 2: Pengguna yang sah mencoba memperbarui data pengguna lain (Forbidden)
	t.Run("ValidUserUpdateOtherUserData", func(t *testing.T) {
		reqBody := `{
			"name": "Muhammad Aimar Rizki Utama",
			"email": "aimrzki@gmail.com",
			"password": "admin123"
		}`

		req := httptest.NewRequest(http.MethodPut, "/users/2", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("2")

		token, err := generateAuthToken()
		if err != nil {
			t.Fatalf("Failed to generate auth token: %v", err)
		}

		req.Header.Set("Authorization", "Bearer "+token)

		// Assertions
		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, http.StatusForbidden, rec.Code)
		}
	})

	// Kasus 3: Token tidak valid (Unauthorized)
	t.Run("InvalidToken", func(t *testing.T) {
		reqBody := `{
			"name": "Muhammad Aimar Rizki Utama",
			"email": "aimrzki@gmail.com",
			"password": "admin123"
		}`

		req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

	// Kasus 4: Pengguna tidak memiliki akses token (Unauthorized)
	t.Run("NoToken", func(t *testing.T) {
		reqBody := `{
			"name": "Muhammad Aimar Rizki Utama",
			"email": "aimrzki@gmail.com",
			"password": "admin123"
		}`

		req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

	// Kasus 5: ID yang salah (Bad Request)
	t.Run("InvalidID", func(t *testing.T) {
		reqBody := `{
			"name": "Muhammad Aimar Rizki Utama",
			"email": "aimrzki@gmail.com",
			"password": "admin123"
		}`

		req := httptest.NewRequest(http.MethodPut, "/users/invalid_id", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("invalid_id")

		token, err := generateAuthToken()
		if err != nil {
			t.Fatalf("Failed to generate auth token: %v", err)
		}

		req.Header.Set("Authorization", "Bearer "+token)

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}

func TestCreateBookController(t *testing.T) {
	e := echo.New()

	// Kasus 1: Pengguna yang sah berhasil membuat buku baru
	t.Run("ValidUserCreateBook", func(t *testing.T) {
		reqBody := `{
			"title": "John Wick",
			"author": "Francine Jay",
			"publisher": "Gramedia"
		}`

		req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token, err := generateAuthToken()
		if err != nil {
			t.Fatalf("Failed to generate auth token: %v", err)
		}

		req.Header.Set("Authorization", "Bearer "+token)

		if assert.NoError(t, CreateBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// Kasus 2: Token tidak valid (Unauthorized)
	t.Run("InvalidToken", func(t *testing.T) {
		reqBody := `{
			"title": "John Wick",
			"author": "Francine Jay",
			"publisher": "Gramedia"
		}`

		req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, CreateBookController(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

	// Kasus 3: Pengguna tidak memiliki akses token (Unauthorized)
	t.Run("NoToken", func(t *testing.T) {
		reqBody := `{
			"title": "John Wick",
			"author": "Francine Jay",
			"publisher": "Gramedia"
		}`

		req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Assertions
		if assert.NoError(t, CreateBookController(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

	// Kasus 4: Data buku tidak valid (Bad Request)
	t.Run("InvalidBookData", func(t *testing.T) {
		reqBody := `{
			"title": 1,
			"author": 1,
			"publisher": 1
		}`

		req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token, err := generateAuthToken()
		if err != nil {
			t.Fatalf("Failed to generate auth token: %v", err)
		}

		req.Header.Set("Authorization", "Bearer "+token)

		if assert.NoError(t, CreateBookController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}

func TestGetAllBooksController(t *testing.T) {
	e := echo.New()

	// Kasus 1: Pengguna yang sah mendapatkan semua data buku
	t.Run("ValidUserGetAllBooks", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token, err := generateAuthToken()
		if err != nil {
			t.Fatalf("Failed to generate auth token: %v", err)
		}

		req.Header.Set("Authorization", "Bearer "+token)

		if assert.NoError(t, GetAllBooksController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// Kasus 2: Token tidak valid (Unauthorized)
	t.Run("InvalidToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, GetAllBooksController(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

	// Kasus 3: Pengguna tidak memiliki akses token (Unauthorized)
	t.Run("NoToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, GetAllBooksController(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})
}

func TestGetBookController(t *testing.T) {
	e := echo.New()

	// Kasus 1: Pengguna yang sah mendapatkan data buku berdasarkan ID yang valid
	t.Run("ValidUserGetBook", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		token, err := generateAuthToken()
		if err != nil {
			t.Fatalf("Failed to generate auth token: %v", err)
		}

		req.Header.Set("Authorization", "Bearer "+token)

		if assert.NoError(t, GetBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// Kasus 2: Token tidak valid (Unauthorized)
	t.Run("InvalidToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, GetBookController(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

	// Kasus 3: Pengguna tidak memiliki akses token (Unauthorized)
	t.Run("NoToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, GetBookController(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

	// Kasus 4: ID buku yang salah (Not Found)
	t.Run("InvalidBookID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books/999", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("999")

		token, err := generateAuthToken()
		if err != nil {
			t.Fatalf("Failed to generate auth token: %v", err)
		}

		req.Header.Set("Authorization", "Bearer "+token)

		if assert.NoError(t, GetBookController(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})

	// Kasus 5: ID buku yang tidak valid (Bad Request)
	t.Run("InvalidID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books/invalid_id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("invalid_id")

		token, err := generateAuthToken()
		if err != nil {
			t.Fatalf("Failed to generate auth token: %v", err)
		}

		req.Header.Set("Authorization", "Bearer "+token)

		if assert.NoError(t, GetBookController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}

func TestUpdateBookController(t *testing.T) {
	e := echo.New()

	// Kasus 1: Pengguna yang sah berhasil mengupdate buku
	t.Run("ValidUserUpdateBook", func(t *testing.T) {
		reqBody := `{
			"title": "John Wick - Sub Indo",
			"author": "Francine Jay",
			"publisher": "Gramedia"
		}`

		req := httptest.NewRequest(http.MethodPut, "/books/1", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		token, err := generateAuthToken()
		if err != nil {
			t.Fatalf("Failed to generate auth token: %v", err)
		}

		req.Header.Set("Authorization", "Bearer "+token)

		if assert.NoError(t, UpdateBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// Kasus 2: Token tidak valid (Unauthorized)
	t.Run("InvalidToken", func(t *testing.T) {
		reqBody := `{
			"title": "John Wick - Sub Indo",
			"author": "Francine Jay",
			"publisher": "Gramedia"
		}`

		req := httptest.NewRequest(http.MethodPut, "/books/1", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, UpdateBookController(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

	// Kasus 3: Pengguna tidak memiliki akses token (Unauthorized)
	t.Run("NoToken", func(t *testing.T) {
		reqBody := `{
			"title": "John Wick - Sub Indo",
			"author": "Francine Jay",
			"publisher": "Gramedia"
		}`

		req := httptest.NewRequest(http.MethodPut, "/books/1", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, UpdateBookController(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

	// Kasus 4: ID buku yang salah (Not Found)
	t.Run("InvalidBookID", func(t *testing.T) {
		reqBody := `{
			"title": "John Wick - Sub Indo",
			"author": "Francine Jay",
			"publisher": "Gramedia"
		}`

		req := httptest.NewRequest(http.MethodPut, "/books/999", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("999")

		token, err := generateAuthToken()
		if err != nil {
			t.Fatalf("Failed to generate auth token: %v", err)
		}

		req.Header.Set("Authorization", "Bearer "+token)

		if assert.NoError(t, UpdateBookController(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})

	// Kasus 5: Data buku tidak valid (Bad Request)
	t.Run("InvalidBookData", func(t *testing.T) {
		reqBody := `{
			"title": 1,
			"author": 1,
			"publisher": 1
		}`

		req := httptest.NewRequest(http.MethodPut, "/books/1", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		token, err := generateAuthToken()
		if err != nil {
			t.Fatalf("Failed to generate auth token: %v", err)
		}

		req.Header.Set("Authorization", "Bearer "+token)

		if assert.NoError(t, UpdateBookController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}

func TestDeleteBookController(t *testing.T) {
	e := echo.New()

	// Kasus 1: Pengguna yang sah berhasil menghapus buku
	t.Run("ValidUserDeleteBook", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/books/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		token, err := generateAuthToken()
		if err != nil {
			t.Fatalf("Failed to generate auth token: %v", err)
		}

		req.Header.Set("Authorization", "Bearer "+token)

		if assert.NoError(t, DeleteBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// Kasus 2: Token tidak valid (Unauthorized)
	t.Run("InvalidToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/books/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, DeleteBookController(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

	// Kasus 3: Pengguna tidak memiliki akses token (Unauthorized)
	t.Run("NoToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/books/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, DeleteBookController(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

	t.Run("InvalidBookID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/books/999", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("999")

		token, err := generateAuthToken()
		if err != nil {
			t.Fatalf("Failed to generate auth token: %v", err)
		}

		req.Header.Set("Authorization", "Bearer "+token)

		if assert.NoError(t, DeleteBookController(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})
}

func TestDeleteUserController(t *testing.T) {
	e := echo.New()

	// Kasus 1: Pengguna yang sah berhasil menghapus akun
	t.Run("ValidUserDeleteUser", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		token, err := generateAuthToken()
		if err != nil {
			t.Fatalf("Failed to generate auth token: %v", err)
		}

		req.Header.Set("Authorization", "Bearer "+token)

		if assert.NoError(t, DeleteUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// Kasus 2: Token tidak valid (Unauthorized)
	t.Run("InvalidToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, DeleteUserController(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

	// Kasus 3: Pengguna tidak memiliki akses token (Unauthorized)
	t.Run("NoToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, DeleteUserController(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

	// Kasus 4: ID pengguna yang salah (Not Found)
	t.Run("InvalidUserID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/users/999", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("999")

		token, err := generateAuthToken()
		if err != nil {
			t.Fatalf("Failed to generate auth token: %v", err)
		}

		req.Header.Set("Authorization", "Bearer "+token)

		if assert.NoError(t, DeleteUserController(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})
}
