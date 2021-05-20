package status

import (
	"errors"
	"fmt"
	"net/http"

	. "github.com/florinutz/liquid-api/pkg"
	"github.com/spf13/cobra"
)

var productsOpts struct {
	perpetual bool
}

func init() {
	ProductsCmd.Flags().BoolVarP(&productsOpts.perpetual, "perpetual", "p", false, "only show perpetual products")
}

// ProductsCmd works shows products
var ProductsCmd = &cobra.Command{
	Use:          "products",
	Short:        "shows products",
	PreRunE:      InitAppE(&App),
	SilenceUsage: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("this commands accepts maximum one argument: the product id (e.g. ETHEUR)")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			product, res, err := App.GetProduct(args[0])
			if err != nil {
				if res != nil {
					if res.StatusCode == http.StatusNotFound {
						return fmt.Errorf("%s was not found", args[0])
					}
				}
				return err
			}
			fmt.Printf("%+v\n", product)
			return nil
		}

		products, _, err := App.GetProducts(productsOpts.perpetual)
		if err != nil {
			return err
		}
		for _, product := range products {
			fmt.Printf("%4d %s\n", product.ID, product.CurrencyPairCode)
		}
		return nil
	},
}

// FiatAccountsCmd represents the fiat accounts command
var FiatAccountsCmd = &cobra.Command{
	Use:          "fiat",
	Short:        "your account's status",
	SilenceUsage: true,
	PreRunE:      InitAppE(&App),
	RunE: func(cmd *cobra.Command, args []string) error {
		accounts, _, err := App.GetFiatAccounts()
		if err != nil {
			return err
		}
		for _, account := range accounts {
			if account.Balance < 0.009 {
				continue
			}
			fmt.Printf("%s: %.2f %s\n", account.Currency, account.Balance, account.CurrencySymbol)
		}
		return nil
	},
}

// CryptoAccountsCmd represents the crypto accounts command
var CryptoAccountsCmd = &cobra.Command{
	Use:          "crypto",
	Short:        "your account's status",
	SilenceUsage: true,
	PreRunE:      InitAppE(&App),
	RunE: func(cmd *cobra.Command, args []string) error {
		accounts, _, err := App.GetCryptoAccounts()
		if err != nil {
			return err
		}
		for _, account := range accounts {
			if account.Balance < 0.009 {
				continue
			}
			fmt.Printf("%s: %.2f %s\n", account.Currency, account.Balance, account.CurrencySymbol)
		}
		return nil
	},
}
