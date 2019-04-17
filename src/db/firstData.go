package db

const fillStatus = `INSERT INTO status(name) VALUES ("Навчається"), 
													("Відрахованний"),
													("Працює"), 
													("Академічна відпустка"), 
													("Відпустка"), 
													("Дикретна відпустка");`

const fillRanks = `INSERT INTO ranks(name) VALUES ("Асистент"), 
													("Викладач"), 
													("Старший викладач"), 
													("Завідувач кафедри"), 
													("Доцент"), 
													("Професор"), 
													("Декан"), 
													("Директор"), 
													("Проректор"), 
													("Ректор");`

var FillDataQueries = [...]string{
	fillStatus,
	fillRanks,
}
