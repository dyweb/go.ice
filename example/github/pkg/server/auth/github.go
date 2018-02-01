package auth

import (
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"

	"context"
	"encoding/base64"
	"fmt"
	"github.com/google/go-github/github"
	"math/rand"
	"net/http"
	"os"
)

const authUrl = "https://github.com/login/oauth/authorize"
const tokenUrl = "https://github.com/login/oauth/access_token"

// gh oauth is based on https://github.com/andrewtian/golang-github-oauth-example/blob/master/main.go
type Gh struct {
	conf  *oauth2.Config
	store sessions.Store
}

func NewGh() *Gh {
	gh := &Gh{}
	gh.conf = &oauth2.Config{
		ClientID:     os.Getenv("ICEHUB_GH_CLIENT_ID"),
		ClientSecret: os.Getenv("ICEHUB_GH_CLIENT_SECRET"),
		Scopes:       []string{"public_repo", "user:email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  authUrl,
			TokenURL: tokenUrl,
		},
		//RedirectURL: TODO: what is redirect URL for ?
	}
	gh.store = sessions.NewCookieStore([]byte("a very secret key ..."))
	return gh
}

func (gh *Gh) Start(w http.ResponseWriter, r *http.Request) {
	b := make([]byte, 16)
	rand.Read(b)

	state := base64.URLEncoding.EncodeToString(b)

	session, _ := gh.store.Get(r, "sess")
	session.Values["state"] = state
	session.Save(r, w)

	url := gh.conf.AuthCodeURL(state)
	http.Redirect(w, r, url, 302)
}

func (gh *Gh) Cb(w http.ResponseWriter, r *http.Request) {
	session, err := gh.store.Get(r, "sess")
	if err != nil {
		fmt.Fprint(w, "session not founf")
		return
	}
	if r.URL.Query().Get("state") != session.Values["state"] {
		fmt.Fprint(w, "state does not match")
		return
	}
	// TODO: use not default http client ... it would use http.DefaultClient if it does not find it from ctx ....
	tkn, err := gh.conf.Exchange(context.Background(), r.URL.Query().Get("code"))
	if err != nil {
		fmt.Fprint(w, "there was an issue getting your token")
		return
	}
	if !tkn.Valid() {
		fmt.Println(w, "invalid token")
		return
	}

	client := github.NewClient(gh.conf.Client(context.Background(), tkn))

	user, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		fmt.Println(w, "error getting name")
		return
	}

	fmt.Fprintf(w, "%s", user)
	return

}
