package domain

type GuestFileOpen struct {
	Execute   string `json:"execute"`
	Arguments struct {
		Path string `json:"path"`
		Mode string `json:"mode"`
	} `json:"arguments"`
}

type GuestFileOpenReply struct {
	Return int `json:"return"`
}

type GuestFileWrite struct {
	Execute   string `json:"execute"`
	Arguments struct {
		Handle int    `json:"handle"`
		BufB64 string `json:"buf-b64"`
	} `json:"arguments"`
}

type GuestFileWriteReply struct {
	Return struct {
		Count int  `json:"count"`
		EOF   bool `json:"eof"`
	} `json:"return"`
}

type GuestFileClose struct {
	Execute   string `json:"execute"`
	Arguments struct {
		Handle int `json:"handle"`
	} `json:"arguments"`
}

type GuestFileCloseReply struct {
	Return struct{} `json:"return"`
}
