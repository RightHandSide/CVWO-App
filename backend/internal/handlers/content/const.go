package content

const (
	ListThreads    = "content.HandleThread"
	CreateThreads  = "content.CreateThread"
	ListPosts      = "content.HandlePost"
	CreatePosts    = "content.CreatePost"
	ListComments   = "content.HandleComment"
	CreateComments = "content.CreateComment"

	SuccessfulListThreadsMessage    = "Successfully Listed Threads"
	SuccessfulCreateThreadsMessage  = "Successfully Created Thread"
	SuccessfulListPostsMessage      = "Successfully Listed Posts"
	SuccessfulCreatePostsMessage    = "Successfully Created Posts"
	SuccessfulListCommentsMessage   = "Successfully Listed Comments"
	SuccessfulCreateCommentsMessage = "Successfully Created Comments"

	ErrRetrieveDatabase = "Failed to Retrieve Database in %s"
	ErrRetrieveThreads  = "Failed to Retrieve Threads in %s"
	ErrCreateThreads    = "Failed to Create Thread in %s"
	ErrRetrievePosts    = "Failed to Retrieve Posts in %s"
	ErrCreatePosts      = "Failed to Create Post in %s"
	ErrRetrieveComments = "Failed to Retrieve Comments in %s"
	ErrCreateComments   = "Failed to Create Comment in %s"

	ErrEncodeView = "Failed to Retrieve in %s"
)
