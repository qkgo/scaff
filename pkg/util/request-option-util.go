package util

type RequestOption struct {
	Page        int      `json:"page" description:"页码"`
	PageSize    int      `json:"pageSize" description:"分页大小"`
	Relation    bool     `json:"relation" description:"决定是否查询关联类"`
	State       string   `json:"state" description:"查询：强制查询state的值、修改和插入："`
	Specify     bool     `json:"specify" description:"精确查询"`
	AutoRank    bool     `json:"autoRank" description:"自动排序"`
	Order       string   `json:"order" description:"排序方式"`
	Count       string   `json:"count"  description:"按条件计算个数"`
	Strict      bool     `json:"strict" description:"使用add创建时，代表如果第一次创建失败则失败，在v2查询时，代表查询state为1的数据，今天也将state 设置为非空default1"`
	Aggregate   string   `json:"aggregate" description:"聚合方式,left/right/full"`
	ReturnField []string `json:"field" description:"返回的字段"`
	Timeout     int      `json:"timeout" description:"查询超时"`
}
