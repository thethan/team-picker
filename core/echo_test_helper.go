package core

import (
	"github.com/labstack/echo"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

type EchoContextMock struct {
	Assertable EchoAssertable
	Mockable EchoMockable
}

type EchoAssertable struct {
	HttpCode   int
	I          interface{}
	Name       string
	String     string
	BoundValue interface{}
}

type EchoMockable struct {
	FormValue        string
	Error            error
	Request          *http.Request
	ParamNames       []string
	ParamValues      []string
	QueryParamNames  []string
	QueryParamValues []string
	ContextData      map[string]interface{}
}

func (ecm *EchoContextMock) Request() *http.Request {
	return ecm.Mockable.Request
}

func (ecm *EchoContextMock) SetRequest(r *http.Request) {
	panic("implement me")
}

func (ecm *EchoContextMock) Response() *echo.Response {
	panic("implement me")
}

func (ecm *EchoContextMock) IsTLS() bool {
	panic("implement me")
}

func (ecm *EchoContextMock) IsWebSocket() bool {
	panic("implement me")
}

func (ecm *EchoContextMock) Scheme() string {
	panic("implement me")
}

func (ecm *EchoContextMock) RealIP() string {
	panic("implement me")
}

func (ecm *EchoContextMock) Path() string {
	panic("implement me")
}

func (ecm *EchoContextMock) SetPath(p string) {
	panic("implement me")
}

func (ecm *EchoContextMock) Param(name string) string {
	panic("implement me")
}

func (ecm *EchoContextMock) ParamNames() []string {
	panic("implement me")
}

func (ecm *EchoContextMock) SetParamNames(names ...string) {
	ecm.Mockable.ParamNames = names
}

func (ecm *EchoContextMock) ParamValues() []string {
	panic("implement me")
}

func (ecm *EchoContextMock) SetParamValues(values ...string) {
	ecm.Mockable.ParamValues = values
}

func (ecm *EchoContextMock) QueryParam(name string) string {
	panic("implement me")
}

func (ecm *EchoContextMock) QueryParams() url.Values {
	panic("implement me")
}

func (ecm *EchoContextMock) QueryString() string {
	panic("implement me")
}

func (ecm *EchoContextMock) FormValue(name string) string {
	panic("implement me")
}

func (ecm *EchoContextMock) FormParams() (url.Values, error) {
	panic("implement me")
}

func (ecm *EchoContextMock) FormFile(name string) (*multipart.FileHeader, error) {
	panic("implement me")
}

func (ecm *EchoContextMock) MultipartForm() (*multipart.Form, error) {
	panic("implement me")
}

func (ecm *EchoContextMock) Cookie(name string) (*http.Cookie, error) {
	panic("implement me")
}

func (ecm *EchoContextMock) SetCookie(cookie *http.Cookie) {
	panic("implement me")
}

func (ecm *EchoContextMock) Cookies() []*http.Cookie {
	panic("implement me")
}

func (ecm *EchoContextMock) Get(key string) interface{} {
	return ecm.Mockable.ContextData[key]
}

func (ecm *EchoContextMock) Set(key string, val interface{}) {
	ecm.Mockable.ContextData[key] = val
}

func (ecm *EchoContextMock) Bind(i interface{}) error {
	panic("implement me")
}

func (ecm *EchoContextMock) Validate(i interface{}) error {
	panic("implement me")
}

func (ecm *EchoContextMock) Render(code int, name string, data interface{}) error {
	panic("implement me")
}

func (ecm *EchoContextMock) HTML(code int, html string) error {
	panic("implement me")
}

func (ecm *EchoContextMock) HTMLBlob(code int, b []byte) error {
	panic("implement me")
}

func (ecm *EchoContextMock) String(code int, s string) error {
	panic("implement me")
}

func (ecm *EchoContextMock) JSON(code int, i interface{}) error {
	ecm.Assertable.HttpCode = code
	ecm.Assertable.I = i

	return ecm.Mockable.Error
}

func (ecm *EchoContextMock) JSONPretty(code int, i interface{}, indent string) error {
	panic("implement me")
}

func (ecm *EchoContextMock) JSONBlob(code int, b []byte) error {
	panic("implement me")
}

func (ecm *EchoContextMock) JSONP(code int, callback string, i interface{}) error {
	panic("implement me")
}

func (ecm *EchoContextMock) JSONPBlob(code int, callback string, b []byte) error {
	panic("implement me")
}

func (ecm *EchoContextMock) XML(code int, i interface{}) error {
	panic("implement me")
}

func (ecm *EchoContextMock) XMLPretty(code int, i interface{}, indent string) error {
	panic("implement me")
}

func (ecm *EchoContextMock) XMLBlob(code int, b []byte) error {
	panic("implement me")
}

func (ecm *EchoContextMock) Blob(code int, contentType string, b []byte) error {
	panic("implement me")
}

func (ecm *EchoContextMock) Stream(code int, contentType string, r io.Reader) error {
	panic("implement me")
}

func (ecm *EchoContextMock) File(file string) error {
	panic("implement me")
}

func (ecm *EchoContextMock) Attachment(file string, name string) error {
	panic("implement me")
}

func (ecm *EchoContextMock) Inline(file string, name string) error {
	panic("implement me")
}

func (ecm *EchoContextMock) NoContent(code int) error {
	panic("implement me")
}

func (ecm *EchoContextMock) Redirect(code int, url string) error {
	panic("implement me")
}

func (ecm *EchoContextMock) Error(err error) {
	panic("implement me")
}

func (ecm *EchoContextMock) Handler() echo.HandlerFunc {
	panic("implement me")
}

func (ecm *EchoContextMock) SetHandler(h echo.HandlerFunc) {
	panic("implement me")
}

func (ecm *EchoContextMock) Logger() echo.Logger {
	panic("implement me")
}

func (ecm *EchoContextMock) Echo() *echo.Echo {
	panic("implement me")
}

func (ecm *EchoContextMock) Reset(r *http.Request, w http.ResponseWriter) {
	panic("implement me")
}
