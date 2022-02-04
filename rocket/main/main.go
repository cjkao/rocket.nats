package main

import (
	"fmt"

	"github.com/cjkao/training/rocket"
)

var rclient *rocket.Client

func Setup(url, acct, pwd string) error {
	if rclient == nil {
		rclient = rocket.NewClient(url)
		login := rocket.LoginPayload{
			User:     acct,
			Password: pwd,
		}
		lg, err := rclient.Login(&login)
		if err != nil {
			fmt.Printf("Error: %+v", err)
			return err
		}
		fmt.Printf("I'm %s\n", lg.Data.Me.Username)
	}
	return nil
}
func CallMeeting(userNt, msgbody string) error {
	resp, err := rclient.CreateIm(&rocket.CreateImRequest{Username: userNt})
	if err != nil {
		fmt.Printf("Error to create IM: %+v", err)
		return err
	}
	//should post meeting message & trigger client open popup window
	postResp, err := rclient.PostMessage(&rocket.Message{
		RoomID: resp.Room.Rid,
		Text:   msgbody,
	})
	if err != nil {
		fmt.Printf("Error to send special meeting message: %+v", err)
		return err
	} else {
		fmt.Printf("Success: %+v", postResp)
	}
	fmt.Printf("%+v\n", resp)
	return nil
}
func main() {
	Setup("http://localhost:3000", "a", "a")
	CallMeeting("c", "[Hello32](https://google.com)")
}
