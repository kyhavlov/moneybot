package main

import (
	"log"
	"net/url"
	"os/exec"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	sc2api "github.com/kyhavlov/moneybot/SC2APIProtocol"
)

// This is a basic, no-op bot that starts a game against a built-in AI and
// runs until the game finishes, then sends a leave.
func main() {
	// Launch sc2
	game := exec.Command(`..\Versions\Base55958\SC2_x64.exe`,
		"-listen", "127.0.0.1",
		"-port", "8888",
		"-displayMode", "0",
	)
	game.Dir = `E:\StarCraft II\Support64`
	go func() {
		err := game.Run()
		if err != nil {
			log.Fatalf("error launching game: %v", err)
		}
	}()
	defer game.Process.Kill()
	time.Sleep(5 * time.Second)

	// Connect to sc2
	u := url.URL{Scheme: "ws", Host: "localhost:8888", Path: "/sc2api"}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	// Make a CreateGame request
	r := &sc2api.Request{Request: &sc2api.Request_CreateGame{CreateGame: &sc2api.RequestCreateGame{}}}
	req := r.GetCreateGame()
	req.Map = &sc2api.RequestCreateGame_BattlenetMapName{BattlenetMapName: "Acolyte LE"}
	req.DisableFog = proto.Bool(true)
	req.Realtime = proto.Bool(true)

	us := &sc2api.PlayerSetup{
		Type:       sc2api.PlayerType_Participant.Enum(),
		Race:       sc2api.Race_Terran.Enum(),
		Difficulty: sc2api.Difficulty_CheatInsane.Enum(),
	}
	opponent := &sc2api.PlayerSetup{
		Type: sc2api.PlayerType_Computer.Enum(),
		Race: sc2api.Race_Zerg.Enum(),
	}

	req.PlayerSetup = []*sc2api.PlayerSetup{us, opponent}

	// Send the create game request
	sendMessage(r, conn)

	// Decode the response
	decodeResponse(conn)

	// Make a JoinGame request
	r = &sc2api.Request{Request: &sc2api.Request_JoinGame{JoinGame: &sc2api.RequestJoinGame{}}}
	joinReq := r.GetJoinGame()
	joinReq.Participation = &sc2api.RequestJoinGame_Race{Race: sc2api.Race_Terran}
	joinReq.Options = &sc2api.InterfaceOptions{Raw: proto.Bool(true)}

	// Send the join game request
	sendMessage(r, conn)
	decodeResponse(conn)

	// Loop until game end
	for {
		// Send an observation request to get the game state
		obsReq := &sc2api.Request{Request: &sc2api.Request_Observation{Observation: &sc2api.RequestObservation{}}}
		sendMessage(obsReq, conn)

		// Read the response
		resp := decodeResponse(conn)

		// If the game ended, break out
		obs := resp.GetObservation()
		if len(obs.PlayerResult) > 0 {
			break
		}

		// If a command center is idle, build an SCV
		actionReq := &sc2api.Request{Request: &sc2api.Request_Action{Action: &sc2api.RequestAction{}}}
		action := actionReq.GetAction()
		for _, unit := range obs.Observation.RawData.Units {
			if *unit.UnitType == 18 {
				if len(unit.Orders) == 0 {
					command := &sc2api.ActionRaw_UnitCommand{
						UnitCommand: &sc2api.ActionRawUnitCommand{
							AbilityId: proto.Int32(524),
							UnitTags:  []uint64{*unit.Tag},
						},
					}
					a := &sc2api.Action{ActionRaw: &sc2api.ActionRaw{Action: command}}
					action.Actions = append(action.Actions, a)
					log.Println("queueing up SCV")
				}
			}
		}
		sendMessage(actionReq, conn)
		resp = decodeResponse(conn)

		//log.Printf("action response: %v, %v", resp, resp.Response)

		/*
		// Advance the game by a step
		stepReq := &sc2api.Request{Request: &sc2api.Request_Step{Step: &sc2api.RequestStep{}}}
		sendMessage(stepReq, conn)

		// Read the response
		decodeResponse(conn)*/
		time.Sleep(100 * time.Millisecond)
	}

	// Leave
	leaveReq := &sc2api.Request{Request: &sc2api.Request_LeaveGame{LeaveGame: &sc2api.RequestLeaveGame{}}}
	sendMessage(leaveReq, conn)
	decodeResponse(conn)
}

// Send a message over the connection
func sendMessage(message proto.Message, conn *websocket.Conn) {
	raw, err := proto.Marshal(message)
	if err != nil {
		log.Fatalf("error serializing protobuf struct: %v", err)
	}
	if err := conn.WriteMessage(websocket.TextMessage, raw); err != nil {
		log.Fatalf("error sending request: %v", err)
	}
}

// Read a response object from the connection
func decodeResponse(conn *websocket.Conn) sc2api.Response {
	response := sc2api.Response{}
	_, respBytes, err := conn.ReadMessage()

	if err != nil {
		log.Fatal("got error receiving response:", err)
	}
	proto.Unmarshal(respBytes, &response)

	//log.Printf("response: %v, %v", response, response.Response)

	return response
}
