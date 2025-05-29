package classroomdao

import (
	model "server/internal/models"

	"gorm.io/gorm"
)

type classroomDAOImpl struct {
	db *gorm.DB
}

func NewClassroomDAO(db *gorm.DB) ClassroomDAO {
	return &classroomDAOImpl{db: db}
}

func (dao *classroomDAOImpl) GetByID(classID string) (*model.Classroom, error) {
	var classroom model.Classroom
	err := dao.db.Preload("Teacher").First(&classroom, "class_id = ?", classID).Error
	return &classroom, err
}

func (dao *classroomDAOImpl) GetAll() ([]*model.Classroom, error) {
	var classrooms []*model.Classroom
	err := dao.db.Preload("Teacher").Find(&classrooms).Error
	return classrooms, err
}

func (dao *classroomDAOImpl) GetByTeacherID(teacherID string, limit int, offset int) ([]*model.Classroom, error) {
	var classrooms []*model.Classroom
	err := dao.db.
		Where("user_created = ?", teacherID).
		Preload("User").
		Preload("StudentClasses").Preload("StudentClasses.User").
		Limit(limit).
		Offset(offset).
		Find(&classrooms).Error
	return classrooms, err
}

func (dao *classroomDAOImpl) Create(classroom *model.Classroom) error {
	return dao.db.Create(classroom).Error
}

func (dao *classroomDAOImpl) Update(classroom *model.Classroom) error {
	return dao.db.Save(classroom).Error
}

func (dao *classroomDAOImpl) Delete(classID string) error {
	return dao.db.Delete(&model.Classroom{}, "class_id = ?", classID).Error
}

func (dao *classroomDAOImpl) GetCountClassroomsByUser(userId string) (int64, error) {
	var count int64
	err := dao.db.Model(&model.Classroom{}).Where("user_created = ?", userId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (dao *classroomDAOImpl) GetClassroomsWithNewScheduler(userId string) ([]*model.Classroom, error) {
	var classrooms []*model.Classroom

	query := `
    SELECT DISTINCT c.*
	FROM classrooms c
	LEFT JOIN (
		SELECT s1.class_id, MIN(s1.start_time) AS nearest_start
		FROM schedulers s1
		WHERE DATE(s1.start_time) = CURDATE()  -- chỉ lấy lịch học trong hôm nay
		GROUP BY s1.class_id
	) smin ON smin.class_id = c.class_id
	WHERE c.user_created = ?
	AND smin.nearest_start IS NOT NULL;
    `

	err := dao.db.Raw(query, userId).Scan(&classrooms).Error
	if err != nil {
		return nil, err
	}

	return classrooms, nil
}

func (dao *classroomDAOImpl) GetTeacherFromClass(classId string) (*model.User, error) {
	var classroom model.Classroom
	err := dao.db.Where("class_id = ?", classId).Preload("User").Find(&classroom).Error
	if err != nil {
		return nil, err
	}
	return &classroom.User, nil
}
