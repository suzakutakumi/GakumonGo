package repository

import "gakumon_go/model"

type GakumonRepositoryImpl struct {
}

func (r GakumonRepositoryImpl) FetchAllGakumonId() []model.GakumonId {
	gakumon_list := []model.GakumonId{ // ダミーデータ
		{
			GakumonId: "s1280149",
		}, {
			GakumonId: "s1290017",
		}}
	return gakumon_list
}
