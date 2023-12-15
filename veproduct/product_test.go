package veproduct

import "testing"

func TestProduct_String(t *testing.T) {
	if expect, got := "BMV-700", BMV700.String(); expect != got {
		t.Errorf("expect %#v but got %#v", expect, got)
	}

	if expect, got := "BMV-702", BMV702.String(); expect != got {
		t.Errorf("expect %#v but got %#v", expect, got)
	}

	if expect, got := "Phoenix Smart IP43 Charger 24|16 (3)", PhoenixSmartIP43Charger24_16_3.String(); expect != got {
		t.Errorf("expect %#v but got %#v", expect, got)
	}
}

func Benchmark_String(b *testing.B) {
	errCnt := 0
	for n := 0; n < b.N; n++ {
		if expect, got := "BMV-700", BMV700.String(); expect != got {
			errCnt++
		}

		if expect, got := "BMV-702", BMV702.String(); expect != got {
			errCnt++
		}

		if expect, got := "Phoenix Smart IP43 Charger 24|16 (3)", PhoenixSmartIP43Charger24_16_3.String(); expect != got {
			errCnt++
		}
	}

	b.Logf("errCnt: %d", errCnt)
}
