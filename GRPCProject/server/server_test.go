package server

import (
	"GRPCProject/mocks"
	pb "GRPCProject/proto"
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestServer_CreateMovie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMovieSer := mocks.NewMockMovieServiceServer(ctrl)
	mockRepository := &mocks.RepositoryI{}

	mockRepository.On("Create", mock.Anything).Return(nil)

	server := Server{
		UnimplementedMovieServiceServer: pb.UnimplementedMovieServiceServer{},
		DB:                              mockRepository,
	}

	mockMovieSer.EXPECT().CreateMovie(gomock.Any(), gomock.Any()).Return(&pb.CreateMovieResponse{}, nil).AnyTimes()

	req := &pb.Movie{
		Title: "kgf",
		Genre: "action",
	}

	convey.Convey("", t, func() {
		res, _ := server.CreateMovie(context.Background(), &pb.CreateMovieRequest{
			Movie: req,
		})
		convey.Convey("assert req Title with recvd title", func() {
			convey.So(req.Title, convey.ShouldEqual, res.GetMovie().GetTitle())
		})
	})

}

func TestServer_GetMovie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMovieSer := mocks.NewMockMovieServiceClient(ctrl)
	mockRepository := &mocks.RepositoryI{}

	mockRepository.On("Query", mock.Anything).Return(nil)

	server := Server{
		UnimplementedMovieServiceServer: pb.UnimplementedMovieServiceServer{},
		DB:                              mockRepository,
	}

	mockMovieSer.EXPECT().GetMovie(gomock.Any(), gomock.Any()).Return(&pb.ReadMovieResponse{
		Movie: &pb.Movie{
			Id:    "12345",
			Title: "abc",
			Genre: "dono",
		},
	}, nil).AnyTimes()

	convey.Convey("", t, func() {
		res, _ := server.GetMovie(context.Background(), &pb.ReadMovieRequest{
			Id: "12345",
		})
		fmt.Println(res.GetMovie().GetTitle())
		convey.Convey("assert req Title with recvd title", func() {
			convey.So("12345", convey.ShouldEqual, res.GetMovie().GetId())
		})
	})

}

func TestServer_UpdateMovie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMovieSer := mocks.NewMockMovieServiceClient(ctrl)
	mockRepository := &mocks.RepositoryI{}

	mockRepository.On("Query", mock.Anything).Return(nil)

	server := Server{
		UnimplementedMovieServiceServer: pb.UnimplementedMovieServiceServer{},
		DB:                              mockRepository,
	}

	_, _ = server.GetMovie(context.Background(), &pb.ReadMovieRequest{
		Id: "12345",
	})

	mockRepository.On("Update", mock.Anything).Return(nil)

	server = Server{
		UnimplementedMovieServiceServer: pb.UnimplementedMovieServiceServer{},
		DB:                              mockRepository,
	}

	mockMovieSer.EXPECT().UpdateMovie(gomock.Any(), gomock.Any()).Return(&pb.UpdateMovieResponse{
		Movie: &pb.Movie{
			Id:    "12345",
			Title: "abc",
			Genre: "dono",
		},
	}, nil).AnyTimes()

	convey.Convey("", t, func() {
		res, _ := server.UpdateMovie(context.Background(), &pb.UpdateMovieRequest{
			Movie: &pb.Movie{
				Id:    "12345",
				Title: "xyz",
			},
		})
		fmt.Println(res.GetMovie().GetTitle())
		convey.Convey("assert req Title with recvd title", func() {
			convey.So("xyz", convey.ShouldEqual, res.GetMovie().GetTitle())
		})
	})

}

func TestServer_DeleteMovie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMovieSer := mocks.NewMockMovieServiceClient(ctrl)
	mockRepository := &mocks.RepositoryI{}

	mockRepository.On("Delete", mock.Anything).Return(nil)

	server := Server{
		UnimplementedMovieServiceServer: pb.UnimplementedMovieServiceServer{},
		DB:                              mockRepository,
	}

	mockMovieSer.EXPECT().DeleteMovie(gomock.Any(), gomock.Any()).Return(&pb.DeleteMovieResponse{
		Sucess: true,
	}, nil).AnyTimes()

	convey.Convey("", t, func() {
		res, _ := server.DeleteMovie(context.Background(), &pb.DeleteMovieRequest{
			Id: "12345",
		})
		fmt.Println(res.Sucess)
		convey.Convey("assert req Title with recvd title", func() {
			convey.So(true, convey.ShouldEqual, res.GetSucess())
		})
	})

}
