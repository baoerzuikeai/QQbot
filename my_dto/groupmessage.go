package my_dto

type GroupMessage struct {
	GroupOpenid string      `json:"group_openid"`      // 机器人访问群的openid
	Content     string      `json:"content,omitempty"` // 文本内容
	MsgType     int         `json:"msg_type"`          // 消息类型：0 文本，2 markdown，3 ark 消息，4 embed，7 media 富媒体
	MarkDown    interface{} `json:"markdown"`
	EventID     string      `json:"event_id,omitempty"` // 前置收到的事件ID，用于发送被动消息，支持的事件："INTERACTION_CREATE"、"GROUP_ADD_ROBOT"、"GROUP_MSG_RECEIVE"
	MsgID       string      `json:"msg_id,omitempty"`   // 前置收到的用户发送过来的消息ID，用于发送被动消息（回复）
	MsgSeq      int         `json:"msg_seq,omitempty"`  // 回复消息的序号，与msg_id联合使用，默认1
	Media       interface{} `json:"media,omitempty"`
}

// file_uuid	string	文件 ID
// file_info	string	文件信息，用于发消息接口的 media 字段使用
// ttl	int	有效期，表示剩余多少秒到期，到期后 file_info 失效，当等于 0 时，表示可长期使用
// id	string	发送消息的唯一ID，当srv_send_msg设置为true时返回
type Media struct {
	FileUuid string `json:"file_uuid"`
	FileInfo string `json:"file_info"`
	Ttl      int    `json:"ttl"`
	Id       string `json:"id"`
}

// file_type	int	是	媒体类型：1 图片，2 视频，3 语音，4 文件（暂不开放） 资源格式要求 图片：png/jpg，视频：mp4，语音：silk
// url	string	是	需要发送媒体资源的url
// srv_send_msg	bool	是	设置 true 会直接发送消息到目标端，且会占用主动消息频次
// file_data	string   base64
type PostMedia struct {
	FileType   int    `json:"file_type"`
	Url        string `json:"url,omitempty"`
	SrvSendMsg bool   `json:"srv_send_msg"`
	FileData   string `json:"file_data,omitempty"`
}
