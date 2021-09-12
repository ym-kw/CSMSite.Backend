package usecase

import "CSMSite.Backend/domain/dtos"

type NotificationUsecase interface {
	createNotification(dtos.NotificationRequest) dtos.NotificationResponse
	getNotificationList() []dtos.NotificationResponse
	getNotification(int) dtos.NotificationResponse
}
