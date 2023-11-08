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

package authorize

import (
    "strings"
)

const (
    // 用户相关
    
    // ScopeUserInfo 获取用户公开信息
    ScopeUserInfo = "user_info"
    // ScopeFansList 粉丝列表
    ScopeFansList = "fans.list"
    // ScopeFollowingList 关注列表
    ScopeFollowingList = "following.list"
    // ScopeFansData 查询创作者粉丝数据
    ScopeFansData = "fans.data"
    
    // 视频相关
    
    // ScopeVideoCreate 上传视频到文件服务器 - 创建抖音视频 -上传图片到文件服务器 - 发布图片
    ScopeVideoCreate = "video.create"
    // ScopeVideoList 列出已发布的视频
    ScopeVideoList = "video.list"
    // ScopeVideoData 查询指定视频数据
    ScopeVideoData = "video.data"
    // ScopeAwemeShare 抖音分享 id 机制
    ScopeAwemeShare = "aweme.share"
    // ScopeVideoDelete 删除抖音视频
    ScopeVideoDelete = "video.delete"
    // ScopeHotSearch 获取实时热点词 --获取热点词聚合的视频
    ScopeHotSearch = "hotsearch"
    
    // 互动
    
    // ScopeVideoComment 评论列表 ---评论回复列表 ---回复视频评论 ---置顶视频评论 (企业号)
    ScopeVideoComment = "video.comment"
    // ScopeIm 给抖音用户发送消息  --- 上传素材
    ScopeIm = "im"
)

// GetUserScope 获取用户相关 Scope.
func GetUserScope() string {
    scopes := []string{ScopeUserInfo, ScopeFansList, ScopeFollowingList, ScopeFansData}
    return strings.Join(scopes, ",")
}

// GetVideoScope 获取视频相关 Scope.
func GetVideoScope() string {
    scopes := []string{ScopeVideoCreate, ScopeVideoList, ScopeVideoData, ScopeAwemeShare, ScopeVideoDelete, ScopeHotSearch}
    return strings.Join(scopes, ",")
}

// GetInteractScope 获取互动相关 Scope.
func GetInteractScope() string {
    scopes := []string{ScopeVideoComment, ScopeIm}
    return strings.Join(scopes, ",")
}

// GetAllScope 获取所有 Scope.
func GetAllScope() string {
    scopes := []string{GetInteractScope(), GetVideoScope(), GetUserScope()}
    return strings.Join(scopes, ",")
}
