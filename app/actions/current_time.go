package actions

import (
	"context"
	"fmt"
	"my_app/app/models"
	"net/http"
	"strings"

	"github.com/go-rel/rel/where"
)

func (axn *Action) CurrentTime(w http.ResponseWriter, r *http.Request) {
	account := models.Account{}

	ctx := context.Background()

	axn.repo.Find(ctx, &account, where.Eq("id", 1))

	var res strings.Builder

	res.WriteString(fmt.Sprintln("================================================="))
	res.WriteString(fmt.Sprintln("Created At:", account.CreatedAt))
	res.WriteString(fmt.Sprintln("Updated At:", account.UpdatedAt))
	res.WriteString(fmt.Sprintln("  Nickname:", account.Nickname))
	res.WriteString(fmt.Sprintln("     Email:", account.Email))
	res.WriteString(fmt.Sprintln("================================================="))

	w.Write([]byte(res.String()))
}
