package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"sandbox/pkg"
)

type CryptRequest struct {
	V string `json:"v,omitempty"`
	K string `json:"k,omitempty"`
}

type CryptResponse struct {
	V []byte `json:"v,omitempty"`
}

type DecryptRequest struct {
	V []byte `json:"v,omitempty"`
	K string `json:"k,omitempty"`
}

type DecryptResponse struct {
	V string `json:"v,omitempty"`
}

func CryptCFB(c echo.Context) error {
	r := new(CryptRequest)
	if err := c.Bind(r); err != nil {
		return err
	}

	v, err := pkg.CryptCFB([]byte(r.V), []byte(r.K))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &CryptResponse{V: v})
}

func DecryptCFB(c echo.Context) error {
	r := new(DecryptRequest)
	if err := c.Bind(r); err != nil {
		return err
	}

	v, err := pkg.DecryptCFB(r.V, []byte(r.K))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &DecryptResponse{V: string(v)})
}

func CryptCBC(c echo.Context) error {
	r := new(CryptRequest)
	if err := c.Bind(r); err != nil {
		return err
	}

	v, err := pkg.CryptCBC([]byte(r.V), []byte(r.K))
	if err != nil {
		return err
	}
	/*
		{
		    "v": "crypt",
		    "k": "astaxie12798akljzmknm.ahkjkljl;k"
		}
		これで暗号化したbyte配列の返却値は
		{
			"v": "iwhaCZa80/9ncfpUgOJ68w=="
		}
		になるので、aes.BlockSizeの倍数になってない
	*/
	c.Logger().Info("返却するbyte配列はaes.BlockSizeの倍数にならないのでそのまま復号できない")
	return c.JSON(http.StatusOK, &CryptResponse{V: v})
}

func DecryptCBC(c echo.Context) error {
	r := new(CryptRequest)
	if err := c.Bind(r); err != nil {
		return err
	}

	c.Logger().Info("暗号化したbyte配列はaes.BlockSizeの倍数にならないのでそのまま復号できない")
	v, err := pkg.DecryptCBC([]byte(r.V), []byte(r.K))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &DecryptResponse{V: string(v)})
}
