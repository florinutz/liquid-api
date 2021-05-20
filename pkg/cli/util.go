package cli

import (
	"fmt"
	"strconv"
	"strings"
)

func NormalizeProductID(idOrPairCode string) (id int, err error) {
	if id, err = strconv.Atoi(idOrPairCode); err != nil { // not int
		var exists bool
		if id, exists = productIDsMap[strings.ToUpper(idOrPairCode)]; !exists {
			err = fmt.Errorf("%s should either be a numeric id that you got from listing products "+
				"or a product code like ETHEUR", idOrPairCode)
			return
		}
	}
	err = nil
	numericIDIsValid := false
	for _, codeID := range productIDsMap {
		if codeID == id {
			numericIDIsValid = true
			break
		}
	}
	if !numericIDIsValid {
		err = fmt.Errorf("id %d does not correspond to any known product", id)
	}

	return
}
