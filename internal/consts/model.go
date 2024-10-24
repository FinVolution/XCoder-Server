package consts

type ConversationType string

const (
	ConversationUT       ConversationType = "ut"
	ConversationChat     ConversationType = "chat"
	ConversationEdit     ConversationType = "edit"
	ConversationExplain  ConversationType = "explain"
	ConversationComment  ConversationType = "comment"
	ConversationOptimize ConversationType = "optimize"
)

func (s ConversationType) String() string {
	return string(s)
}

const (
	ProjectName    = "XCoder"
	ProjectVersion = "1.0.0"
)

type ContextFileType string

const (
	ContextUrl           ContextFileType = "url"
	ContextFile          ContextFileType = "file"
	ContextFileLocal     ContextFileType = "file@local"
	ContextFileLocalTest ContextFileType = "file@localTest"
	ContextFileOtherTest ContextFileType = "file@otherTest"
)

func (s ContextFileType) String() string {
	return string(s)
}
