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

func (dao *studentClassDAOImpl) Delete(sc *model.StudentClass) error {
	return dao.db.Where("user_id = ? AND class_id = ?", sc.UserId, sc.ClassId).
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
func (dao *studentClassDAOImpl) GetClassroomsWithNewScheduler(userId string) ([]*model.Classroom, error) {
	var classrooms []*model.Classroom

	query := `
	SELECT DISTINCT c.*
	FROM classrooms c
	JOIN student_classes sc ON sc.class_id = c.class_id
	JOIN (
		SELECT class_id, MIN(start_time) AS nearest_start
		FROM schedulers
		WHERE DATE(start_time) = CURDATE()  -- chỉ lấy buổi học trong hôm nay
		GROUP BY class_id
	) smin ON smin.class_id = c.class_id
	JOIN schedulers s ON s.class_id = smin.class_id AND s.start_time = smin.nearest_start
	WHERE sc.user_id = ?
	ORDER BY s.start_time ASC;
    `

	if err := dao.db.Raw(query, userId).Scan(&classrooms).Error; err != nil {
		return nil, err
	}
	return classrooms, nil
}

func (dao *studentClassDAOImpl) GetUserJoinedWithClassrooms(classId string, limit int, offset int) ([]*model.User, error) {
	var studentClasses []*model.StudentClass
	err := dao.db.
		Where("class_id = ? and state = ?", classId, "joined").
		Preload("User").
		Limit(limit).Offset(offset).
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
func (dao *studentClassDAOImpl) GetUserWaitingWithClassrooms(classId string, limit int, offset int) ([]*model.User, error) {
	var studentClasses []*model.StudentClass
	err := dao.db.
		Where("class_id = ? and state = ?", classId, "waiting").
		Preload("User").
		Limit(limit).Offset(offset).
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

func (dao *studentClassDAOImpl) GetCountUsersByClassroom(userId string) (int64, int64, error) {
	var countJoined, countWaiting int64
	err := dao.db.Model(&model.StudentClass{}).Where("class_id = ? and state = ?", userId, "joined").Count(&countJoined).Error
	errW := dao.db.Model(&model.StudentClass{}).Where("class_id = ? and state = ?", userId, "waiting").Count(&countWaiting).Error

	if err != nil {
		return 0, 0, err
	}
	if errW != nil {
		return 0, 0, err
	}
	return countJoined, countWaiting, nil
}
func (dao *studentClassDAOImpl) GetInfo(classId string, userId string) (*model.StudentClass, error) {
	var stdCls model.StudentClass
	err := dao.db.Where("class_id = ? and user_id = ?", classId, userId).
		Preload("User").
		Preload("Classroom").Find(&stdCls).Error
	if err != nil {
		return nil, err
	}
	return &stdCls, err
}
