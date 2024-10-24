package consts

var (
	MultiLineGenerateType = []string{"YN?NN", "YN?NY"}
)

type AcceptStatus string

const (
	CodeAcceptStatusAccept    AcceptStatus = "accept"
	CodeAcceptStatusCancelled AcceptStatus = "cancelled"
	CodeAcceptStatusReject    AcceptStatus = "reject"
)

func (a AcceptStatus) String() string {
	return string(a)
}
