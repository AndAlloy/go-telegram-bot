package main

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
	"log"
	"time"
)

const (
	WeatherCommand = "/weather"
)

var (
	BusyState     = CreateState("BusyState")
	DayState      = CreateState("WeatherDayState")
	MethodState   = CreateState("WeatherMethodState")
	TownState     = CreateState("WeatherTownState")
	LocationState = CreateState("WeatherLocationState")
	RmMarkup      = &tele.ReplyMarkup{RemoveKeyboard: true}
)

func main() {
	pref := tele.Settings{
		Token:     BotToken,
		Poller:    &tele.LongPoller{Timeout: 310 * time.Second},
		ParseMode: "Markdown",
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(WeatherCommand, func(ctx tele.Context) error {
		Enable(BusyState)
		menu := &tele.ReplyMarkup{ResizeKeyboard: true}
		menu.Reply(menu.Row(menu.Text("Сьогодні"), menu.Text("Завтра")),
			menu.Row(menu.Text("На тиждень")))
		err := ctx.Send("Бажаєте переглянути погоду на:", menu)
		Enable(DayState)
		if err != nil {
			return err
		}
		return err
	})

	b.Handle(tele.OnText, func(ctx tele.Context) error {
		if IsOn(BusyState) {
			if IsOn(DayState) {
				switch ctx.Message().Text {
				case "Сьогодні", "Завтра", "На тиждень":
					{
						SetStateContext(DayState, ctx)
						Disable(DayState)
						menu := &tele.ReplyMarkup{ResizeKeyboard: true}
						menu.Reply(menu.Row(menu.Text("За координатами")),
							menu.Row(menu.Text("По місту")))
						err := ctx.Send("Оберіть спосіб:", menu)
						Enable(MethodState)
						if err != nil {
							return err
						}
					}
				default:
					err := ctx.Send("Будь-ласка, оберіть опції з клавіатури.", RmMarkup)
					if err != nil {
						return err
					}
					Disable(BusyState)
				}
			} else if IsOn(MethodState) {
				switch ctx.Message().Text {
				case "За координатами":
					{
						Enable(LocationState)
						Disable(MethodState)
						menu := &tele.ReplyMarkup{ResizeKeyboard: true}
						menu.Reply(menu.Row(menu.Location("Поділитися місцезнаходженням")))
						err := ctx.Send("Надайте доступ до своєї локації:", menu)
						if err != nil {
							return err
						}
					}
				case "По місту":
					{
						Enable(TownState)
						Disable(MethodState)
						err := ctx.Send("Напишіть місто:", RmMarkup)
						if err != nil {
							return err
						}
					}
				default:
					err := ctx.Send("Будь-ласка, оберіть опції з клавіатури.", RmMarkup)
					if err != nil {
						return err
					}
					Disable(BusyState)
				}

			} else if IsOn(TownState) {
				Disable(BusyState)
				SetStateContext(TownState, ctx)
				Disable(TownState)
				log.Println(fmt.Sprintf("%s", ctx.Message().Text))
				//TODO weather func
			}
		}
		return nil
	})

	b.Handle(tele.OnLocation, func(ctx tele.Context) error {
		if IsOn(LocationState) {
			Disable(MethodState)
			SetStateContext(LocationState, ctx)
			log.Println(fmt.Sprintf("%f", ctx.Message().Location.Lng))
			log.Println(fmt.Sprintf("%f", ctx.Message().Location.Lat))
			Disable(LocationState)
			//TODO weather func
		}
		return nil
	})

	b.Start()
}
