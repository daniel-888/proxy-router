package msgbus

import (
	"context"
	"fmt"

	"gitlab.com/TitanInd/lumerin/lumerinlib"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
	"gitlab.com/TitanInd/lumerin/msgbus"
)

type ValidateID IDString

type Validate struct {
	ID      ValidateID // "Random ID value"
	MinerID string
	DestID  string
	Data    interface{}
}

type Submit struct {
	JobID     string
	Extraonce string
	NTime     string
	NOnce     string
}

type Notify struct {
	JobID           string
	PrevBlockHash   string
	GenTransaction1 string
	GenTransaction2 string
	MerkelBranches  []string
	Version         string
	Nbits           string
	Ntime           string
	Clean           bool
}

type SetDifficulty struct {
	Diff int
}

func NewValidate(id ValidateID, minerID string, destID string, data interface{}) (v *Validate) {

	v = &Validate{
		ID:      id,
		MinerID: minerID,
		DestID:  destID,
		Data:    data,
	}

	return v
}

func NewSubmit(jobid string, extranonce string, ntime string, nonce string) (s *Submit) {

	s = &Submit{
		JobID:     jobid,
		Extraonce: extranonce,
		NTime:     ntime,
		NOnce:     nonce,
	}

	return s
}

func NewNotify(jobid string, prevblock string, gen1 string, gen2 string, merkel []string, version string, nbits string, ntime string, clean bool) (n *Notify) {

	n = &Notify{

		JobID:           jobid,
		PrevBlockHash:   prevblock,
		GenTransaction1: gen1,
		GenTransaction2: gen2,
		MerkelBranches:  merkel,
		Version:         version,
		Nbits:           nbits,
		Ntime:           ntime,
		Clean:           clean,
	}

	return n
}

func NewSetDiff(diff int) (d *SetDifficulty) {

	d = &SetDifficulty{
		Diff: diff,
	}

	return d
}

//
//
//
func SendValidateSubmit() {

}

//
//
//
func SendValidateNotify() {

}

//
//
//
func SendValidateSetDiff(ctx context.Context, ps *PubSub, m MinerID, d DestID, diff int) {

	id := fmt.Sprintf("SubmitID:%d", <-SubmitCountChan)
	setdiff := msgbus.NewSetDiff(diff)
	validate := msgbus.NewValidate(msgbus.ValidateID(id), string(m), string(d), setdiff)

	_, e = ps.Pub(ValidateMsg, IDString(id), validate)
	if e != nil {
		contextlib.Logf(ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Pub() error:%s", e)
	}
}
