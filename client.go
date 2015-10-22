package billforward

import (
	"github.com/authclub/billforward/types"
	"golang.org/x/net/context"

	"net/url"
	// TOOD(pquerna): remove fmt.Errorf, use correct error types
	"fmt"
	"strings"
)

func New(c Config) (Client, error) {
	u, err := url.Parse(c.Endpoint)
	if err != nil {
		return nil, err
	}

	if !u.IsAbs() {
		return nil, fmt.Errorf("config.Endpoint must be an absolute url: %s", c.Endpoint)
	}

	config := Config(c)
	config.Endpoint = strings.TrimSuffix(u.String(), "/")
	return &client{
		config: config,
		// TODO(pquerna): should we allow the client to pass the context?
		ctx: context.Background(),
	}, nil
}

// Each major sub-system is a separate Interface which makes it easier to mock in test cases.
type Client interface {
	Accounts() AccountClient
	Invoices() InvoiceClient
	Products() ProductClient

	httpClientDo
}

type AccountIterateFn func(*types.Account) (bool, error)

type AccountClient interface {
	CreateAccount(profile *types.Profile) (*types.Account, error)
	GetAccountById(accountId string) (*types.Account, error)
	RetireAccountById(accountId string) error
	ListAccounts(filter *types.AccountsFilter, fn AccountIterateFn) error
}

type InvoiceIterateFn func(*types.Invoice) (bool, error)

type InvoiceClient interface {
	ListInvoicesByAccountId(accountId string, fn InvoiceIterateFn) error
}

type ProductIterateFn func(*types.Product) (bool, error)

type ProductClient interface {
	ListProducts(filter *types.ProductsFilter, fn ProductIterateFn) error
}

type client struct {
	config Config
	ctx    context.Context
}

func (c *client) Accounts() AccountClient {
	return c
}

func (c *client) Invoices() InvoiceClient {
	return c
}

func (c *client) Products() ProductClient {
	return c
}
