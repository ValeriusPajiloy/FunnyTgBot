package notification

// notification
const (
	//first answer
	MSG_EMPTY_SPELL    = `Что ты колдуешь, дебич? Прочитай "/колдую help"`
	MSG_NOT_FULL_SPELL = `Ты что очко свое заколдовываешь? Укажи через сколько и на кого ты колдуешь, баран, блять!
	(Правильный пример: "/колдую понос 45 @Valera")`
	MSG_BIG_TIME = "Уух... Какой большой... Масимальный таймер - 1440"
	MSG_SPELL    = `Начинаем отсчет! Через %d min у %s %s`

	//main MSG_
	MSG_SPELL_NOW_PREFIX             = `Сейчас у %s %s`
	MSG_SPELL_THIRTY_MINUTS_PREFIX   = `Через 30 минут у %s будет %s`
	MSG_SPELL_FIVETEEN_MINUTS_PREFIX = `Через 15 минут у %s будет %s`
	MSG_SPELL_SOME_MINUTS_PREFIX     = `Через %s у %s будет %s`
	//minuts
	FIVE  = "5 минут"
	FOUR  = "4 минуты"
	THREE = "3 минуты"
	TWO   = "2 минуты"
	ONE   = "1 минуту"

	//help
	MSG_HELP_SPELL = `Чтобы заняться колдунством, введи:
	"/колдую *заклинание* *через сколько минут подействует* *цель* "
	Например:
	"/колдую понос 45 @Valera "`
)
