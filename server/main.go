package main

import (
	music "github.com/tozastation/go-grpc-streaming-example/idl/output"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
)

// MusicService is ...
type MusicService struct{}

// Start is ...
func (s *MusicService) Start(req *music.Empty, stream music.Music_StartServer) error {
	// ファイル用意
	file, err := os.Open("./src/canon_in_d.mp3")
	if err != nil {
		return err
	}
	defer file.Close()
	// 読み込み & ストリーミング
	buf := make([]byte, 1024)
	for {
		_, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		stream.Send(&music.StreamingMusic{Data: buf})
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":11111")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	music.RegisterMusicServer(server, &MusicService{})
	server.Serve(lis)
}
