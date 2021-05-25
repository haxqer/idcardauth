package idcardauth

type AuthCheckRequest struct {
	AuthId string `json:"ai"`
	Name   string `json:"name"`
	IdCard string `json:"idNum"`
}

type AuthQueryRequest struct {
	AuthId string `json:"ai"`
}

type BehaviorRequest struct {
	Collections []*BehaviorParamCollections `json:"collections"`
}

type BehaviorParamCollections struct {
	No                       int    `json:"no"` //条目编码; 在批量模式中标识一条行为数据，取值范围1-128
	SessionId                string `json:"si"` //游戏内部会话标识
	BehaviorType             int    `json:"bt"` //用户行为类型
	OccurTime                int64  `json:"ot"` //行为发生时间
	CollectionType           int    `json:"ct"` //上报类型
	DeviceId                 string `json:"di,omitempty"` //设备标识
	GovernmentPlatformUserId string `json:"pi,omitempty"` //用户唯一标识
}
