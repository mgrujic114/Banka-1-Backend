package exchange

import (
	"banka1.com/controllers"
	"banka1.com/types"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func setupApp() *fiber.App {
	app := fiber.New()
	controllers.InitExchangeRoutes(app)
	return app
}

func TestGetAllExchanges(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/exchanges", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)

	var response types.Response
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)
	assert.NotNil(t, response.Data)
}

func TestGetExchangeByID(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/exchanges/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)

	var response types.Response
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)
	assert.NotNil(t, response.Data)
}

func TestGetExchangeByMIC(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/exchanges/mic/XNAS", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)

	var response types.Response
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)
	assert.NotNil(t, response.Data)
}

func TestGetExchangeByAcronym(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/exchanges/acronym/NASDAQ", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)

	var response types.Response
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)
	assert.NotNil(t, response.Data)
}
