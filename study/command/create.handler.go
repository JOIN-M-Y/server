package command

import (
	"github.com/JOIN-M-Y/server/study/model"
)

func (bus *Bus) handleCreateCommand(
	command *CreateStudyCommand,
) (*model.Study, error) {
	// uuid, _ := uuid.NewRandom()
	return &model.Study{}, nil
}
