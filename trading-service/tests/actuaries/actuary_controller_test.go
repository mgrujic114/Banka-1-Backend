package actuary

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
	controllers.InitActuaryRoutes(app)
	return app
}

func jsonBody(data interface{}) *bytes.Reader {
	body, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	return bytes.NewReader(body)
}

func TestCreateActuary(t *testing.T) {
	app := setupApp()

	actuaryRequest := map[string]interface{}{
		"name": "John Doe",
		"role": "Senior Actuary",
	}

	req := httptest.NewRequest("POST", "/actuaries", jsonBody(actuaryRequest))
	resp, _ := app.Test(req)

	assert.Equal(t, 201, resp.StatusCode)

	var response types.Response
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)
	assert.NotNil(t, response.Data)
}

func TestGetAllActuaries(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/actuaries", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)

	var response types.Response
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)
	assert.NotNil(t, response.Data)
}
