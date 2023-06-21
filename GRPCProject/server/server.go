package server

import (
	"GRPCProject/database"
	"GRPCProject/model"
	pb "GRPCProject/proto"
	"context"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() database.RepositoryI {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(new(model.Movie))
	repo := database.Repository{Db: db}
	return &repo
}

var err error

type Server struct {
	pb.UnimplementedMovieServiceServer
	DB database.RepositoryI
}

func (s *Server) CreateMovie(ctx context.Context, req *pb.CreateMovieRequest) (*pb.CreateMovieResponse, error) {
	fmt.Println("Create Movie")
	movie := req.GetMovie()
	movie.Id = uuid.New().String()

	data := model.Movie{
		ID:    movie.GetId(),
		Title: movie.GetTitle(),
		Genre: movie.GetGenre(),
	}

	err := s.DB.Create(&data)
	if err != nil {
		return &pb.CreateMovieResponse{Msg: "add movie failed"}, err
	}
	return &pb.CreateMovieResponse{
		Movie: &pb.Movie{
			Id:    movie.GetId(),
			Title: movie.GetTitle(),
			Genre: movie.GetGenre(),
		},
	}, nil
}

func (s *Server) GetMovie(ctx context.Context, req *pb.ReadMovieRequest) (*pb.ReadMovieResponse, error) {
	fmt.Println("Read Movie", req.GetId())
	var movie model.Movie
	movie.ID = req.GetId()
	err := s.DB.Query(&movie)
	if err != nil {
		return &pb.ReadMovieResponse{}, err
	}
	return &pb.ReadMovieResponse{
		Movie: &pb.Movie{
			Id:    movie.ID,
			Title: movie.Title,
			Genre: movie.Genre,
		},
	}, nil
}

func (s *Server) GetMovies(ctx context.Context, req *pb.ReadMoviesRequest) (*pb.ReadMoviesResponse, error) {
	fmt.Println("Read Movies")
	var movies []*pb.Movie
	err := s.DB.QueryAll(&movies)
	if err != nil {
		return &pb.ReadMoviesResponse{}, err
	}

	return &pb.ReadMoviesResponse{
		Movies: movies,
	}, nil
}

func (s *Server) UpdateMovie(ctx context.Context, req *pb.UpdateMovieRequest) (*pb.UpdateMovieResponse, error) {
	fmt.Println("Update Movie")
	reqMovie := req.GetMovie()
	var movie pb.Movie
	movie.Id = req.GetMovie().GetId()
	err := s.DB.Query(&movie)
	if err != nil {
		return &pb.UpdateMovieResponse{}, err
	}
	err = s.DB.Update(&reqMovie)
	if err != nil {
		return &pb.UpdateMovieResponse{}, err
	}
	fmt.Println("request movie", reqMovie)
	return &pb.UpdateMovieResponse{
		Movie: &pb.Movie{
			Id:    reqMovie.GetId(),
			Title: reqMovie.GetTitle(),
			Genre: reqMovie.GetGenre(),
		},
	}, nil
}

func (S *Server) DeleteMovie(ctx context.Context, req *pb.DeleteMovieRequest) (*pb.DeleteMovieResponse, error) {
	fmt.Println("Delete Movie")
	var movie model.Movie
	movie.ID = req.GetId()
	err := S.DB.Delete(&movie)
	if err != nil {
		return &pb.DeleteMovieResponse{}, err
	}
	return &pb.DeleteMovieResponse{
		Sucess: true,
	}, nil
}
