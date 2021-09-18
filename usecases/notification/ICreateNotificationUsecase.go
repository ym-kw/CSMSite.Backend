package Usecase

import "CSMSite.Backend/Domain/Dtos"

type ICreateNotificationUsecase interface {
	createNotification(Dtos.NotificationRequest) Dtos.NotificationResponse
}
