package actions_test

import (
	"inventoryApp/app"
	"testing"

	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/suite/v3"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	action, err := suite.NewActionWithFixtures(app.New(), packr.New("Test_ActionSuite", "../fixtures"))
	if err != nil {
		t.Fatal(err)
	}

	as := &ActionSuite{
		Action: action,
	}
	suite.Run(t, as)
}

func (as *ActionSuite) TableChange(table string, count int, fn func()) {
	scount, err := as.DB.Count(table)
	as.NoError(err)

	fn()

	ecount, err := as.DB.Count(table)
	as.NoError(err)
	as.Equal(count, ecount-scount)
}
