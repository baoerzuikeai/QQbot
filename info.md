### Guild

| 字段名       | 类型    | 描述               | 
| ------------ | ------- | ------------------ | 
| ID           | string  | 频道 ID            | 
| Name         | string  | 频道名称           | 
| Icon         | string  | 频道头像           | 
| OwnerID      | string  | 拥有者ID          |  
| IsOwner      | bool | 当前人是否是创建人 |    
| MemberCount  | int  | 成员数             |    
| MaxMembers   | int64  | 最大成员数         |  
| Desc         | string  | 描述               | 
| JoinedAt     | Timestamp | 当前用户加入群的时间，此字段只在GUILD_CREATE事件中使用|   

```json
{
    "id":"12148371228066669507",
    "name":"机器人测试",
    "owner_id":"17167000634799849449",
    "owner":false,
    "member_count":3,
    "max_members":5000000,
    "description":""
}
```

```json
{
    "id": "661006928",
    "guild_id": "12148371228066669507",
    "name": "测试",
    "type": 0,
    "position": 1,
    "parent_id": "660748018",
    "owner_id": "17167000634799849449",
    "op_user_id": null,
    "sub_type": 0,
    "private_type": 0,
    "private_user_ids": [],
    "speak_permission": 1,
    "application_id": "0",
    "permissions": "3"
}
```