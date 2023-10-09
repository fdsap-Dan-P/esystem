package client

import (
	"context"
	"log"
	dump "simplebank/db/datastore/esystemdump"
	local "simplebank/db/datastore/esystemlocal"
	pb "simplebank/pb"

	timestamp "google.golang.org/protobuf/types/known/timestamppb"
)

func (dumpClient *DumpClient) CreateDumpBranchList(req local.OrgParmsInfo) (*pb.CreateDumpBranchListResponse, error) {
	brList := &pb.CreateDumpBranchListRequest{
		BranchList: &pb.DumpBranchList{
			BrCode:      req.BrCode,
			EbSysDate:   timestamp.New(req.EbSysDate),
			RunState:    req.RunState,
			OrgAddress:  req.OrgAddress,
			TaxInfo:     req.TaxInfo,
			DefCity:     req.DefCity,
			DefProvince: req.DefProvince,
			DefCountry:  req.DefCountry,
			DefZip:      req.DefZip,
			WaivableInt: req.WaivableInt,
			DBVersion:   req.DBVersion,
			ESystemVer:  req.ESystemVer,
			NewBrCode:   req.NewBrCode,
		},
	}
	log.Printf("CreateDumpBranchList: req.ESystemVer [%v]", string(req.ESystemVer))
	log.Printf("CreateDumpBranchList: %v", brList)

	br, err := dumpClient.service.CreateDumpBranchList(context.Background(), brList)
	log.Printf("CreateDumpBranchList: Error[%v]", err)
	return br, err
}

func (dumpClient *DumpClient) GetDumpBranchList(brCode string) (dump.BranchList, error) {

	br, err := dumpClient.service.GetDumpBranchList(context.Background(), &pb.GetDumpBranchListRequest{BrCode: brCode})
	if err != nil {
		return dump.BranchList{}, err
	}

	return dump.BranchList{
		BrCode:      br.BrCode,
		EbSysDate:   br.EbSysDate.AsTime(),
		RunState:    br.RunState,
		OrgAddress:  br.OrgAddress,
		TaxInfo:     br.TaxInfo,
		DefCity:     br.DefCity,
		DefProvince: br.DefProvince,
		DefCountry:  br.DefCountry,
		DefZip:      br.DefZip,
		WaivableInt: br.WaivableInt,
		DBVersion:   br.DBVersion,
		ESystemVer:  br.ESystemVer,
		NewBrCode:   br.NewBrCode,
	}, nil
}
