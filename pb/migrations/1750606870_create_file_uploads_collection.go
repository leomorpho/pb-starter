package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `{
			"id": "pbc_file_uploads",
			"created": "2022-01-01 00:00:00.000Z",
			"updated": "2022-01-01 00:00:00.000Z",
			"name": "file_uploads",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "file_field",
					"name": "file",
					"type": "file",
					"required": false,
					"presentable": true,
					"unique": false,
					"options": {
						"mimeTypes": [],
						"thumbs": ["100x100", "300x300"],
						"maxSelect": 1,
						"maxSize": 104857600,
						"protected": false
					}
				},
				{
					"system": false,
					"id": "upload_id",
					"name": "upload_id",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": true,
					"options": {
						"min": null,
						"max": null,
						"pattern": "^[a-zA-Z0-9+/=_-]+$"
					}
				},
				{
					"system": false,
					"id": "metadata",
					"name": "metadata",
					"type": "json",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSize": 2000000
					}
				},
				{
					"system": false,
					"id": "processing_status",
					"name": "processing_status",
					"type": "select",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"pending",
							"processing",
							"completed",
							"failed"
						]
					}
				},
				{
					"system": false,
					"id": "file_type",
					"name": "file_type",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": "^(avatar|document|media|temp)$"
					}
				},
				{
					"system": false,
					"id": "category",
					"name": "category",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": 100,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "user",
					"name": "user",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": ["email"]
					}
				},
				{
					"system": false,
					"id": "visibility",
					"name": "visibility",
					"type": "select",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"public",
							"private",
							"shared"
						]
					}
				},
				{
					"system": false,
					"id": "processed_variants",
					"name": "processed_variants",
					"type": "json",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSize": 2000000
					}
				},
				{
					"system": false,
					"id": "original_name",
					"name": "original_name",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": 255,
						"pattern": ""
					}
				}
			],
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_file_uploads_upload_id` + "`" + ` ON ` + "`" + `file_uploads` + "`" + ` (` + "`" + `upload_id` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_file_uploads_user` + "`" + ` ON ` + "`" + `file_uploads` + "`" + ` (` + "`" + `user` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_file_uploads_file_type` + "`" + ` ON ` + "`" + `file_uploads` + "`" + ` (` + "`" + `file_type` + "`" + `)"
			],
			"listRule": "@request.auth.id = user.id",
			"viewRule": "@request.auth.id = user.id || visibility = \"public\"",
			"createRule": "@request.auth.id != \"\"",
			"updateRule": "@request.auth.id = user.id",
			"deleteRule": "@request.auth.id = user.id",
			"options": {}
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("file_uploads")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
