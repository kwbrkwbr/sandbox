package handler

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	type args struct {
		c echo.Context
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "正常",
			args:    args{c: c},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Hello(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Hello() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				assert.Equal(t, http.StatusOK, c.Response().Status)
				assert.Equal(t, "Hello, World!", res.Body.String())
			}
		})
	}
}
