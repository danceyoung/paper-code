

---

## Rest ful API

### 创建活动

**简要描述：**

- 举办方创建活动

**请求URL：**

- `/events/new`

**请求方式：**

- POST

**参数：**


| 参数名 | 必选 | 类型 | 说明 |
| :- | :- | :- | - |
| name | 是 | string | 活动名称 |
| start_date | 是 | string | 开始日期 |
| expired_on | 是 | string | 结束日期 |
| member_count_limint | 是 | int | 报名人数上限 |
| address | 是 | string | 活动地址 |
| desc | 是 | string | 活动描述 |

**请求示例**

```
{
    "name":"极限编程大赛",
    "start_date":"2020-10-10",
    "expired_on":"2021-01-01",
    "member_count_limit":20,
    "address":"红湖酒店咖啡厅某个角落",
    "desc":"什么是极限编程，什么是敏捷开发，什么是架构设计，什么是设计模式，什么是领域驱动开发"
}
```

**返回示例**

```
  {
    "code": 0,
	"msg":"",
    "data": null
  }
```

### 参加活动

**简要描述：**

- 学生参加活动

**请求URL：**

- `/event/:id/join`

**请求方式：**

- POST

**参数：**


| 参数名 | 必选 | 类型 | 说明 |
| :- | :- | :- | - |
| name | 是 | string | 姓名 |
| g_m | 是 | string | 性别 |
| student_id | 是 | string | 学号 |
| college | 是 | string | 学院名称 |
| level | 是 | string | 年级 |
| profession | 是 | string | 专业 |

**请求示例**

```
{
    "name":"Young Ding",
    "g_m":"g",
    "student_id":"20042222555",
    "college":"information college",
    "level":"2",
	"profession":"info manage"
}
```

**返回示例**

```
{
"code": 0,
"msg":"",
"data": null
}
```

### 活动列表



**简要描述：**

- 获取活动列表

**请求URL：**

- `/events?student-id=111`

**请求方式：**

- GET

**参数：**


| 参数名 | 必选 | 类型 | 说明 |
| :- | :- | :- | - |
| student_id | 是 | string | 学号 |

**返回示例**

```
{
    "code":0,
    "msg":"",
    "data":[
        {
            "id":1,
            "name":"极限编程大赛",
            "start_date":"2020-10-10",
            "expired_on":"2021-01-01",
            "member_count_limit":20,
            "address":"红湖酒店咖啡厅某个角落",
            "desc":"什么是极限编程，什么是敏捷开发，什么是架构设计，什么是设计模式，什么是领域驱动开发"
        }
    ]
}
```

欢迎使用ShowDoc！
