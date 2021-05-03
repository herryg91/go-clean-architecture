package video_repository_v1

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/herryg91/go-clean-architecture/examples/video-rest-api/entity"
	"github.com/herryg91/go-clean-architecture/examples/video-rest-api/pkg/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {
	fmt.Println("init db for test")
	var err error

	loop := 0
	maxLoop := 100
	for {
		db, err = mysql.Connect("localhost", "33061", "root", "password", "gca", logger.LogLevel(3))
		if err == nil {
			break
		}
		loop++
		if loop > maxLoop {
			panic("Failed to init db, 100 times test already")
		}
		time.Sleep(time.Second)
	}
}

func Test_repository_Get(t *testing.T) {
	type fields struct {
		db        *gorm.DB
		tableName string
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Video
		wantErr bool
	}{
		{
			name:   "success: standart scenario",
			fields: fields{db: db, tableName: "videos"},
			args:   args{id: 1},
			want: &entity.Video{
				Id:    1,
				Title: "Learn Go Programming - Golang Tutorial for Beginners",
				Url:   "https://www.youtube.com/watch?v=YS4e4q9oBaU",
			},
			wantErr: false,
		},
		{
			name:    "failed: not found",
			fields:  fields{db: db, tableName: "videos"},
			args:    args{id: 100},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "failed: wrong tablename (human/dev error)",
			fields:  fields{db: db, tableName: "videos2"},
			args:    args{id: 100},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db:        tt.fields.db,
				tableName: tt.fields.tableName,
			}
			got, err := r.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_GetAll(t *testing.T) {
	type fields struct {
		db        *gorm.DB
		tableName string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*entity.Video
		wantErr bool
	}{
		{
			name:   "success: standart scenario",
			fields: fields{db: db, tableName: "videos"},
			want: []*entity.Video{
				&entity.Video{Id: 1, Title: "Learn Go Programming - Golang Tutorial for Beginners", Url: "https://www.youtube.com/watch?v=YS4e4q9oBaU"},
				&entity.Video{Id: 2, Title: "7 Habits of Highly Effective Programmers (ft. ex-Google TechLead)", Url: "https://www.youtube.com/watch?v=W8ykZNSLDqE"},
				&entity.Video{Id: 3, Title: "Building a Bank with Go", Url: "https://www.youtube.com/watch?v=y2j_TB3NsRc"},
				&entity.Video{Id: 4, Title: "ITkonekt 2019 | Robert C. Martin (Uncle Bob), Clean Architecture and Design", Url: "https://www.youtube.com/watch?v=2dKZ-dWaCiU"},
				&entity.Video{Id: 5, Title: "The Principles of Clean Architecture by Uncle Bob Martin", Url: "https://www.youtube.com/watch?v=o_TH-Y78tt4"},
				&entity.Video{Id: 6, Title: "Making Architecture Matter - Martin Fowler Keynote", Url: "https://www.youtube.com/watch?v=DngAZyWMGR0"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db:        tt.fields.db,
				tableName: tt.fields.tableName,
			}
			got, err := r.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, g := range got {
				fmt.Println(*g)
				// if !reflect.DeepEqual(g, tt.want[i]) {
				// 	t.Errorf("repository.GetAll() = %v, want %v", got, tt.want)
				// }
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_Create(t *testing.T) {
	type fields struct {
		db        *gorm.DB
		tableName string
	}
	type args struct {
		in entity.Video
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Video
		wantErr bool
	}{
		{
			name:    "success: standart scenario",
			fields:  fields{db: db, tableName: "videos"},
			args:    args{in: entity.Video{Title: "Test Video", Url: "http://localhost/test-video"}},
			want:    &entity.Video{Id: 7, Title: "Test Video", Url: "http://localhost/test-video"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db:        tt.fields.db,
				tableName: tt.fields.tableName,
			}
			got, err := r.Create(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_Update(t *testing.T) {
	type fields struct {
		db        *gorm.DB
		tableName string
	}
	type args struct {
		in entity.Video
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Video
		wantErr bool
	}{
		{
			name:    "success: standart scenario",
			fields:  fields{db: db, tableName: "videos"},
			args:    args{in: entity.Video{Id: 2, Title: "Test Update Video", Url: "http://localhost/test-update-video"}},
			want:    &entity.Video{Id: 2, Title: "Test Update Video", Url: "http://localhost/test-update-video"},
			wantErr: false,
		},
		{
			name:    "success: nothing to update",
			fields:  fields{db: db, tableName: "videos"},
			args:    args{in: entity.Video{Id: 101, Title: "Test Update Video", Url: "http://localhost/test-update-video"}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db:        tt.fields.db,
				tableName: tt.fields.tableName,
			}
			got, err := r.Update(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_Delete(t *testing.T) {
	type fields struct {
		db        *gorm.DB
		tableName string
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "success: standart scenario",
			fields:  fields{db: db, tableName: "videos"},
			args:    args{id: 6},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db:        tt.fields.db,
				tableName: tt.fields.tableName,
			}
			if err := r.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("repository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
