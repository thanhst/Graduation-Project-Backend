package studentclassdao

import (
	model "server/internal/models"

	"gorm.io/gorm"
)

type studentClassDAOImpl struct {
	db *gorm.DB
}

func NewStudentClassDAO(db *gorm.DB) StudentClassDAO {
	return &studentClassDAOImpl{db: db}
}

func (dao *studentClassDAOImpl) Get(userID, classID string) (*model.StudentClass, error) {
	var sc model.StudentClass
	err := dao.db.
		Preload("User").
		Preload("Classroom").
		Where("user_id = ? AND class_id = ?", userID, classID).
		First(&sc).Error
	return &sc, err
}

func (dao *studentClassDAOImpl) ListByUser(userID string) ([]model.StudentClass, error) {
	var scs []model.StudentClass
	err := dao.db.
		Preload("Classroom").
		Where("user_id = ?", userID).
		Find(&scs).Error
	return scs, err
}

func (dao *studentClassDAOImpl) ListByClass(classID string) ([]model.StudentClass, error) {
	var scs []model.StudentClass
	err := dao.db.
		Preload("User").
		Where("class_id = ?", classID).
		Find(&scs).Error
	return scs, err
}

func (dao *studentClassDAOImpl) JoinClass(sc *model.StudentClass) error {
	return dao.db.Create(sc).Error
}

func (dao *studentClassDAOImpl) Update(sc *model.StudentClass) error {
	return dao.db.Save(sc).Error
}

func (dao *studentClassDAOImpl) Delete(userID, classID string) error {
	return dao.db.Where("user_id = ? AND class_id = ?", userID, classID).
		Delete(&model.StudentClass{}).Error
}
func (dao *studentClassDAOImpl) GetAllClassroomsByUser(userId string) ([]*model.Classroom, error) {
	var studentClasses []*model.StudentClass
	err := dao.db.
		Where("user_id = ?", userId).
		Preload("Classroom").
		Order("created_at DESC").
		Find(&studentClasses).Error

	if err != nil {
		return nil, err
	}

	classrooms := make([]*model.Classroom, 0, len(studentClasses))
	for _, sc := range studentClasses {
		classrooms = append(classrooms, &sc.Classroom)
	}
	return classrooms, nil
}
func (dao *studentClassDAOImpl) GetClassroomsJoinedByUser(userId string) ([]*model.Classroom, error) {
	var studentClasses []*model.StudentClass
	err := dao.db.
		Where("user_id = ? and state = ?", userId, "joined").
		Preload("Classroom").
		Order("created_at DESC").
		Find(&studentClasses).Error

	if err != nil {
		return nil, err
	}

	classrooms := make([]*model.Classroom, 0, len(studentClasses))
	for _, sc := range studentClasses {
		classrooms = append(classrooms, &sc.Classroom)
	}
	return classrooms, nil
}
func (dao *studentClassDAOImpl) GetClassroomsWaitingByUser(userId string) ([]*model.Classroom, error) {
	var studentClasses []*model.StudentClass
	err := dao.db.
		Where("user_id = ? and state = ?", userId, "waiting").
		Preload("Classroom").
		Order("created_at DESC").
		Find(&studentClasses).Error
	if err != nil {
		return nil, err
	}

	classrooms := make([]*model.Classroom, 0, len(studentClasses))
	for _, sc := range studentClasses {
		classrooms = append(classrooms, &sc.Classroom)
	}
	return classrooms, nil
}

func (dao *studentClassDAOImpl) GetAllStudentWaitingWithClassroom(classId string) ([]*model.User, error) {
	var studentClasses []*model.StudentClass
	err := dao.db.
		Where("class_id = ? and state = ?", classId, "waiting").
		Preload("User").
		Order("created_at DESC").
		Find(&studentClasses).Error
	if err != nil {
		return nil, err
	}
	users := make([]*model.User, 0, len(studentClasses))
	for _, sc := range studentClasses {
		users = append(users, &sc.User)
	}
	return users, nil
}
func (dao *studentClassDAOImpl) GetClassroomsByUser(userId string, limit int, offset int) ([]*model.Classroom, error) {
	var studentClasses []*model.StudentClass
	err := dao.db.
		Where("user_id = ? and state = ?", userId, "joined").
		Preload("User").
		Preload("Classroom").Preload("Classroom.StudentClasses.User").
		Limit(limit).Offset(offset).
		Order("created_at DESC").
		Find(&studentClasses).Error
	if err != nil {
		return nil, err
	}
	classrooms := make([]*model.Classroom, 0, len(studentClasses))
	for _, sc := range studentClasses {
		classrooms = append(classrooms, &sc.Classroom)
	}
	return classrooms, nil
}
func (dao *studentClassDAOImpl) GetCountClassroomsByUser(userId string) (int64, error) {
	var count int64
	err := dao.db.Model(&model.StudentClass{}).Where("user_id = ? and state = ?", userId, "joined").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
