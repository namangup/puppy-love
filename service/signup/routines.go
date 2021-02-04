package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	gomail "gopkg.in/mail.v2"
)

func MarkNotDirtyAlt(u User) mgo.Change {
	return mgo.Change{
		Update: bson.M{"$set": bson.M{
			"dirty": false,
		}},
		ReturnNew: true,
	}
}

func QueueService(listen_channel chan string, signup_channel chan string) {
	for request := range listen_channel {
		go func(request string) {
			signup_channel <- request[1:]
		}(request)
	}
}

func SignupService(
	Db PuppyDb,
	signup_channel chan string,
	mail_channel chan User) {

	for id := range signup_channel {
		log.Println("Signing up: " + id)

		u := User{}

		// If no such user
		if err := Db.GetById("user", id).One(&u); err != nil {
			log.Print(err)
			continue
		}

		// If user has already been computed
		if u.Dirty == false {
			log.Print("User ", id, " is not dirty. Skipping.")

			// Mailing should be async
			go func(user User) {
				log.Println("Sending mail now")
				mail_channel <- user
			}(u)

			continue
		}

		// Mark user as not dirty
		if _, err := Db.GetById("user", id).
			Apply(MarkNotDirtyAlt(u), &u); err != nil {

			log.Println("ERROR: Could not mark ", id, " as not dirty")
			log.Println(err)
		}

		// Mailing should be async
		go func(user User) {
			mail_channel <- user
			log.Println("Sending mail now")
		}(u)
	}
}

func MailerService(Db PuppyDb, mail_channel chan User) {
	mailCounter := 0

	for u := range mail_channel {
		log.Println("Setting up smtp")
		msg := "Use this token while signing up, and don't share it with anyone:\r\n" +
			"    Token: " + u.AuthC + " \r\n\r\n" +
			"Please ensure that you don't forget your password, as it will not be possible to recover your account if the password is lost.\r\n\r\n" +
			"We sincerely wish that you find your puppy love!\r\n\r\n" +
			"Regards,\r\nProgramming Club\r\n"

		m := gomail.NewMessage()
		m.SetHeader("From", EmailUser)
		m.SetHeader("To", u.Email+"@iitk.ac.in")
		m.SetHeader("Subject", "Puppy Love authentication code")
		m.SetBody("text/plain", msg)

		d := gomail.Dialer{Host: EmailHost, Port: EmailPortInt}
		if err := d.DialAndSend(m); err != nil {
			log.Println("ERROR: while mailing user ", u.Email, " ", u.Id)
			log.Println(err)
		} else {
			mailCounter += 1
			log.Println("Mailed " + u.Id)
			log.Println("Mails sent since inception", mailCounter)
		}
	}
}
