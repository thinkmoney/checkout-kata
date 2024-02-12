package checkout

import "errors"

type ScannedItem struct {
	quantityScanned uint
}

type Scanner struct {
	scannedItems map[rune]ScannedItem
}

func New() Scanner {
    return Scanner{
        scannedItems: make(map[rune]ScannedItem),
    }
}

var (
    ErrItemDoesntExist = errors.New("Item SKU does not exist in system")
    emptyScannedItem = ScannedItem{}
)

func (s *Scanner) Scan(SKU rune) (ScannedItem, error) {
	if _, itemExists := allItems[SKU]; !itemExists {
		return emptyScannedItem, ErrItemDoesntExist
    }
        
	if scannedItem, scannedItemExists := s.scannedItems[SKU]; scannedItemExists {
		scannedItem.quantityScanned++
		s.scannedItems[SKU] = scannedItem
        return scannedItem, nil
	} else {
        toReturn := ScannedItem{
            quantityScanned: 1,
        }
		s.scannedItems[SKU] = toReturn
        return toReturn, nil
	}
}
