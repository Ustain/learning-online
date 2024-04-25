package Po

//import (
//	"time"
//)
//
//// MediaProcessHistory 代表媒资处理历史信息
//type MediaProcessHistory struct {
//	ID         int        `json:"id" gorm:"primaryKey"`             // 主键
//	FileID     string     `json:"fileId"`                           // 文件标识
//	Filename   string     `json:"filename"`                         // 文件名称
//	Bucket     string     `json:"bucket"`                           // 存储源
//	FilePath   string     `json:"filePath"`                         // 文件路径
//	Status     string     `json:"status"`                           // 状态，1:未处理，视频处理完成更新为2
//	CreateDate time.Time  `json:"createDate" gorm:"autoCreateTime"` // 上传时间
//	FinishDate *time.Time `json:"finishDate"`                       // 完成时间
//	URL        string     `json:"url"`                              // 媒资文件访问地址
//	ErrorMsg   string     `json:"errormsg"`                         // 失败原因
//	FailCount  int        `json:"failCount"`                        // 失败次数
//}
