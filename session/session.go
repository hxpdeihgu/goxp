package session

import (
	"sync"
	"time"
	"fmt"
	"encoding/hex"
	"crypto/rand"
	"net/http"
	"os"
	"path"
	"encoding/json"
	"goxp"
	"io/ioutil"
	"path/filepath"
	"net/url"
)

//cookiename
const sessionName = "_goxpid"
const sessionPath = "/tmp/session"
const aesKey = "goxpGoxpGOxpGOXpGOXP"

const cookieLifeTime  = time.Hour*2
var sessionMap map[string]time.Time
var sessionId string

type session struct {
	lock sync.RWMutex
	id string
	value map[string]string
	expire time.Time
}

func SessionInit(){
	sessionMap = make(map[string]time.Time)
	flush()
	GC()
}
func SessionStart(this *goxp.Xp) *session{
	sessionId,_ = getId(this.Rq)
	s:=new(session)
	s.value = make(map[string]string)
	if _,ok:=sessionMap[sessionId];!ok{
		cookieSession:= &http.Cookie{
			Name:     sessionName,
			Value:    sessionId,
			Path:     "/",
			HttpOnly: true,
			Secure:   false,
		}
		cookieSession.Expires = time.Now().Add(cookieLifeTime)
		http.SetCookie(this.Rs,cookieSession)
	}
	s.id = sessionId
	s.sessionRead()
	return s
}


const SessionIDLength = 26

func getSessionId() (string,error) {
	b := make([]byte, SessionIDLength)
	n, err := rand.Read(b)
	if n != len(b) || err != nil {
		return "", fmt.Errorf("Could not successfully read from the system CSPRNG.%v",err)
	}
	return hex.EncodeToString(b), nil
}

func (this *session) Add(key string,value string) *session{
	this.lock.Lock()
	defer this.lock.Unlock()
	this.value[key] = value
	sessionMap[this.id] = time.Now()
	return this
}
func (this *session) Get(key string) string{
	this.lock.RLock()
	defer this.lock.RUnlock()
	if value,ok:=this.value[key];ok{
		return value
	}
	return ""
}


func (this *session) Commit() {
	b,err:=json.Marshal(this.value)
	if err !=nil {
		goxp.Error(err)
		return
	}
	aesEncrypt:=goxp.AesEncrypt{Key:aesKey}
	strValue,err := aesEncrypt.EncryptBytes(b)
	if err!=nil{
		goxp.Error(err)
		return
	}
	pathSession:=path.Join(sessionPath, string(this.id[0]), string(this.id[1]), this.id)
	_, err = os.Stat(pathSession)
	var f *os.File
	if err == nil {
		f, err = os.OpenFile(pathSession, os.O_RDWR, 0777)
	} else if os.IsNotExist(err) {
		f, err = os.Create(pathSession)
	} else {
		return
	}
	f.Truncate(0)
	f.Seek(0, 0)
	f.Write(strValue)
	f.Close()
}

//清空所有session文件
func flush()  {
	os.RemoveAll(sessionPath)
}

func (this *session) sessionRead() error {
	this.lock.Lock()
	defer this.lock.Unlock()
	pathSession:=path.Join(sessionPath, string(this.id[0]), string(this.id[1]), this.id)
	err := os.MkdirAll(path.Join(sessionPath, string(this.id[0]), string(this.id[1])), 0777)
	if err != nil {
		goxp.Error(err)
		return err
	}
	
	_, err = os.Stat(pathSession)
	var f *os.File
	if err == nil {
		f, err = os.OpenFile(pathSession, os.O_RDWR, 0777)
	} else if os.IsNotExist(err) {
		f, err = os.Create(pathSession)
	} else {
		goxp.Error(err)
		return err
	}
	
	os.Chtimes(pathSession, time.Now(), time.Now())
	b, err := ioutil.ReadAll(f)
	if err != nil {
		goxp.Error(err)
		return err
	}
	
	aesEncrypt:=goxp.AesEncrypt{Key:aesKey}
	strValue,err := aesEncrypt.DecryptBytes(b)
	if err !=nil {
		goxp.Error(err)
		return err
	}
	
	err=json.Unmarshal(strValue,&this.value)
	if err != nil {
		return  err
	}
	f.Close()
	return nil
}

func (this *session) Id() string {
	return this.id
}
func gcpath(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	if (info.ModTime().Add(cookieLifeTime).Unix()) < time.Now().Unix() {
		err=os.Remove(path)
		return err
	}
	return nil
}

func getId(r *http.Request) (string, error) {
	cookie, errs := r.Cookie(sessionName)
	if errs != nil || cookie.Value == "" {
		return getSessionId()
	}
	return url.QueryUnescape(cookie.Value)
}

func GC(){
	var mutex sync.Mutex
	tmp := sessionMap
	mutex.Lock()
	for id,t:=range tmp{
		if (t.Add(cookieLifeTime).Unix()) < time.Now().Unix() {
			pathSession:=path.Join(sessionPath, string(id[0]), string(id[1]),id)
			err:=filepath.Walk(pathSession, gcpath)
			if err==nil{
				delete(sessionMap,pathSession)
			}
		}
	}
	mutex.Unlock()
	time.AfterFunc(time.Hour*2,GC)
}