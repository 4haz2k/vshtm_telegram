package resources

var data map[string]string

func init() {

	data = map[string]string{
		"subscribeLabel":      "Подписаться",
		"unsubscribeLabel":    "Отписаться",
		"scheduleLabel":       "Расписание на неделю",
		"startMessage":        "Привет.\n\n<b>У этого бота 3 функции:</b>\n1. Подписаться на уведомления о начале занятия. Такое уведомление придёт за 2 часа до начала занятий.\n2. Отписаться от уведомлений. Уведомления приходить не будут (это мой любимый вариант).\n3. Получить расписание на неделю.\n\nВ общем то и всё. <b>Выбирай одну из трёх кнопок.</b>",
		"subscribedMessage":   "<b>Вы подписались на уведомления.</b> Уведомление о начале занятий будет приходить за 2 часа.",
		"unsubscribedMessage": "<b>Вы отписались от уведомлений.</b>",
	}
}

func Lang(key string) string {
	if name, ok := data[key]; ok {
		return name
	}

	return key // todo: мб возвращать здесь что-то другое вместо ключа
}
