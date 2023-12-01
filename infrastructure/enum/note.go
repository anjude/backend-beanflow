package enum

type NotePublic int8

const (
	PrivateNote NotePublic = iota
	PublicNote
)
