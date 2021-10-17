package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"pizza-app/internal/data"
	"strings"
	"testing"
)

func ReadResponse(w *httptest.ResponseRecorder) (*Response, error) {
	r := &Response{}
	if err := json.Unmarshal(w.Body.Bytes(), r); err != nil {
		return nil, err
	}
	return r, nil
}

func TestApp_orderPizzaHandler(t *testing.T) {
	a := App{}
	var addOrderDetailsFn = func(order *data.OrderData) error {
		return nil
	}
	var addOrderDetailsErrFn = func(order *data.OrderData) error {
		return errors.New("error")
	}
	tests := []struct {
		name                string
		body                string
		status              string
		code                int
		httpStatusCode      int
		mockAddOrderDetails func(order *data.OrderData) error
	}{
		{
			name:   "success",
			status: StatusOK,
			body: `
				{
					"pizza_id": 2,
					"pizza_size": "medium",
					"user_id": 1
				}
			`,
			code:                http.StatusOK,
			httpStatusCode:      http.StatusOK,
			mockAddOrderDetails: addOrderDetailsFn,
		},
		{
			name:   "400: Bad request",
			status: StatusFail,
			body: `
				{
					"pizza_id": 2,
					"pizza_size": "medium",
					"user_id": 1
				}
			`,
			code:                http.StatusInternalServerError,
			httpStatusCode:      http.StatusInternalServerError,
			mockAddOrderDetails: addOrderDetailsErrFn,
		},
		{
			name:   "500: Internal server error",
			status: StatusFail,
			body: `
				{
					"pizza_id": 2,
					"pizza_size": "medium"
					"user_id": 1
				}
			`,
			code:                http.StatusBadRequest,
			httpStatusCode:      http.StatusBadRequest,
			mockAddOrderDetails: addOrderDetailsFn,
		},
	}
	for _, tc := range tests {
		path := "/api/buy_pizza"
		a.Repo = &data.MockDB{AddOrderDetailsFn: tc.mockAddOrderDetails}
		r, err := http.NewRequest("POST", path, strings.NewReader(tc.body))
		if err != nil {
			t.Fatalf(
				"%s: failed to create request %s",
				tc.name,
				err.Error(),
			)
		}
		w := httptest.NewRecorder()
		http.HandlerFunc(a.orderPizzaHandler).ServeHTTP(w, r)

		if w.Result().StatusCode != tc.httpStatusCode {
			t.Fatalf("expected http status code %d, got http status code %d", tc.httpStatusCode, w.Result().StatusCode)
		}

		if tc.status == StatusOK {
			_, err := ReadResponse(w)
			if err != nil {
				t.Fatalf(
					"%s: %s",
					tc.name,
					err.Error(),
				)
			}
		} else {
			response, err := ReadResponse(w)
			if err != nil {
				t.Fatalf(
					"%s: failure: %s",
					tc.name,
					err.Error(),
				)
			}
			if response.Error.Code != tc.code {
				t.Errorf(
					"%s: expected error code %d but got %d",
					tc.name,
					tc.code,
					response.Error.Code,
				)
			}
		}
	}
}
