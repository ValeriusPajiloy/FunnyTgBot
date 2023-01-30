package telegram

// TODO: Add message here
const (

	//default
	msgHelp     = `Привет, я бот Абоба, пока что я ничего не умею, но в будущем собираюсь научиться колдунству, считать ваш социальный рейтинг и многое другое`
	msgRealHelp = `ПОМОГИТЕ, ЭТОТ ЕБЛАН ВАЛЕРА ДЕРЖИТ МЕНЯ В ПОДВАЛЕ И ЗАСТАВЛЯЕТ ПИСАТЬ ЭТУ ХУЙНЮ!`
	msgUnknown  = `я не знаю что делать...`
)

////game
// const (
//
//	msgGoDota    = `го в доту`
//	msgGoMain    = `го в майнкрафт`
//	msgGoValheim = `го в вальхейм`
//
// )

// notification
const (
	msgEmptySpell   = `Что ты колдуешь, дебич?`
	msgNotFullSpell = `Ты что очко свое заколдовываешь? Укажи через сколько и на кого ты колдуешь, баран, блять!?`
	msgSpell        = `Начинаем отсчет! Через %d min у %s %s`

	//main msg
	msgSpellNowPrefix            = `Сейчас у %s %s`
	msgSpellThirtyMinutsPrefix   = `Через 30 минут у %s будет %s`
	msgSpellFiveteenMinutsPrefix = `Через 15 минут у %s будет %s`
	msgSpellSomeMinutsPrefix     = `Через %s у %s будет %s`
	//minuts
	five  = "5 минут"
	four  = "4 минуты"
	three = "3 минуты"
	two   = "2 минуты"
	one   = "1 минуту"

	//help
	msgHelpSpell = `Чтобы заняться колдунством, введи:
	"/колдую *заклинание* *через сколько минут подействует* *цель* "
	Например:
	"/колдую понос 45 @Valera "`
)
