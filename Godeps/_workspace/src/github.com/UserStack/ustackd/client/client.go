package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
	"net/textproto"
	"strconv"
	"strings"
	"sync"

	"github.com/UserStack/ustackd/backends"
)

/* parts of the code are taken from smtp.go from the core library */

type Client struct {
	mutex   sync.Mutex
	Text    *textproto.Conn
	conn    net.Conn
	tlsConn *tls.Conn
	host    string
}

// Dial returns a new Client connected to an ustack server at addr.
// The addr must include a port number.
func Dial(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	host, _, _ := net.SplitHostPort(addr)
	return NewClient(conn, host)
}

// NewClient returns a new Client using an existing connection and host as a
// server name to be used when authenticating.
func NewClient(conn net.Conn, host string) (*Client, error) {
	text := textproto.NewConn(conn)
	line, err := text.ReadLine()
	if err != nil {
		text.Close()
		return nil, err
	}
	if !strings.Contains(line, "ustack") {
		text.Close()
		return nil, fmt.Errorf("Not a ustackd server")
	}
	return &Client{Text: text, conn: conn, host: host}, nil
}

func (client *Client) StartTls(config *tls.Config) error {
	_, err := client.Text.Cmd("starttls")
	if err != nil {
		return err
	}
	client.tlsConn = tls.Client(client.conn, config)
	if err != nil {
		return err
	}
	client.conn = client.tlsConn
	client.Text = textproto.NewConn(client.conn)
	return nil
}

func (client *Client) StartTlsWithoutCertCheck() error {
	return client.StartTls(&tls.Config{InsecureSkipVerify: true})
}

func (client *Client) StartTlsWithCert(cert string) error {
	pemCerts, err := ioutil.ReadFile(cert)
	if err != nil {
		return fmt.Errorf("Unable to open cert at %s: %s\n", cert, err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemCerts) {
		return fmt.Errorf("File doesn't contain pem certificates: %s\n", cert)
	}
	return client.StartTls(&tls.Config{
		RootCAs:            certPool,
		InsecureSkipVerify: false,
		ServerName:         client.host,
	})
}

func (client *Client) ClientAuth(password string) *backends.Error {
	return client.simpleCmd("client auth %s", password)
}

func (client *Client) CreateUser(name string, password string) (uid int64, err *backends.Error) {
	uid, err = client.simpleIntCmd("user %s %s", name, password)
	return
}

func (client *Client) DisableUser(nameuid string) *backends.Error {
	return client.simpleCmd("disable %s", nameuid)
}

func (client *Client) EnableUser(nameuid string) *backends.Error {
	return client.simpleCmd("enable %s", nameuid)
}

func (client *Client) SetUserData(nameuid string, key string, value string) *backends.Error {
	return client.simpleCmd("set %s %s %s", nameuid, key, value)
}

func (client *Client) GetUserData(nameuid string, key string) (string, *backends.Error) {
	list, err := client.listCmd("get %s %s", nameuid, key)
	if err != nil {
		return "", err
	}
	return list[0], nil
}

func (client *Client) LoginUser(name string, password string) (uid int64, err *backends.Error) {
	uid, err = client.simpleIntCmd("login %s %s", name, password)
	return
}

func (client *Client) ChangeUserPassword(nameuid string, password string, newpassword string) *backends.Error {
	return client.simpleCmd("change password %s %s %s", nameuid, password, newpassword)
}

func (client *Client) ChangeUserName(nameuid string, password string, newname string) *backends.Error {
	return client.simpleCmd("change name %s %s %s", nameuid, password, newname)
}

func (client *Client) UserGroups(nameuid string) (list []backends.Group, err *backends.Error) {
	list, err = client.listGroupCmd("user groups %s", nameuid)
	return
}

func (client *Client) DeleteUser(nameuid string) *backends.Error {
	return client.simpleCmd("delete user %s", nameuid)
}

func (client *Client) Users() (list []backends.User, err *backends.Error) {
	list, err = client.listUserCmd("users")
	return
}

func (client *Client) CreateGroup(name string) (gid int64, err *backends.Error) {
	gid, err = client.simpleIntCmd("group %s", name)
	return
}

func (client *Client) AddUserToGroup(nameuid string, groupgid string) *backends.Error {
	return client.simpleCmd("add %s %s", nameuid, groupgid)
}

func (client *Client) RemoveUserFromGroup(nameuid string, groupgid string) *backends.Error {
	return client.simpleCmd("remove %s %s", nameuid, groupgid)
}

func (client *Client) DeleteGroup(groupgid string) *backends.Error {
	return client.simpleCmd("delete group %s", groupgid)
}

func (client *Client) Groups() (list []backends.Group, err *backends.Error) {
	list, err = client.listGroupCmd("groups")
	return
}

func (client *Client) GroupUsers(groupgid string) (list []backends.User, err *backends.Error) {
	list, err = client.listUserCmd("group users %s", groupgid)
	return
}

func (client *Client) Stats() (stats map[string]int64, err *backends.Error) {
	stats = make(map[string]int64)
	list, err := client.listCmd("stats")
	if err != nil {
		return
	}

	for _, line := range list {
		args := strings.Split(line, ": ")
		if len(args) != 2 {
			err = &backends.Error{Code: "EFAULT", Message: "Expected two values: " + line}
			return
		}
		value, verr := strconv.Atoi(args[1])
		if verr != nil {
			err = &backends.Error{Code: "EFAULT", Message: "Expected number: " + args[1]}
			return
		}
		stats[strings.TrimSpace(args[0])] = int64(value)
	}
	return
}

func (client *Client) Close() {
	client.Text.Cmd("quit")
	client.Text.Close()
}

// Helpers

func (client *Client) handleIntResponse() (int64, *backends.Error) {
	line, rerr := client.Text.ReadLine()
	if rerr != nil {
		return 0, &backends.Error{Code: "EFAULT", Message: rerr.Error()}
	}
	ret := strings.Split(line, " ")
	if ret[0] == "-" {
		return 0, &backends.Error{Code: ret[1], Message: "Remote failure"}
	}
	val, perr := strconv.ParseInt(ret[2], 10, 64)
	if perr != nil {
		return 0, &backends.Error{Code: "EFAULT", Message: perr.Error()}
	}
	return val, nil
}

func (client *Client) handleResponse() *backends.Error {
	line, rerr := client.Text.ReadLine()
	if rerr != nil {
		return &backends.Error{Code: "EFAULT", Message: rerr.Error()}
	}
	ret := strings.Split(line, " ")
	if ret[0] == "-" {
		return &backends.Error{Code: ret[1], Message: "Remote failure"}
	}
	return nil
}

func (client *Client) simpleCmd(format string, args ...interface{}) *backends.Error {
	client.mutex.Lock()
	defer client.mutex.Unlock()
	_, err := client.Text.Cmd(format, args...)
	if err != nil {
		return &backends.Error{Code: "EFAULT", Message: err.Error()}
	}
	return client.handleResponse()
}

func (client *Client) simpleIntCmd(format string, args ...interface{}) (int64, *backends.Error) {
	client.mutex.Lock()
	defer client.mutex.Unlock()
	_, err := client.Text.Cmd(format, args...)
	if err != nil {
		return 0, &backends.Error{Code: "EFAULT", Message: err.Error()}
	}
	return client.handleIntResponse()
}

func (client *Client) listUserCmd(format string, args ...interface{}) ([]backends.User, *backends.Error) {
	list, err := client.listCmd(format, args...)
	if err != nil {
		return nil, err
	}

	var users []backends.User
	for _, line := range list {
		args := strings.Split(line, ":")
		if len(args) != 3 {
			return nil, &backends.Error{Code: "EFAULT", Message: "Expected three values: " + line}
		}
		uid, perr := strconv.ParseInt(args[1], 10, 64)
		if perr != nil {
			return nil, &backends.Error{Code: "EFAULT", Message: perr.Error()}
		}
		users = append(users, backends.User{
			Uid:    uid,
			Name:   args[0],
			Active: (args[2] == "Y"),
		})
	}
	return users, nil
}

func (client *Client) listGroupCmd(format string, args ...interface{}) ([]backends.Group, *backends.Error) {
	list, err := client.listCmd(format, args...)
	if err != nil {
		return nil, err
	}

	var groups []backends.Group
	for _, line := range list {
		args := strings.Split(line, ":")
		if len(args) != 2 {
			return nil, &backends.Error{Code: "EFAULT", Message: "Expected two values: " + line}
		}
		gid, perr := strconv.ParseInt(args[1], 10, 64)
		if perr != nil {
			return nil, &backends.Error{Code: "EFAULT", Message: perr.Error()}
		}
		groups = append(groups, backends.Group{
			Gid:  gid,
			Name: args[0],
		})
	}
	return groups, nil
}

func (client *Client) listCmd(format string, args ...interface{}) ([]string, *backends.Error) {
	client.mutex.Lock()
	defer client.mutex.Unlock()
	_, err := client.Text.Cmd(fmt.Sprintf(format, args...))
	if err != nil {
		return nil, &backends.Error{Code: "EFAULT", Message: err.Error()}
	}
	var list []string
	for {
		line, rerr := client.Text.ReadLine()
		if rerr != nil {
			return nil, &backends.Error{Code: "EFAULT", Message: rerr.Error()}
		}
		if strings.HasPrefix(line, "- E") {
			ret := strings.Split(line, " ")
			return nil, &backends.Error{Code: ret[1], Message: "Remote failure"}
		} else if strings.HasPrefix(line, "+ ") {
			return list, nil
		}
		list = append(list, line)
	}
}
