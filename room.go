package main

import ()

type Room struct {
	users     []*Client
	nextIndex int
}

func NewRoom() *Room {
	return &Room{nextIndex: 0}
}

func (r *Room) NextUser() *Client {
	println("Next user")
	for {
		if len(r.users) > 0 {
			next := r.nextIndex
			r.nextIndex = (r.nextIndex + 1) % (len(r.users))
			println("next index is ")
			println(r.nextIndex)
			return r.users[next]
		}
	}

}

func (r *Room) AddUser(c *Client) {
	r.users = append(r.users, c)
	r.nextIndex = r.FindUser(c)
}

func (r *Room) DeleteUser(c *Client) {
	i := r.FindUser(c)
	r.users = append(r.users[:i], r.users[i+1:]...)
}

func (r *Room) FindUser(c *Client) int {

	for p, v := range r.users {
		if v == c {
			return p
		}
	}
	return -1
}

func (r *Room) run() {
	sentence := NewSentence()
	println("Exquisite corpse...")
	for {

		user := r.NextUser()
		user.SendMessage("Waiting for input:\n")
		text, err := user.ReadString('\n')
		if err != nil {
			println("we received an error.")
			println(err)
			r.DeleteUser(user)
		} else {
			if sentence.HandleText(text) {
				println("Handled text")
			} else {
				println("Couldnlt handle")
				return
			}
		}
	}
}
