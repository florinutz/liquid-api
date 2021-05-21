package products

import (
	"fmt"
	"net/http"
	"strings"

	. "github.com/florinutz/liquid-api/pkg"

	"github.com/florinutz/liquid-api/pkg/cli"
	"github.com/spf13/cobra"
)

var productsOpts struct {
	perpetual bool
	idsMap    bool
}

func init() {
	ProductsCmd.Flags().BoolVarP(&productsOpts.perpetual, "perpetual", "p", false, "only show perpetual products")
	ProductsCmd.Flags().BoolVarP(&productsOpts.idsMap, "map", "m", false, "generate ProductIDsMap")
	_ = ProductsCmd.Flags().MarkHidden("map")
}

// ProductsCmd works shows products
var ProductsCmd = &cobra.Command{
	Use:          "products",
	Short:        "lists products",
	PreRunE:      InitAppE(&App),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			product, res, err := App.GetProduct(strings.Join(args, ""))
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

		if productsOpts.idsMap {
			idsMap := map[string]int{}
			for _, product := range products {
				idsMap[product.CurrencyPairCode] = product.ID
			}

			printedMap, err := cli.JsonPrint(idsMap, true)
			if err != nil {
				return fmt.Errorf("can't print products map: %w", err)
			}

			fmt.Printf("var ProductIDsMap = map[string]int%s", printedMap)

			return nil
		}

		for _, product := range products {
			fmt.Printf("%4d %s\n", product.ID, product.CurrencyPairCode)
		}

		return nil
	},
}
