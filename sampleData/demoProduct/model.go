package demoProduct

type Product struct {
	ID                int
	ArticleNumber     string
	Name              string
	Description       string
	Color             string
	Size              string
	StockAvailability int
	PriceCents        int
	OnSale            bool
}

type ProductChangeSet struct {
	ArticleNumber     *string
	Name              *string
	Description       *string
	Color             *string
	Size              *string
	StockAvailability *int
	PriceCents        *int
	OnSale            *bool
}

func (c ProductChangeSet) toMap() map[string]interface{} {
	m := make(map[string]interface{})
	if c.ArticleNumber != nil {
		m["article_number"] = c.ArticleNumber
	}
	if c.Name != nil {
		m["name"] = c.Name
	}
	if c.Description != nil {
		m["description"] = c.Description
	}
	if c.Color != nil {
		m["color"] = c.Color
	}
	if c.Size != nil {
		m["size"] = c.Size
	}
	if c.StockAvailability != nil {
		m["stock_availability"] = c.StockAvailability
	}
	if c.PriceCents != nil {
		m["price_cents"] = c.PriceCents
	}
	if c.OnSale != nil {
		m["on_sale"] = c.OnSale
	}
	return m
}
