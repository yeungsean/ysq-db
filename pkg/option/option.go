package option

// ForEach 遍历
func ForEach[TStruct any, TOption ~func(TStruct)](t TStruct, opts []TOption) {
	for _, opt := range opts {
		opt(t)
	}
}

// Option 公共option
type Option struct {
	// Quote 是否用引号包含
	Quote bool
	// Alias 别名
	Alias string
}

// Options 可选项
type Options func(*Option)

// WithQuote 引用设置
func WithQuote(qs ...bool) Options {
	return func(o *Option) {
		q := true
		if len(qs) > 0 {
			q = qs[0]
		}
		o.Quote = q
	}
}

// WithAlias 别名设置
func WithAlias(a string) Options {
	return func(o *Option) {
		o.Alias = a
	}
}
