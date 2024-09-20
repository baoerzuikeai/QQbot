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


```json
{
    "guild_id":"AD529F50A408B63D98B4B1094F3EDCD2"
}

```


```json

{
    "op": 0,
    "s": 2,
    "t": "GROUP_AT_MESSAGE_CREATE",
    "id": "GROUP_AT_MESSAGE_CREATE:bysicoipxy6vn57zlepzoqxtfrgpjmnyqxmsinxqkgxwpynhysl6q6ws849",
    "d": {
        "author": {
            "id": "B71229F2171C9566660DDC6E9BA0D920",
            "member_openid": "B71229F2171C9566660DDC6E9BA0D920"
        },
        "content": " ",
        "group_id": "AD529F50A408B63D98B4B1094F3EDCD2",
        "group_openid": "AD529F50A408B63D98B4B1094F3EDCD2",
        "id": "ROBOT1.0_BYSicoiPXY6vn57zlepZoreV7AkmYLM6ifqzjIWZ1SpjYTfjTwc9l2Hk0o0Rfy4mU.E7B1aAyR.3zI2POwazactG81ovPjw88HwjHppK6Gc!",
        "timestamp": "2024-09-13T15:28:34+08:00"
    }
}
```