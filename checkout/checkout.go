package checkout

import "errors"

type ScannedItem struct {
    quantityScanned int
}

type Scanner struct {
    scannedItems map[rune]ScannedItem
}

func (s *Scanner) Scan(SKU rune) error {
    if _, itemExists := allItems[SKU]; !itemExists {
        return errors.New("Item SKU does not exist in system")
    }

    if scannedItem, scannedItemExists := s.scannedItems[SKU]; scannedItemExists { 
        scannedItem.quantityScanned++
        s.scannedItems[SKU] = scannedItem
    } else {
        s.scannedItems[SKU] = ScannedItem{}
    }

    return nil
}
