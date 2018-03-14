package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// User is the basic structure of user
type User struct {
	Username string `name:"用户名" json:"username" binding:"required"`
	Email    string `name:"Email地址" json:"email" binding:"required"`
	Password string `name:"密码" password:"yes" json:"password" binding:"required"`
	// Detail   Detail `json:"detail"`
}

// Detail is addtional info user may add later
type Detail struct {
	Name                     string    `name:"真实姓名" json:"name" binding:"required"`
	Gentle                   uint      `name:"性别" option:"男,女" placeholder:"请输入你的性别" json:"gentle" binding:"required"`
	Birthday                 time.Time `name:"生日" time:"yes" json:"birthday" binding:"required"`
	Phone                    string    `name:"手机号码" json:"phone" binding:"required"`
	TshirtSize               uint      `name:"T-Shirt尺码" option:"[XS,S,M,L,XL,XXL,XXXL]" json:"t_shirt_size" binding:"required"`
	City                     string    `name:"所在城市" placeholder:"请输入你所在的城市, 方便我们报销" json:"city" binding:"required"`
	Alipay                   string    `name:"支付宝账户" placeholder:"请输入你的支付宝账号, 方便我们报销" json:"alipay" binding:"required"`
	School                   string    `name:"学校名称" json:"school" binding:"required"`
	College                  string    `name:"学院名称" json:"college" binding:"required"`
	Profession               string    `name:"专业名称" json:"profession" binding:"required"`
	Grade                    string    `name:"年级" json:"grade" binding:"required"`
	GraduateTime             time.Time `name:"预计毕业时间" time:"yes" json:"graduate_time" binding:"required"`
	SpecialNeeds             string    `name:"特殊需要" placeholder:"如过敏源、饮食、出行方面的特殊需要" json:"special_needs"`
	UrgentConcatName         string    `name:"紧急联系人" json:"urgent_concat_name" binding:"required"`
	UrgentConcatPhone        string    `name:"紧急联系人电话" json:"urgent_concat_phone" binding:"required"`
	UrgentConcatRelationship string    `name:"与紧急联系人的关系" placeholder:"如父亲、母亲、女友" json:"urgent_concat_relationship" binding:"required"`
	Github                   string    `name:"Github 账号" json:"github"`
	LinkedIn                 string    `name:"LinkedIn 账号" json:"linkedin"`
	CodingDotNet             string    `name:"Coding.net 账号" placeholder:"会有Coding.net大奖送出哦！" json:"codeing"`
	Blog                     string    `name:"个人博客" json:"blog"`
	OtherAccount             string    `name:"其他账户" placeholder:"例如：微信：uniqu" json:"other_account"`
}

func main() {

}

// RegHandler is the handler for user reg endpoint
func RegHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err == nil {
		res, _ := json.Marshal(user)
		fmt.Println(string(res))
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// DetailHandler is a handler
func DetailHandler(c *gin.Context) {

}
