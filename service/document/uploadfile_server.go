package service

import (
	"bufio"
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
	pb "simplebank/pb"
	"simplebank/util"

	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	// eSys "simplebank/db/datastore/esystemdump"
)

const maxFileSize = 1 << 20

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (server *DocumentServer) UploadFile(stream pb.DocumentService_UploadFileServer) error {

	fileData := bytes.Buffer{}
	fileSize := 0
	checkSumStr := ""

	req, err := stream.Recv()
	if err != nil {
		return util.LogError(status.Errorf(codes.Unknown, "cannot receive image info"))
	}

	log.Printf("UploadFile: --> %v", util.Struct2Json(req.GetInfo()))

	brCode := req.GetInfo().GetReferenceCode()
	// fileType := req.GetInfo().GetFileType()
	fileName := fmt.Sprintf("%v-%v", brCode, req.GetInfo().GetFileName())
	targetTable := req.GetInfo().GetTargetTable()
	serverPath := req.GetInfo().GetServerPath()
	dockerImgID := req.GetInfo().GetDockerImgID()
	dockerPath := req.GetInfo().GetDockerPath()

	f, err := os.Create(fileName)
	check(err)
	defer f.Close()
	f.Sync()

	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		log.Print("waiting to receive more data")

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}

		chunk := req.GetChunkData()
		b := md5.Sum([]byte(chunk[:]))
		checkSumStr = util.Concatinate(checkSumStr, hex.EncodeToString([]byte(hex.EncodeToString(b[:]))))
		// log.Printf("chunk: %v", checkSumStr)

		size := len(chunk)

		log.Printf("received a chunk with size: %d-> chunk: %v ", size, b)

		fileSize += size
		// if fileSize > maxFileSize {
		// 	return util.LogError(status.Errorf(codes.InvalidArgument, "file is too large: %d > %d", fileSize, maxFileSize))
		// }

		// write slowly
		// time.Sleep(time.Second)

		w := bufio.NewWriter(f)
		_, err = w.Write(chunk)
		check(err)
		w.Flush()

		_, err = fileData.Write(chunk)
		if err != nil {
			return util.LogError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
		}
	}

	b := md5.Sum([]byte(checkSumStr))
	// log.Printf("chunk: %v", checkSumStr)
	res := &pb.UploadFileResponse{CheckSum: hex.EncodeToString([]byte(hex.EncodeToString(b[:])))}

	err = stream.SendAndClose(res)
	if err != nil {
		return util.LogError(status.Errorf(codes.Unknown, "cannot send response: %v", err))
	}

	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(dir)

	// dockerImgID := "263d578de57f"
	// serverPath := "/eSystem/simplebank/cmd/server/"
	// serverPath := "/Users/rhickmercado/Documents/Programming/go/src/simplebank/cmd/server"
	// dockerPath := "/var/lib/postgresql/"
	// dockerImgID := "57792c1eb13f"
	cmdStr := fmt.Sprintf(
		"docker cp %s%s %s:%s%s ; docker exec %s chmod 777 %v%v ",
		serverPath,
		fileName,
		dockerImgID,
		dockerPath,
		fileName, dockerImgID, dockerPath, fileName,
	)
	log.Printf("cmdStr: %v", cmdStr)
	cmd := exec.Command("/bin/sh", "-c", cmdStr)

	output, err := cmd.CombinedOutput()
	if err != nil {
		// The command failed.
		fmt.Printf("err:%v - output: %s", err, output)
	}

	// err = cmd.Start()
	// if err != nil {
	// 	// The command failed to start.
	// 	util.LogError(err)
	// }

	// // Wait for the command to finish.
	// err = cmd.Wait()
	// if err != nil {
	// 	util.LogError(err)
	// }

	// err = cmd.Run()
	// if err != nil {
	// 	return util.LogError(status.Errorf(codes.Unknown, "cannot copy to Postgres Server: %v", err))
	// }

	err = server.document.LoadData(context.Background(),
		targetTable, fmt.Sprintf("/var/lib/postgresql/%v", fileName), brCode)

	if err != nil {
		return util.LogError(status.Errorf(codes.Unknown, "cannot Load Data to Table: %v", err))
	}
	return nil
}

// docker cp 57792c1eb13f:/var/lib/postgresql/data/postgresql.conf .
// check the memmory
// docker stats 57792c1eb13f
// docker restart 57792c1eb13f
// docker cp ./postgresql.conf 57792c1eb13f:/var/lib/postgresql/data/postgresql.conf
// docker exec -it 57792c1eb13f bash
