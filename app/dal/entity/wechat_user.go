package entity

type WechatUser struct {
	BaseEntity
	OpenId         string `xorm:'open_id' varchar(45) notnull unqiue`
	NickName       string `xorm:'nick_name' varchar(45) notnull`
	AvatarUrl      string `xorm:'avatar_url' varchar(1000)`
	Gender         string `xorm:'gender' varchar(45) default('0')`
	City           string `xorm:'city' varchar(60)`
	Province       string `xorm:'province varchar(60)'`
	Country        string `xorm:'country' varchar(45)`
	Language       string `xorm:'language' varchar(45) notnull default('zh_CN')`
	IsDel          string `xorm:'is_del' varchar(4) notnull default('NO')`
	CategoryOrders string `xorm:'category_orders' varchar(1000)`
}
