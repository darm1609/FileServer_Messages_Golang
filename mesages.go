package FileServer_Messages_Golang

type Messages struct {
}

func (m *Messages) Message(required string) string {

	if required == "MAX_FILESIZE" {
		return "512000"
	}

	if required == "HOST_Error_Convert_Max_FileSize" {
		return "error converting max file size"
	}
	if required == "HOST_HEAD_Unsoported" {
		return "unsupported network protocol:"
	}
	if required == "HOST_HEAD_failed_create_listener" {
		return "failed to create listener:"
	}
	if required == "HOST_HEAD_failed_close_listener" {
		return "failed to close listener:"
	}
	if required == "HOST_HEAD_FileServer" {
		return "**** FileServer ****"
	}
	if required == "HOST_HEAD_ServiceStart" {
		return "Service started:"
	}
	if required == "HOST_HEAD_Connected_to" {
		return "Connected to"
	}
	if required == "HOST_HEAD_Close" {
		return "q <close connection>"
	}
	if required == "HOST_CLIENT_Missing_Parameters" {
		return "Missing parameters"
	}
	if required == "HOST_CLIENT_Invalid_Syntax" {
		return "Invalid syntax"
	}
	if required == "HOST_CLIENT_Must_Subscribe" {
		return "You must subscribe to a channel"
	}
	if required == "HOST_CLIENT_No_Client_On_Channel" {
		return "There are no active clients on the channel in receive mode"
	}
	if required == "HOST_CLIENT_Subscribe_To_Channel" {
		return "Subscribe to the channel"
	}
	if required == "HOST_CLIENT_Channel_Not_Exist" {
		return "Channel does not exist"
	}
	if required == "HOST_CLIENT_Client_In_Mode" {
		return "Client in mode "
	}
	if required == "HOST_CLIENT_Invalid_Mode" {
		return "Invalid mode"
	}
	if required == "HOST_CLIENT_Must_Subscribe_Before_Mode" {
		return "You must subscribe to a channel before setting the mode"
	}
	if required == "HOST_CLIENT_Server_Failed_To_Receive" {
		return "Server failed to receive file"
	}
	if required == "HOST_CLIENT_The_File_Was_Send" {
		return "The file was sent to the clients in the channel"
	}
	if required == "HOST_CLIENT_Fail_To_Send_File" {
		return "Failed to send file to clients in channel"
	}
	if required == "HOST_CLIENT_Goodbye" {
		return "Goodbye"
	}
	if required == "HOST_CLIENT_Must_Set_Mode_Send" {
		return "You must set the mode to "
	}
	if required == "HOST_CLIENT_Error_Writing" {
		return "error writing: "
	}
	if required == "HOST_GENERAL_Send" {
		return "send"
	}
	if required == "HOST_GENERAL_Receive" {
		return "receive"
	}

	if required == "GUEST_GENERAL_Mjs_Form_Server" {
		return "Message from server: "
	}
	if required == "GUEST_GENERAL_Error_In_Receive_Mode" {
		return "Error connecting in receive mode"
	}
	if required == "GUEST_GENERAL_File_Is_Received" {
		return "The file is received: "
	}
	if required == "GUEST_GENERAL_Text_To_Send" {
		return "Text to send: "
	}
	if required == "GUEST_Error_Max_File_Size" {
		return "The file exceeds the size of: "
	}
	return ""

}
