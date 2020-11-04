package xomologisou

type PublishableError int

const (
	PublishableErrorNoQuota PublishableError = iota
)

func (p PublishableError) Error() string {
	return string(p)
}
