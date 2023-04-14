package open_im_sdk


func GetWorkMomentsUnReadCount(callback open_im_sdk_callback.Base, operationID string) {
	call(callback, operationID, userForSDK.WorkMoments().GetWorkMomentsUnReadCount)
}


func GetWorkMomentsNotification(callback open_im_sdk_callback.Base, operationID string, offset int, count int) {
	call(callback, operationID, userForSDK.WorkMoments().GetWorkMomentsNotification, offset,count)
}


func ClearWorkMomentsNotification(callback open_im_sdk_callback.Base, operationID string) {
	call(callback, operationID, userForSDK.WorkMoments().ClearWorkMomentsNotification)
}
