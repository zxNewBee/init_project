package dataLib

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	Age       int64
	Grade     string
}

func CreateStudentInfo(db *gorm.DB) {
	db.AutoMigrate(&Student{})

	// 检查是否已有数据
	var count int64
	db.Model(&Student{}).Count(&count)
	if count > 0 {
		fmt.Printf("数据库中已有 %d 个学生记录，跳过创建\n", count)
		return
	}

	// 创建测试学生数据
	students := []Student{
		{Name: "张三", Age: 20, Grade: "大一"},
		{Name: "李四", Age: 21, Grade: "大二"},
		{Name: "王五", Age: 22, Grade: "大三"},
	}

	result := db.Create(&students)

	if result.Error != nil {
		fmt.Printf("创建学生记录时出错: %v\n", result.Error)
		panic(result.Error)
	}

	fmt.Printf("成功创建 %d 个学生记录\n", len(students))
	for i, student := range students {
		fmt.Printf("学生 %d: ID=%d, Name=%s, Age=%d, Grade=%s\n",
			i+1, student.ID, student.Name, student.Age, student.Grade)
	}
}

// QueryStudentsOver18 查询年龄大于18岁的学生
func QueryStudentsOver18(db *gorm.DB) {
	var students []Student

	// 查询年龄大于18岁的学生
	result := db.Where("age > ?", 18).Find(&students)
	if result.Error != nil {
		fmt.Printf("查询出错: %v\n", result.Error)
		return
	}

	fmt.Printf("\n年龄大于18岁的学生共有 %d 人:\n", len(students))
	for i, student := range students {
		fmt.Printf("学生 %d: ID=%d, Name=%s, Age=%d, Grade=%s\n",
			i+1, student.ID, student.Name, student.Age, student.Grade)
	}
}
