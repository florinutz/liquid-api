package accounts

import (
	"fmt"

	. "github.com/florinutz/liquid-api/pkg"
	"github.com/spf13/cobra"
)

var accountOpts struct {
	fiat bool
}

func init() {
	AccountsCmd.Flags().BoolVarP(&accountOpts.fiat, "fiat", "f", false, "show fiat instead of crypto accounts")
}

// AccountsCmd represents the fiat accounts command
var AccountsCmd = &cobra.Command{
	Use:          "accounts",
	Short:        "your account's status",
	SilenceUsage: true,
	PreRunE:      InitAppE(&App),
	RunE: func(cmd *cobra.Command, args []string) error {
		getter := App.GetCryptoAccounts
		if accountOpts.fiat {
			getter = App.GetFiatAccounts
		}

		accounts, _, err := getter()
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
