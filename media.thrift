namespace go xuetang

struct MediaFiles {
    1: required string id,           // 主键
    2: required i64 companyId,       // 机构ID
    3: required string companyName, // 机构名称
    4: required string filename,    // 文件名称
    5: required string fileType,    // 文件类型（文档，音频，视频）
    6: required string tags,        // 标签
    7: required string bucket,      // 存储目录
    8: required string filePath,    // 存储路径
    9: required string fileId,      // 文件标识
    10: required string url,        // 媒资文件访问地址
    11: required string username,   // 上传人
    12: required string createDate,    // 上传时间
    13: required string changeDate,    // 修改时间
    14: required string status,     // 状态，1:未处理，视频处理完成更新为2
    15: required string remark,     // 备注
    16: required string auditStatus,// 审核状态
    17: required string auditMind,  // 审核意见
    18: required i64 fileSize       // 文件大小
}

//媒资查询Model
struct PageParams {
    1: required i64 pageNo, // 当前页码
    2: required i64 pageSize // 每页显示记录数
}

struct PageResult {
    1: required list<MediaFiles> items, // 数据列表
    2: required i64 counts,              // 总记录数
    3: required i64 page,                // 当前页码
    4: required i64 pageSize             // 每页记录数
}

// 媒资上传Model
struct UploadFileParamsDto{
    1: required string filename  // 文件名称
    2: required string fileType // 文件类型（文档，音频，视频）
    3: required i64    fileSize // 文件大小
    4: required string tags     // 标签
    5: required string username // 上传人
    6: required string remark   // 备注
}

struct UploadFileResultDto {
    1: MediaFiles mediaFiles
}

service Media{
     PageResult QueryMediaFiles(1:PageParams req)
     UploadFileResultDto UploadMediaFiles(1:UploadFileParamsDto req, 2: string filePath)
}