package application

import (
	"context"

	me "github.com/octoposprime/op-be-dlr/internal/domain/model/entity"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
)

// This is the event listener handler of the application layer.
func (a *Service) EventListen() *Service {
	go a.Listen(context.Background(), smodel.ChannelCreateDlr, a.EventListenerCallBack)
	go a.Listen(context.Background(), smodel.ChannelDeleteDlr, a.EventListenerCallBack)
	return a
}

// This is a call-back function of the event listener handler of the application layer.
func (a *Service) EventListenerCallBack(channelName string, dlr me.Dlr) {
	if channelName == smodel.ChannelCreateDlr {
		a.CreateDlr(context.Background(), dlr)
	} else if channelName == smodel.ChannelDeleteDlr {
		a.DeleteDlr(context.Background(), dlr)
	} else {
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "EventListenerCallBack", channelName, smodel.ErrorChannelNameNotValid.Error()))
	}
}
