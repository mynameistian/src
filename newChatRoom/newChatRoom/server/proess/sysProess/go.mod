module sysProess

go 1.13

replace chatroom.com/message => ../../../message

replace chatroom.com/server/model => ../../model

require (
	chatroom.com/message v0.0.0-00010101000000-000000000000
	chatroom.com/server/model v0.0.0-00010101000000-000000000000
)
