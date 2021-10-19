package timelock

import (
	"testing"

	"github.com/drand/kyber"
	bls "github.com/drand/kyber-bls12381"
	"github.com/drand/kyber/util/random"
	"github.com/stretchr/testify/require"
)

func TestTimelock(t *testing.T) {
	suite := bls.NewBLS12381Suite()
	P := suite.G1().Point().Base()
	s := suite.G1().Scalar().Pick(random.New())
	Ppub := suite.G1().Point().Mul(s, P)

	ID := []byte("passtherand")
	IDP := suite.G2().Point().(kyber.HashablePoint)
	Qid := IDP.Hash(ID)     // public key
	sQid := Qid.Mul(s, Qid) // secret key

	msg := []byte("Hello World\n")
	c, err := Encrypt(suite, Ppub, ID, msg)
	require.NoError(t, err)
	msg2, err := Decrypt(suite, Ppub, sQid, c)
	require.NoError(t, err)
	require.Equal(t, msg, msg2)
}
