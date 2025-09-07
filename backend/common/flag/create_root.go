package flag

import (
	"fmt"
	"go-admin-server/api/entity"
	"go-admin-server/global"
	"go-admin-server/pkg/encrypt"
	"syscall"

	"golang.org/x/term"
)

func CreateRootAccount() error {
	var root entity.SysAdmin

	var username string
	fmt.Println("请输入账号: ")
	fmt.Scanln(&username)

	fmt.Println("请输入密码: ")
	pwdBytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return err
	}
	pwd := string(pwdBytes)

	fmt.Println("请确认密码: ")
	repwdBytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return err
	}
	repwd := string(repwdBytes)

	if pwd != repwd {
		return err
	}

	hashPwd, _ := encrypt.EncryptPassword(pwd)

	root.Username = username
	root.Password = hashPwd
	if err := global.DB.Create(&root).Error; err != nil {
		return err
	}
	return nil
}
