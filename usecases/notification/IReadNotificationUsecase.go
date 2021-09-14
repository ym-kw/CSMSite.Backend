package usecase

import "CSMSite.Backend/domain/dtos"

type IReadNotificationUsecase interface {
	getNotificationList() []dtos.NotificationResponse
	getNotification(int) dtos.NotificationResponse
}
