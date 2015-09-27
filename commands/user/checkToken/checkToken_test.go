package checkToken

import (
	t "testing"
	"github.com/AitorGuerrero/UserGo/commands/manager/newUser"
	"github.com/AitorGuerrero/UserGo/commands/user/login"
	"github.com/AitorGuerrero/UserGo/implementation/services"
)

var com = Command{services.TokenChecker(), services.UserSource()}
var req = Request{}

func TestIfTheTokenDoesNotExistsShouldReturnAnError (t *t.T) {
	req.Token = "PisToken"
	err := com.Execute(req)
	if nil == err {
		t.Error("S ror")
	}
}

func TestIfUserDoNotOwnTheTokenShouldReturnAnError (t *t.T) {
	userAId := "userA";
	userAPasskey := "passA"
	userBId := "userB"
	commandNewUser := newUser.Command{services.UserSource(), services.UserFactory()}
	commandNewUser.Execute(newUser.Request{userAId, userAPasskey})
	commandNewUser.Execute(newUser.Request{userBId, "passB"})

	loginCommand := login.Command{services.UserLogin(), services.UserSource()}
	res, _ := loginCommand.Execute(login.Request{userAId, userAPasskey})
	userAToken := res.SessionToken

	err := com.Execute(Request{userBId, userAToken})

	if (nil == err) {
		t.Error("Should throw an error")
	} else if ("Incorrect Token" != err.Error()) {
		t.Error("Should throw 'Incorrect Token' error")
	}
}
