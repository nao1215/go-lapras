package lapras

// Person is the information of one LAPRAS user
type Person struct {
	// Name is user name.
	Name string `json:"name"`
	// Description is personal information to be filled in by the user.
	Description string `json:"description"`
	// EScore is Engineering score
	EScore float64 `json:"e_score"`
	// BScore is Business score
	BScore float64 `json:"b_score"`
	// IScore is Influense score
	IScore float64 `json:"i_score"`
	// QiitaArticles is slice of Qiita article
	QiitaArticles []QiitaArticle `json:"qiita_articles"`
	// ZennArticles is slice of Zenn article
	ZennArticles []ZennArticle `json:"zenn_articles"`
	// BlogArticles is slice of Blog article
	BlogArticles []BlogArticle `json:"blog_articles"`
	// NoteArticles is slice of Note article
	NoteArticles []NoteArticle `json:"note_articles"`
	// SpeakerDeckSlides is slice of SpeakerDeck slide information.
	SpeakerDeckSlides []SpeakerDeckSlide `json:"speaker_deck_slides"`
	// GithubRepositories is slice of github repository information
	GithubRepositories []GithubRepository `json:"github_repositories"`
	// TeratailReplies is slice of teratail reply information
	TeratailReplies []TeratailReply `json:"teratail_replies"`
	// TeratailReplies is slice of event
	Events []Event `json:"events"`
	// Activities is slice of activity log
	Activities []Activity `json:"activities"`
}

// QiitaArticle is article in qiita.
type QiitaArticle struct {
	// Title is article title
	Title string `json:"title"`
	// URL is article url
	URL string `json:"url"`
	// Tags attached to the article
	Tags []string `json:"tags"`
	// Headlines is article headlines
	Headlines []string `json:"headlines"`
	// StockersCount is number of how many users stocked
	StockersCount int `json:"stockers_count"`
	// UpdateAt is update date & time
	UpdatedAt string `json:"updated_at"`
}

// ZennArticle is article in zenn.
type ZennArticle struct {
	// Title is article title
	Title string `json:"title"`
	// URL is article url
	URL string `json:"url"`
	// Tags attached to the article
	Tags []string `json:"tags"`
	// PostedAt is post date & time
	PostedAt string `json:"posted_at"`
}

// BlogArticle is article in blog.
type BlogArticle struct {
	// Title is article title
	Title string `json:"title"`
	// URL is article url
	URL string `json:"url"`
	// Tags attached to the article
	Tags []string `json:"tags"`
	// PostedAt is posted date & time
	PostedAt string `json:"posted_at"`
}

// NoteArticle is article in note.
type NoteArticle struct {
	// URL is article url
	URL string `json:"url"`
	// Title is article title
	Title string `json:"title"`
	// Tags attached to the article
	Tags []string `json:"tags"`
	// LikeCount is number of how many users liked
	LikeCount int `json:"like_count"`
	// PublishedAt is published date & time
	PublishedAt string `json:"published_at"`
}

// SpeakerDeckSlide is SpeakerDeck slide information.
type SpeakerDeckSlide struct {
	// Title is SpeakerDeck Slide title
	Title string `json:"title"`
	// URL is SpeakerDeck Slide url
	URL string `json:"url"`
	// Tags attached to the SpeakerDeck Slide
	Tags []string `json:"tags"`
	// Description is description for slide
	Description string `json:"description"`
	// ViewCount is number of how many users viewed
	ViewCount int `json:"view_count"`
	// ViewCount is number of how many users stared
	StarCount int `json:"star_count"`
	// PresentationDate is presentation date
	PresentationDate string `json:"presentation_date"`
}

// GithubRepository is github repository information
type GithubRepository struct {
	// ID is github repository id
	ID int `json:"id"`
	// Title is github repository title
	Title string `json:"title"`
	// URL is github repository url
	URL string `json:"url"`
	// IsOSS is whether repository is open source or not
	IsOSS bool `json:"is_oss"`
	// IsFork is whether repository is possible to fork
	IsFork bool `json:"is_fork"`
	// IsOwner is whether the user is the owner of the repository
	IsOwner bool `json:"is_owner"`
	// Description is github repository description
	Description string `json:"description"`
	// StargazersCount is stargazers count
	StargazersCount int `json:"stargazers_count"`
	// StargazersURL is stargazers url
	StargazersURL string `json:"stargazers_url"`
	// Forks is how many forks the repository
	Forks int `json:"forks"`
	// ContributorsCount is how many contributors the repository
	ContributorsCount int `json:"contributors_count"`
	// ContributorsURL is contributors url
	ContributorsURL string `json:"contributors_url"`
	// Contributions is how many contributions the repository
	Contributions int `json:"contributions"`
	// ContributionsURL is contributions url
	ContributionsURL string `json:"contributions_url"`
	// Language is programing language name
	Language string `json:"language"`
	// Languages is how many bytes of a certain programming language are used
	Languages []struct {
		// Name is programing language name
		Name string `json:"name"`
		// Bytes is how many bytes of a certain programming language are used
		Bytes int `json:"bytes"`
	} `json:"languages"`
}

// TeratailReply is teratail reply information
type TeratailReply struct {
	// Title is teratail title
	Title string `json:"title"`
	// URL is teratail url
	URL string `json:"url"`
	// IsBestAnswer is whether best answer or not
	IsBestAnswer bool
	// Tags attached to the teratail reply
	Tags []string `json:"tags"`
	// CreatedAt is create date & time
	CreatedAt string `json:"created_at"`
}

// Event is information about events attended by users
type Event struct {
	// Title is event title
	Title string `json:"title"`
	// URL is event url
	URL string `json:"url"`
	// Status is TODO: add documentation comment
	Status int `json:"status"`
	// IsPresenter is whether user is presenter or not
	IsPresenter bool `json:"is_presenter"`
	// IsOrganizer is whether user is organizer or not
	IsOrganizer bool `json:"is_organizer"`
}

// Activity is activity log
type Activity struct {
	// Title is activity title
	Title string `json:"title"`
	// URL is activity url
	URL string `json:"url"`
	// Date is activity log date
	Date string `json:"date"`
	// Type is activity type
	Type string `json:"type"`
}
