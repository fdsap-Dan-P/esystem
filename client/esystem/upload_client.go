package client

import (
	"context"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"crypto/md5"
	"encoding/hex"
	pb "simplebank/pb"
	"simplebank/util"
	// "github.com/docker/docker/api/types/container"
)

type UploadFileRequest struct {
	FilePath    string
	RefCode     string
	Remarks     string
	TargetTable string
	ServerPath  string
	DockerImgID string
	DockerPath  string
}

var maxRetryCount int = 5
var retryDelay = time.Duration(time.Duration.Seconds(1))

// UploadFile calls upload file RPC
func (dumpClient *DumpClient) UploadFile(ctx context.Context, parm UploadFileRequest) (*pb.UploadFileResponse, error) {
	file, err := os.Open(parm.FilePath)
	if err != nil {
		log.Fatal("cannot open file: ", err)
	}
	defer file.Close()

	bufferSize := 1024
	checkSumStr := ""

	stream, err := dumpClient.document.UploadFile(ctx)
	if err != nil {
		return &pb.UploadFileResponse{}, err
	}

	req := &pb.UploadFileRequest{
		Data: &pb.UploadFileRequest_Info{
			Info: &pb.FileInfo{
				FileType:      filepath.Ext(parm.FilePath),
				FileName:      filepath.Base(parm.FilePath),
				ReferenceCode: parm.RefCode,
				Remarks:       parm.Remarks,
				TargetTable:   parm.TargetTable,
				ServerPath:    parm.ServerPath,
				DockerImgID:   parm.DockerImgID,
				DockerPath:    parm.DockerPath,
			},
		},
	}

	err = stream.Send(req)
	if err != nil {
		log.Fatal("cannot send image info to server: ", err, stream.RecvMsg(nil))
	}

	chunkChan := make(chan []byte, bufferSize)

	// Start a goroutine to read chunks from the file and send them to the channel
	go func() {
		defer close(chunkChan)

		for {
			buffer := make([]byte, bufferSize)
			n, err := file.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("error reading file: %v", err)
				return
			}

			chunk := make([]byte, n)
			copy(chunk, buffer[:n])

			select {
			case chunkChan <- chunk:
			case <-ctx.Done():
				return
			}
		}
	}()

	for {
		select {
		case chunk, ok := <-chunkChan:
			if !ok {
				// Chunk channel closed, all chunks sent
				goto Done
			}

			req := &pb.UploadFileRequest{
				Data: &pb.UploadFileRequest_ChunkData{
					ChunkData: chunk,
				},
			}

			b := md5.Sum(chunk)
			// checkSumStr = util.Concatinate(checkSumStr, hex.EncodeToString(b[:]))
			checkSumStr = util.Concatinate(checkSumStr, hex.EncodeToString([]byte(hex.EncodeToString(b[:]))))
			// log.Printf("chunk: %v", checkSumStr)

			err = stream.Send(req)
			if err != nil {
				log.Printf("error sending chunk: %v", err)
				goto Done
			}

		case <-ctx.Done():
			// Context canceled or timeout occurred
			log.Printf("file transfer canceled: %v", ctx.Err())
			return &pb.UploadFileResponse{}, ctx.Err()
		}

	}

Done:
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("cannot receive response: %v", err)
		return &pb.UploadFileResponse{}, err
	}

	b := md5.Sum([]byte(checkSumStr))

	// log.Printf("chunk: %v", checkSumStr)
	ret := &pb.UploadFileResponse{CheckSum: hex.EncodeToString([]byte(hex.EncodeToString(b[:])))}
	if res.GetCheckSum() != ret.GetCheckSum() {
		log.Printf("checksum not match %v-%v ", res.GetCheckSum(), ret.GetCheckSum())
		return &pb.UploadFileResponse{}, errors.New("checksum not match")
	}

	// b := md5.Sum([]byte(checkSumStr))
	// ret := &pb.UploadFileResponse{CheckSum: hex.EncodeToString(b[:])}
	// if res.GetCheckSum() != ret.GetCheckSum() {
	// 	log.Printf("checksum not match %v-%v ", res.GetCheckSum(), ret.GetCheckSum())
	// 	return &pb.UploadFileResponse{}, errors.New("checksum not match")
	// }

	log.Printf("final Response: err[%v] %v", err, res)
	return ret, err
}

// Function to check if the error is a transient error that can be retried
func isTransientError(err error) bool {
	// Implement your own logic here to determine if the error is transient
	// For example, network-related errors, server-side congestion, etc.
	// You can check the error type or message to make the determination

	// Return true for transient errors
	return true
}

// func Copy2Docker(ctx context.Context, filePath string) {
// 	// Create a Docker client
// 	cli, err := client.NewClientWithOpts(client.FromEnv)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Open the file to be copied
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	// Create a TAR archive of the file
// 	var buf bytes.Buffer
// 	tarWriter := tar.NewWriter(&buf)
// 	fileInfo, err := file.Stat()
// 	if err != nil {
// 		panic(err)
// 	}
// 	header := &tar.Header{
// 		Name: fileInfo.Name(),
// 		Size: fileInfo.Size(),
// 	}
// 	if err := tarWriter.WriteHeader(header); err != nil {
// 		panic(err)
// 	}
// 	if _, err := io.Copy(tarWriter, file); err != nil {
// 		panic(err)
// 	}
// 	if err := tarWriter.Close(); err != nil {
// 		panic(err)
// 	}

// 	// Create the container
// 	resp, err := cli.ContainerCreate(ctx, &container.Config{
// 		Image: "263d578de57f",
// 		Cmd:   []string{"command_to_run"},
// 	}, nil, nil, "")
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Copy the TAR archive to the container
// 	err = cli.CopyToContainer(ctx, resp.ID, "/var/lib/postgresql/.", &buf, types.CopyToContainerOptions{
// 		AllowOverwriteDirWithFile: true,
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Start the container
// 	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
// 		panic(err)
// 	}

// 	// Wait for the container to exit
// 	_, err = cli.ContainerWait(ctx, resp.ID)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Remove the container
// 	if err := cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{}); err != nil {
// 		panic(err)
// 	}
// }
