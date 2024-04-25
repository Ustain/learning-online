package Po

//import (
//	"time"
//)
//
//// MediaFiles 代表媒资文件信息
//type MediaFiles struct {
//	ID          string    `json:"id" gorm:"primaryKey"`             // 主键
//	CompanyID   int       `json:"companyId"`                        // 机构ID
//	CompanyName string    `json:"companyName"`                      // 机构名称
//	Filename    string    `json:"filename"`                         // 文件名称
//	FileType    string    `json:"fileType"`                         // 文件类型（文档，音频，视频）
//	Tags        string    `json:"tags"`                             // 标签
//	Bucket      string    `json:"bucket"`                           // 存储目录
//	FilePath    string    `json:"filePath"`                         // 存储路径
//	FileID      string    `json:"fileId"`                           // 文件标识
//	URL         string    `json:"url"`                              // 媒资文件访问地址
//	Username    string    `json:"username"`                         // 上传人
//	CreateDate  time.Time `json:"createDate" gorm:"autoCreateTime"` // 上传时间
//	ChangeDate  time.Time `json:"changeDate" gorm:"autoUpdateTime"` // 修改时间
//	Status      string    `json:"status"`                           // 状态，1:未处理，视频处理完成更新为2
//	Remark      string    `json:"remark"`                           // 备注
//	AuditStatus string    `json:"auditStatus"`                      // 审核状态
//	AuditMind   string    `json:"auditMind"`                        // 审核意见
//	FileSize    int       `json:"fileSize"`                         // 文件大小
//}
