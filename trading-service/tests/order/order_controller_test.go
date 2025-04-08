package order

import (
	"banka1.com/controllers"
	"banka1.com/types"
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http/httptest"
	"testing"
)

func setupApp() *fiber.App {
	app := fiber.New()
	controllers.InitOrderRoutes(app)
	return app
}

func jsonBody(data interface{}) *bytes.Reader {
	body, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	return bytes.NewReader(body)
}

func TestCreateOrder(t *testing.T) {
	app := setupApp()

	orderRequest := map[string]interface{}{
		"securityID":        1,
		"accountID":         1001,
		"orderType":         "LIMIT",
		"quantity":          20,
		"limitPricePerUnit": 95.50,
		"direction":         "BUY",
		"contractSize":      1,
	}

	req := httptest.NewRequest("POST", "/orders", jsonBody(orderRequest))
	resp, _ := app.Test(req)

	assert.Equal(t, 201, resp.StatusCode)

	var response types.Response
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)
	assert.NotNil(t, response.Data)
}

func TestGetOrders(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/orders", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)

	var response types.Response
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)
	assert.NotNil(t, response.Data)
}
