package rate

import (
	"context"
	"fmt"
	"log"
	"net"
	"sort"
	"time"

	"github.com/opentracing/opentracing-go"
	pb "github.com/ucy-coast/hotel-app/internal/rate/proto"
)

// Rate implements the rate service
type Rate struct {
	port      int
	addr      string
	dbsession *DatabaseSession
	tracer    opentracing.Tracer
}

// NewRate returns a new server
func NewRate(a string, p int, db *DatabaseSession, tr opentracing.Tracer) *Rate {
	return &Rate{
		addr:      a,
		port:      p,
		dbsession: db,
		tracer:    tr,
	}
}

// Run starts the server
func (s *Rate) Run() error {
	// TODO: Implement me

	if s.port == 0 {
		return fmt.Errorf("server port must be set")
	}

	opts := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Timeout: 120 * time.Second,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			PermitWithoutStream: true,
		}),
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(s.tracer),
		),
	}

	// Create an instance of the gRPC server
	srv := grpc.NewServer(opts...)

	// Register our service implementation with the gRPC server
	pb.RegisterRateServer(srv, s)

	// Register reflection service on gRPC server.
	reflection.Register(srv)

	// Listen for client requests
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Accept and serve incoming client requests
	log.Printf("Start Profile server. Addr: %s:%d\n", s.addr, s.port)
	return srv.Serve(lis)
}

func inTimeSpan(start, end, check time.Time) bool {
	return (check.Equal(start) || check.After(start)) && (check.Equal(end) || check.Before(end))
}

// GetRates gets rates for hotels for specific date range.
func (s *Rate) GetRates(ctx context.Context, req *pb.Request) (*pb.Result, error) {
	// TODO: Implement me
	// HINT: Reuse the implementation from the monolithic implementation
	// HINT: and modify as needed.

	res := new(pb.Result)

	ratePlans, err := s.dbsession.GetRates(req.HotelIds)
	if err != nil {
		return nil, err
	}
	finalRatePlans := make(RatePlans, 0)

	start, _ := time.Parse("2006-01-02", req.InDate)
	end, _ := time.Parse("2006-01-02", req.OutDate)

	sort.Sort(ratePlans)
	for _, rateplan := range ratePlans {
		in, _ := time.Parse("2006-01-02", rateplan.InDate)
		out, _ := time.Parse("2006-01-02", rateplan.OutDate)
		if inTimeSpan(in, out, start) && inTimeSpan(in, out, end) {
			finalRatePlans = append(finalRatePlans, rateplan)
		}
	}

	res.RatePlans = finalRatePlans

	return res, nil

}

type RatePlans []*pb.RatePlan

func (r RatePlans) Len() int {
	return len(r)
}

func (r RatePlans) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r RatePlans) Less(i, j int) bool {
	return r[i].RoomType.TotalRate > r[j].RoomType.TotalRate
}
