package incident

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusIncident(t *testing.T) {
	t.Run("Тест инцидент http коды 200 и 500", func(t *testing.T) {
		handlerHttp := &TestHandler{}

		ts := httptest.NewServer(handlerHttp)
		defer ts.Close()

		result := StatusIncident(ts.URL + "/incident/ok")
		require.Equal(t, data, result)

		empty := StatusIncident(ts.URL + "/incident/fail")
		require.Equal(t, []IncidentData{}, empty)
	})
}

var data = []IncidentData{
	{Topic: "Доставка SMS в EU", Status: "closed"},
	{Topic: "Стабильность MMS-соединения", Status: "closed"},
	{Topic: "Чистота подключения к голосовому вызову", Status: "closed"},
	{Topic: "Страница оформления заказа закрыта", Status: "closed"},
	{Topic: "Перезапуск поддержки", Status: "active"},
	{Topic: "Номер телефона не работает в US", Status: "closed"},
	{Topic: "API с низкой задержкой", Status: "closed"},
}

type TestHandler struct{}

func (m *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/incident/ok":
		response, _ := json.Marshal(data)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(response)
	case "/incident/fail":
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(""))
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
