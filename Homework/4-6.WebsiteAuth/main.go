package static

import (
  	"fmt"
  	"html/template"
      "net/http"
  	"time"
  	
  	"appengine"
  	"appengine/datastore"
  	"appengine/user"
  )
  
  type Post struct {
  	Author string
  	User string
  	Content string
  	Date time.Time
  	Key string
}

var imageboardTemplate = template.Must(template.ParseFiles("public/index.html"))

func init() {
    http.HandleFunc("/", root)
    http.HandleFunc("/posts", posts)
  	http.HandleFunc("/edit", edit)
  	http.HandleFunc("/submitedit", submitedit)
  	http.HandleFunc("/destroy",destroy)
}

func imageboardKey(c appengine.Context) *datastore.Key {
	  return datastore.NewKey(c, "Post", "post_listing", 0, nil)
}

func root(w http.ResponseWriter, r *http.Request) {
  	conroot := appengine.NewContext(r)
  	userroot := user.Current(conroot)
  	if userroot == nil 
  	{
  		url, err := user.LoginURL(conroot, r.URL.String())
  		if err != nil {
  			http.Error(w, err.Error(), http.StatusInternalServerError)
  			return
  		}
          w.Header().Set("Location", url)
          w.WriteHeader(http.StatusFound)
          return
	  }
	q := datastore.NewQuery("Post").Ancestor(imageboardKey(c)).Order("-Date").Limit(10)
	
	postlist := make([]Post, 0, 10)
	
	if _, err := q.GetAll(c, &postlist); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	return
	}

	if err := imageboardTemplate.Execute(w, postlist); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func posts(w http.ResponseWriter, r *http.Request) {
	conpost := appengine.NewContext(r)
	p := Post{
		Content: r.FormValue("content"),
		Date: time.Now(),
	}
	if u := user.Current(conpost); u!= nil {
		p.Author = u.String()
		p.User = u.ID
	} else {
		p.User = ""
	}
	nkey := datastore.NewKey(conpost, "Post", "ohbanana", 0, imageboardKey(conpost))
	p.Key = nkey.Encode()
	_, err := datastore.Put(conpost, nkey, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

var editTemplate = template.Must(template.ParseFiles("public/page.html"))

func edit(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	p := new(Post)
	// add something to stop people from just dumping it when they want
	keyURL := r.FormValue("skey")
	nkey, err := datastore.DecodeKey(keyURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := datastore.Get(c, nkey, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if u := user.Current(c); u != nil {
		if  p.User == u.ID {
			if err := editTemplate.Execute(w, p); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	} else if p.User == "" {
		if err := editTemplate.Execute(w, p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		fmt.Fprintf(w,"Access denied.")
	}
}

func submitedit(w http.ResponseWriter, r *http.Request) {
	csubmit := appengine.NewContext(r)
	psubmit := new(Post)
	// add something to stop people from just dumping it when they want
	nkey, err := datastore.DecodeKey(r.FormValue("skey"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := datastore.Get(csubmit, nkey, psubmit); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p.Content = r.FormValue("content")
	_, err2 := datastore.Put(csubmit, nkey, psubmit)
	if err2 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func destroy(w http.ResponseWriter, r *http.Request) {
	cdelete := appengine.NewContext(r)
	keyURL := r.FormValue("skey")
	nkey, err := datastore.DecodeKey(keyURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	if err := datastore.Delete(cdelete, nkey); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
