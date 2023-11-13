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

const (
    // GetVideoList 获取视频列表
    getVideoList = "https://open.douyin.com/api/dyc_voc/get_video_list?access_token="
    
    // deleteVideo 删除视频
    deleteVideo = "https://open.douyin.com/api/dyc_voc/delete_video?access_token="
    
    // getVideoURLByVID 获取视频地址
    getVideoByVID = "https://open.douyin.com/api/dyc_voc/get_video_by_vid?access_token="
    
    // uploadVideoByURLs 上传视频
    uploadVideoByURLs = "https://open.douyin.com/api/dyc_voc/upload_video_by_urls?access_token="
    
    // getVideoUploadJobInfo 查询 URL 批量上传任务状态
    getVideoUploadJobInfo = "https://open.douyin.com/api/dyc_voc/get_upload_job_info?access_token="
    
    // startWorkFlow 开始工作流
    startWorkFlow = "https://open.douyin.com/api/dyc_voc/start_work_flow?access_token="
    
    // getWorkFlowStatus 获取工作流状态 通过 run_id 查询工作流状态
    getWorkFlowStatus = "https://open.douyin.com/api/dyc_voc/get_work_flow_exection?access_token="
)
