package data

import "testing"

func TestChecksCalidation(t *testing.T) {
	t.Run("without required", func(t *testing.T) {
		p := Product{}

		err := p.Validate()

		if err == nil {
			t.Error("Should return error")
		}
	})

	t.Run("with only name", func(t *testing.T) {
		p := Product{Name: "tea"}

		err := p.Validate()

		if err == nil {
			t.Error("Should return error")
		}
	})

	t.Run("with name and price", func(t *testing.T) {
		p := Product{Name: "tea", Price: 10}

		err := p.Validate()

		if err == nil {
			t.Error("Should return error")
		}
	})

	t.Run("with name and price and invalid sku", func(t *testing.T) {
		p := Product{Name: "tea", Price: 10, SKU: "ab-dre"}

		err := p.Validate()

		if err == nil {
			t.Error("Should return error")
		}
	})

	t.Run("with name and price and valid sku", func(t *testing.T) {
		p := Product{Name: "tea", Price: 10, SKU: "ab-dre-gaef"}

		err := p.Validate()

		if err != nil {
			t.Error("Should not return error")
		}
	})
}
