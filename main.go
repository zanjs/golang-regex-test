package main

import (
	"fmt"

	cregex "github.com/mingrammer/commonregex"
)

func main() {

	text := `John, please get that 
	article on www.linkedin.com to me by 
	5:00PM on Jan 9th 2012. 4:00 would be ideal, actually. 
	If you have any questions, 
	You can reach me at (519)-236-2723x341 or get in touch with my 
	associate at harold.smith@gmail.com`

	date_list := cregex.Date(text)

	// ['Jan 9th 2012']
	time_list := cregex.Time(text)
	// ['5:00PM', '4:00']
	link_list := cregex.Links(text)
	// ['www.linkedin.com', 'harold.smith@gmail.com']
	phone_list := cregex.PhonesWithExts(text)
	// ['(519)-236-2723x341']
	email_list := cregex.Emails(text)
	// ['harold.smith@gmail.com']

	fmt.Println(date_list)
	fmt.Println(time_list)
	fmt.Println(link_list)
	fmt.Println(phone_list)
	fmt.Println(email_list)
}
