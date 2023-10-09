package sample

import (
	pb "simplebank/pb"
)

// NewArea returns a new sample keyboard
func NewDumpArea() *pb.DumpArea {
	area := &pb.DumpArea{
		ModCtr:   1,
		BrCode:   "01",
		AreaCode: 1,
		Area:     &pb.NullString{Value: "Area Name", Valid: true},
	}

	return area
}
