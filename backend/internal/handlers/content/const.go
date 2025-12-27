package content

const (
	ListThreads   = "content.HandleThread"
	CreateThreads = "content.CreateThread"
	ListPosts     = "content.HandlePost"
	ListComments  = "content.HandleComment"

	SuccessfulListThreadsMessage   = "Successfully Listed Threads"
	SuccessfulCreateThreadsMessage = "Successfully Created Thread"
	SuccessfulListPostsMessage     = "Successfully Listed Posts"
	SuccessfulListCommentsMessage  = "Successfully Listed Comments"

	ErrRetrieveDatabase = "Failed to Retrieve Database in %s"
	ErrRetrieveThreads  = "Failed to Retrieve Threads in %s"
	ErrCreateThreads    = "Failed to Create Thread in %s"
	ErrRetrievePosts    = "Failed to Retrieve Posts in %s"
	ErrRetrieveComments = "Failed to Retrieve Comments in %s"

	ErrEncodeView = "Failed to Retrieve in %s"
)
