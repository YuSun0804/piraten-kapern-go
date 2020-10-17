package main

import (
	"github.com/pknetwork"
	"net"
	"pkserver/model"
	"strconv"
)

type server struct {
	piratenKapern model.PiratenKapern
	playRecord    *model.PlayRecord
	scoreRecord   *model.ScoreRecord
	conn          net.Conn
}

// 处理函数
func process(conn net.Conn, playerIndex int) {
	defer conn.Close()

	msg, err := pknetwork.ReadMessage(conn)
	if err != nil {
		return
	}

	piratenKapern := model.NewPiratenKapern()
	player := model.Player{PlayerIndex: playerIndex, PlayerName: msg}
	playRecord := model.PlayRecord{Player: player, TurnNum: 0}
	scoreRecord := model.ScoreRecord{Player: player, Score: 0}
	server := server{piratenKapern: *piratenKapern, playRecord: &playRecord, conn: conn, scoreRecord: &scoreRecord}

	pknetwork.WriteMessage(server.conn, "Hi, "+msg+", enjoy your game!.")
	for {
		msg, err = pknetwork.ReadMessage(conn)
		if err != nil {
			break
		}
		server.processMessage(msg)
	}

}

func (server server) processMessage(msg string) {
	if "f" == msg {
		card := server.piratenKapern.DrawFortuneCard(server.playRecord)
		pknetwork.WriteMessage(server.conn, "The fortuneCard you draw is ："+card.FortuneCardType.String())
	} else if "r" == msg {
		rollResult := server.piratenKapern.Roll(server.playRecord)
		pknetwork.WriteMessage(server.conn, "The roll result is ："+rollResult.String())
	} else if "s" == msg {
		score := server.piratenKapern.ComputeScore(server.playRecord)
		updatedScore := server.piratenKapern.UpdateScore(server.scoreRecord, score)
		pknetwork.WriteMessage(server.conn, "The score in total is ："+strconv.Itoa(updatedScore)+", and the score of current turn is "+strconv.Itoa(score))
	}
}

func main() {
	pknetwork.StartServer(process)
}
