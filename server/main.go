package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "github.com/im-raguram/ticket-reservation/protos"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

var (
	seatInfo []*pb.SeatInfo
)

type ticketServer struct {
	pb.UnimplementedTicketServiceServer
}

func main() {
	initSeats()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterTicketServiceServer(s, &ticketServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initSeats() {

	seat1 := &pb.SeatInfo{
		Ticket: &pb.PurchaseRequest{
			From: "Pudukkottai",
			To:   "Chennai",
			User: &pb.User{
				FirstName: "Raguram",
				LastName:  "Shanmugam",
				Email:     "raguram806@gmail.com",
			},
			Price: 390,
		},
		Seat: &pb.Seat{
			Section:    "A",
			SeatNumber: 1,
		},
	}

	seat2 := &pb.SeatInfo{
		Ticket: &pb.PurchaseRequest{
			From: "Karaikudi",
			To:   "Chennai",
			User: &pb.User{
				FirstName: "Cloud",
				LastName:  "Bees",
				Email:     "cloudbees@gmail.com",
			},
			Price: 420,
		},
		Seat: &pb.Seat{
			Section:    "B",
			SeatNumber: 1,
		},
	}

	seat3 := &pb.SeatInfo{
		Ticket: &pb.PurchaseRequest{
			From: "Trichy",
			To:   "Chennai",
			User: &pb.User{
				FirstName: "Ram",
				LastName:  "Sathya",
				Email:     "ramsathya@gmail.com",
			},
			Price: 310,
		},
		Seat: &pb.Seat{
			Section:    "B",
			SeatNumber: 2,
		},
	}

	seatInfo = append(seatInfo, seat1)
	seatInfo = append(seatInfo, seat2)
	seatInfo = append(seatInfo, seat3)
}

func (s *ticketServer) PurchaseTicket(ctx context.Context, input *pb.PurchaseRequest) (*pb.PurchaseResponse, error) {
	log.Printf("Received: %v", input)

	seat := pb.Seat{
		Section:    "A",
		SeatNumber: 2,
	}

	res := pb.SeatInfo{
		Ticket: input,
		Seat:   &seat,
	}

	seatInfo = append(seatInfo, &res)

	output := pb.PurchaseResponse{
		Status: "Confirmed",
		Seat:   &seat,
	}

	return &output, nil
}

func (s *ticketServer) ViewSeatBySection(input *pb.SectionID, stream pb.TicketService_ViewSeatBySectionServer) error {
	log.Printf("Received: %v", input)

	for _, seat := range seatInfo {
		if seat.Seat.Section == input.Id {
			if err := stream.Send(seat); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *ticketServer) ViewSeatByEmail(ctx context.Context, input *pb.Email) (*pb.SeatInfo, error) {
	log.Printf("Received: %v", input)

	var output *pb.SeatInfo

	for _, seat := range seatInfo {
		if seat.Ticket.User.Email == input.Email {
			output = seat
		}
	}

	return output, nil
}

func (s *ticketServer) RemoveUser(ctx context.Context, input *pb.Email) (*pb.Status, error) {
	log.Printf("Received: %v", input)

	var output *pb.Status

	for i, seat := range seatInfo {
		if seat.Ticket.User.Email == input.Email {
			seatInfo = append(seatInfo[:i], seatInfo[i+1:]...)
			output = &pb.Status{
				Message: "success",
			}
		}
	}

	return output, nil
}

func (s *ticketServer) ModifySeat(ctx context.Context, input *pb.ModifySeatRequest) (*pb.PurchaseResponse, error) {
	log.Printf("Received: %v", input)

	var (
		output      *pb.PurchaseResponse
		oldSeatInfo *pb.SeatInfo
	)

	for _, seat := range seatInfo {
		if input.NewSeat.SeatNumber == seat.Seat.SeatNumber {
			return output, errors.New("seat already occupied")
		}
		if seat.Ticket.User.Email == input.Email {

			oldSeatInfo = seat

			email := pb.Email{
				Email: input.Email,
			}

			_, err := s.RemoveUser(ctx, &email)
			if err != nil {
				return output, err
			}

			newSeatInfo := pb.SeatInfo{
				Ticket: oldSeatInfo.Ticket,
				Seat:   input.NewSeat,
			}

			seatInfo = append(seatInfo, &newSeatInfo)
			output = &pb.PurchaseResponse{
				Status: "confirmed",
				Seat:   input.NewSeat,
			}

		}
	}

	return output, nil
}
