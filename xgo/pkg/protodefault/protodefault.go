package protodefault

type Default interface {
	SetDefault()
}

func AssertAndInvoke(candidate interface{}) {
	if candidate == nil {
		return
	}
	if dt, ok := candidate.(Default); ok {
		dt.SetDefault()
	}
}
