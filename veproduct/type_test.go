package veproduct

import "testing"

func TestType_String(t *testing.T) {
	if expect, got := "BMV", TypeBMV.String(); expect != got {
		t.Errorf("expect %#v but got %#v", expect, got)
	}

	if expect, got := "BlueSolar MPPT", TypeBlueSolarMPPT.String(); expect != got {
		t.Errorf("expect %#v but got %#v", expect, got)
	}
}

func TestType_IsBMV(t *testing.T) {
	if expect, got := true, TypeBMV.IsBMV(); expect != got {
		t.Errorf("expect %#v but got %#v", expect, got)
	}

	if expect, got := false, TypeBlueSolarMPPT.IsBMV(); expect != got {
		t.Errorf("expect %#v but got %#v", expect, got)
	}
}

func TestType_IsSolar(t *testing.T) {
	if expect, got := false, TypeBMV.IsSolar(); expect != got {
		t.Errorf("expect %#v but got %#v", expect, got)
	}

	if expect, got := true, TypeBlueSolarMPPT.IsSolar(); expect != got {
		t.Errorf("expect %#v but got %#v", expect, got)
	}
}

func TestType_IsInverter(t *testing.T) {
	if expect, got := false, TypeBMV.IsInverter(); expect != got {
		t.Errorf("expect %#v but got %#v", expect, got)
	}

	if expect, got := true, TypePhoenixInverter.IsInverter(); expect != got {
		t.Errorf("expect %#v but got %#v", expect, got)
	}
}
