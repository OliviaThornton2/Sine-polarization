import (
	"testing"

	util "github.com/Pylons-tech/pylons/testutil/cli"
	"github.com/Pylons-tech/pylons/testutil/network"
	"github.com/Pylons-tech/pylons/x/pylons/client/cli")


func TestCreate(t *testing.T) {
	preTestValidate(t)

	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

  		args := []string{"NewUser0", goodPLC}
		cmd := DevCreate()
		_, err = clitestutil.ExecTestCLICmd(ctx, cmd, args)
		assert.Nil(t, err)
	}

func NewAccountCreationDecorator(ak types.AccountKeeper) AccountCreationDecorator {
	return AccountCreationDecorator{
		ak: ak,
	}
}

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Pylons-tech/pylons/x/pylons/types"
)

	}
	return cmd
}
