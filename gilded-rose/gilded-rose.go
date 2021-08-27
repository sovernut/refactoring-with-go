package main

type Item struct {
	name            string
	sellIn, quality int
}

func isIn(name string, items []string) bool {
	for _, item := range items {
		if name == item {
			return true
		}
	}
	return false
}

func increaseQuality(item *Item) {
	if item.quality < 50 {
		item.quality++
	}
}

func decreaseQuality(item *Item) {
	if item.quality > 0 {
		item.quality--
	}
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		if isIn(items[i].name, []string{"Aged Brie", "Backstage passes to a TAFKAL80ETC concert", "Sulfuras, Hand of Ragnaros"}) {
			increaseQuality(items[i])
		} else {
			decreaseQuality(items[i])
		}

		if items[i].name == "Backstage passes to a TAFKAL80ETC concert" && items[i].quality < 50 {
			if items[i].sellIn < 11 {
				items[i].quality++
			}
			if items[i].sellIn < 6 {
				items[i].quality++
			}
		}

		if items[i].name != "Sulfuras, Hand of Ragnaros" {
			items[i].sellIn--
		}

		if items[i].sellIn < 0 {
			if items[i].name == "Backstage passes to a TAFKAL80ETC concert" {
				items[i].quality = 0
				continue
			}

			if isIn(items[i].name, []string{"Aged Brie", "Sulfuras, Hand of Ragnaros"}) {
				increaseQuality(items[i])
				continue
			}

			decreaseQuality(items[i])
		}
	}

}
