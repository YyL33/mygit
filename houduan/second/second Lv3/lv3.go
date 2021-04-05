package main

type Video struct {
	Name      string  //视频名
	Author    string  //视频作者
	Icon      Author  //作者头像
	Click     int     //点击量
	Recommend []Video //推荐视频
	Givelikes int     //点赞
	Collects  int     //收藏
	Coins     int     //硬币
}
type Author struct {
	Name string //昵称
	VIP  bool   //是否大会员

	Foucs string //关注人数
}

func (v *Video) Givelike() {
	v.Givelikes++
}
func (v *Video) Collect() {
	v.Collects++
}
func (v *Video) Coin() {
	v.Coins++
}
func (v *Video) Sanlian() {
	v.Givelike()
	v.Collect()
	v.Coin()
}
func publish(upname string, videoname string) Video {
	return Video{
		Author: upname,
		Name:    videoname,
	}
}
