package public

const (
	LoadTypeHTTP = 0

	HTTPRuleTypePrefixURL = 0
	HTTPRuleTypeDomain    = 1
)

var (
	LoadTypeMap = map[int]string{
		LoadTypeHTTP: "HTTP",
	}
)
