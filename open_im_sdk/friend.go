package open_im_sdk

import "open_im_sdk/open_im_sdk_callback"

func GetDesignatedFriendsInfo(callback open_im_sdk_callback.Base, operationID string, userIDList string) {
	call(callback, operationID, UserForSDK.Friend().GetDesignatedFriendsInfo, userIDList)
}

func GetFriendList(callback open_im_sdk_callback.Base, operationID string) {
	call(callback, operationID, UserForSDK.Friend().GetFriendList)
}

func SearchFriends(callback open_im_sdk_callback.Base, operationID string, searchParam string) {
	call(callback, operationID, UserForSDK.Friend().SearchFriends, searchParam)
}

func CheckFriend(callback open_im_sdk_callback.Base, operationID string, userIDList string) {
	call(callback, operationID, UserForSDK.Friend().CheckFriend, userIDList)
}

func AddFriend(callback open_im_sdk_callback.Base, operationID string, userIDReqMsg string) {
	call(callback, operationID, UserForSDK.Friend().AddFriend, userIDReqMsg)
}

func SetFriendRemark(callback open_im_sdk_callback.Base, operationID string, userIDRemark string) {
	call(callback, operationID, UserForSDK.Friend().SetFriendRemark, userIDRemark)
}

func DeleteFriend(callback open_im_sdk_callback.Base, operationID string, friendUserID string) {
	call(callback, operationID, UserForSDK.Friend().DeleteFriend, friendUserID)
}

func GetRecvFriendApplicationList(callback open_im_sdk_callback.Base, operationID string) {
	call(callback, operationID, UserForSDK.Friend().GetRecvFriendApplicationList)
}

func GetSendFriendApplicationList(callback open_im_sdk_callback.Base, operationID string) {
	call(callback, operationID, UserForSDK.Friend().GetSendFriendApplicationList)
}

func AcceptFriendApplication(callback open_im_sdk_callback.Base, operationID string, userIDHandleMsg string) {
	call(callback, operationID, UserForSDK.Friend().AcceptFriendApplication, userIDHandleMsg)
}

func RefuseFriendApplication(callback open_im_sdk_callback.Base, operationID string, userIDHandleMsg string) {
	call(callback, operationID, UserForSDK.Friend().RefuseFriendApplication, userIDHandleMsg)
}

func AddBlack(callback open_im_sdk_callback.Base, operationID string, blackUserID string) {
	call(callback, operationID, UserForSDK.Friend().AddBlack, blackUserID)
}

func GetBlackList(callback open_im_sdk_callback.Base, operationID string) {
	call(callback, operationID, UserForSDK.Friend().GetBlackList)
}

func RemoveBlack(callback open_im_sdk_callback.Base, operationID string, removeUserID string) {
	call(callback, operationID, UserForSDK.Friend().RemoveBlack, removeUserID)
}