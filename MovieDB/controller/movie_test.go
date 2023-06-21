package controller

import (
	"MovieDB/mocks"
	"MovieDB/models"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
)

var rPath = "/test"

func TestMoviecontroller_AddMovie(t *testing.T) {
	{
		convey.Convey("given a movie to add", t, func() {
			body := models.Movie{Name: "Titanic", Language: "Hindi", Length: 1}
			mockRepository := &mocks.RepositoryI{}
			mockRepository.On("Create", mock.Anything).Return(nil)

			router := gin.Default()
			m := Moviecontroller{}
			router.POST(rPath, m.AddMovie)

			m.Repository = mockRepository
			b, _ := json.Marshal(body)
			req, _ := http.NewRequest("POST", rPath, bytes.NewReader(b))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			t.Logf("status: %d", w.Code)
			t.Logf("response:%s", w.Body.String())

			convey.Convey("assert 201 OK", func() {
				convey.So(w.Code, convey.ShouldEqual, http.StatusCreated)
			})
		})
	}
}
func TestMoviecontroller_NegAddMovie(t *testing.T) {
	{
		convey.Convey("given a movie to add with a missing item", t, func() {
			body := models.Movie{Name: "Titanic", Length: 1}
			mockRepository := &mocks.RepositoryI{}
			mockRepository.On("Create", mock.Anything).Return(http.StatusBadRequest, errors.New("err"))

			router := gin.Default()
			m := Moviecontroller{}
			router.POST(rPath, m.AddMovie)

			m.Repository = mockRepository
			b, _ := json.Marshal(body)
			req, _ := http.NewRequest("POST", rPath, bytes.NewReader(b))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			t.Logf("status: %d", w.Code)
			t.Logf("response:%s", w.Body.String())

			convey.Convey("assert 400", func() {
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
			})
		})
	}
}
func TestMoviecontroller_GetMovieslist(t *testing.T) {
	{
		convey.Convey("given a movie get all the movie", t, func() {
			body := models.Movie{}
			mockRepository := &mocks.RepositoryI{}
			mockRepository.On("QueryAll", mock.Anything).Return(nil)

			router := gin.Default()
			m := Moviecontroller{}
			router.GET(rPath, m.GetMovies)

			m.Repository = mockRepository
			b, _ := json.Marshal(body)
			req, _ := http.NewRequest("GET", rPath, bytes.NewReader(b))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			t.Logf("status: %d", w.Code)
			t.Logf("response:%s", w.Body.String())

			if w.Code <= int(body.Length) {
				convey.Convey("assert 200 OK", func() {
					convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
				})
			}
		})
	}
}
func TestMoviecontroller_GetMovies(t *testing.T) {
	{
		convey.Convey("given a movie to get", t, func() {
			body := models.Movie{Id: "1"}
			mockRepository := &mocks.RepositoryI{}
			mockRepository.On("Query", mock.Anything).Return(nil)

			router := gin.Default()
			m := Moviecontroller{}
			router.GET(rPath, m.GetByIdMovies)

			m.Repository = mockRepository
			b, _ := json.Marshal(body)
			req, _ := http.NewRequest("GET", rPath, bytes.NewReader(b))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			t.Logf("status: %d", w.Code)
			t.Logf("response:%s", w.Body.String())

			convey.Convey("assert 200 OK", func() {
				convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
			})
		})
	}
}

func TestMoviecontroller_UpdateMovie(t *testing.T) {
	{
		convey.Convey("given a movie to update", t, func() {
			body := models.Movie{Id: "1", Name: "Titanic", Language: "english", Length: 2}
			mockRepository := &mocks.RepositoryI{}
			mockRepository.On("Update", mock.Anything).Return(nil)

			router := gin.Default()
			m := Moviecontroller{}
			router.PATCH(rPath, m.UpdateMovie)

			m.Repository = mockRepository
			b, _ := json.Marshal(body)
			req, _ := http.NewRequest("PATCH", rPath, bytes.NewReader(b))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			t.Logf("status: %d", w.Code)
			t.Logf("response:%s", w.Body.String())

			convey.Convey("assert 200 OK", func() {
				convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
			})
		})
	}
}

func TestMoviecontroller_NegUpdateMovie(t *testing.T) {
	{
		convey.Convey("given a movie to get with a missing value", t, func() {
			body := models.Movie{}
			mockRepository := &mocks.RepositoryI{}
			mockRepository.On("Update", mock.Anything).Return(http.StatusBadRequest, errors.New("err"))

			router := gin.Default()
			m := Moviecontroller{}
			router.PATCH(rPath, m.UpdateMovie)

			m.Repository = mockRepository
			b, _ := json.Marshal(body)
			req, _ := http.NewRequest("PATCH", rPath, bytes.NewReader(b))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			t.Logf("status: %d", w.Code)
			t.Logf("response:%s", w.Body.String())

			convey.Convey("assert 400", func() {
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
			})
		})
	}
}
func TestMoviecontroller_DeleteMovie(t *testing.T) {
	{
		convey.Convey("given a movie to delete by id", t, func() {
			body := models.Movie{Id: "1", Name: "titanic", Language: "english", Length: 2}
			mockRepository := &mocks.RepositoryI{}
			mockRepository.On("Delete", mock.Anything).Return(nil)

			router := gin.Default()
			m := Moviecontroller{}
			router.DELETE(rPath, m.DeleteMovie)

			m.Repository = mockRepository
			b, _ := json.Marshal(body)
			req, _ := http.NewRequest("DELETE", rPath, bytes.NewReader(b))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			t.Logf("status: %d", w.Code)
			t.Logf("response:%s", w.Body.String())

			convey.Convey("assert 200 OK", func() {
				convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
			})
		})
	}
}

func TestMoviecontroller_NegDeleteMovie(t *testing.T) {
	{
		convey.Convey("given a movie to delete without id", t, func() {
			body := models.Movie{}
			mockRepository := &mocks.RepositoryI{}
			mockRepository.On("Delete", mock.Anything).Return(errors.New("err"))

			router := gin.Default()
			m := Moviecontroller{
				Repository: mockRepository,
			}
			router.DELETE(rPath, m.DeleteMovie)

			b, _ := json.Marshal(body)
			req, _ := http.NewRequest("DELETE", rPath, bytes.NewReader(b))

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			t.Logf("status: %d", w.Code)
			t.Logf("response:%s", w.Body.String())

			convey.Convey("assert 400", func() {
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
			})
		})
	}
}

func TestMoviecontroller_NegGetMovies(t *testing.T) {
	{
		convey.Convey("given a movie to get without id", t, func() {
			body := models.Movie{Id: ""}
			mockRepository := &mocks.RepositoryI{}
			mockRepository.On("Query", mock.Anything).Return(errors.New("err"))

			router := gin.Default()
			m := Moviecontroller{}
			router.GET(rPath, m.GetByIdMovies)

			m.Repository = mockRepository
			b, _ := json.Marshal(body)
			req, _ := http.NewRequest("GET", rPath, bytes.NewReader(b))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			t.Logf("status: %d", w.Code)
			t.Logf("response:%s", w.Body.String())

			convey.Convey("assert 400", func() {
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
			})
		})
	}
}

// func TestMoviecontroller_GetMovies(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		body models.Movie
// 		mock func() repository.RepositoryI
// 		want int
// 	}{{
// 		name: "pass",
// 		body: models.Movie{Id: "test"},
// 		mock: func() repository.RepositoryI {
// 			ri := mocks.RepositoryI{}
// 			ri.On("QueryAll", mock.Anything).Return(nil)
// 			return &ri
// 		},
// 		want: http.StatusOK,
// 	}}

// 	router := gin.Default()
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			m := &Moviecontroller{Repository: tt.mock()}
// 			router.GET(rPath, m.GetMovies)
// 			b, _ := json.Marshal(tt.body)
// 			req, _ := http.NewRequest("GET", rPath, bytes.NewReader(b))
// 			w := httptest.NewRecorder()
// 			router.ServeHTTP(w, req)
// 			t.Logf("status: %d", w.Code)
// 			t.Logf("response: %s", w.Body.String())
// 			assert.Equal(t, tt.want, w.Code)
// 		})
// 	}
// }

// func TestMoviecontroller_UpdateMovie(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		body models.Movie
// 		mock func() repository.RepositoryI
// 		want int
// 	}{{
// 		name: "pass",
// 		body: models.Movie{
// 			Id:       "id",
// 			Name:     "Titanic",
// 			Language: "Hindi",
// 			Length:   2,
// 		},
// 		mock: func() repository.RepositoryI {
// 			ri := mocks.RepositoryI{}
// 			ri.On("Update", mock.Anything).Return(nil)
// 			return &ri
// 		},
// 		want: http.StatusOK,
// 	}}

// 	router := gin.Default()
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			m := &Moviecontroller{Repository: tt.mock()}
// 			router.GET(rPath, m.UpdateMovie)
// 			b, _ := json.Marshal(tt.body)
// 			req, _ := http.NewRequest("GET", rPath, bytes.NewReader(b))
// 			w := httptest.NewRecorder()
// 			router.ServeHTTP(w, req)
// 			t.Logf("status: %d", w.Code)
// 			t.Logf("response: %s", w.Body.String())
// 			assert.Equal(t, tt.want, w.Code)
// 		})
// 	}
// }

// func TestMoviecontroller_DeleteMovie(t *testing.T) {

// 	tests := []struct {
// 		name string
// 		body models.Movie
// 		mock func() repository.RepositoryI
// 		want int
// 	}{{
// 		name: "pass",
// 		body: models.Movie{
// 			Id:       "id",
// 			Name:     "Titanic",
// 			Language: "Hindi",
// 			Length:   2,
// 		},
// 		mock: func() repository.RepositoryI {
// 			ri := mocks.RepositoryI{}
// 			ri.On("Delete", mock.Anything).Return(nil)
// 			return &ri
// 		},
// 		want: http.StatusOK,
// 	}}

// 	router := gin.Default()
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			m := &Moviecontroller{Repository: tt.mock()}
// 			router.GET(rPath, m.DeleteMovie)
// 			b, _ := json.Marshal(tt.body)
// 			req, _ := http.NewRequest("GET", rPath, bytes.NewReader(b))
// 			w := httptest.NewRecorder()
// 			router.ServeHTTP(w, req)
// 			t.Logf("status: %d", w.Code)
// 			t.Logf("response: %s", w.Body.String())
// 			assert.Equal(t, tt.want, w.Code)
// 		})
// 	}
// }
