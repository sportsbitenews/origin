package raven

// http://sentry.readthedocs.org/en/latest/developer/interfaces/index.html#sentry.interfaces.Message
type Message struct {
	// Required
	Message string `json:"message"`

	// Optional
	Params []interface{} `json:"params,omitempty"`
}

func (m *Message) Class() string { return "sentry.interfaces.Message" }

// http://sentry.readthedocs.org/en/latest/developer/interfaces/index.html#sentry.interfaces.Template
type Template struct {
	// Required
	Filename    string `json:"filename"`
	Lineno      int    `json:"lineno"`
	ContextLine string `json:"context_line"`

	// Optional
	PreContext   []string `json:"pre_context,omitempty"`
	PostContext  []string `json:"post_context,omitempty"`
	AbsolutePath string   `json:"abs_path,omitempty"`
}

func (t *Template) Class() string { return "sentry.interfaces.Template" }

// http://sentry.readthedocs.org/en/latest/developer/interfaces/index.html#sentry.interfaces.User
type User struct {
	ID       string `json:"id"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

func (h *User) Class() string { return "sentry.interfaces.User" }

// http://sentry.readthedocs.org/en/latest/developer/interfaces/index.html#sentry.interfaces.Query
type Query struct {
	// Required
	Query string `json:"query"`

	// Optional
	Engine string `json:"engine,omitempty"`
}

func (q *Query) Class() string { return "sentry.interfaces.Query" }
