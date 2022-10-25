package temp_struct

type LocalChatLog struct {
	ServerMsgID      string ` json:"serverMsgID,omitempty"`
	SendID           string ` json:"sendID,omitempty"`
	RecvID           string ` json:"recvID,omitempty"`
	SenderPlatformID int32  ` json:"senderPlatformID,omitempty"`
	SenderNickname   string ` json:"senderNickname,omitempty"`
	SenderFaceURL    string ` json:"senderFaceURL,omitempty"`
	SessionType      int32  ` json:"sessionType,omitempty"`
	MsgFrom          int32  ` json:"msgFrom,omitempty"`
	ContentType      int32  ` json:"contentType,omitempty"`
	Content          string ` json:"content,omitempty"`
	IsRead           bool   ` json:"isRead,omitempty"`
	Status           int32  ` json:"status,omitempty"`
	Seq              uint32 ` json:"seq,omitempty"`
	SendTime         int64  ` json:"sendTime,omitempty"`
	CreateTime       int64  ` json:"createTime,omitempty"`
	AttachedInfo     string ` json:"attachedInfo,omitempty"`
	Ex               string ` json:"ex,omitempty"`
}
type LocalConversation struct {
	ConversationID        string ` json:"conversationID,omitempty"`
	ConversationType      int32  ` json:"conversationType,omitempty"`
	UserID                string ` json:"userID,omitempty"`
	GroupID               string ` json:"groupID,omitempty"`
	ShowName              string ` json:"showName,omitempty"`
	FaceURL               string ` json:"faceURL,omitempty"`
	RecvMsgOpt            int32  ` json:"recvMsgOpt,omitempty"`
	UnreadCount           int32  ` json:"unreadCount,omitempty"`
	GroupAtType           int32  ` json:"groupAtType,omitempty"`
	LatestMsg             string ` json:"latestMsg,omitempty"`
	LatestMsgSendTime     int64  ` json:"latestMsgSendTime,omitempty"`
	DraftText             string ` json:"draftText,omitempty"`
	DraftTextTime         int64  ` json:"draftTextTime,omitempty"`
	IsPinned              bool   ` json:"isPinned,omitempty"`
	IsPrivateChat         bool   ` json:"isPrivateChat,omitempty"`
	IsNotInGroup          bool   ` json:"isNotInGroup,omitempty"`
	UpdateUnreadCountTime int64  ` json:"updateUnreadCountTime,omitempty"`
	AttachedInfo          string ` json:"attachedInfo,omitempty"`
	Ex                    string ` json:"ex,omitempty"`
}
type LocalSuperGroup struct {
	GroupID                string `json:"groupID,omitempty"`
	GroupName              string `json:"groupName,omitempty"`
	Notification           string `json:"notification,omitempty"`
	Introduction           string `json:"introduction,omitempty"`
	FaceURL                string `json:"faceURL,omitempty"`
	CreateTime             uint32 `json:"createTime,omitempty"`
	Status                 int32  `json:"status,omitempty"`
	CreatorUserID          string `json:"creatorUserID,omitempty"`
	GroupType              int32  `json:"groupType,omitempty"`
	OwnerUserID            string `json:"ownerUserID,omitempty"`
	MemberCount            int32  `json:"memberCount,omitempty"`
	Ex                     string `json:"ex,omitempty"`
	AttachedInfo           string `json:"attachedInfo,omitempty"`
	NeedVerification       int32  `json:"needVerification,omitempty"`
	LookMemberInfo         int32  `json:"lookMemberInfo,omitempty"`
	ApplyMemberFriend      int32  `json:"applyMemberFriend,omitempty"`
	NotificationUpdateTime uint32 `json:"notificationUpdateTime,omitempty"`
	NotificationUserID     string `json:"notificationUserID,omitempty"`
}