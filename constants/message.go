package constants

// MessageType 消息类型
type MessageType uint

const (
	MessageTypeGroup     MessageType = iota + 1 // 群聊
	MessageTypePrivate                          // 私聊
	MessageTypeBroadcast                        // 广播
)

// MediaType 媒体类型
type MediaType uint

const (
	MediaTypeText  MediaType = iota + 1 // 文字
	MediaTypeImage                      // 图片
	MediaTypeVideo                      // 视频
	MediaTypeAudio                      // 音频
	MediaTypeFile                       // 文件
)
