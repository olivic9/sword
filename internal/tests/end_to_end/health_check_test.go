package end_to_end

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sword-project/internal/handlers"
	"sword-project/pkg/configs"
	"sword-project/pkg/helpers/test_helpers"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHealthcheckHandler(t *testing.T) {

	expectedResponse := "{\"status\":\"available\",\"system_info\":{\"environment\":\"\"}}"

	configs.InitializeConfigs()
	w := httptest.NewRecorder()

	gin.SetMode(gin.ReleaseMode)
	var params []gin.Param
	u := url.Values{}
	e, ctx := test_helpers.BuildGinTestEngine(w)

	e.GET("/ping", handlers.HealthcheckHandler)
	test_helpers.MockJsonGet(ctx, params, u)
	req, _ := http.NewRequestWithContext(ctx, "GET", "/ping", nil)
	e.ServeHTTP(w, req)
	responseData, _ := io.ReadAll(w.Body)

	test_helpers.IsObjectEqual(t, string(responseData), expectedResponse)

}
