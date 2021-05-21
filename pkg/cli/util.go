package cli

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func GetProductID(idOrPairCode string) (id int, err error) {
	if id, err = strconv.Atoi(idOrPairCode); err != nil { // not int
		var exists bool
		if id, exists = ProductIDsMap[strings.ToUpper(idOrPairCode)]; !exists {
			err = fmt.Errorf("%s should either be a numeric id that you got from listing products "+
				"or a product code like ETHEUR", idOrPairCode)
			return
		}
	}
	err = nil
	numericIDIsValid := false
	for _, codeID := range ProductIDsMap {
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

func JsonPrint(v interface{}, pretty bool) (str string, err error) {
	var b []byte
	if pretty {
		b, err = json.MarshalIndent(v, "", "  ")
	} else {
		b, err = json.Marshal(v)
	}
	if err != nil {
		return "", fmt.Errorf("can't json marshal data: %w", err)
	}
	return string(b), nil
}
