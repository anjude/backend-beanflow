package enum

type NotePublic int8

const (
	PrivateNote NotePublic = iota
	PublicNote
)

type NoteLike int8

const (
	DislikeNote NoteLike = iota
	LikeNote
)
