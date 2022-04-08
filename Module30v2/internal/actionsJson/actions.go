package actionsJson

type MakeFriends struct {
	SourceId string `json:"source_id"`
	TargetId string `json:"target_id"`
}

type UpdateUser struct {
	NewAge string `json:"new_age"`
}
