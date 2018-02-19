package auth

import (
	"context"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"os"

	"github.com/google/go-github/github"
	"github.com/gorilla/sessions"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/oauth2"
)

const authUrl = "https://github.com/login/oauth/authorize"
const tokenUrl = "https://github.com/login/oauth/access_token"

type Handler struct {
	conf  *oauth2.Config
	store sessions.Store
}

func New() *Handler {
	h := &Handler{}
	h.conf = &oauth2.Config{
		ClientID:     os.Getenv("ICEHUB_GH_CLIENT_ID"),
		ClientSecret: os.Getenv("ICEHUB_GH_CLIENT_SECRET"),
		Scopes:       []string{"public_repo", "user:email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  authUrl,
			TokenURL: tokenUrl,
		},
		//RedirectURL: TODO: what is redirect URL for ?
	}
	h.store = sessions.NewCookieStore([]byte("a very secret key ..."))
	return h
}

// gh oauth is based on https://github.com/andrewtian/golang-github-oauth-example/blob/master/main.go

func (h *Handler) GitHubLogin(w http.ResponseWriter, r *http.Request) {
	b := make([]byte, 16)
	rand.Read(b)

	state := base64.URLEncoding.EncodeToString(b)

	session, _ := h.store.Get(r, "sess")
	session.Values["state"] = state
	session.Save(r, w)

	url := h.conf.AuthCodeURL(state)
	http.Redirect(w, r, url, 302)
}

func (h *Handler) GitHubLoginCallback(w http.ResponseWriter, r *http.Request) {
	session, err := h.store.Get(r, "sess")
	if err != nil {
		fmt.Fprint(w, "session not founf")
		return
	}
	if r.URL.Query().Get("state") != session.Values["state"] {
		fmt.Fprint(w, "state does not match")
		return
	}

	// TODO: we should check if span is null
	rootSpan := opentracing.SpanFromContext(r.Context())
	span := rootSpan.Tracer().StartSpan("github-exchange", opentracing.ChildOf(rootSpan.Context()))
	// TODO: use not default http client ... it would use http.DefaultClient if it does not find it from ctx ....
	tkn, err := h.conf.Exchange(context.Background(), r.URL.Query().Get("code"))
	if err != nil {
		fmt.Fprint(w, "there was an issue getting your token")
		span.Finish()
		return
	}
	if !tkn.Valid() {
		fmt.Println(w, "invalid token")
		span.Finish()
		return
	}
	span.Finish()

	span = rootSpan.Tracer().StartSpan("github-info", opentracing.ChildOf(rootSpan.Context()))
	defer span.Finish()
	client := github.NewClient(h.conf.Client(context.Background(), tkn))

	user, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		fmt.Println(w, "error getting name")
		return
	}

	fmt.Fprintf(w, "%s", user)
	return
}
