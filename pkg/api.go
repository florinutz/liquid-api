package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/http/httputil"

	"github.com/florinutz/liquid-api/pkg/cli"

	"github.com/prometheus/common/log"

	"github.com/florinutz/liquid-api/pkg/model"
)

type Provider interface {
	GetFiatAccounts() (accounts []*model.Account, res *http.Response, err error)
	GetCryptoAccounts() (accounts []*model.Account, res *http.Response, err error)
	GetProducts(perpetual bool) (products []*model.Product, res *http.Response, err error)
	GetProduct(idOrpairCode string) (product *model.Product, res *http.Response, err error)
}

func (a *Application) GetFiatAccounts() (accounts []*model.Account, res *http.Response, err error) {
	if res, err = RequestAndDeserialize(http.MethodGet, "/fiat_accounts", nil, &accounts); err != nil {
		err = fmt.Errorf("can't get fiat accounts: %w", err)
	}
	return
}

func (a *Application) GetCryptoAccounts() (accounts []*model.Account, res *http.Response, err error) {
	if res, err = RequestAndDeserialize(http.MethodGet, "/crypto_accounts", nil, &accounts); err != nil {
		err = fmt.Errorf("can't get products: %w", err)
	}
	return
}

func (a *Application) GetProducts(perpetual bool) (products []*model.Product, res *http.Response, err error) {
	path := "/products"
	if perpetual {
		path += "?perpetual=1"
	}
	if res, err = RequestAndDeserialize(http.MethodGet, path, nil, &products); err != nil {
		err = fmt.Errorf("can't get products: %w", err)
	}
	return
}

func (a *Application) GetProduct(idOrPairCode string) (product *model.Product, res *http.Response, err error) {
	var id int
	if id, err = cli.GetProductID(idOrPairCode); err != nil {
		return
	}
	product = new(model.Product)
	if res, err = RequestAndDeserialize(http.MethodGet, fmt.Sprintf("/products/%d", id), nil, product); err != nil {
		err = fmt.Errorf("can't get product: %w", err)
	}
	return
}

func RequestAndDeserialize(httpMethod string, path string, body io.Reader, into interface{}) (*http.Response, error) {
	res, err := App.Request(httpMethod, path, body)
	if err != nil {
		return nil, fmt.Errorf("can't perform request: %w", err)
	}
	defer res.Body.Close()

	if err = CheckResponseStatus(res.StatusCode); err != nil {
		return res, fmt.Errorf("bad response status: %w", err)
	}

	if err = json.NewDecoder(res.Body).Decode(into); err != nil {
		return res, fmt.Errorf("can't deserialize response: %w", err)
	}

	log.Debugf("got response:\n%s", SprintResponse(res))

	return res, nil
}

func SprintResponse(res *http.Response) string {
	r, _ := httputil.DumpResponse(res, true)
	return fmt.Sprintf("got response: %s", r)
}

func CheckResponseStatus(statusCode int) error {
	if math.Floor(float64(statusCode)/100) != 2 {
		return fmt.Errorf("got %d: %s", statusCode, http.StatusText(statusCode))
	}
	return nil
}
