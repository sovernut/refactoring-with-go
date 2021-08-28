package main
type Item struct {
	name string
	sellIn, quality int
}
func quality(it Item, q int) Item {
	it.quality = it.quality + q
	if it.quality > 50 {
		it.quality = 50
		return it
	}
	if it.quality < 0 {
		it.quality = 0
		return it
	}
	return it
}
type AgedBrie struct {
	item Item
}
func (ab AgedBrie) qualify() Genre {
	ab.item = quality(ab.item, +1)
	return ab
}
func (ab AgedBrie) updateSellIn() Genre {
	ab.item.sellIn = ab.item.sellIn - 1
	return ab
}
func (ab AgedBrie) update() Item {
	if ab.item.sellIn < 0 {
		ab.item = quality(ab.item, +1)
	}
	return ab.item
}
type Backstage struct {
	item Item
}
func (bs Backstage) qualify() Genre {
	item := bs.item
	if item.sellIn > 10 {
		item = quality(item, +1)
	}
	if item.sellIn > 5 && item.sellIn <= 10 {
		item = quality(item, +2)
	}
	if item.sellIn >= 0 && item.sellIn <= 5 {
		item = quality(item, +3)
	}
	return Backstage{item: item}
}
func (bs Backstage) updateSellIn() Genre {
	item := bs.item
	item.sellIn = item.sellIn - 1
	return Backstage{item: item}
}
func (bs Backstage) update() Item {
	item := bs.item
	if item.sellIn < 0 {
		item.quality = item.quality - item.quality
	}
	return item
}
type Genre interface {
	qualify() Genre
	updateSellIn() Genre
	update() Item
}
type Normal struct {
	item Item
}
func (ab Normal) qualify() Genre {
	ab.item = quality(ab.item, -1)
	return ab
}
func (ab Normal) updateSellIn() Genre {
	ab.item.sellIn = ab.item.sellIn - 1
	return ab
}
func (ab Normal) update() Item {
	if ab.item.sellIn < 0 {
		ab.item = quality(ab.item, -1)
	}
	return ab.item
}
func updateState(gn Genre) Item {
	it := gn.qualify().updateSellIn().update()
	return it
}
/*
 type AgedBrie struct {}
 type Backstage struct {}
 type Normal struct {}
 type Sulfuras struct {}
*/
func GenreFactory(item Item) Genre {
	switch item.name {
	case "Aged Brie":
		return AgedBrie{item: item}
	case "Backstage passes to a TAFKAL80ETC concert":
		return Backstage{item: item}
	// case "Sulfuras, Hand of Ragnaros":
	// return Sulfuras{item: item}
	default:
		return Normal{item: item}
	}
}
func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		item := items[i]
		if item.name == "Sulfuras, Hand of Ragnaros" {
			continue
		}
		it := updateState(GenreFactory(*item))
		items[i] = &Item{
			name: it.name,
			quality: it.quality,
			sellIn: it.sellIn,
		}
	}
}
