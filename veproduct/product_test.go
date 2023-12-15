package veproduct

import "testing"

func TestProduct_Exists(t *testing.T) {
	if expect, got := true, BMV700.Exists(); expect != got {
		t.Errorf("expect %#v but got %#v", expect, got)
	}

	if expect, got := true, PhoenixInverter12V1200VA230Vac64kHS.Exists(); expect != got {
		t.Errorf("expect %#v but got %#v", expect, got)
	}

	if expect, got := false, Product(0x1234).Exists(); expect != got {
		t.Errorf("expect %#v but got %#v", expect, got)
	}
}

func TestProduct_Model(t *testing.T) {
	if expect, got := "700", BMV700.Model(); expect != got {
		t.Errorf("expect %#v but got %#v", expect, got)
	}

	if expect, got := "24|16 (3)", PhoenixSmartIP43Charger24_16_3.Model(); expect != got {
		t.Errorf("expect %#v but got %#v", expect, got)
	}

	if expect, got := "", Product(0x1234).Model(); expect != got {
		t.Errorf("expect %#v but got %#v", expect, got)
	}
}

func TestProduct_Type(t *testing.T) {
	if expect, got := TypeBMV, BMV700.Type(); expect != got {
		t.Errorf("expect %#v but got %#v", expect, got)
	}

	if expect, got := TypePhoenixSmartIP43Charger, PhoenixSmartIP43Charger24_16_3.Type(); expect != got {
		t.Errorf("expect %#v but got %#v", expect, got)
	}

	if expect, got := TypeUnknown, Product(0x1234).Type(); expect != got {
		t.Errorf("expect %#v but got %#v", expect, got)
	}
}

func TestProduct_String(t *testing.T) {
	if expect, got := "BMV 700", BMV700.String(); expect != got {
		t.Errorf("expect %#v but got %#v", expect, got)
	}

	if expect, got := "Phoenix Smart IP43 Charger 24|16 (3)", PhoenixSmartIP43Charger24_16_3.String(); expect != got {
		t.Errorf("expect %#v but got %#v", expect, got)
	}
}

func Benchmark_String(b *testing.B) {
	errCnt := 0
	for n := 0; n < b.N; n++ {
		if expect, got := "BMV 700", BMV700.String(); expect != got {
			errCnt++
		}

		if expect, got := "Phoenix Smart IP43 Charger 24|16 (3)", PhoenixSmartIP43Charger24_16_3.String(); expect != got {
			errCnt++
		}
	}

	b.Logf("errCnt: %d", errCnt)
}

func TestGetStringMap(t *testing.T) {
	if expect, got := "BMV 700", GetStringMap()[BMV700]; expect != got {
		t.Errorf("expect %#v but got %#v", expect, got)
	}

	if expect, got := "Phoenix Smart IP43 Charger 24|16 (3)", GetStringMap()[PhoenixSmartIP43Charger24_16_3]; expect != got {
		t.Errorf("expect %#v but got %#v", expect, got)
	}
}
