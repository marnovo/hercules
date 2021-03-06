package hercules

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func fixtureDaysSinceStart() *DaysSinceStart {
	dss := DaysSinceStart{}
	dss.Initialize(testRepository)
	return &dss
}

func TestDaysSinceStartMeta(t *testing.T) {
	dss := fixtureDaysSinceStart()
	assert.Equal(t, dss.Name(), "DaysSinceStart")
	assert.Equal(t, len(dss.Provides()), 1)
	assert.Equal(t, dss.Provides()[0], "day")
	assert.Equal(t, len(dss.Requires()), 0)
}

func TestDaysSinceStartFinalize(t *testing.T) {
	dss := fixtureDaysSinceStart()
	r := dss.Finalize()
	assert.Nil(t, r)
}

func TestDaysSinceStartConsume(t *testing.T) {
	dss := fixtureDaysSinceStart()
	deps := map[string]interface{}{}
	commit, _ := testRepository.CommitObject(plumbing.NewHash(
		"cce947b98a050c6d356bc6ba95030254914027b1"))
	deps["commit"] = commit
	deps["index"] = 0
	res, err := dss.Consume(deps)
	assert.Nil(t, err)
	assert.Equal(t, res["day"].(int), 0)
	assert.Equal(t, dss.previousDay, 0)

	commit, _ = testRepository.CommitObject(plumbing.NewHash(
		"fc9ceecb6dabcb2aab60e8619d972e8d8208a7df"))
	deps["commit"] = commit
	deps["index"] = 10
	res, err = dss.Consume(deps)
	assert.Nil(t, err)
	assert.Equal(t, res["day"].(int), 1)
	assert.Equal(t, dss.previousDay, 1)

	commit, _ = testRepository.CommitObject(plumbing.NewHash(
		"a3ee37f91f0d705ec9c41ae88426f0ae44b2fbc3"))
	deps["commit"] = commit
	deps["index"] = 20
	res, err = dss.Consume(deps)
	assert.Nil(t, err)
	assert.Equal(t, res["day"].(int), 1)
	assert.Equal(t, dss.previousDay, 1)

	commit, _ = testRepository.CommitObject(plumbing.NewHash(
		"a8b665a65d7aced63f5ba2ff6d9b71dac227f8cf"))
	deps["commit"] = commit
	deps["index"] = 20
	res, err = dss.Consume(deps)
	assert.Nil(t, err)
	assert.Equal(t, res["day"].(int), 2)
	assert.Equal(t, dss.previousDay, 2)

	commit, _ = testRepository.CommitObject(plumbing.NewHash(
		"186ff0d7e4983637bb3762a24d6d0a658e7f4712"))
	deps["commit"] = commit
	deps["index"] = 30
	res, err = dss.Consume(deps)
	assert.Nil(t, err)
	assert.Equal(t, res["day"].(int), 2)
	assert.Equal(t, dss.previousDay, 2)
}
