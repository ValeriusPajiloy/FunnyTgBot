package tagall

import (
	"context"
	"fmt"
	mongodb "tgbot/internal/storage/mongo_db"

	client "tgbot/internal/clients"
)

type TagAll struct {
	tg   client.Client
	repo *mongodb.Repository
}

func NewTagAll(tg client.Client, repo *mongodb.Repository) *TagAll {
	return &TagAll{
		tg:   tg,
		repo: repo,
	}
}

func (t *TagAll) CreateAnswer(ctx context.Context, hasKeys bool, keys []string, chatId int) {
	if hasKeys {
		if len(keys) < 2 {
			switch keys[0] {
			case "help":
				t.tg.SendMessage(chatId, MSG_HELP_TEG_ALL)
			case "groups":
				groupListString := t.groupListString(ctx, chatId)
				t.tg.SendMessage(chatId, groupListString)
			default:
				groupList := t.groupList(ctx, chatId)
				if contains(*groupList, keys[0]) {
					tagAll := t.tagListString(ctx, keys[0], chatId)
					t.tg.SendMessage(chatId, tagAll)
				} else {
					t.tg.SendMessage(chatId, MSG_UNDEFIND_GROUP_TEG_ALL)
				}
			}
		} else {
			switch keys[0] {
			case "addgroup":
				newGroup := keys[1]
				t.createGroup(ctx, newGroup, chatId)
			case "add":
				if len(keys) > 2 {
					gropName := keys[1]
					for i := 2; i < len(keys); i++ {
						userTag := keys[i]
						t.addUserInGroup(ctx, userTag, gropName, chatId)
					}
				} else {
					t.tg.SendMessage(chatId, MSG_EMPTY_TEG_ALL)
				}
			default:
				t.tg.SendMessage(chatId, MSG_EMPTY_TEG_ALL)
			}
		}
	} else {
		t.tg.SendMessage(chatId, MSG_EMPTY_TEG_ALL)
	}
}
func (t *TagAll) groupListString(ctx context.Context, chatId int) string {
	groupListString := "Пока нет доступных групп"
	groupList, err := t.repo.TagGroup.GetAll(ctx, chatId)
	if err != nil {
		return groupListString
	}
	if len(groupList) > 0 {
		groupListString = "Доступные группы:\n"
		for _, group := range groupList {
			groupListString += fmt.Sprintf("%s\n", group)
		}
	}
	return groupListString
}
func (t *TagAll) groupList(ctx context.Context, chatId int) *[]string {
	groupList, err := t.repo.TagGroup.GetAll(ctx, chatId)
	if err != nil {
		return &[]string{}
	}
	return &groupList
}
func (t *TagAll) createGroup(ctx context.Context, newGroup string, chatId int) {
	t.repo.TagGroup.Create(ctx, newGroup, chatId)
}

func (t *TagAll) tagListString(ctx context.Context, group string, chatId int) string {
	tagListString := "В группе пусто"
	tagList, err := t.repo.TagUser.GetAllForGroup(ctx, group, chatId)
	if err != nil {
		return tagListString
	}
	if len(tagList) > 0 {
		tagListString = fmt.Sprintf("го в %s", group)
		for _, tag := range tagList {
			tagListString += fmt.Sprintf(" %s", tag)
		}
	}
	return tagListString
}
func (t *TagAll) addUserInGroup(ctx context.Context, userTag string, gropName string, chatId int) {
	t.repo.TagUser.Create(ctx, userTag, gropName, chatId)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
