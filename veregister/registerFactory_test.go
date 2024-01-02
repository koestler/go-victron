package veregister

import (
	"github.com/koestler/go-victron/veproduct"
	"testing"
)

func TestGetRegisterListByProductType(t *testing.T) {
	t.Run("unsupported type", func(t *testing.T) {
		_, err := GetRegisterListByProductType(veproduct.TypeUnknown)
		if err == nil {
			t.Errorf("expected an error for unknown product type")
		}
	})

	t.Run("supported types", func(t *testing.T) {
		countPerType := map[veproduct.Type]int{
			veproduct.TypeBMV:                  27,
			veproduct.TypeBMVSmart:             31,
			veproduct.TypeBlueSolarMPPT:        42,
			veproduct.TypeSmartSolarMPPT:       42,
			veproduct.TypePhoenixInverter:      36,
			veproduct.TypePhoenixInverterSmart: 36,
			veproduct.TypeSmartShunt:           31,
		}

		for productType, expectedCount := range countPerType {
			rl, err := GetRegisterListByProductType(productType)
			if err != nil {
				t.Errorf("unexpected error for %s: %s", productType, err)
			}

			count := len(rl.NumberRegisters)
			count += len(rl.TextRegisters)
			count += len(rl.EnumRegisters)
			count += len(rl.FieldListRegisters)

			if count != expectedCount {
				t.Errorf("unexpected count for %s: expected %d but got %d", productType, expectedCount, count)
			}
		}
	})
}
