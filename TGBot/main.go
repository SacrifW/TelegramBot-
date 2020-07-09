package main

import (
	"TGBot/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main()  {
	botToken := "1192406046:AAEfLYwSqVFHhJ0bsPC903ogdxHVMNvBIUs"
//https://api.telegram.org/bot<token>/METHOD_NAME
	botApi := "https://api.telegram.org/bot"
	botUrl := botApi + botToken
	offset := 0

	for ; ; { //бесконечный цикл без условий
		updates, err := getUpdates(botUrl, offset)
		if err != nil {
			log.Println("Smth went wrong:", err.Error())
		}
		for _, update := range updates{
			err = respond(botUrl, update)
			offset = update.UpdateId + 1
		}

		fmt.Println(updates)
	}
}


//запрашивает обновления
func getUpdates(botUrl string, offset int)([]models.Update, error)  {
resp, err := http.Get(botUrl+ "/getUpdates" + "?offset=" + strconv.Itoa(offset))
if err != nil{
	return nil, err
}
defer resp.Body.Close()

body, err := ioutil.ReadAll(resp.Body)
if err != nil{
	return nil, err
}

var restResponse models.RestResponse
err = json.Unmarshal(body, &restResponse)
if err != nil{
	return nil, err
}
return restResponse.Result, nil
}

//отвечает на обновление
func respond(botUrl string, update models.Update) (err error)  {
var botMessage models.BotMessage
botMessage.ChatId = update.Message.Chat.ChatId
botMessage.Text = "Привет, давно не виделись. Как ты?)"//текст, который будет давать в ответ на любое  присланное сообщение наш бот

//наш бот будет отвечать на слово "грустно" поддержкой
//только слово "грустно" в таком же регистре
if update.Message.Text == "грустно"{
	botMessage.Text = "Мур-мур, не грусти :3"
	}

//бот должен уметь за себя постоять, поэтому вот
if update.Message.Text == "пидор"{
	botMessage.Text = "Сам ты пидор!"
}

buf, err := json.Marshal(botMessage)
if err != nil{
return err
}
_, err = http.Post(botUrl + "/sendMessage", "application/json", bytes.NewBuffer(buf))
if err != nil{
	return err
}
	return nil
}