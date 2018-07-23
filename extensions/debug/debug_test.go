package debug_test

import (
	"testing"

	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	examplecert "github.com/s7techlab/cckit/examples/cert"
	"github.com/s7techlab/cckit/extensions/debug"
	"github.com/s7techlab/cckit/extensions/owner"
	"github.com/s7techlab/cckit/identity"
	"github.com/s7techlab/cckit/router"
	testcc "github.com/s7techlab/cckit/testing"
	expectcc "github.com/s7techlab/cckit/testing/expect"
)

func TestDebug(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Debug suite")
}

type DebuggableChaincode struct {
	router *router.Group
}

func (cc *DebuggableChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return owner.SetFromCreator(cc.router.Context(`init`, stub))
}

func (cc *DebuggableChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// delegate handling to router
	return cc.router.Handle(stub)
}

func New() *DebuggableChaincode {
	r := router.New(`debuggable`)
	debug.AddHandlers(r, `debug`, owner.Only)

	return &DebuggableChaincode{r}
}

var _ = Describe(`Debuggable`, func() {

	//Create chaincode mock
	cc := testcc.NewMockStub(`debuggable`, New())
	actors, err := identity.ActorsFromPemFile(`SOME_MSP`, map[string]string{
		`owner`: `s7techlab.pem`,
	}, examplecert.Content)
	if err != nil {
		panic(err)
	}

	owner := actors[`owner`]
	cc.From(owner).Init()

	Describe("Debug", func() {

		It("Allow to clean empty state", func() {
			emptyResult := expectcc.PayloadIs(
				cc.From(owner).Invoke(`debugStateClean`, []string{`some`, `non existent`, `keys`}), new(map[string]int)).(map[string]int)

			Expect(emptyResult[`some`]).To(Equal(0))
			Expect(len(emptyResult)).To(Equal(3))
		})

		It("Allow put value in state", func() {
			for i := 0; i < 5; i++ {
				expectcc.ResponseOk(cc.From(owner).Invoke(`debugStatePut`, []string{`prefixA`, `key` + strconv.Itoa(i)}, []byte(`value`+strconv.Itoa(i))))
			}

			for i := 0; i < 7; i++ {
				expectcc.ResponseOk(cc.From(owner).Invoke(`debugStatePut`, []string{`prefixB`, `subprefixA`, `key` + strconv.Itoa(i)}, []byte(`value`+strconv.Itoa(i))))
				expectcc.ResponseOk(cc.From(owner).Invoke(`debugStatePut`, []string{`prefixB`, `subprefixB`, `key` + strconv.Itoa(i)}, []byte(`value`+strconv.Itoa(i))))
			}

			cc.From(owner).Invoke(`debugStatePut`, []string{`keyA`}, []byte(`valueKeyA`))
			cc.From(owner).Invoke(`debugStatePut`, []string{`keyB`}, []byte(`valueKeyB`))
			cc.From(owner).Invoke(`debugStatePut`, []string{`keyC`}, []byte(`valueKeyC`))
		})

		It("Allow to get value in state", func() {
			Expect(cc.From(owner).Invoke(`debugStateGet`, []string{`prefixA`, `key1`}).Payload).To(Equal([]byte(`value1`)))
			Expect(cc.From(owner).Invoke(`debugStateGet`, []string{`keyA`}).Payload).To(Equal([]byte(`valueKeyA`)))
		})

		It("Allow to get keys", func() {
			keys := expectcc.PayloadIs(cc.From(owner).Invoke(`debugStateKeysList`, []string{`prefixA`}), &[]string{}).([]string)
			Expect(len(keys)).To(Equal(5))

			key0, key0rest, _ := cc.SplitCompositeKey(keys[0])
			Expect(key0).To(Equal(`prefixA`))
			Expect(key0rest).To(Equal([]string{`key0`}))

			keys = expectcc.PayloadIs(cc.From(owner).Invoke(`debugStateKeysList`, []string{`prefixB`, `subprefixB`}), &[]string{}).([]string)
			Expect(len(keys)).To(Equal(7))
		})

		It("Allow to delete state entry", func() {
			expectcc.ResponseOk(cc.From(owner).Invoke(`debugStateDelete`, []string{`prefixA`, `key0`}))
			keys := expectcc.PayloadIs(cc.From(owner).Invoke(`debugStateKeysList`, []string{`prefixA`}), &[]string{}).([]string)
			Expect(len(keys)).To(Equal(4))

			expectcc.ResponseOk(cc.From(owner).Invoke(`debugStateDelete`, []string{`prefixA`, `key4`}))
			keys = expectcc.PayloadIs(cc.From(owner).Invoke(`debugStateKeysList`, []string{`prefixA`}), &[]string{}).([]string)
			Expect(len(keys)).To(Equal(3))
		})

		It("Allow to clean state", func() {
			cleanResult := expectcc.PayloadIs(
				cc.From(owner).Invoke(`debugStateClean`, []string{`prefixA`}), new(map[string]int)).(map[string]int)

			Expect(cleanResult[`prefixA`]).To(Equal(3))
			Expect(len(cleanResult)).To(Equal(1))

			keys := expectcc.PayloadIs(cc.From(owner).Invoke(`debugStateKeysList`, []string{`prefixA`}), &[]string{}).([]string)
			Expect(len(keys)).To(Equal(0))
		})

	})
})