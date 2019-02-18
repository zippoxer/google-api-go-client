// Package admin provides access to the Email Migration API v2.
//
// See https://developers.google.com/admin-sdk/email-migration/v2/
//
// Usage example:
//
//   import "github.com/zippoxer/google-api-go-client/admin/email_migration/v2"
//   ...
//   adminService, err := admin.New(oauthHttpClient)
package admin // import "github.com/zippoxer/google-api-go-client/admin/email_migration/v2"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "github.com/zippoxer/google-api-go-client/gensupport"
	googleapi "github.com/zippoxer/google-api-go-client/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "admin:email_migration_v2"
const apiName = "admin"
const apiVersion = "email_migration_v2"
const basePath = "https://www.googleapis.com/email/v2/users/"

// OAuth2 scopes used by this API.
const (
	// Manage email messages of users on your domain
	EmailMigrationScope = "https://www.googleapis.com/auth/email.migration"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Mail = NewMailService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Mail *MailService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewMailService(s *Service) *MailService {
	rs := &MailService{s: s}
	return rs
}

type MailService struct {
	s *Service
}

// MailItem: JSON template for MailItem object in Email Migration API.
type MailItem struct {
	// IsDeleted: Boolean indicating if the mail is deleted (used in Vault)
	IsDeleted bool `json:"isDeleted,omitempty"`

	// IsDraft: Boolean indicating if the mail is draft
	IsDraft bool `json:"isDraft,omitempty"`

	// IsInbox: Boolean indicating if the mail is in inbox
	IsInbox bool `json:"isInbox,omitempty"`

	// IsSent: Boolean indicating if the mail is in 'sent mails'
	IsSent bool `json:"isSent,omitempty"`

	// IsStarred: Boolean indicating if the mail is starred
	IsStarred bool `json:"isStarred,omitempty"`

	// IsTrash: Boolean indicating if the mail is in trash
	IsTrash bool `json:"isTrash,omitempty"`

	// IsUnread: Boolean indicating if the mail is unread
	IsUnread bool `json:"isUnread,omitempty"`

	// Kind: Kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// Labels: List of labels (strings)
	Labels []string `json:"labels,omitempty"`

	// ForceSendFields is a list of field names (e.g. "IsDeleted") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *MailItem) MarshalJSON() ([]byte, error) {
	type noMethod MailItem
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// method id "emailMigration.mail.insert":

type MailInsertCall struct {
	s                   *Service
	userKey             string
	mailitem            *MailItem
	urlParams_          gensupport.URLParams
	media_              io.Reader
	mediaType_          string
	resumable_          googleapi.SizeReaderAt
	resumableMediaType_ string
	protocol_           string
	progressUpdater_    googleapi.ProgressUpdater
	ctx_                context.Context
}

// Insert: Insert Mail into Google's Gmail backends
func (r *MailService) Insert(userKey string, mailitem *MailItem) *MailInsertCall {
	c := &MailInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.userKey = userKey
	c.mailitem = mailitem
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *MailInsertCall) QuotaUser(quotaUser string) *MailInsertCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *MailInsertCall) UserIP(userIP string) *MailInsertCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Media specifies the media to upload in a single chunk. At most one of
// Media and ResumableMedia may be set.
func (c *MailInsertCall) Media(r io.Reader, options ...googleapi.MediaOption) *MailInsertCall {
	opts := googleapi.ProcessMediaOptions(options)
	c.media_, c.mediaType_ = gensupport.DetermineContentType(r, opts.ContentType)
	c.protocol_ = "multipart"
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be
// canceled with ctx. At most one of Media and ResumableMedia may be
// set. mediaType identifies the MIME media type of the upload, such as
// "image/png". If mediaType is "", it will be auto-detected. The
// provided ctx will supersede any context previously provided to the
// Context method.
func (c *MailInsertCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *MailInsertCall {
	c.ctx_ = ctx
	c.resumable_ = io.NewSectionReader(r, 0, size)
	c.resumableMediaType_ = mediaType
	c.protocol_ = "resumable"
	return c
}

// ProgressUpdater provides a callback function that will be called
// after every chunk. It should be a low-latency function in order to
// not slow down the upload operation. This should only be called when
// using ResumableMedia (as opposed to Media).
func (c *MailInsertCall) ProgressUpdater(pu googleapi.ProgressUpdater) *MailInsertCall {
	c.progressUpdater_ = pu
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MailInsertCall) Fields(s ...googleapi.Field) *MailInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
// This context will supersede any context previously provided to the
// ResumableMedia method.
func (c *MailInsertCall) Context(ctx context.Context) *MailInsertCall {
	c.ctx_ = ctx
	return c
}

func (c *MailInsertCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.mailitem)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{userKey}/mail")
	if c.media_ != nil || c.resumable_ != nil {
		urls = strings.Replace(urls, "https://www.googleapis.com/", "https://www.googleapis.com/upload/", 1)
		c.urlParams_.Set("uploadType", c.protocol_)
	}
	urls += "?" + c.urlParams_.Encode()
	if c.protocol_ != "resumable" && c.media_ != nil {
		var combined io.ReadCloser
		combined, ctype = gensupport.CombineBodyMedia(body, ctype, c.media_, c.mediaType_)
		defer combined.Close()
		body = combined
	}
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"userKey": c.userKey,
	})
	if c.protocol_ == "resumable" {
		if c.resumableMediaType_ == "" {
			c.resumableMediaType_ = gensupport.DetectMediaType(c.resumable_)
		}
		req.Header.Set("X-Upload-Content-Type", c.resumableMediaType_)
	}
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "emailMigration.mail.insert" call.
func (c *MailInsertCall) Do() error {
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	if c.protocol_ == "resumable" {
		loc := res.Header.Get("Location")
		rx := &gensupport.ResumableUpload{
			Client:        c.s.client,
			UserAgent:     c.s.userAgent(),
			URI:           loc,
			Media:         c.resumable_,
			MediaType:     c.resumableMediaType_,
			ContentLength: c.resumable_.Size(),
			Callback: func(curr int64) {
				if c.progressUpdater_ != nil {
					c.progressUpdater_(curr, c.resumable_.Size())
				}
			},
		}
		res, err = rx.Upload(c.ctx_)
		if err != nil {
			return err
		}
		defer res.Body.Close()
	}
	return nil
	// {
	//   "description": "Insert Mail into Google's Gmail backends",
	//   "httpMethod": "POST",
	//   "id": "emailMigration.mail.insert",
	//   "mediaUpload": {
	//     "accept": [
	//       "message/rfc822"
	//     ],
	//     "maxSize": "35MB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/email/v2/users/{userKey}/mail"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/email/v2/users/{userKey}/mail"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "userKey"
	//   ],
	//   "parameters": {
	//     "userKey": {
	//       "description": "The email or immutable id of the user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userKey}/mail",
	//   "request": {
	//     "$ref": "MailItem"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/email.migration"
	//   ],
	//   "supportsMediaUpload": true
	// }

}
