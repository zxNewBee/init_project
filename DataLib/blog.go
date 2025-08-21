package dataLib

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BlogUser struct {
	ID        int            `gorm:"primarykey"`        // 用户ID
	Name      string         `json:"name"`              // 用户名
	Email     string         `json:"email"`             // 用户邮箱
	Password  string         `json:"password"`          // 用户密码
	PostCount int            `gorm:"default:0"`         // 文章数量统计
	CreatedAt time.Time      `json:"created_at"`        // 创建时间
	UpdatedAt time.Time      `json:"updated_at"`        // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index"`             // 删除时间
	Posts     []Post         `gorm:"foreignKey:UserID"` // 用户的帖子
}

type Post struct {
	ID            int            `gorm:"primarykey"`        // 帖子ID
	UserID        int            `gorm:"not null"`          // 用户ID外键
	Title         string         `json:"title"`             // 帖子标题
	Content       string         `json:"content"`           // 帖子内容
	CommentStatus string         `gorm:"default:'有评论'"`     // 评论状态
	CreatedAt     time.Time      `json:"created_at"`        // 创建时间
	UpdatedAt     time.Time      `json:"updated_at"`        // 更新时间
	DeletedAt     gorm.DeletedAt `gorm:"index"`             // 删除时间
	Comments      []Comment      `gorm:"foreignKey:PostID"` // 帖子的评论
}
type Comment struct {
	ID        int            `gorm:"primarykey"` // 评论ID
	PostID    int            `gorm:"not null"`   // 帖子ID外键
	Content   string         `json:"content"`    // 评论内容
	CreatedAt time.Time      `json:"created_at"` // 创建时间
	UpdatedAt time.Time      `json:"updated_at"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index"`      // 删除时间
}

func CreateBlogTable(db *gorm.DB) {
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&BlogUser{})
	//CreateBlogData(db)
	QueryData(db)

	// 测试钩子函数
	TestHooks(db)
}

// 创建示例数据
func CreateBlogData(db *gorm.DB) {
	// 1. 先创建用户
	user := BlogUser{
		Name:     "张三",
		Email:    "zhangsan@example.com",
		Password: "123456",
	}
	db.Create(&user)
	fmt.Printf("创建用户: ID=%d, Name=%s\n", user.ID, user.Name)

	// 2. 再创建帖子
	post := Post{
		UserID:  user.ID, // 使用刚创建的用户ID
		Title:   "我的第一篇博客",
		Content: "这是博客内容...",
	}
	db.Create(&post)
	fmt.Printf("创建帖子: ID=%d, Title=%s, UserID=%d\n", post.ID, post.Title, post.UserID)
	post2 := Post{
		UserID:  user.ID, // 使用刚创建的用户ID
		Title:   "我的第二篇博客",
		Content: "这是博客内容...",
	}
	db.Create(&post2)
	fmt.Printf("创建帖子: ID=%d, Title=%s, UserID=%d\n", post2.ID, post2.Title, post2.UserID)

	// 3. 最后创建评论
	comment := Comment{
		PostID:  post.ID, // 使用刚创建的帖子ID
		Content: "很好的博客！",
	}
	comment2 := Comment{
		PostID:  post.ID, // 使用刚创建的帖子ID
		Content: "很差的博客！",
	}
	db.Create(&comment)
	db.Create(&comment2)
	fmt.Printf("创建评论: ID=%d, Content=%s, PostID=%d\n", comment.ID, comment.Content, comment.PostID)
	fmt.Printf("创建评论: ID=%d, Content=%s, PostID=%d\n", comment2.ID, comment2.Content, comment2.PostID)
	comment3 := Comment{
		PostID:  post2.ID, // 使用刚创建的帖子ID
		Content: "第二篇很好的博客！",
	}
	comment4 := Comment{
		PostID:  post2.ID, // 使用刚创建的帖子ID
		Content: "第二篇很差的博客！",
	}
	db.Create(&comment3)
	db.Create(&comment4)
	fmt.Printf("创建评论: ID=%d, Content=%s, PostID=%d\n", comment3.ID, comment3.Content, comment3.PostID)
	fmt.Printf("创建评论: ID=%d, Content=%s, PostID=%d\n", comment4.ID, comment4.Content, comment4.PostID)
}

func QueryData(db *gorm.DB) {
	//查询某个用户所有文章
	var posts []Post
	db.Select("title").Where("user_id=?", 1).Find(&posts)

	for _, post := range posts {
		fmt.Println("title is: ", post.Title)
	}
	//查询评论数最多文章
	queryPostWithMostCommentsGorm(db)
}

// 使用GORM关联查询的方式查询评论数最多的文章
func queryPostWithMostCommentsGorm(db *gorm.DB) {
	var posts []Post

	// 预加载评论关联，然后按评论数量排序
	err := db.Preload("Comments").
		Joins("LEFT JOIN comments ON posts.id = comments.post_id").
		Group("posts.id").
		Order("COUNT(comments.id) DESC").
		Limit(1).
		Find(&posts).Error

	if err != nil {
		fmt.Println("GORM关联查询失败:", err)
		return
	}

	if len(posts) > 0 {
		post := posts[0]
		fmt.Printf("\nGORM关联查询结果:\n")
		fmt.Printf("  文章ID: %d\n", post.ID)
		fmt.Printf("  标题: %s\n", post.Title)
		fmt.Printf("  内容: %s\n", post.Content)
		fmt.Printf("  评论数: %d\n", len(post.Comments))

		// 显示评论内容
		fmt.Printf("  评论内容:\n")
		for i, comment := range post.Comments {
			fmt.Printf("    %d. %s\n", i+1, comment.Content)
		}
	}
}

// 为 Post 模型添加 AfterCreate 钩子函数
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	// 更新用户的文章数量统计
	var user BlogUser
	if err := tx.First(&user, p.UserID).Error; err != nil {
		return err
	}

	// 计算用户的总文章数
	var postCount int64
	if err := tx.Model(&Post{}).Where("user_id = ?", p.UserID).Count(&postCount).Error; err != nil {
		return err
	}

	// 更新用户的文章数量字段
	if err := tx.Model(&user).Update("post_count", postCount).Error; err != nil {
		return err
	}

	fmt.Printf("用户 %s (ID: %d) 的文章数量已更新为: %d\n", user.Name, user.ID, postCount)

	return nil
}

// 为 Comment 模型添加 BeforeDelete 钩子函数
func (c *Comment) BeforeDelete(tx *gorm.DB) (err error) {
	// 获取文章信息
	var post Post
	if err := tx.First(&post, c.PostID).Error; err != nil {
		return err
	}

	// 计算文章的评论数量
	var commentCount int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount).Error; err != nil {
		return err
	}

	// 如果删除这条评论后评论数量为0，则更新文章状态
	if commentCount <= 1 { // 当前评论还在，删除后就是0
		// 更新文章状态为"无评论"
		if err := tx.Model(&post).Update("comment_status", "无评论").Error; err != nil {
			return err
		}
		fmt.Printf("文章 '%s' (ID: %d) 删除评论后状态已更新为: 无评论\n", post.Title, post.ID)
	}

	return nil
}

// 测试钩子函数的演示函数
func TestHooks(db *gorm.DB) {
	fmt.Println("\n=== 测试钩子函数 ===")

	// 1. 测试 Post 创建钩子
	fmt.Println("1. 创建新文章，测试 AfterCreate 钩子...")
	newPost := Post{
		UserID:  1,
		Title:   "测试钩子函数的文章",
		Content: "这篇文章用来测试钩子函数",
	}

	if err := db.Create(&newPost).Error; err != nil {
		fmt.Println("创建文章失败:", err)
		return
	}
	fmt.Printf("文章创建成功，ID: %d\n", newPost.ID)

	// 2. 测试 Comment 删除钩子
	fmt.Println("\n2. 删除评论，测试 BeforeDelete 钩子...")
	var comment Comment
	if err := db.Where("post_id = ?", 17).First(&comment).Error; err != nil {
		fmt.Println("未找到评论，跳过删除测试")
		return
	}

	if err := db.Delete(&comment).Error; err != nil {
		fmt.Println("删除评论失败:", err)
		return
	}
	fmt.Println("评论删除成功")

	// 3. 验证结果
	fmt.Println("\n3. 验证钩子函数效果...")
	var user BlogUser
	db.First(&user, 1)
	fmt.Printf("用户文章数量: %d\n", user.PostCount)

	var post Post
	db.First(&post, newPost.ID)
	fmt.Printf("文章评论状态: %s\n", post.CommentStatus)
}
