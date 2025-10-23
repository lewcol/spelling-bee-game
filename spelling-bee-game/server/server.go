package main

import (
	"context"
	"log"
	"net"
	managerpb "spelling-bee-game/server/api/spellingbee/v1"
	"spelling-bee-game/server/manager"

	"google.golang.org/grpc"
)

type server struct {
	managerpb.UnimplementedManagerServer
	manager manager.Manager
}

func (s *server) CreateGame(ctx context.Context, req *managerpb.CreateGameRequest) (*managerpb.CreateGameResponse, error) {
	id, g := s.manager.Create()
	return &managerpb.CreateGameResponse{Id: int32(id), Letters: g.PrintableLettersWithCentre()}, nil
}

func (s *server) EndGame(ctx context.Context, req *managerpb.EndGameRequest) (*managerpb.EndGameResponse, error) {
	g, ok := s.manager.GetGame(int(req.GetId()))
	if ok {
		err := s.manager.End(int(req.GetId()))
		if err != nil {
			return nil, err
		}
		return &managerpb.EndGameResponse{Score: int32(g.Score())}, nil
	}
	return &managerpb.EndGameResponse{Score: int32(-1)}, nil
}

func (s *server) Score(ctx context.Context, req *managerpb.ScoreRequest) (*managerpb.ScoreResponse, error) {
	g, ok := s.manager.GetGame(int(req.GetId()))
	if ok {
		return &managerpb.ScoreResponse{Score: int32(g.Score())}, nil
	}
	return &managerpb.ScoreResponse{Score: int32(-1)}, nil
}

func (s *server) Submit(ctx context.Context, req *managerpb.SubmitRequest) (*managerpb.SubmitResponse, error) {
	g, ok := s.manager.GetGame(int(req.GetId()))
	if ok {
		message, score := g.Submit(req.GetGuess())
		return &managerpb.SubmitResponse{Message: message, Score: int32(score)}, nil
	}
	return &managerpb.SubmitResponse{Message: "", Score: int32(-1)}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := &server{manager: manager.GetManager()}
	grpcServer := grpc.NewServer()
	managerpb.RegisterManagerServer(grpcServer, s)

	log.Println("Server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
		//log.Fatal(err)
	}
}
