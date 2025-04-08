package portfolio

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
	controllers.InitPortfolioRoutes(app)
	return app
}

func jsonBody(data interface{}) *bytes.Reader {
	body, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	return bytes.NewReader(body)
}

func TestGetPortfolio(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/portfolios/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)

	var response types.Response
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)
	assert.NotNil(t, response.Data)
}

func TestCreatePortfolio(t *testing.T) {
	app := setupApp()

	portfolioRequest := map[string]interface{}{
		"name":   "My Portfolio",
		"userID": 1001,
	}

	req := httptest.NewRequest("POST", "/portfolios", jsonBody(portfolioRequest))
	resp, _ := app.Test(req)

	assert.Equal(t, 201, resp.StatusCode)

	var response types.Response
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)
	assert.NotNil(t, response.Data)
}

func TestUpdatePortfolio(t *testing.T) {
	app := setupApp()

	portfolioRequest := map[string]interface{}{
		"name": "Updated Portfolio",
	}

	req := httptest.NewRequest("PUT", "/portfolios/1", jsonBody(portfolioRequest))
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)

	var response types.Response
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)
	assert.NotNil(t, response.Data)
}

func TestDeletePortfolio(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("DELETE", "/portfolios/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 204, resp.StatusCode)
}
