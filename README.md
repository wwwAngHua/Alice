# 爱丽丝 Alice

<img src="https://cdn.jsdelivr.net/gh/wwwanghua/Alice@main/image/alice.png" style="width: 128px">

这是一个高度可定制化的<span style="color: red;">机器人</span>，您可以通过编码实现任意功能！配合移动端产品（准备中<span style="color: green;">Beta</span>），打造专属于您的定制化服务；您可以使她接入强大的AI处理模型，完成各类任务、特定需求任务的数据处理、24h自动售卖服务、智能客服、在线充值服务等等……

### 其它编程语言实现规范
#### Chat Request Body
```json
{
    "uid": "123",
    "version": 1.0,
    "message": "/start"
}
```
#### Chat Response Body
```json
{
    "code": 200,
    "data": null,
    "message": "<img src=\"https://db.t1y.net/alice.jpg\"><br/>我叫<span style=\"color: yellow;\">爱丽丝</span>。是您的私人助手，我会尽力帮助您完成各种任务。请问有什么我可以帮助您的吗？<a href=\"https://github.com/wwwAngHua/Alice\"><br/><br/>进一步了解我们！</a>"
}
```
#### Upload Request Body
```text
form-data -> file
```
#### Upload Response Body
```json
{
    "code": 200,
    "data": null,
    "message": "上传成功啦！"
}
```