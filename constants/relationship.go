package constants

type RelationshipType uint

const (
	RelationshipTypeFriend RelationshipType = iota + 1 // 好友
	RelationshipTypeGroup                              // 群
)
