package handler

import (
	pb "api-gateway/genproto"
)

type Handler struct {
	Account pb.AccountServiceClient
    Budget pb.BudgetServiceClient
    Category pb.CategoryServiceClient
	Goal pb.GoalServiceClient
	Transaction pb.TransactionServiceClient
}

func NewHandler(account pb.AccountServiceClient, budget pb.BudgetServiceClient, category pb.CategoryServiceClient, goal pb.GoalServiceClient, transaction pb.TransactionServiceClient) *Handler {
	return &Handler{Account: account, Budget: budget, Category : category, Goal : goal, Transaction : transaction}
}
