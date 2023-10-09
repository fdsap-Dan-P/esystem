package service

import (
	"context"
	"log"

	pb "simplebank/pb"
	"simplebank/util"

	dump "simplebank/db/datastore/esystemdump"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	timestamp "google.golang.org/protobuf/types/known/timestamppb"
)

// CreateDumpBranchList(ctx context.Context, in *CreateDumpBranchListRequest, opts ...grpc.CallOption) (*CreateDumpBranchListResponse, error)
func (server *DumpServer) CreateDumpBranchList(ctx context.Context, req *pb.CreateDumpBranchListRequest) (*pb.CreateDumpBranchListResponse, error) {

	log.Printf("received BranchList request: CreateDumpBranchList id = %s", "req")

	branchList := req.BranchList

	log.Printf("received BranchList request: string(branchList.ESystemVer) = %s", string(branchList.ESystemVer))

	branchListReq := dump.BranchList{
		BrCode:         branchList.BrCode,
		EbSysDate:      branchList.EbSysDate.AsTime(),
		RunState:       branchList.RunState,
		OrgAddress:     branchList.OrgAddress,
		TaxInfo:        branchList.TaxInfo,
		DefCity:        branchList.DefCity,
		DefProvince:    branchList.DefProvince,
		DefCountry:     branchList.DefCountry,
		DefZip:         branchList.DefZip,
		WaivableInt:    branchList.WaivableInt,
		DBVersion:      branchList.DBVersion,
		ESystemVer:     branchList.ESystemVer,
		NewBrCode:      branchList.NewBrCode,
		LastConnection: branchList.LastConnection.AsTime(),
	}
	err := server.dump.CreateBranchList(ctx, branchListReq)
	if err != nil {
		return nil, util.LogError(status.Errorf(codes.Unknown, "Unable to create BranchList: %v", err))
	}
	log.Printf("branchListReq: %v", branchListReq)
	return &pb.CreateDumpBranchListResponse{BrCode: branchListReq.BrCode}, nil
}

func (server *DumpServer) GetDumpBranchList(ctx context.Context, req *pb.GetDumpBranchListRequest) (*pb.DumpBranchList, error) {
	br, err := server.dump.GetBranchList(ctx, req.BrCode)
	if err != nil {
		return &pb.DumpBranchList{}, err
	}
	return &pb.DumpBranchList{
		BrCode:         br.BrCode,
		EbSysDate:      timestamp.New(br.EbSysDate),
		RunState:       br.RunState,
		OrgAddress:     br.OrgAddress,
		TaxInfo:        br.TaxInfo,
		DefCity:        br.DefCity,
		DefProvince:    br.DefProvince,
		DefCountry:     br.DefCountry,
		DefZip:         br.DefZip,
		WaivableInt:    br.WaivableInt,
		DBVersion:      br.DBVersion,
		ESystemVer:     []byte(br.ESystemVer),
		NewBrCode:      br.NewBrCode,
		LastConnection: timestamp.New(br.LastConnection),
	}, nil
}
