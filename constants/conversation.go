package constants

type ConversationType uint

const (
	ConversationTypeGroup   ConversationType = iota + 1 // 群聊
	ConversationTypePrivate                             // 私聊
)
