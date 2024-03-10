package resources

var data map[string]string

func init() {
	data = map[string]string{
		"subscribeLabel":         "Подписаться",
		"unsubscribeLabel":       "Отписаться",
		"scheduleLabel":          "Расписание на неделю",
		"startMessage":           "Привет.\n\n<b>У этого бота 3 функции:</b>\n1. Подписаться на уведомления о начале занятия. Такое уведомление придёт за 2 часа до начала занятий.\n2. Отписаться от уведомлений. Уведомления приходить не будут (это мой любимый вариант).\n3. Получить расписание на неделю.\n\nВ общем то и всё. <b>Выбирай одну из трёх кнопок.</b>",
		"subscribedMessage":      "<b>Вы подписались на уведомления.</b> Уведомление о начале занятий будет приходить за 2 часа.",
		"unsubscribedMessage":    "<b>Вы отписались от уведомлений.</b>",
		"onlineScheduleMessage":  "<b>Предмет:</b> %s\n<b>Тип занятия:</b> %s\n<b>Дата и время:</b> %s\n<b>Преподаватель:</b> %s\n<b>Ссылка на занятие:</b> %s\n\n",
		"lanScheduleMessage":     "<b>Предмет:</b> %s\n<b>Тип занятия:</b> %s\n<b>Дата и время:</b> %s\n<b>Преподаватель:</b> %s\n<b>Здание КАИ:</b> %s\n<b>Аудитория:</b> %s\n\n",
		"emptyScheduleMessage":   "Расписание на ближайшую неделю отсутствует",
		"ScheduleForWeekMessage": "Расписание на неделю:\n\n",
		"unknownMessage":         "<b>Не знаю, что тебе ответить, поэтому, вот тебе анекдот для людей за 40:</b>\n\n%s",
	}
}

func Lang(key string) string {
	if name, ok := data[key]; ok {
		return name
	}

	return key // todo: мб возвращать здесь что-то другое вместо ключа
}
