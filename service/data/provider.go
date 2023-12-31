package data

type PostIndex struct {
	Index                   map[string]PostHeader
	IndexTime               []string
	IndexDrafts             map[string]int
	IndexDeleted            map[string]int
	IndexTags               map[string]map[int]string
	IndexPopularityLikes    map[int][]string
	IndexPopularityViews    map[int][]string
	IndexPopularityComments map[int][]string
	IndexPopularityTags     map[string]map[string]int
	IndexCountLikes         map[string]int
	IndexCountComments      map[string]int
	IndexCountViews         map[string]int
	IndexUsers              map[string]User
	IndexUsersHandles       map[string]string
	IndexUsersFollowers     map[string]map[string]string
	IndexUsersFollowing     map[string]map[string]string
	IndexUsersPosts         map[string][]int
}

type PostHeader struct {
	Id                string   `json:"id"`
	Title             string   `json:"title"`
	Summary           string   `json:"summary"`
	Image             string   `json:"image"`
	Draft             bool     `json:"draft"`
	Deleted           bool     `json:"deleted"`
	Tags              []string `json:"tags"`
	AuthorId          string   `json:"authorId"`
	AuthorDisplayName string   `json:"authorDisplayName"`
	AuthorProfilePic  string   `json:"authorProfilePic"`
	Author            User     `json:"author"`
	UserFollowing     bool     `json:"userFollowing"`
	Created           string   `json:"created"`
	Index             int      `json:"index"`
	Updated           string   `json:"updated"`
	Upvotes           int      `json:"upvotes"`
	CommentCount      int      `json:"commentCount"`
	FileCount         int      `json:"fileCount"`
}

type PostComment struct {
	Id                string        `json:"id"`
	Created           string        `json:"created"`
	Updated           string        `json:"updated"`
	AuthorId          string        `json:"authorId"`
	AuthorDisplayName string        `json:"authorDisplayName"`
	AuthorProfilePic  string        `json:"authorProfilePic"`
	Content           string        `json:"content"`
	Upvotes           int           `json:"upvotes"`
	Children          []PostComment `json:"children"`
}

type Post struct {
	Header  PostHeader `json:"header"`
	Content string     `json:"content"`
	Files   []string   `json:"files"`
}

type User struct {
	UID             string `json:"uid"`
	Handle          string `json:"handle"`
	HandleSetByUser bool   `json:"handleSetByUser"`
	DisplayName     string `json:"displayName"`
	Email           string `json:"-"`
	ProfileText     string `json:"profileText"`
	JoinDate        string `json:"joinDate"`
	PhotoURL        string `json:"photoURL"`
	ProviderId      string `json:"-"`
	EmailVerified   bool   `json:"-"`
	IsAnonymous     bool   `json:"-"`
	Registered      string `json:"registered"`
	FollowersCount  int    `json:"followers"`
	FollowingCount  int    `json:"following"`
	PostsCount      int    `json:"postCount"`
}

type UserFollow struct {
	UID string `json:"uid"`
}

type UserHandle struct {
	OldHandle string `json:"oldHandle"`
	NewHandle string `json:"newHandle"`
}

type SearchResult struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Count int    `json:"count"`
}

type Metadata struct {
	PostCount    int `json:"postCount"`
	DeletedCount int `json:"deletedCount"`
	DraftCount   int `json:"draftCount"`
	UserCount    int `json:"userCount"`
}

type ImageUploadResult struct {
	Url string `json:"url"`
}

type Provider interface {
	Initialize()
	Finalize(peristMode PersistMode, index PostIndex)
	CreateDir(dirName string) error
	UploadFile(fileName string, content []byte) error
	DownloadFile(fileName string) ([]byte, error)
	DeleteFile(fileName string) error
}

type PersistMode int

const (
	PersistAll = iota
	PersistOnlyHeaders
	PersistOnlyTime
	PersistOnlyTags
	PersistOnlyPopularityLikes
	PersistOnlyPopularityComments
	PersistOnlyPopularityViews
	PersistOnlyCountLikes
	PersistOnlyCountComments
	PersistOnlyCountViews
	PersistOnlyDrafts
	PersistOnlyDeleted
	PersistOnlyUsers
	PersistOnlyUsersFollowers
)
