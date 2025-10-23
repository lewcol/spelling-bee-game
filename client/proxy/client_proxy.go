package proxy

import (
	"context"
	"fmt"
	managerpb "spelling-bee-game/server/api/spellingbee/v1"
	"strings"
	"time"

	"google.golang.org/grpc"
)

type ClientProxy interface {
	MaxReqPerSecond() int32
	CreateGame(ctx context.Context, in *managerpb.CreateGameRequest, opts ...grpc.CallOption) (*managerpb.CreateGameResponse, error)
	EndGame(ctx context.Context, in *managerpb.EndGameRequest, opts ...grpc.CallOption) (*managerpb.EndGameResponse, error)
	Score(ctx context.Context, in *managerpb.ScoreRequest, opts ...grpc.CallOption) (*managerpb.ScoreResponse, error)
	Submit(ctx context.Context, in *managerpb.SubmitRequest, opts ...grpc.CallOption) (*managerpb.SubmitResponse, error)
}

type clientProxy struct {
	inner           managerpb.ManagerClient
	maxReqPerSecond int32
	minGap          time.Duration
	lastCall        time.Time
}

func (c *clientProxy) MaxReqPerSecond() int32 {
	return c.maxReqPerSecond
}

func (c *clientProxy) throttle() {
	now := time.Now()
	wait := c.minGap - now.Sub(c.lastCall)
	if wait > 0 {
		fmt.Println("Rate Limited: Waiting for ", wait.Milliseconds(), " milliseconds.")
		time.Sleep(wait)
	}
	c.lastCall = now
}

func (c *clientProxy) sanitiseInput(s string) string {
	var sanitisedInput []string
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			sanitisedInput = append(sanitisedInput, string(letter))
		} else if letter >= 'A' && letter <= 'Z' {
			sanitisedInput = append(sanitisedInput, string(letter+32))
		}
	}
	return strings.Join(sanitisedInput, "")
}

func (c *clientProxy) CreateGame(ctx context.Context, in *managerpb.CreateGameRequest, opts ...grpc.CallOption) (*managerpb.CreateGameResponse, error) {
	c.throttle()
	return c.inner.CreateGame(ctx, in, opts...)
}

func (c *clientProxy) EndGame(ctx context.Context, in *managerpb.EndGameRequest, opts ...grpc.CallOption) (*managerpb.EndGameResponse, error) {
	c.throttle()
	return c.inner.EndGame(ctx, in, opts...)
}

func (c *clientProxy) Score(ctx context.Context, in *managerpb.ScoreRequest, opts ...grpc.CallOption) (*managerpb.ScoreResponse, error) {
	c.throttle()
	return c.inner.Score(ctx, in, opts...)
}

func (c *clientProxy) Submit(ctx context.Context, in *managerpb.SubmitRequest, opts ...grpc.CallOption) (*managerpb.SubmitResponse, error) {
	c.throttle()
	in.Guess = c.sanitiseInput(in.Guess)
	return c.inner.Submit(ctx, in, opts...)
}

func NewClientProxy(conn *grpc.ClientConn, maxReqPerSecond int32) ClientProxy {
	inner := managerpb.NewManagerClient(conn)
	minGap := time.Second
	if maxReqPerSecond > 0 {
		minGap = time.Second / time.Duration(maxReqPerSecond)
	}
	return &clientProxy{inner: inner, maxReqPerSecond: maxReqPerSecond, minGap: minGap}
}
