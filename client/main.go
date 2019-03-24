package main

import (
	"context"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	music "github.com/tozastation/go-grpc-streaming-example/idl/output"
	"google.golang.org/grpc"
	"io"
	"os"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:11111")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := music.NewMusicClient(conn)

	stream, err := client.Start(context.Background(), &music.Empty{})
	if err != nil {
		panic(err)
	}
	for {
		musicData, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		d, err := mp3.Decoder(musicData.GetData())
		if err != nil {
			return err
		}
		defer d.Close()
		p, err := oto.NewPlayer(d.SampleRate(), 2, 2, 65536)
		if err != nil {
			return err
		}
		defer p.Close()
		if _, err := io.Copy(p, d); err != nil {
			return err
		}
		//fmt.Println(person)
	}
}
