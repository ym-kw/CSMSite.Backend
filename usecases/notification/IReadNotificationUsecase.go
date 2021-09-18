package Usecase

import "CSMSite.Backend/Domain/Dtos"

type IReadNotificationUsecase interface {
	getNotificationList() []Dtos.NotificationResponse
	getNotification(int) Dtos.NotificationResponse
}
