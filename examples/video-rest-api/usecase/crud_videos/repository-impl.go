package crud_videos

import (
	"fmt"

	video_datasource "github.com/herryg91/go-clean-architecture/examples/video-rest-api/drivers/datasource/mysql/videos"

	"github.com/herryg91/go-clean-architecture/examples/video-rest-api/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type repository struct {
	db *gorm.DB
	ds *video_datasource.MysqlDatasource
}

func NewRepository(db *gorm.DB, ds *video_datasource.MysqlDatasource) Repository {
	return &repository{db, ds}
}

func (r *repository) Get(id int) (*entity.Video, error) {
	out, err := r.ds.Get(nil, id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, ErrRecordNotFound
		}
		return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}

	return out.ToVideoEntity(), err
}

func (r *repository) GetAll() ([]*entity.Video, error) {
	datas, err := r.ds.GetAll(nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}
	out := []*entity.Video{}
	for _, data := range datas {
		out = append(out, data.ToVideoEntity())
	}
	return out, err
}

func (r *repository) Create(in entity.Video) (*entity.Video, error) {
	out, err := r.ds.Create(nil, *video_datasource.VideoModel{}.FromVideoEntity(in))
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}
	return out.ToVideoEntity(), err
}

func (r *repository) Update(in entity.Video) (*entity.Video, error) {
	out, err := r.ds.Update(nil, *video_datasource.VideoModel{}.FromVideoEntity(in))
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}
	return out.ToVideoEntity(), err
}

func (r *repository) Delete(id int) error {
	err := r.ds.Delete(nil, id)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}
	return err
}
