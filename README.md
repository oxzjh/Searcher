# Searcher

## 接口
##### 1. 搜索引擎搜索
- 说明：搜索引擎搜索
- 路由地址：/search
- 请求方式：POST
- Content-Type: application/json
- 参数：

|名称|类型|默认值|描述|
|-|-|-|-|
|q|string||查询关键字|
|news|bool|false|是否查询新闻|
|engine|string|bing|搜索引擎（baidu, bing)|

##### 2. 提取网页内容
- 说明：提取网页内容
- 路由地址：/search/sanitize
- 请求方式：POST
- Content-Type: application/json
- 参数：

|名称|类型|默认值|描述|
|-|-|-|-|
|url|string||网页地址|

## 测试
![image](https://github.com/user-attachments/assets/a27bd57f-fe90-43b5-80ad-ac84212ccb2f)
![image](https://github.com/user-attachments/assets/903cd88c-293b-4baa-a62b-1e435bf5fa1e)
