package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

// TODO: a lot needs to be fixed here
// better handling when the quotes db is empty/0 for the ID
// lastID needs to not be init with 1, this is a temp fix because of how i spaghetti'd this together
// need a way to list IDs and quotes to delete
// select a quote by ID

var lastID int

func init() {
	NewCommand("quote", false, getQuote).SetHelp("get a random quote").Add()
	NewCommand("addquote", false, addQuote).SetHelp("add quote").ReqAdmin().Add()
	NewCommand("remquote", false, remQuote).SetHelp("remove quote by id").ReqAdmin().Add()
	initDB()
	lastID = 1
}

func getQuote(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	database, _ := sql.Open("sqlite3", "./quotes.db")
	var quote string

	rows, _ := database.Query("SELECT MAX(rowid) FROM quotes") // LAST_INSERT_ROWID() would be better but this is fine
	for rows.Next() {                                          // is this how to do this?
		rows.Scan(&lastID)
	}

	randomID := GenRandomNum(lastID)
	if randomID == 0 { // feels gross
		randomID++
	}

	rows, _ = database.Query("SELECT quote FROM quotes WHERE rowid=?", randomID)
	for rows.Next() {
		rows.Scan(&quote)
		s.ChannelMessageSend(m.ChannelID, quote)
	}
}

func addQuote(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	database, _ := sql.Open("sqlite3", "./quotes.db")
	quoteToAdd := strings.Join(msgList[1:], " ")

	statement, _ := database.Prepare("INSERT INTO quotes (quote) VALUES (?)")
	statement.Exec(quoteToAdd)
}

func remQuote(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	if len(msgList) < 2 {
		s.ChannelMessageSend(m.ChannelID, "usage: >remquote ID")
		return
	}
	database, _ := sql.Open("sqlite3", "./quotes.db")
	ID, _ := StrArrayToInt(msgList[1:]) // this and the following line *need* to be changed, this isn't a good way to do this
	if ID[0] > lastID || ID[0] == 0 {
		s.ChannelMessageSend(m.ChannelID, "ID not found, try again.")
	}

	statement, _ := database.Prepare("DELETE FROM quotes WHERE rowid=?")
	statement.Exec(ID[0])
}

func listQuote(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	// placeholder
}

// checking if the quote db exists and if not, create it
func initDB() {
	filename := "quotes.db"
	if _, err := os.Stat(filename); err == nil {
		fmt.Println("db already exists")
	} else if os.IsNotExist(err) {
		fmt.Println("db does not exist, creating")
		database, _ := sql.Open("sqlite3", "./quotes.db")
		statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS quotes (quote TEXT)")
		statement.Exec()
	} else {
		// typically it shouldnt reach this code but in the event that os.Stat() does something weird this should 'catch' it
		// should probably make this send a message in the discord that a weird error occurred to help debugging but im lazy /shrug
		fmt.Println("some wild kind of error has occurred")
		fmt.Println(err)
	}
}