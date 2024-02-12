package checkout

type SpecialPrice struct {
    Quantity uint
    Price int
}

type Item struct {
    SKU rune
    UnitPrice int
    SpecialPrice SpecialPrice
}

var allItems []Item = make([]Item, 0, 26)

func init() {

    allItems = append(allItems, Item{
        SKU: 'A',
        UnitPrice: 50,
        SpecialPrice: SpecialPrice{
            Quantity: 3,
            Price: 130,
        },
    })

    allItems = append(allItems, Item{
        SKU: 'B',
        UnitPrice: 30,
        SpecialPrice: SpecialPrice{
            Quantity: 2,
            Price: 45,
        },
    })

    allItems = append(allItems, Item{
        SKU: 'C',
        UnitPrice: 20,
    })

    allItems = append(allItems, Item{
        SKU: 'D',
        UnitPrice: 15,
    })
}
