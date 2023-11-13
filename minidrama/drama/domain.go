/*
 * Copyright Bytedance Author(https://houseme.github.io/bytedance/). All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * You can obtain one at https://github.com/houseme/bytedance.
 *
 */

package drama

// ResourceUploadImageRequest 资源上传 图片
type ResourceUploadImageRequest struct {
    ResourceType int        `json:"resource_type"`
    MaAppID      string     `json:"ma_app_id"`
    ImageMeta    *ImageMeta `json:"image_meta"`
}

// ImageMeta 图片元数据
type ImageMeta struct {
    URL string `json:"url"`
}

// ResourceUploadImageResponse 资源上传 图片
type ResourceUploadImageResponse struct {
    Data   *ImageData `json:"data"`
    ErrMsg string     `json:"err_msg"`
    ErrNo  int        `json:"err_no"`
    LogID  string     `json:"log_id"`
}

// ImageResult 图片结果
type ImageResult struct {
    OpenPicID string `json:"open_pic_id"`
}

// ImageData 资源上传 图片
type ImageData struct {
    ImageResult  *ImageResult `json:"image_result"`
    ResourceType int          `json:"resource_type"`
}

// ResourceUploadVideoRequest 资源上传 视频
type ResourceUploadVideoRequest struct {
    ResourceType int        `json:"resource_type"`
    MaAppID      string     `json:"ma_app_id"`
    VideoMeta    *VideoMeta `json:"video_meta"`
}

// VideoMeta 视频元数据
type VideoMeta struct {
    Title       string `json:"title"`
    Description string `json:"description,omitempty"`
    URL         string `json:"url"`
    Format      string `json:"format"`
    UseDyCloud  bool   `json:"use_dy_cloud"`
    DyCloudID   string `json:"dy_cloud_id,omitempty" description:"抖音云的视频 ID，用于从抖音云同步视频到内容库，该字段有值时，URL、use_dy_cloud 无实际作用"`
}

// ResourceUploadVideoResponse 资源上传 视频
type ResourceUploadVideoResponse struct {
    ErrMsg string     `json:"err_msg"`
    ErrNo  int        `json:"err_no"`
    LogID  string     `json:"log_id"`
    Data   *VideoData `json:"data"`
}

// VideoData 资源上传 视频
type VideoData struct {
    ResourceType int          `json:"resource_type"`
    VideoResult  *VideoResult `json:"video_result"`
}

// VideoResult 视频结果
type VideoResult struct {
    OpenVideoID string `json:"open_video_id"`
}

// QueryVideoRequest 查询视频
type QueryVideoRequest struct {
    MaAppID     string `json:"ma_app_id"`
    VideoIdType int    `json:"video_id_type" description:"视频 ID 类型，1：抖音视频 ID，2：抖音云视频 ID"`
    OpenVideoID string `json:"open_video_id,omitempty" description:"video_id_type 为 1 必填"`
    DyCloudID   string `json:"dy_cloud_id,omitempty" description:"video_id_type 为 2 必填"`
}

// QueryVideoResponse 查询视频
type QueryVideoResponse struct {
    Data   *QueryVideoData `json:"data"`
    ErrMsg string          `json:"err_msg"`
    ErrNo  int             `json:"err_no"`
    LogID  string          `json:"log_id"`
}

// QueryVideoData 查询视频
type QueryVideoData struct {
    OpenVideoID string `json:"open_video_id"`
    Status      string `json:"status"`
    DyCloudID   string `json:"dy_cloud_id"`
}

// Tag 标签
type Tag struct {
    ID    string `json:" id "`
    Value string `json:" value "`
}

// CreateVideoRequest 创建视频
type CreateVideoRequest struct {
    MaAppID   string     `json:"ma_app_id"`
    AlbumInfo *AlbumInfo `json:"album_info"`
}

// AlbumInfo 专辑信息
type AlbumInfo struct {
    Title          string      `json:"title" description:"短剧标题"`
    SeqNum         int         `json:"seq_num" description:"总集数"`
    Year           int         `json:"year" description:"短剧发行年份"`
    AlbumStatus    int         `json:"album_status" description:"短剧更新状态：1-未上映，2-更新中，3-已完结"`
    Qualification  int         `json:"qualification" description:"资质状态：1-未报审，2-报审通过，3-报审不通过，4-不建议报审"`
    RecordInfo     *RecordInfo `json:"record_info" description:"备案信息"`
    Desp           string      `json:"desp" description:"短剧简介（200 汉字以内）"`
    Recommendation string      `json:"recommendation" description:"短剧推荐语（12 汉字以内）"`
    TagList        []int       `json:"tag_list" description:"短剧标签，最多 3 个"`
    CoverList      []string    `json:"cover_list" description:"专辑封面图，最多 3 张"`
}

// RecordInfo 记录信息
type RecordInfo struct {
    LicenseNum        string `json:"license_num,omitempty"`
    RegistrationNum   string `json:"registration_num,omitempty"`
    OrdinaryRecordNum string `json:"ordinary_record_num,omitempty"`
    KeyRecordNum      string `json:"key_record_num,omitempty"`
}

// CreateVideoResponse 创建视频
type CreateVideoResponse struct {
    Data   *CreateVideoData `json:"data"`
    ErrMsg string           `json:"err_msg"`
    ErrNo  int              `json:"err_no"`
    LogID  string           `json:"log_id"`
}

// CreateVideoData 创建视频
type CreateVideoData struct {
    AlbumID int64 `json:"album_id"`
}

// EditVideoRequest 编辑视频
type EditVideoRequest struct {
    MaAppID         string         `json:"ma_app_id"`
    AlbumID         int64          `json:"album_id"`
    AlbumInfo       *AlbumInfo     `json:"album_info,omitempty" description:"短剧信息，会覆盖之前的短剧信息"`
    EpisodeInfoList []*EpisodeInfo `json:"episode_info_list,omitempty" description:"剧集信息列表，剧集信息 单次（需要小于 100），总数没有限制"`
}

// EpisodeInfo 专辑信息
type EpisodeInfo struct {
    Title       string   `json:"title"`
    Seq         int      `json:"seq"`
    CoverList   []string `json:"cover_list"`
    OpenVideoID string   `json:"open_video_id" description:"剧集对应的视频，抖音开放平台的视频 ID"`
}

// EditVideoResponse 编辑视频
type EditVideoResponse struct {
    LogID  string         `json:"log_id"`
    ErrNo  int            `json:"err_no"`
    Data   *EditVideoData `json:"data"`
    ErrMsg string         `json:"err_msg"`
}

// EditVideoData  编辑视频
type EditVideoData struct {
    AlbumID      int64             `json:"album_id"`
    Version      int               `json:"version" description:"短剧的版本号（审核后再编辑版本会增加）"`
    EpisodeIdMap map[string]string `json:"episode_id_map" description:"剧集 ID 映射，key 为 seq，value 为剧集 ID"`
}

// QueryVideoAlbumRequest 查询视频专辑
type QueryVideoAlbumRequest struct {
    MaAppID     string       `json:"ma_app_id"`
    QueryType   int          `json:"query_type"`
    DetailQuery *DetailQuery `json:"detail_query,omitempty"`
    BatchQuery  *BatchQuery  `json:"batch_query,omitempty"`
    SingleQuery *SingleQuery `json:"single_query,omitempty"`
}

// BatchQuery 批量查询
type BatchQuery struct {
    Limit  int `json:"limit"`
    Offset int `json:"offset"`
}

// SingleQuery 单个查询
type SingleQuery struct {
    AlbumID int64 `json:"album_id"`
    Limit   int   `json:"limit"`
    Offset  int   `json:"offset"`
}

// DetailQuery 详情查询
type DetailQuery struct {
    AlbumID int64 `json:"album_id"`
    Limit   int   `json:"limit"`
    Offset  int   `json:"offset"`
    Version int   `json:"version"`
}

// QueryVideoAlbumResponse 查询视频专辑
type QueryVideoAlbumResponse struct {
    ErrNo int `json:"err_no"`
    Data  struct {
        BatchData  *BatchData  `json:"batch_data,omitempty"`
        SingleData *SingleData `json:"single_data,omitempty"`
        DetailData *DetailData `json:"detail_data,omitempty"`
    } `json:"data"`
    ErrMsg string `json:"err_msg"`
    LogID  string `json:"log_id"`
}

// BatchData 批量数据
type BatchData struct {
    Total         int              `json:"total"`
    AlbumInfoList []*AlbumInfoResp `json:"album_info_list"`
}

type SingleData struct {
    Total         int              `json:"total"`
    AlbumInfoList []*AlbumInfoResp `json:"album_info_list"`
}

// AlbumAuditInfo 专辑审核信息
type AlbumAuditInfo struct {
    Status  int   `json:"status"`
    AlbumID int64 `json:"album_id"`
    Version int   `json:"version"`
}

// AlbumInfoResp 专辑信息
type AlbumInfoResp struct {
    AlbumStatus    int             `json:"album_status"`
    Recommendation string          `json:"recommendation"`
    SeqNum         int             `json:"seq_num"`
    Title          string          `json:"title"`
    CoverList      []string        `json:"cover_list"`
    TagList        []int           `json:"tag_list"`
    AlbumAuditInfo *AlbumAuditInfo `json:"album_audit_info"`
    Year           int             `json:"year"`
    Qualification  int             `json:"qualification"`
    Desp           string          `json:"desp"`
    RecordInfo     *RecordInfoResp `json:"record_info"`
}

// RecordInfoResp 记录信息
type RecordInfoResp struct {
    LicenseNum        string `json:"license_num"`
    RegistrationNum   string `json:"registration_num"`
    OrdinaryRecordNum string `json:"ordinary_record_num"`
    KeyRecordNum      string `json:"key_record_num"`
}

// EpisodeAuditInfo 剧集审核信息
type EpisodeAuditInfo struct {
    EpisodeID int64 `json:"episode_id"`
    Version   int   `json:"version"`
    Status    int   `json:"status"`
}

// EpisodeInfoResp 剧集信息
type EpisodeInfoResp struct {
    Title            string            `json:"title"`
    CoverList        []string          `json:"cover_list"`
    OpenVideoID      string            `json:"open_video_id"`
    EpisodeAuditInfo *EpisodeAuditInfo `json:"episode_audit_info"`
    Seq              int               `json:"seq"`
}

// DetailData 详情数据
type DetailData struct {
    Total           int                `json:"total"`
    EpisodeInfoList []*EpisodeInfoResp `json:"episode_info_list"`
}

// ReviewVideoRequest 审核视频 短剧送审
type ReviewVideoRequest struct {
    MaAppID string `json:"ma_app_id"`
    AlbumID int64  `json:"album_id"`
}

// ReviewVideoResponse 审核视频
type ReviewVideoResponse struct {
    ErrNo  int    `json:"err_no"`
    ErrMsg string `json:"err_msg"`
    LogID  string `json:"log_id"`
}

// AuthorizeVideoRequest 授权视频
type AuthorizeVideoRequest struct {
    MaAppID   string   `json:"ma_app_id"`
    AlbumID   int64    `json:"album_id"`
    AppIdList []string `json:"app_id_list"`
}

// AuthorizeVideoResponse 授权视频
type AuthorizeVideoResponse struct {
    ErrNo  int    `json:"err_no"`
    ErrMsg string `json:"err_msg"`
    LogID  string `json:"log_id"`
}

// OnlineAlbumRequest 上线专辑
type OnlineAlbumRequest struct {
    AlbumID int64 `json:"album_id"`
    Operate int   `json:"operate" desc:"操作类型 1-查询 2-修改"`
    Version int   `json:"version" desc:"上线的版本，操作类型为 2 时必填，默认为 0"`
}

// OnlineAlbumResponse 上线专辑
type OnlineAlbumResponse struct {
    Data   *OnlineAlbumData `json:"data"`
    ErrNo  int              `json:"err_no"`
    LogID  string           `json:"log_id"`
    ErrMsg string           `json:"err_msg"`
}

// OnlineAlbumData 上线专辑
type OnlineAlbumData struct {
    AlbumID int64 `json:"album_id"`
    Version int   `json:"version"`
}

// BindAlbumRequest 绑定专辑
type BindAlbumRequest struct {
    SchemaBindType   int               `json:"schema_bind_type" description:"绑定类型：1-单集绑定"`
    SingleSchemaBind *SingleSchemaBind `json:"single_schema_bind"`
}

// SingleSchemaBind 单个绑定
type SingleSchemaBind struct {
    AlbumID   int64     `json:"album_id"`
    EpisodeID int64     `json:"episode_id"`
    Path      string    `json:"path"`
    Params    []*Params `json:"params,omitempty"`
}

// Params 参数
type Params struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}

// BindAlbumResponse 绑定专辑
type BindAlbumResponse struct {
    ErrNo  int    `json:"err_no"`
    LogID  string `json:"log_id"`
    ErrMsg string `json:"err_msg"`
}

// PlayInfoRequest 播放信息
type PlayInfoRequest struct {
    MaAppID   string `json:"ma_app_id"`
    AlbumID   int64  `json:"album_id"`
    EpisodeID int64  `json:"episode_id"`
}

// PlayInfoResponse 播放信息
type PlayInfoResponse struct {
    Data   *PlayInfoData `json:"data"`
    ErrMsg string        `json:"err_msg"`
    ErrNo  int           `json:"err_no"`
    LogID  string        `json:"log_id"`
}

// PlayInfoData 播放信息
type PlayInfoData struct {
    Definition string `json:"definition"`
    Format     string `json:"format"`
    PlayURL    string `json:"play_url"`
    Size       int    `json:"size"`
    UrlExpire  string `json:"url_expire"`
    Bitrate    int    `json:"bitrate"`
    Codec      string `json:"codec"`
}

// AsyncRequest 异步请求
type AsyncRequest struct {
    Msg     string `json:"msg"`
    Type    string `json:"type" description:"消息类型 album_audit,episode_audit,upload_video"`
    Version string `json:"version" description:"版本号，默认为 2.0"`
}

// AsyncAlbumAudit 异步专辑审核
type AsyncAlbumAudit struct {
    MaAppID     string `json:"ma_app_id"`
    AlbumID     int64  `json:"album_id"`
    Version     int    `json:"version"`
    AuditStatus int    `json:"audit_status" description:"审核状态：99-未审核：98-审核中，1-不可播放，2-可播放"`
    AuditMsg    string `json:"auditMsg" description:"审核备注"`
}

// AsyncEpisodeAudit 剧集审核
type AsyncEpisodeAudit struct {
    MaAppID     string `json:"ma_app_id"`
    AlbumID     int64  `json:"album_id"`
    EpisodeID   int64  `json:"episode_id"`
    Version     int    `json:"version"`
    AuditStatus int    `json:"audit_status" description:"审核状态：99-未审核：98-审核中，1-不可播放，2-可播放"`
    AuditMsg    string `json:"auditMsg" description:"审核备注"`
}

// AsyncUploadVideo 上传视频
type AsyncUploadVideo struct {
    MaAppID     string `json:"ma_app_id"`
    OpenVideoID string `json:"open_video_id"`
    Success     bool   `json:"success"`
}

// AsyncResponse 异步响应
type AsyncResponse struct {
    ErrNo   int    `json:"err_no"`
    ErrTips string `json:"err_tips"`
}
