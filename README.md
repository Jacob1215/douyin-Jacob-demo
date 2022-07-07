# tiktok-Jacob
基于 gRPC微服务 + Gin HTTP 服务完成的第三届字节跳动青训营-极简抖音后端项目

# 微服务模块测试
## 一、user服务：8081
1.User-register: successed,密码过长不会成功。

2.User-login:successed,错误的密码登录失败，jwt还没过期。

3.User-Info:Jwt没有出问题，注意，测试的时候用的是用户6的ID和token。那么其他服务的验证token都还是要加上
懂了，token是在router那儿验的，这时已经能够获取用户id了。
## 二、publish服务：8084-服务能正常起，逻辑还要改一下。
1、publish_list要去修改，还要再去关联一下表，这个晚上来做。

2、