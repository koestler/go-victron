package veregister

import "fmt"

func (r RegisterStruct) testString() string {
	return fmt.Sprintf("category=%s, name=%s, description=%s, sort=%d, address=0x%X, static=%t, writable=%t",
		r.category, r.name, r.description, r.sort, r.address, r.static, r.writable,
	)
}

func (r NumberRegisterStruct) testString() string {
	return fmt.Sprintf("Number: %s, signed=%t, factor=%d, offset=%f, unit=%s",
		r.RegisterStruct.testString(), r.signed, r.factor, r.offset, r.unit,
	)
}

func (r TextRegisterStruct) testString() string {
	return fmt.Sprintf("Text: %s", r.RegisterStruct.testString())
}

func (r EnumRegisterStruct) testString() string {
	return fmt.Sprintf("Enum: %s, bit=%d, enum=%v", r.RegisterStruct.testString(), r.bit, r.enum)
}
