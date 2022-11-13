package common

// OptionForEach 遍历
func OptionForEach[TStruct any, TOption ~func(TStruct)](t TStruct, opts []TOption) {
	for _, opt := range opts {
		opt(t)
	}
}
