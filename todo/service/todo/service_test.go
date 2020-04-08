package todo

import (
	mock_repository "github.com/cjcjcj/todo/mocks/github.com/cjcjcj/todo/todo/repository"
	"github.com/cjcjcj/todo/todo/repository"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestNewTodoService(t *testing.T) {
	type args struct {
		todoRepo repository.TodoRepository
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockTodoRepository(ctrl)

	tests := []struct {
		name string
		args args
		want *todoService
	}{
		{
			name: "OK",
			args: args{
				todoRepo: mockRepo,
			},
			want: &todoService{
				repo: mockRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTodoService(tt.args.todoRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTodoService() = %v, want %v", got, tt.want)
			}
		})
	}
}