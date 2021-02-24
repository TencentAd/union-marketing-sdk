package orm

import (
	"testing"

	"git.code.oa.com/tme-server-component/kg_growth_open/api/sdk"
	"github.com/stretchr/testify/assert"
)

func TestAuthAccount(t *testing.T) {
	db, err := setupTestDB(t)
	assert.NoError(t, err)

	{
		account := &sdk.AuthAccount{
			AccountId:   1,
			AccessToken: "a",
			ScopeList: []string{"a"},
		}
		assert.NoError(t, AuthAccountUpsert(db, account))
	}
	{
		account, err := AuthAccountGetAll(db)
		assert.NoError(t, err)
		assert.Len(t, account, 1)
		assert.EqualValues(t, 1, account[0].AccountId)
	}

	{
		account := &sdk.AuthAccount{
			AccountId:   1,
			AccessToken: "b",
		}
		assert.NoError(t, AuthAccountUpdate(db, account))
	}
	{
		account, err := AuthAccountGetAll(db)
		assert.NoError(t, err)
		assert.Len(t, account, 1)
		assert.EqualValues(t, "b", account[0].AccessToken)
	}
}