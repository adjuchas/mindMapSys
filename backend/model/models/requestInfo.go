package models

type RequestInfo struct {
	AccessToken string `json:"accessToken"`
}

type RequestGetDraftInfo struct {
	Id string `json:"Id"`
}

type RequestSetDraft struct {
	StuId string `json:"StuId"`
}

type RequestGetDot struct {
	Id    string `json:"Id"`
	DotId string `json:"dotId"`
}

type RequestSetNoteState struct {
	NoteId string `json:"noteId"`
	State  string `json:"state"`
}

type RequestStartNote struct {
	StuId  string `json:"StuId"`
	NoteId string `json:"noteId"`
}
