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

package voc

// QueryListRequest 查询列表请求
type QueryListRequest struct {
    PageSize       int    `json:"page_size"`
    PageNumber     int    `json:"page_number"`
    EnvId          string `json:"env_id"`
    BusinessStatus string `json:"business_status,omitempty"`
    VideoName      string `json:"video_name,omitempty"`
}

// QueryListResponse 查询列表响应
type QueryListResponse struct {
    ErrMsg   string           `json:"err_msg"`
    ErrNo    int              `json:"err_no"`
    LogID    string           `json:"log_id"`
    TotalNum int              `json:"total_num"`
    Data     []*QueryDataItem `json:"data"`
}

// QueryDataItem 查询数据项
type QueryDataItem struct {
    BusinessStatus int    `json:"business_status"`
    Size           int    `json:"size"`
    UploadTime     int    `json:"upload_time"`
    Vid            string `json:"vid"`
    VideoName      string `json:"video_name"`
}

// DeleteVideoRequest 删除视频请求
type DeleteVideoRequest struct {
    EnvId string `json:"env_id"`
    Vid   string `json:"vid"`
}

// DeleteVideoResponse 删除视频响应
type DeleteVideoResponse struct {
    ErrMsg string `json:"err_msg"`
    ErrNo  int    `json:"err_no"`
    LogID  string `json:"log_id"`
}

// QueryVideoURLRequest 查询视频地址请求
type QueryVideoURLRequest struct {
    Vid string `json:"vid"`
}

// QueryVideoURLResponse 查询视频地址响应
type QueryVideoURLResponse struct {
    LogID  string               `json:"log_id"`
    Data   []*QueryVideoURLData `json:"data"`
    ErrMsg string               `json:"err_msg"`
    ErrNo  int                  `json:"err_no"`
}

// QueryVideoURLData 查询视频地址数据
type QueryVideoURLData struct {
    Format        string `json:"format"`
    MainPlayUrl   string `json:"main_play_url"`
    Size          int    `json:"size"`
    BackUpPlayUrl string `json:"back_up_play_url"`
    Bitrate       int    `json:"bitrate"`
    Codec         string `json:"codec"`
    Definition    string `json:"definition"`
    BackUrlExpire string `json:"back_url_expire"`
    MainUrlExpire string `json:"main_url_expire"`
}

// UploadByURLRequest 上传视频请求
// 源视频的 urlSet 列表，一次最多传入 20 条
type UploadByURLRequest struct {
    UrlSets []*URLSets `json:"url_sets"`
}

// URLSets 上传视频请求
type URLSets struct {
    SourceURL string `json:"source_url"`
    FileName  string `json:"file_name"`
}

// UploadByURLResponse 上传视频响应
type UploadByURLResponse struct {
    LogID  string            `json:"log_id"`
    Data   *UploadByURLRData `json:"data"`
    ErrMsg string            `json:"err_msg"`
    ErrNo  int               `json:"err_no"`
}

// UploadByURLRData 上传视频响应
type UploadByURLRData struct {
    UrlJobs []*URLJob `json:"url_jobs"`
}

// URLJob 上传视频响应
type URLJob struct {
    JobID     string `json:"job_id"`
    SourceURL string `json:"source_url"`
}

// QueryUploadByURLRequest 查询上传视频请求
type QueryUploadByURLRequest struct {
    AccessToken string   `json:"access_token"`
    JobIds      []string `json:"job_ids"`
}

// QueryUploadByURLResponse 查询上传视频响应
type QueryUploadByURLResponse struct {
    Data   *QueryUploadByURLData `json:"data"`
    ErrMsg string                `json:"err_msg"`
    ErrNo  int                   `json:"err_no"`
    LogID  string                `json:"log_id"`
}

// QueryUploadByURLData 查询上传视频响应
type QueryUploadByURLData struct {
    NotExistJobIds []string     `json:"not_exist_job_ids"`
    VideoInfos     []*VideoInfo `json:"video_infos"`
}

// VideoInfo 上传视频响应
type VideoInfo struct {
    State      string      `json:"state"`
    ViD        string      `json:"vid"`
    JobID      string      `json:"job_id"`
    RequestID  string      `json:"request_id"`
    SourceInfo *SourceInfo `json:"source_info"`
    SourceURL  string      `json:"source_url"`
}

// SourceInfo 上传视频响应
type SourceInfo struct {
    FileType string  `json:"file_type"`
    Format   string  `json:"format"`
    Size     int     `json:"size"`
    StoreURI string  `json:"store_uri"`
    Width    int     `json:"width"`
    Bitrate  int     `json:"bitrate"`
    Duration float64 `json:"duration"`
    FilePath string  `json:"file_path"`
    Height   int     `json:"height"`
    Md5      string  `json:"md5"`
}

// StartWorkFlowRequest 发起媒资转码请求
type StartWorkFlowRequest struct {
    Vid string `json:"vid"`
}

// StartWorkFlowResponse 发起媒资转码响应
type StartWorkFlowResponse struct {
    ErrMsg string `json:"err_msg"`
    ErrNo  int    `json:"err_no"`
    LogID  string `json:"log_id"`
    RunID  string `json:"run_id"`
}

// QueryWorkFlowRequest 查询媒资转码请求
type QueryWorkFlowRequest struct {
    RunID string `json:"run_id"`
}

// QueryWorkFlowResponse 查询媒资转码响应
type QueryWorkFlowResponse struct {
    ErrMsg       string `json:"err_msg"`
    ErrNo        int    `json:"err_no"`
    LogID        string `json:"log_id"`
    RunStatus    int    `json:"run_status" description:"转码任务运行状态（1:未转码  2:转码中 9:转码失败  10:转码成功）"`
    RunStatusMsg string `json:"run_status_msg" description:"转码任务运行状态描述"`
}
