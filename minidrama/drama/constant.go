/*
 * Copyright Bytedance Author(https://houseme.github.io/bytedance/). All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the"License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an"AS IS"BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * You can obtain one at https://github.com/houseme/bytedance.
 *
 */

package drama

const (
    // resourceUpload 资源上传
    resourceUpload = "https://open.douyin.com/api/playlet/v2/resource/upload/?access_token="
    // queryVideo 查询视频
    queryVideo = "https://open.douyin.com/api/playlet/v2/video/query/?access_token="
    
    // createVideo 创建视频
    createVideo = "https://open.douyin.com/api/playlet/v2/video/create/?access_token="
    
    // editVideo 编辑视频
    editVideo = "https://open.douyin.com/api/playlet/v2/video/edit/?access_token="
    
    // queryVideoAlbum https://open.douyin.com/api/playlet/v2/album/fetch
    queryVideoAlbum = "https://open.douyin.com/api/playlet/v2/album/fetch/?access_token="
    
    // reviewVideo https://open.douyin.com/api/playlet/v2/video/review
    reviewVideo = "https://open.douyin.com/api/playlet/v2/video/review/?access_token="
    
    // authorizeVideo https://open.douyin.com/api/playlet/v2/auth/authorize
    authorizeVideo = "https://open.douyin.com/api/playlet/v2/auth/authorize/?access_token="
    
    // onlineAlbum https://open.douyin.com/api/playlet/v2/album/online/
    onlineAlbum = "https://open.douyin.com/api/playlet/v2/album/online/?access_token="
    
    // bindAlbum https://open.douyin.com/api/playlet/v2/album/bind
    bindAlbum = "https://open.douyin.com/api/playlet/v2/album/bind/?access_token="
    
    // playInfo https://open.douyin.com/api/playlet/v2/video/play_info
    playInfo = "https://open.douyin.com/api/playlet/v2/video/play_info/?access_token="
)

const (
    // ResourceTypeVideo 资源类型 1 视频，2 图片
    ResourceTypeVideo = 1
    // ResourceTypeImage 资源类型 1 视频，2 图片
    ResourceTypeImage = 2
    
    // ErrNoSuccess = 0
    ErrNoSuccess = 0
    
    // ErrTipsSuccess = "success"
    ErrTipsSuccess = "success"
)

var (
    // VideoTagList 视频标签
    VideoTagList = []Tag{
        {
            ID:    "1",
            Value: "医神",
        },
        {
            ID:    "2",
            Value: "赘婿",
        },
        {
            ID:    "3",
            Value: "鉴宝",
        },
        {
            ID:    "4",
            Value: "战神",
        },
        {
            ID:    "5",
            Value: "娱乐明星",
        },
        {
            ID:    "6",
            Value: "神医",
        },
        {
            ID:    "7",
            Value: "重生",
        },
        {
            ID:    "8",
            Value: "职场",
        },
        {
            ID:    "9",
            Value: "逆袭",
        },
        {
            ID:    "10",
            Value: "复仇",
        },
        {
            ID:    "11",
            Value: "青春",
        },
        {
            ID:    "12",
            Value: "官场",
        },
        {
            ID:    "13",
            Value: "家庭情感",
        },
        {
            ID:    "14",
            Value: "乡村",
        },
        {
            ID:    "15",
            Value: "正能量",
        },
        {
            ID:    "16",
            Value: "成长",
        },
        {
            ID:    "17",
            Value: "伦理",
        },
        {
            ID:    "18",
            Value: "都市情感",
        },
        {
            ID:    "19",
            Value: "社会话题",
        },
        {
            ID:    "20",
            Value: "灵异",
        },
        {
            ID:    "21",
            Value: "悬疑推理",
        },
        {
            ID:    "22",
            Value: "虐恋",
        },
        {
            ID:    "23",
            Value: "甜宠",
        },
        {
            ID:    "24",
            Value: "高干军婚",
        },
        {
            ID:    "25",
            Value: "年代",
        },
        {
            ID:    "26",
            Value: "萌宝",
        },
        {
            ID:    "27",
            Value: "腹黑",
        },
        {
            ID:    "28",
            Value: "总裁",
        },
        {
            ID:    "29",
            Value: "宫斗宅斗",
        },
        {
            ID:    "30",
            Value: "穿越",
        },
        {
            ID:    "31",
            Value: "种田经商",
        },
        {
            ID:    "33",
            Value: "民俗",
        },
        {
            ID:    "34",
            Value: "古装",
        },
        {
            ID:    "35",
            Value: "穿越战争",
        },
        {
            ID:    "36",
            Value: "现代军事",
        },
        {
            ID:    "37",
            Value: "奇幻",
        },
        {
            ID:    "38",
            Value: "科幻",
        },
        {
            ID:    "39",
            Value: "架空玄幻",
        },
        {
            ID:    "40",
            Value: "热血",
        },
        {
            ID:    "41",
            Value: "历史",
        },
        {
            ID:    "42",
            Value: "搞笑",
        },
        {
            ID:    "43",
            Value: "仙侠",
        },
        {
            ID:    "44",
            Value: "武侠",
        },
        {
            ID:    "45",
            Value: "二次元",
        },
        {
            ID:    "46",
            Value: "其他",
        },
    }
)
