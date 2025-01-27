package veregister

import (
	"errors"
	"github.com/koestler/go-victron/veproduct"
	"testing"
)

func TestGetRegisterListByProduct(t *testing.T) {
	t.Run("unsupported type", func(t *testing.T) {
		_, err := GetRegisterListByProduct(veproduct.Product(0x1234))
		if !errors.Is(err, ErrUnsupportedType) {
			t.Errorf("expected an error for unknown product type")
		}
	})

	t.Run("supported types", func(t *testing.T) {
		countPerType := map[veproduct.Product]int{
			veproduct.BMV700:                                 27,
			veproduct.BMV712Smart:                            31,
			veproduct.BlueSolarMPPT70_15:                     48, // Panel current not available
			veproduct.BlueSolarMPPT75_50:                     43,
			veproduct.SmartSolarMPPT250_100:                  43,
			veproduct.PhoenixInverter12V250VA230V:            36,
			veproduct.PhoenixInverterSmart12V5000VA230Vac64k: 36,
			veproduct.SmartShunt500A_50mV:                    31,
		}

		for product, expectedCount := range countPerType {
			rl, err := GetRegisterListByProduct(product)
			if err != nil {
				t.Errorf("unexpected error for %s: %s", product, err)
			}

			count := len(rl.NumberRegisters)
			count += len(rl.TextRegisters)
			count += len(rl.EnumRegisters)
			count += len(rl.FieldListRegisters)

			if count != expectedCount {
				t.Errorf("unexpected count for %s: expected %d but got %d", product, expectedCount, count)
			}
		}
	})
}
