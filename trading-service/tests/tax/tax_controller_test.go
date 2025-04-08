package taxpackage

import (
	"banka1.com/controllers"
	"banka1.com/types"
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func setupApp() *fiber.App {
	app := fiber.New()
	controllers.InitTaxRoutes(app)
	return app
}

func jsonBody(data interface{}) *bytes.Reader {
	body, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	return bytes.NewReader(body)
}

func TestGetAllTaxes(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/taxes", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)

	var response types.Response
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)
	assert.NotNil(t, response.Data)
}

func TestGetTaxByID(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/taxes/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)

	var response types.Response
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)
	assert.NotNil(t, response.Data)
}

func TestCreateTax(t *testing.T) {
	app := setupApp()

	taxRequest := map[string]interface{}{
		"name": "VAT",
		"rate": 0.20,
	}

	req := httptest.NewRequest("POST", "/taxes", jsonBody(taxRequest))
	resp, _ := app.Test(req)

	assert.Equal(t, 201, resp.StatusCode)

	var response types.Response
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)
	assert.NotNil(t, response.Data)
}

func TestUpdateTax(t *testing.T) {
	app := setupApp()

	taxRequest := map[string]interface{}{
		"rate": 0.25,
	}

	req := httptest.NewRequest("PUT", "/taxes/1", jsonBody(taxRequest))
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)

	var response types.Response
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)
	assert.NotNil(t, response.Data)
}

func TestDeleteTax(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("DELETE", "/taxes/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 204, resp.StatusCode)
}
