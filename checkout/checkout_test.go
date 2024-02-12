package checkout

import (
	"errors"
	"testing"
)

func TestScan(t *testing.T) {
    testCases := []struct{
        sku rune
        scannedItem ScannedItem
        scannedError error
    } {
        {
            sku: 'A',
            scannedItem: ScannedItem{
                quantityScanned: 1,
            },
            scannedError: nil,
        },
        {
            sku: 'A',
            scannedItem: ScannedItem{
                quantityScanned: 2,
            },
            scannedError: nil,
        },
        {
            sku: '|',
            scannedError: ErrItemDoesntExist,
        },
    }

    scanner := New()

    for _, tc := range testCases {
        item, err := scanner.Scan(tc.sku)
        if item.quantityScanned != tc.scannedItem.quantityScanned {
            t.Fatalf(
                "quantity scanned does not match expected. Wanted: %v, got: %v",
                tc.scannedItem.quantityScanned,
                item.quantityScanned,
            )
        }

        if !errors.Is(tc.scannedError, err) {
            t.Fatalf(
                "received mismatch of expected error value. Wanteed: %v, got: %v",
                tc.scannedError,
                err,
            )
        }
    }
}
