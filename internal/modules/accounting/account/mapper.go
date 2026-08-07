package account

func ToAccountResponse(a *Account) AccountResponse {
	if a == nil {
		return AccountResponse{}
	}
	return AccountResponse{
		ID:            a.ID,
		Code:          a.Code,
		Name:          a.Name,
		Type:          a.Type,
		NormalBalance: a.NormalBalance,
		ParentID:      a.ParentID,
		Currency:      a.Currency,
		IsActive:      a.IsActive,
		IsSystem:      a.IsSystem,
		Description:   a.Description,
		CreatedAt:     a.CreatedAt,
		UpdatedAt:     a.UpdatedAt,
	}
}

func ToAccountResponseList(accounts []Account) []AccountResponse {
	res := make([]AccountResponse, len(accounts))
	for i, a := range accounts {
		res[i] = ToAccountResponse(&a)
	}
	return res
}

func ToAccountTreeResponse(a *Account) AccountTreeResponse {
	if a == nil {
		return AccountTreeResponse{}
	}
	tree := AccountTreeResponse{
		AccountResponse: ToAccountResponse(a),
	}
	if len(a.Children) > 0 {
		tree.Children = make([]AccountTreeResponse, len(a.Children))
		for i, child := range a.Children {
			tree.Children[i] = ToAccountTreeResponse(&child)
		}
	}
	return tree
}
