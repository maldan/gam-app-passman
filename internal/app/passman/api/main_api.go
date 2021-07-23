package api

import (
	"github.com/maldan/gam-app-passman/internal/app/passman/core"
	"github.com/maldan/go-cmhp/cmhp_crypto"
	"github.com/maldan/go-cmhp/cmhp_file"
)

type MainApi struct {
}

func (r MainApi) GetIndex() string {
	return "Test"
}

// Add new item
func (r MainApi) PostItem(args core.Item) {
	args.Id = cmhp_crypto.UID(10)
	cmhp_file.WriteJSON(core.DataDir+"/"+"item/"+args.Id+".json", &args)
}

// Update item
func (r MainApi) PatchItem(args core.Item) {
	cmhp_file.WriteJSON(core.DataDir+"/"+"item/"+args.Id+".json", &args)
}

// Delete item
func (r MainApi) DeleteItem(args core.Item) {
	cmhp_file.Delete(core.DataDir + "/" + "item/" + args.Id + ".json")
}

// Get item list
func (r MainApi) GetList() []core.Item {
	items := make([]core.Item, 0)
	files, _ := cmhp_file.List(core.DataDir + "/item")
	for _, file := range files {
		var i core.Item
		cmhp_file.ReadJSON(core.DataDir+"/item/"+file.Name(), &i)
		items = append(items, i)
	}
	return items
}
