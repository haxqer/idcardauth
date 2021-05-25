package idcardauth

const (
	NULL     = ""
	SUCCESS  = "SUCCESS"
	FAIL     = "FAIL"
	OK       = "OK"
	DebugOff = 0
	DebugOn  = 1

	CheckUrl        = "https://api.wlc.nppa.gov.cn/idcard/authentication/check"
	CheckTestUrl    = "https://wlc.nppa.gov.cn/test/authentication/check"
	QueryUrl        = "http://api2.wlc.nppa.gov.cn/idcard/authentication/query"
	QueryTestUrl    = "https://wlc.nppa.gov.cn/test/authentication/query"
	BehaviorUrl     = "http://api2.wlc.nppa.gov.cn/behavior/collection/loginout"
	BehaviorTestUrl = "https://wlc.nppa.gov.cn/test/collection/loginout"
)
