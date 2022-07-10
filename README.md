# tiktok-Jacob
基于 gRPC微服务 + Gin HTTP 服务完成的第三届字节跳动青训营-极简抖音后端项目

# 重构
还是把数据库的连接以及打包给重构一下吧，不知道搞得完不。
感觉这样不好用，因为指不定以后要分库分表什么的。

# 建立自己使用的错误码及描述
要把所有的错误描述都给加进去
# bug，没办法用c.postfrom
# JWT收得到token但是没法验证，离谱。
明白了，bug解决了。Yes,就是jwt会根据token拿到用户id，这样再去对比是否是我们请求的id。

# 微服务模块测试-完全成功
## 一、user服务：8081--卧槽，又改了很多BUG出来
1.User-register: successed,密码过长不会成功。
解决了JWT的问题。

2.User-login:successed,错误的密码登录失败，jwt还没过期。
解决了一些登录的密码验证的问题，就是c.query无法拿到password
必须用c.shouldBind.

3.User-Info:Jwt没有出问题，注意，测试的时候用的是用户6的ID和token。那么其他服务的验证token都还是要加上
懂了，token是在router那儿验的，这时已经能够获取用户id了。
## 二、publish服务：8084-服务能正常起，逻辑还要改一下。
1、publish_action要去修改，还要再去关联一下表，这个晚上来做。
200了，但是还是有问题。感觉要重构一下错误处理。

2、publish_list有bug，查询的是视频id，应该查询的是用户ID

但是现在服务器没什么问题。
感觉现在就完全可以再重构一下了。

## 三、favorite服务：端口8089
1、fav action成功
完全成功，非常简单的接口

2、fav list成功
几种情况都可以
jwt也可认证。

## 四、comment服务：端口
1、comment action 成功
但是有一些地方，比如视频的fav可以更新，relation可以更新。

2、comment list 成功

## relation success
1、action成功

2、follow list 成功

3、follower list 成功


