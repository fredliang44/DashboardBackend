package user

import "time"

// User is the basic structure of user
type User struct {
	Username string `name:"用户名"`
	Name     string `name:"真实姓名"`
	Email    string `name:"Email地址"`
	Password string `name:"密码" password:"yes"`
	Detail   Detail
}

// Detail is addtional info user may add later
type Detail struct {
	Gentle                   uint      `name:"性别" option:"男,女" placeholder:"请输入你的性别"`
	Birthday                 time.Time `name:"生日" time:"yes"`
	Phone                    string    `name:"手机号码"`
	TshirtSize               uint      `name:"T-Shirt尺码" option:"[XS,S,M,L,XL,XXL,XXXL]"`
	City                     string    `name:"所在城市" placeholder:"请输入你所在的城市, 方便我们报销"`
	AlipayAccount            string    `name:"支付宝账户" placeholder:"请输入你的支付宝账号, 方便我们报销"`
	School                   string    `name:"学校名称"`
	College                  string    `name:"学院名称"`
	Profession               string    `name:"专业名称"`
	Grade                    string    `name:"年级"`
	GraduateTime             time.Time `name:"预计毕业时间" time:"yes"`
	SpecialNeed              string    `name:"特殊需要" placeholder:"如过敏源、饮食、出行方面的特殊需要"`
	UrgentConcatName         string    `name:"紧急联系人"`
	UrgentConcatPhone        string    `name:"紧急联系人电话"`
	UrgentConcatRelationship string    `name:"与紧急联系人的关系" placeholder:"如父亲、母亲、女友"`
	Github                   string    `name:"Github 账号"`
	LinkedIn                 string    `name:"LinkedIn 账号"`
	CodingDotNet             string    `name:"Coding.net 账号" placeholder:"会有Coding.net大奖送出哦！"`
	PersonalBlog             string    `name:"个人博客"`
	OtherAccount             string    `name:"其他账户" placeholder:"例如：微信：uniqu"`
}

func main() {

}

// Handler is the handler for user endpoint
func Handler() {

}
