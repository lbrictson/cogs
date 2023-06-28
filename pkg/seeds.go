package pkg

import (
	"context"
	"github.com/lbrictson/cogs/ent"
	"github.com/lbrictson/cogs/ent/schema"
)

func ExecuteDefaultSeeds(ctx context.Context, db *ent.Client) error {
	// Create a user if there isn't one
	users, err := getUsers(ctx, db)
	if err != nil {
		return err
	}
	if len(users) == 0 {
		// Seed default admin user
		_, err := createUser(ctx, db, CreateUserInput{
			Email:    "admin@localhost.com",
			Password: "ChangeMe1234!",
			Role:     "admin",
			APIKey:   generateAPIKey(),
		})
		if err != nil {
			return err
		}
		LogFromCtx(ctx).With("system", "seeder").Info("seeded default admin user")
	}
	// Create a project if there isn't one
	projects, err := getProjects(ctx, db)
	if err != nil {
		return err
	}
	if len(projects) == 0 {
		// Seed default project
		p, err := createProject(ctx, db, NewProjectInput{
			Name:        "Default",
			Description: "Default project that comes with cogs",
		})
		if err != nil {
			return err
		}
		// Create a script for the default project
		_, err = createScript(ctx, db, CreateScriptInput{
			ProjectID: p.ID,
			Name:      "Example Echo Information",
			Script: `
#!/bin/bash
echo $COGS_GREETING $COGS_NAME
echo $HOSTNAME
`,
			Description:    "An example script",
			TimeoutSeconds: 300,
			Parameters: []schema.ScriptInputOptions{
				{
					Name:          "Name",
					Description:   "Who to greet",
					StrictOptions: false,
					Options:       nil,
					VariableType:  "text",
				},
				{
					Name:          "Greeting",
					Description:   "What to say",
					StrictOptions: true,
					Options:       []string{"hello", "hi", "howdy"},
					VariableType:  "text",
				},
			},
		})
		if err != nil {
			return err
		}
		LogFromCtx(ctx).With("system", "seeder").Info("seeded default project")
	}
	return nil
}
