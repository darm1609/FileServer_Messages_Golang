package msj

type Messages struct {
}

func (m *Messages) Message(required string) string {
	if required == "HEAD_Unsoported" {
		return "unsupported network protocol:"
	}
	if required == "HEAD_failed_create_listener" {
		return "failed to create listener:"
	}
	if required == "HEAD_failed_close_listener" {
		return "failed to close listener:"
	}
	if required == "HEAD_FileServer" {
		return "**** FileServer ****"
	}
	if required == "HEAD_ServiceStart" {
		return "Service started:"
	}
	if required == "HEAD_Connected_to" {
		return "Connected to"
	}
	if required == "CLIENT_Missing_Parameters" {
		return "Missing parameters"
	}
	if required == "CLIENT_Invalid_Syntax" {
		return "Invalid syntax"
	}
	if required == "CLIENT_Must_Subscribe" {
		return "You must subscribe to a channel"
	}
	if required == "CLIENT_No_Client_On_Channel" {
		return "There are no active clients on the channel in receive mode"
	}
	if required == "CLIENT_Subscribe_To_Channel" {
		return "Subscribe to the channel"
	}
	if required == "CLIENT_Channel_Not_Exist" {
		return "Channel does not exist"
	}
	if required == "CLIENT_Client_In_Mode" {
		return "Client in mode "
	}
	if required == "CLIENT_Invalid_Mode" {
		return "Invalid mode"
	}
	if required == "CLIENT_Must_Subscribe_Before_Mode" {
		return "You must subscribe to a channel before setting the mode"
	}
	if required == "CLIENT_Server_Failed_To_Receive" {
		return "Server failed to receive file"
	}
	if required == "CLIENT_The_File_Was_Send" {
		return "The file was sent to the clients in the channel"
	}
	if required == "CLIENT_Fail_To_Send_File" {
		return "Failed to send file to clients in channel"
	}
	if required == "CLIENT_Goodbye" {
		return "Goodbye"
	}
	if required == "CLIENT_Must_Set_Mode_Send" {
		return "You must set the mode to "
	}
	if required == "CLIENT_Error_Writing" {
		return "error writing: "
	}
	if required == "GENERAL_Send" {
		return "send"
	}
	if required == "GENERAL_Receive" {
		return "receive"
	}
	return ""

}
