package render

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
	"github.com/justinas/nosurf"
)

type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
	JetViews   *jet.Set
	Session    *scs.SessionManager
}

type TemplateData struct {
	IsAuthenticated bool
	IntMap          map[string]int
	StringMap       map[string]string
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Port            string
	ServerName      string
	Secure          bool
	Error           string
	Flash           string
}

func (g *Render) defaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.Secure = g.Secure
	td.ServerName = g.ServerName
	td.CSRFToken = nosurf.Token(r)
	td.Port = g.Port
	if g.Session.Exists(r.Context(), "userID") {
		td.IsAuthenticated = true
	}
	td.Error = g.Session.PopString(r.Context(), "error")
	td.Flash = g.Session.PopString(r.Context(), "flash")
	return td
}

func (g *Render) Page(w http.ResponseWriter, r *http.Request, view string, variables, data interface{}) error {
	switch strings.ToLower(g.Renderer) {
	case "go":
		return g.GoPage(w, r, view, data)
	case "jet":
		return g.JetPage(w, r, view, variables, data)
	default:

	}
	return errors.New("no rendering engine specified")
}

// GoPage renders a standard Go template
func (g *Render) GoPage(w http.ResponseWriter, r *http.Request, view string, data interface{}) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/views/%s.page.tmpl", g.RootPath, view))
	if err != nil {
		return err
	}

	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}

	err = tmpl.Execute(w, &td)
	if err != nil {
		return err
	}

	return nil
}

// JetPage renders a template using the Jet templating engine
func (g *Render) JetPage(w http.ResponseWriter, r *http.Request, templateName string, variables, data interface{}) error {
	var vars jet.VarMap

	if variables == nil {
		vars = make(jet.VarMap)
	} else {
		vars = variables.(jet.VarMap)
	}

	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}

	td = g.defaultData(td, r)

	t, err := g.JetViews.GetTemplate(fmt.Sprintf("%s.jet", templateName))
	if err != nil {
		log.Println(err)
		return err
	}

	if err = t.Execute(w, vars, td); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
