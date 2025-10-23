package main

import (
	"context"
	"fmt"
	managerpb "spelling-bee-game/server/api/spellingbee/v1"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GameState struct {
	id      int32
	score   int32
	letters string
}

func gameLoop(ctx context.Context, client managerpb.ManagerClient, game GameState) {
	for {
		fmt.Println(game.letters)
		fmt.Print("Enter Word > ")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			panic(err)
		}

		// temporary, quit if input is "q"
		if input == "q" {
			break
		}

		s, err := client.Submit(ctx, &managerpb.SubmitRequest{Id: game.id, Guess: input})
		if err != nil {
			panic(err)
		}

		game.score += s.Score

		fmt.Println(s.Message, "Current score is", game.score)
	}
}

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := managerpb.NewManagerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	// Create Game Session
	cr, err := client.CreateGame(ctx, &managerpb.CreateGameRequest{})
	if err != nil {
		panic(err)
	}
	game := GameState{cr.GetId(), 0, cr.GetLetters()}

	fmt.Println("Spelling Bee!")

	// Main Game Loop
	gameLoop(ctx, client, game)

	// End Game Session
	_, err = client.EndGame(ctx, &managerpb.EndGameRequest{Id: game.id})
	if err != nil {
		panic(err)
	}
	println("Game Over! Score: ", game.score)
}
