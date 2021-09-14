package usecase

import "CSMSite.Backend/domain/dtos"

type ICreateNotificationUsecase interface {
	createNotification(dtos.NotificationRequest) dtos.NotificationResponse
}
