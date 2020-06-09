module newChatRoom

go 1.13

require (
	chatroom.com/message v0.0.0-00010101000000-000000000000
	chatroom.com/server/model v0.0.0-00010101000000-000000000000
	chatroom.com/server/proess/chatProess v0.0.0-00010101000000-000000000000
	chatroom.com/server/proess/sysProess v0.0.0-00010101000000-000000000000
	chatroom.com/utils/pack v0.0.0-00010101000000-000000000000
)

replace chatroom.com/message => ./newChatRoom/message

replace chatroom.com/utils/pack => ./newChatRoom/utils/pack

replace chatroom.com/server/model => ./newChatRoom/server/model

replace chatroom.com/server/proess/sysProess => ./newChatRoom/server/proess/sysProess

replace chatroom.com/server/proess/chatProess => ./newChatRoom/server/proess/chatProess
