package pkg

import (
	"context"
	"errors"
	"github.com/lbrictson/cogs/ent"
	"github.com/lbrictson/cogs/ent/access"
	"github.com/lbrictson/cogs/ent/history"
	"github.com/lbrictson/cogs/ent/project"
	"github.com/lbrictson/cogs/ent/schema"
	"github.com/lbrictson/cogs/ent/script"
	"github.com/lbrictson/cogs/ent/scriptstats"
	"github.com/lbrictson/cogs/ent/secret"
	"github.com/lbrictson/cogs/ent/user"
	"math"
	"regexp"
	"strings"
	"time"
)

func convertEntUserToUserModel(entUser *ent.User) UserModel {
	if entUser == nil {
		return UserModel{}
	}
	return UserModel{
		ID:             entUser.ID,
		Email:          entUser.Email,
		Role:           entUser.Role.String(),
		APIKey:         entUser.APIKey,
		CreatedAt:      entUser.CreatedAt,
		UpdatedAt:      entUser.UpdatedAt,
		HashedPassword: entUser.HashedPassword,
	}
}

type CreateUserInput struct {
	Email    string
	Password string
	Role     string
	APIKey   string
}

func createUser(ctx context.Context, db *ent.Client, input CreateUserInput) (UserModel, error) {
	u, err := db.User.Create().
		SetEmail(input.Email).
		SetHashedPassword(hashAndSaltPassword(input.Password)).
		SetRole(user.Role(input.Role)).
		SetAPIKey(input.APIKey).
		Save(ctx)
	if err != nil {
		return UserModel{}, err
	}
	return convertEntUserToUserModel(u), nil
}

func getUserByEmail(ctx context.Context, db *ent.Client, email string) (UserModel, error) {
	u, err := db.User.Query().Where(user.EmailEQ(email)).First(ctx)
	if err != nil {
		return UserModel{}, err
	}
	return convertEntUserToUserModel(u), nil
}

func getUserByID(ctx context.Context, db *ent.Client, id int) (UserModel, error) {
	u, err := db.User.Query().Where(user.IDEQ(id)).First(ctx)
	if err != nil {
		return UserModel{}, err
	}
	return convertEntUserToUserModel(u), nil
}

func getUsers(ctx context.Context, db *ent.Client) ([]UserModel, error) {
	users, err := db.User.Query().All(ctx)
	if err != nil {
		return []UserModel{}, err
	}
	var userModels []UserModel
	for _, user := range users {
		userModels = append(userModels, convertEntUserToUserModel(user))
	}
	return userModels, nil
}

type UpdateUserInput struct {
	Password *string
	Role     *string
}

func updateUser(ctx context.Context, db *ent.Client, id int, input UpdateUserInput) (UserModel, error) {
	builder := db.User.Update()
	if input.Password != nil {
		builder.SetHashedPassword(hashAndSaltPassword(*input.Password))
	}
	if input.Role != nil {
		builder.SetRole(user.Role(*input.Role))
	}
	_, err := builder.Where(user.IDEQ(id)).Save(ctx)
	if err != nil {
		return UserModel{}, err
	}
	return getUserByID(ctx, db, id)
}

func deleteUser(ctx context.Context, db *ent.Client, id int) error {
	return db.User.DeleteOneID(id).Exec(ctx)
}

func convertEntProjectToProjectModel(entProject *ent.Project) ProjectModel {
	if entProject == nil {
		return ProjectModel{}
	}
	return ProjectModel{
		ID:          entProject.ID,
		Name:        entProject.Name,
		Description: entProject.Description,
		CreatedAt:   entProject.CreatedAt,
		UpdatedAt:   entProject.UpdatedAt,
	}
}

type NewProjectInput struct {
	Name        string
	Description string
}

func createProject(ctx context.Context, db *ent.Client, input NewProjectInput) (ProjectModel, error) {
	p, err := db.Project.Create().
		SetName(input.Name).
		SetDescription(input.Description).
		Save(ctx)
	if err != nil {
		return ProjectModel{}, err
	}
	return convertEntProjectToProjectModel(p), nil
}

func getProjectByID(ctx context.Context, db *ent.Client, id int) (ProjectModel, error) {
	p, err := db.Project.Query().Where(project.IDEQ(id)).First(ctx)
	if err != nil {
		return ProjectModel{}, err
	}
	return convertEntProjectToProjectModel(p), nil
}

func getProjects(ctx context.Context, db *ent.Client) ([]ProjectModel, error) {
	projects, err := db.Project.Query().All(ctx)
	if err != nil {
		return []ProjectModel{}, err
	}
	var projectModels []ProjectModel
	for _, project := range projects {
		projectModels = append(projectModels, convertEntProjectToProjectModel(project))
	}
	return projectModels, nil
}

type UpdateProjectInput struct {
	Name        *string
	Description *string
}

func updateProject(ctx context.Context, db *ent.Client, id int, input UpdateProjectInput) (ProjectModel, error) {
	builder := db.Project.Update()
	if input.Name != nil {
		builder.SetName(*input.Name)
	}
	if input.Description != nil {
		builder.SetDescription(*input.Description)
	}
	_, err := builder.Where(project.IDEQ(id)).Save(ctx)
	if err != nil {
		return ProjectModel{}, err
	}
	return getProjectByID(ctx, db, id)
}

func deleteProject(ctx context.Context, db *ent.Client, id int) error {
	return db.Project.DeleteOneID(id).Exec(ctx)
}

func convertEntScriptToScriptModel(entScript *ent.Script) ScriptModel {
	if entScript == nil {
		return ScriptModel{}
	}
	return ScriptModel{
		ID:             entScript.ID,
		Name:           entScript.Name,
		Script:         entScript.Script,
		ProjectID:      entScript.ProjectID,
		Description:    entScript.Description,
		TimeoutSeconds: entScript.TimeoutSeconds,
		Parameters:     entScript.Parameters,
		CreatedAt:      entScript.CreatedAt,
		UpdatedAt:      entScript.UpdatedAt,
	}
}

type CreateScriptInput struct {
	ProjectID      int
	Name           string
	Script         string
	Description    string
	TimeoutSeconds int
	Parameters     []schema.ScriptInputOptions
}

func createScript(ctx context.Context, db *ent.Client, input CreateScriptInput) (ScriptModel, error) {
	p, err := db.Project.Query().Where(project.IDEQ(input.ProjectID)).First(ctx)
	if err != nil {
		return ScriptModel{}, err
	}
	s, err := db.Script.Create().
		SetName(input.Name).
		SetScript(input.Script).
		SetDescription(input.Description).
		SetTimeoutSeconds(input.TimeoutSeconds).
		SetProjectID(p.ID).
		SetParameters(input.Parameters).
		Save(ctx)
	if err != nil {
		return ScriptModel{}, err
	}
	// Always create a script stat to go with each script
	_, err = createScriptStats(ctx, db, CreateScriptStatsInput{
		ScriptID:  s.ID,
		ProjectID: p.ID,
	})
	return convertEntScriptToScriptModel(s), err
}

func getScriptByID(ctx context.Context, db *ent.Client, id int) (ScriptModel, error) {
	s, err := db.Script.Query().Where(script.IDEQ(id)).First(ctx)
	if err != nil {
		return ScriptModel{}, err
	}
	return convertEntScriptToScriptModel(s), nil
}

func getScripts(ctx context.Context, db *ent.Client) ([]ScriptModel, error) {
	scripts, err := db.Script.Query().All(ctx)
	if err != nil {
		return []ScriptModel{}, err
	}
	var scriptModels []ScriptModel
	for _, script := range scripts {
		scriptModels = append(scriptModels, convertEntScriptToScriptModel(script))
	}
	return scriptModels, nil
}

type UpdateScriptInput struct {
	Name           *string
	Script         *string
	Description    *string
	TimeoutSeconds *int
	Parameters     *[]schema.ScriptInputOptions
}

func updateScript(ctx context.Context, db *ent.Client, id int, input UpdateScriptInput) (ScriptModel, error) {
	builder := db.Script.Update()
	if input.Name != nil {
		builder.SetName(*input.Name)
	}
	if input.Script != nil {
		builder.SetScript(*input.Script)
	}
	if input.Description != nil {
		builder.SetDescription(*input.Description)
	}
	if input.TimeoutSeconds != nil {
		builder.SetTimeoutSeconds(*input.TimeoutSeconds)
	}
	if input.Parameters != nil {
		builder.SetParameters(*input.Parameters)
	}
	_, err := builder.Where(script.IDEQ(id)).Save(ctx)
	if err != nil {
		return ScriptModel{}, err
	}
	return getScriptByID(ctx, db, id)
}

func deleteScript(ctx context.Context, db *ent.Client, id int) error {
	return db.Script.DeleteOneID(id).Exec(ctx)
}

func getProjectScripts(ctx context.Context, db *ent.Client, projectID int) ([]ScriptModel, error) {
	scripts, err := db.Script.Query().Where(script.ProjectIDEQ(projectID)).All(ctx)
	if err != nil {
		return []ScriptModel{}, err
	}
	var scriptModels []ScriptModel
	for _, script := range scripts {
		scriptModels = append(scriptModels, convertEntScriptToScriptModel(script))
	}
	return scriptModels, nil
}

func convertEntSecretToSecretModel(entSecret *ent.Secret) SecretModel {
	if entSecret == nil {
		return SecretModel{}
	}
	return SecretModel{
		ID:        entSecret.ID,
		Name:      entSecret.Name,
		Value:     entSecret.Value,
		CreatedAt: entSecret.CreatedAt,
		UpdatedAt: entSecret.UpdatedAt,
	}
}

type CreateSecretInput struct {
	Name      string
	Value     string
	CreatedBy string
	ProjectID int
}

func validateSecretName(name string) error {
	if len(name) == 0 {
		return errors.New("name cannot be empty")
	}
	if len(name) > 255 {
		return errors.New("name cannot be longer than 255 characters")
	}
	// No spaces
	if strings.Contains(name, " ") {
		return errors.New("name cannot contain spaces")
	}
	// No special characters
	if !regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(name) {
		return errors.New("name can only contain alphanumeric characters")
	}
	return nil
}

func createSecret(ctx context.Context, db *ent.Client, input CreateSecretInput) (SecretModel, error) {
	if err := validateSecretName(input.Name); err != nil {
		return SecretModel{}, err
	}
	s, err := db.Secret.Create().
		SetName(input.Name).
		SetValue(input.Value).
		SetProjectID(input.ProjectID).
		SetCreatedBy(input.CreatedBy).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return SecretModel{}, err
	}
	return convertEntSecretToSecretModel(s), nil
}

func getSecretByID(ctx context.Context, db *ent.Client, id int) (SecretModel, error) {
	s, err := db.Secret.Query().Where(secret.IDEQ(id)).First(ctx)
	if err != nil {
		return SecretModel{}, err
	}
	return convertEntSecretToSecretModel(s), nil
}

func getSecrets(ctx context.Context, db *ent.Client) ([]SecretModel, error) {
	secrets, err := db.Secret.Query().All(ctx)
	if err != nil {
		return []SecretModel{}, err
	}
	var secretModels []SecretModel
	for _, secret := range secrets {
		secretModels = append(secretModels, convertEntSecretToSecretModel(secret))
	}
	return secretModels, nil
}

type UpdateSecretInput struct {
	Name  *string
	Value *string
}

func updateSecret(ctx context.Context, db *ent.Client, id int, input UpdateSecretInput) (SecretModel, error) {
	builder := db.Secret.Update()
	if input.Name != nil {
		if err := validateSecretName(*input.Name); err != nil {
			return SecretModel{}, err
		}
		builder.SetName(*input.Name)
	}
	if input.Value != nil {
		builder.SetValue(*input.Value)
	}
	_, err := builder.Where(secret.IDEQ(id)).Save(ctx)
	if err != nil {
		return SecretModel{}, err
	}
	return getSecretByID(ctx, db, id)
}

func deleteSecret(ctx context.Context, db *ent.Client, id int) error {
	return db.Secret.DeleteOneID(id).Exec(ctx)
}

func getProjectSecrets(ctx context.Context, db *ent.Client, projectID int) ([]SecretModel, error) {
	secrets, err := db.Secret.Query().Where(secret.ProjectIDEQ(projectID)).All(ctx)
	if err != nil {
		return []SecretModel{}, err
	}
	var secretModels []SecretModel
	for _, secret := range secrets {
		secretModels = append(secretModels, convertEntSecretToSecretModel(secret))
	}
	return secretModels, nil
}

func convertEntHistoryToHistoryModel(entHistory *ent.History) HistoryModel {
	if entHistory == nil {
		return HistoryModel{}
	}
	return HistoryModel{
		ID:          entHistory.ID,
		ScriptID:    entHistory.ScriptID,
		Success:     entHistory.Success,
		Output:      entHistory.Output,
		Duration:    entHistory.Duration,
		Trigger:     entHistory.Trigger,
		TriggeredBy: entHistory.TriggeredBy,
		RunID:       entHistory.RunID,
		ExitCode:    entHistory.ExitCode,
		CreatedAt:   entHistory.CreatedAt,
		UpdatedAt:   entHistory.UpdatedAt,
		Arguments:   entHistory.Arguments,
		Status:      entHistory.Status,
	}
}

type CreateHistoryInput struct {
	ScriptID        int
	Success         bool
	Output          string
	DurationSeconds int
	Trigger         string
	TriggeredBy     string
	RunID           string
	ExitCode        int
	Status          string
	Arguments       map[string]string
}

func createHistory(ctx context.Context, db *ent.Client, input CreateHistoryInput) (HistoryModel, error) {
	s, err := db.Script.Query().Where(script.IDEQ(input.ScriptID)).First(ctx)
	if err != nil {
		return HistoryModel{}, err
	}
	h, err := db.History.Create().
		SetSuccess(input.Success).
		SetOutput(input.Output).
		SetDuration(input.DurationSeconds).
		SetTrigger(input.Trigger).
		SetTriggeredBy(input.TriggeredBy).
		SetRunID(input.RunID).
		SetExitCode(input.ExitCode).
		SetScriptID(s.ID).
		SetArguments(input.Arguments).
		SetStatus(input.Status).
		Save(ctx)
	if err != nil {
		return HistoryModel{}, err
	}
	return convertEntHistoryToHistoryModel(h), nil
}

type UpdateHistoryInput struct {
	Success         *bool
	Output          *string
	DurationSeconds *int
	Trigger         *string
	TriggeredBy     *string
	RunID           *string
	ExitCode        *int
	Status          *string
	Arguments       *map[string]string
}

func updateHistory(ctx context.Context, db *ent.Client, id int, input UpdateHistoryInput) (HistoryModel, error) {
	builder := db.History.Update()
	if input.Success != nil {
		builder.SetSuccess(*input.Success)
	}
	if input.Output != nil {
		builder.SetOutput(*input.Output)
	}
	if input.DurationSeconds != nil {
		builder.SetDuration(*input.DurationSeconds)
	}
	if input.Trigger != nil {
		builder.SetTrigger(*input.Trigger)
	}
	if input.TriggeredBy != nil {
		builder.SetTriggeredBy(*input.TriggeredBy)
	}
	if input.RunID != nil {
		builder.SetRunID(*input.RunID)
	}
	if input.ExitCode != nil {
		builder.SetExitCode(*input.ExitCode)
	}
	if input.Status != nil {
		builder.SetStatus(*input.Status)
	}
	if input.Arguments != nil {
		builder.SetArguments(*input.Arguments)
	}
	_, err := builder.Where(history.IDEQ(id)).Save(ctx)
	if err != nil {
		return HistoryModel{}, err
	}
	return getHistoryByID(ctx, db, id)
}

func getHistoryByID(ctx context.Context, db *ent.Client, id int) (HistoryModel, error) {
	h, err := db.History.Query().Where(history.IDEQ(id)).First(ctx)
	if err != nil {
		return HistoryModel{}, err
	}
	return convertEntHistoryToHistoryModel(h), nil
}

func getHistories(ctx context.Context, db *ent.Client) ([]HistoryModel, error) {
	histories, err := db.History.Query().All(ctx)
	if err != nil {
		return []HistoryModel{}, err
	}
	var historyModels []HistoryModel
	for _, history := range histories {
		historyModels = append(historyModels, convertEntHistoryToHistoryModel(history))
	}
	return historyModels, nil
}

type QueryHistoriesInput struct {
	Limit       int
	Offset      int
	ScriptID    *int
	Trigger     *string
	Status      *string
	Duration    *int
	Success     *bool
	TriggeredBy *string
}

func queryHistories(ctx context.Context, db *ent.Client, input QueryHistoriesInput) ([]HistoryModel, error) {
	builder := db.History.Query().Offset(input.Offset).Limit(input.Limit)
	if input.ScriptID != nil {
		builder = builder.Where(history.ScriptIDEQ(*input.ScriptID))
	}
	if input.Trigger != nil {
		builder = builder.Where(history.TriggerEQ(*input.Trigger))
	}
	if input.Status != nil {
		builder = builder.Where(history.StatusEQ(*input.Status))
	}
	if input.Duration != nil {
		builder = builder.Where(history.DurationEQ(*input.Duration))
	}
	if input.Success != nil {
		builder = builder.Where(history.SuccessEQ(*input.Success))
	}
	if input.TriggeredBy != nil {
		builder = builder.Where(history.TriggeredByEQ(*input.TriggeredBy))
	}
	histories, err := builder.All(ctx)
	if err != nil {
		return []HistoryModel{}, err
	}
	var historyModels []HistoryModel
	for _, history := range histories {
		historyModels = append(historyModels, convertEntHistoryToHistoryModel(history))
	}
	return historyModels, nil

}

func getScriptHistories(ctx context.Context, db *ent.Client, scriptID int, limit int, offset int) ([]HistoryModel, error) {
	histories, err := db.History.Query().Where(history.ScriptIDEQ(scriptID)).Order(ent.Desc(history.FieldCreatedAt)).Limit(limit).Offset(offset).All(ctx)
	if err != nil {
		return []HistoryModel{}, err
	}
	var historyModels []HistoryModel
	for _, history := range histories {
		historyModels = append(historyModels, convertEntHistoryToHistoryModel(history))
	}
	return historyModels, nil
}

func deleteHistory(ctx context.Context, db *ent.Client, id int) error {
	return db.History.DeleteOneID(id).Exec(ctx)
}

func convertEntAccessToAccessModel(entAccess *ent.Access) AccessModel {
	if entAccess == nil {
		return AccessModel{}
	}
	return AccessModel{
		ID:        entAccess.ID,
		ProjectID: entAccess.ProjectID,
		UserID:    entAccess.UserID,
		CreatedAt: entAccess.CreatedAt,
		UpdatedAt: entAccess.UpdatedAt,
		Role:      entAccess.Role.String(),
	}
}

type CreateAccessInput struct {
	ProjectID int
	UserID    int
	Role      string
}

func createAccess(ctx context.Context, db *ent.Client, input CreateAccessInput) (AccessModel, error) {
	a, err := db.Access.Create().
		SetProjectID(input.ProjectID).
		SetUserID(input.UserID).
		SetRole(access.Role(input.Role)).
		Save(ctx)
	if err != nil {
		return AccessModel{}, err
	}
	return convertEntAccessToAccessModel(a), nil
}

func getAccessByID(ctx context.Context, db *ent.Client, id int) (AccessModel, error) {
	a, err := db.Access.Query().Where(access.IDEQ(id)).First(ctx)
	if err != nil {
		return AccessModel{}, err
	}
	return convertEntAccessToAccessModel(a), nil
}

func getAccesses(ctx context.Context, db *ent.Client) ([]AccessModel, error) {
	accesses, err := db.Access.Query().All(ctx)
	if err != nil {
		return []AccessModel{}, err
	}
	var accessModels []AccessModel
	for _, access := range accesses {
		accessModels = append(accessModels, convertEntAccessToAccessModel(access))
	}
	return accessModels, nil
}

func getProjectAccesses(ctx context.Context, db *ent.Client, projectID int) ([]AccessModel, error) {
	accesses, err := db.Access.Query().Where(access.ProjectIDEQ(projectID)).All(ctx)
	if err != nil {
		return []AccessModel{}, err
	}
	var accessModels []AccessModel
	for _, access := range accesses {
		accessModels = append(accessModels, convertEntAccessToAccessModel(access))
	}
	return accessModels, nil
}

func getUserAccesses(ctx context.Context, db *ent.Client, userID int) ([]AccessModel, error) {
	accesses, err := db.Access.Query().Where(access.UserIDEQ(userID)).All(ctx)
	if err != nil {
		return []AccessModel{}, err
	}
	var accessModels []AccessModel
	for _, access := range accesses {
		accessModels = append(accessModels, convertEntAccessToAccessModel(access))
	}
	return accessModels, nil
}

func deleteAccess(ctx context.Context, db *ent.Client, id int) error {
	return db.Access.DeleteOneID(id).Exec(ctx)
}

func convertEntScriptStatsToScriptStatsModel(entScriptStats *ent.ScriptStats) ScriptStatsModel {
	if entScriptStats == nil {
		return ScriptStatsModel{}
	}
	successRate := 100.0
	if entScriptStats.TotalRuns > 0 {
		successRate = float64(entScriptStats.TotalSuccess) / float64(entScriptStats.TotalRuns) * 100
	}
	// round to two decimals
	successRate = math.Round(successRate*100) / 100
	return ScriptStatsModel{
		ID:              entScriptStats.ID,
		ScriptID:        entScriptStats.ScriptID,
		ProjectID:       entScriptStats.ProjectID,
		TotalRuns:       entScriptStats.TotalRuns,
		TotalSuccess:    entScriptStats.TotalSuccess,
		TotalError:      entScriptStats.TotalErrors,
		AverageDuration: entScriptStats.AverageRuntime,
		LastRun:         entScriptStats.LastRun,
		LongestRun:      entScriptStats.MaxRuntime,
		ShortestRun:     entScriptStats.MinRuntime,
		TotalDuration:   entScriptStats.TotalRuntime,
		SuccessRate:     successRate,
	}
}

type CreateScriptStatsInput struct {
	ScriptID  int
	ProjectID int
}

func createScriptStats(ctx context.Context, db *ent.Client, input CreateScriptStatsInput) (ScriptStatsModel, error) {
	s, err := db.ScriptStats.Create().
		SetScriptID(input.ScriptID).
		SetProjectID(input.ProjectID).
		SetMinRuntime(0).
		SetMaxRuntime(0).
		SetAverageRuntime(0).
		SetTotalErrors(0).
		SetTotalSuccess(0).
		SetTotalRuns(0).
		SetLastRun(time.Now()).
		SetTotalRuntime(0).
		SetSuccessRate(100).
		Save(ctx)
	if err != nil {
		return ScriptStatsModel{}, err
	}
	return convertEntScriptStatsToScriptStatsModel(s), nil
}

func getScriptStatsByScriptID(ctx context.Context, db *ent.Client, scriptID int) (ScriptStatsModel, error) {
	s, err := db.ScriptStats.Query().Where(scriptstats.ScriptIDEQ(scriptID)).First(ctx)
	if err != nil {
		return ScriptStatsModel{}, err
	}
	return convertEntScriptStatsToScriptStatsModel(s), nil
}

type UpdateScriptStatsInput struct {
	ScriptID         int
	IncrementSuccess bool
	IncrementError   bool
	DurationSeconds  int
}

func updateScriptStats(ctx context.Context, db *ent.Client, input UpdateScriptStatsInput) (ScriptStatsModel, error) {
	s, err := db.ScriptStats.Query().Where(scriptstats.ScriptIDEQ(input.ScriptID)).First(ctx)
	if err != nil {
		return ScriptStatsModel{}, err
	}
	builder := db.ScriptStats.UpdateOneID(s.ID)
	if input.IncrementSuccess {
		builder = builder.SetTotalSuccess(s.TotalSuccess + 1)
	}
	if input.IncrementError {
		builder = builder.SetTotalErrors(s.TotalErrors + 1)
	}
	if input.DurationSeconds > s.MaxRuntime {
		builder = builder.SetMaxRuntime(input.DurationSeconds)
	}
	if input.DurationSeconds < s.MinRuntime {
		builder = builder.SetMinRuntime(input.DurationSeconds)
	}
	builder = builder.SetTotalRuns(s.TotalRuns + 1)
	builder = builder.SetTotalRuntime(s.TotalRuntime + input.DurationSeconds)
	builder = builder.SetAverageRuntime((s.TotalRuntime + input.DurationSeconds) / (s.TotalRuns + 1))
	builder = builder.SetLastRun(time.Now())
	if input.IncrementError {
		builder = builder.SetSuccessRate(s.TotalSuccess / (s.TotalRuns + 1) * 100)
	} else {
		builder = builder.SetSuccessRate((s.TotalSuccess + 1) / (s.TotalRuns + 1) * 100)
	}
	s, err = builder.Save(ctx)
	return convertEntScriptStatsToScriptStatsModel(s), nil
}
