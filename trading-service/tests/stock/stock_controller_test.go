package stock

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
	controllers.InitStockRoutes(app)
	return app
}

func jsonBody(data interface{}) *bytes.Reader {
	body, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	return bytes.NewReader(body)
}

func TestGetAllStocks(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/stocks", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)

	var response types.Response
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)
	assert.NotNil(t, response.Data)
}

func TestGetStockByID(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/stocks/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)

	var response types.Response
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)
	assert.NotNil(t, response.Data)
}

func TestCreateStock(t *testing.T) {
	app := setupApp()

	stockRequest := map[string]interface{}{
		"name":  "AAPL",
		"price": 150.00,
	}

	req := httptest.NewRequest("POST", "/stocks", jsonBody(stockRequest))
	resp, _ := app.Test(req)

	assert.Equal(t, 201, resp.StatusCode)

	var response types.Response
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)
	assert.NotNil(t, response.Data)
}

func TestUpdateStock(t *testing.T) {
	app := setupApp()

	stockRequest := map[string]interface{}{
		"price": 155.00,
	}

	req := httptest.NewRequest("PUT", "/stocks/1", jsonBody(stockRequest))
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)

	var response types.Response
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)
	assert.NotNil(t, response.Data)
}

func TestDeleteStock(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("DELETE", "/stocks/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 204, resp.StatusCode)
}
