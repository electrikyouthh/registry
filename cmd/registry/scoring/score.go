// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package scoring

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/apigee/registry/cmd/registry/core"
	"github.com/apigee/registry/cmd/registry/patch"
	"github.com/apigee/registry/cmd/registry/patterns"
	"github.com/apigee/registry/connection"
	"github.com/apigee/registry/log"
	"github.com/apigee/registry/rpc"
	"github.com/apigee/registry/server/registry/names"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func scoreID(definitionID string) string {
	return fmt.Sprintf("score-%s", definitionID)
}

func FetchScoreDefinitions(
	ctx context.Context,
	client connection.Client,
	resource patterns.ResourceName) ([]*rpc.Artifact, error) {
	defArtifacts := make([]*rpc.Artifact, 0)

	project := fmt.Sprintf("%s/locations/global", resource.Project())
	artifact, err := names.ParseArtifact(fmt.Sprintf("%s/artifacts/-", project))
	if err != nil {
		return nil, err
	}
	listFilter := fmt.Sprintf("mime_type == %q", patch.MimeTypeForKind("ScoreDefinition"))
	err = core.ListArtifacts(ctx, client, artifact, listFilter, true,
		func(artifact *rpc.Artifact) error {
			definition := &rpc.ScoreDefinition{}
			if err := proto.Unmarshal(artifact.GetContents(), definition); err != nil {
				return err
			}

			// Check if ScoreDefinition.TargetResource matches with the supplied resource
			err := matchResourceWithTarget(definition.GetTargetResource(), resource, project)
			if err != nil {
				return err
			}

			defArtifacts = append(defArtifacts, artifact)
			return nil
		})

	if err != nil {
		return nil, err
	}

	return defArtifacts, nil
}

func matchResourceWithTarget(targetPattern *rpc.ResourcePattern, resourceName patterns.ResourceName, project string) error {
	targetPatternName, err := patterns.ParseResourcePattern(fmt.Sprintf("%s/%s", project, targetPattern.GetPattern()))
	if err != nil {
		return err
	}

	switch tp := targetPatternName.(type) {
	case patterns.SpecName:
		// Check if targetPattern and resourceName match in type
		r, ok := resourceName.(patterns.SpecName)
		if !ok {
			return fmt.Errorf("resource %s doesn't match target pattern %s", r, tp)
		}

		// Check if the individual entities match
		if tp.Name.ApiID != "-" && tp.Name.ApiID != r.Name.ApiID {
			return fmt.Errorf("api mismatch in resource %s and target pattern %v", resourceName.String(), targetPattern)
		}
		if tp.Name.VersionID != "-" && tp.Name.VersionID != r.Name.VersionID {
			return fmt.Errorf("version mismatch in resource %s and target pattern %v", resourceName.String(), targetPattern)
		}
		if tp.Name.SpecID != "-" && tp.Name.SpecID != r.Name.SpecID {
			return fmt.Errorf("spec mismatch in resource %s and target pattern %v", resourceName.String(), targetPattern)
		}
	case patterns.VersionName:
		// Check if targetPattern and resourceName match in type
		r, ok := resourceName.(patterns.VersionName)
		if !ok {
			return fmt.Errorf("resource %s doesn't match target pattern %s", r, tp)
		}

		// Check if the individual entities match
		if tp.Name.ApiID != "-" && tp.Name.ApiID != r.Name.ApiID {
			return fmt.Errorf("api mismatch in resource %s and target pattern %v", resourceName.String(), targetPattern)
		}
		if tp.Name.VersionID != "-" && tp.Name.VersionID != r.Name.VersionID {
			return fmt.Errorf("version mismatch in resource %s and target pattern %v", resourceName.String(), targetPattern)
		}
	case patterns.ApiName:
		// Check if targetPattern and resourceName match in type
		r, ok := resourceName.(patterns.ApiName)
		if !ok {
			return fmt.Errorf("resource %s doesn't match target pattern %s", r, tp)
		}

		// Check if the individual entities match
		if tp.Name.ApiID != "-" && tp.Name.ApiID != r.Name.ApiID {
			return fmt.Errorf("api mismatch in resource %s and target pattern %v", resourceName.String(), targetPattern)
		}
	default:
		return fmt.Errorf("unsupported resource type %T", targetPatternName)
	}

	// TODO: Filter check
	return nil
}

func CalculateScore(
	ctx context.Context,
	client connection.Client,
	defArtifact *rpc.Artifact,
	resource patterns.ResourceInstance) error {
	project := fmt.Sprintf("%s/locations/global", resource.ResourceName().Project())

	// Extract definition
	definition := &rpc.ScoreDefinition{}
	if err := proto.Unmarshal(defArtifact.GetContents(), definition); err != nil {
		return err
	}

	var takeAction bool

	// Fetch the to be generated score artifact (if present)
	artifactName := fmt.Sprintf("%s/artifacts/%s", resource.ResourceName().String(), scoreID(definition.GetId()))
	scoreArtifact, err := getArtifact(ctx, client, artifactName, false)
	if err != nil {
		// Calculate score if the score artifact doesn't exist
		if status.Code(err) == codes.NotFound {
			takeAction = true
		} else {
			return fmt.Errorf("failed to fetch artifact %q: %s", artifactName, err)
		}
	}

	// Calculate score if the definition has been updated
	if scoreArtifact != nil && scoreArtifact.GetUpdateTime().AsTime().Before(defArtifact.GetUpdateTime().AsTime()) {
		takeAction = true
	}

	// evaluate the expression and return a scoreValue
	result := processFormula(ctx, client, definition, resource, scoreArtifact, takeAction)
	if result.err != nil {
		return err
	}

	if result.needsUpdate {
		// generate a score proto from the scoreValue
		score, err := processScoreType(definition, result.value, project)
		if err != nil {
			return err
		}

		// TODO: Add dry_run flag

		err = uploadScore(ctx, client, resource, score)
		if err != nil {
			return err
		}
	}

	return nil
}

// Response returned after applying the score_expression on score_formula.artifact s.
type scoreResult struct {
	// Represents the value generated by the expression
	// Supported types are: int64, float64, bool
	value interface{}
	// Represents if the final scoreArtifact needs an updated
	// This is determined based on the timestamps of the existing scoreArtifact and the dependent artifacts in score_formula
	needsUpdate bool
	// Represents the error generated while applying the score_expression.
	err error
}

func processFormula(
	ctx context.Context,
	client connection.Client,
	definition *rpc.ScoreDefinition,
	resource patterns.ResourceInstance,
	scoreArtifact *rpc.Artifact,
	takeAction bool) scoreResult {
	// Apply score formula
	switch formula := definition.GetFormula().(type) {
	case *rpc.ScoreDefinition_ScoreFormula:
		return processScoreFormula(ctx, client, formula.ScoreFormula, resource, scoreArtifact, takeAction)
	case *rpc.ScoreDefinition_RollupFormula:
		return processRollUpFormula(ctx, client, formula.RollupFormula, resource, scoreArtifact, takeAction)
	default:
		return scoreResult{
			value:       nil,
			needsUpdate: false,
			err:         fmt.Errorf("invalid formula in ScoreDefinition: {%v} ", formula),
		}
	}
}

func processScoreFormula(
	ctx context.Context,
	client connection.Client,
	formula *rpc.ScoreFormula,
	resource patterns.ResourceInstance,
	scoreArtifact *rpc.Artifact,
	takeAction bool) scoreResult {
	extendedArtifact, err := patterns.SubstituteReferenceEntity(formula.GetArtifact().GetPattern(), resource.ResourceName())
	if err != nil {
		return scoreResult{
			value:       nil,
			needsUpdate: false,
			err:         fmt.Errorf("invalid score_formula.artifact.pattern: %s for {%v}, %s", formula.GetArtifact().GetPattern(), formula, err),
		}
	}
	if formula.GetScoreExpression() == "" {
		return scoreResult{
			value:       nil,
			needsUpdate: false,
			err:         fmt.Errorf("missing score_formula.score_expression for {%v}", formula),
		}
	}

	// Fetch the artifact
	artifact, err := getArtifact(ctx, client, extendedArtifact.String(), true)
	if err != nil {
		return scoreResult{
			value:       nil,
			needsUpdate: false,
			err:         fmt.Errorf("failed to fetch artifact %s: %s", extendedArtifact.String(), err),
		}
	}

	// Update required tells the calling function if the score artifact needs to be updated
	updateRequired := takeAction || scoreArtifact.GetUpdateTime().AsTime().Before(artifact.GetUpdateTime().AsTime())

	// Apply the scoreExpression by default. This value will be required by the rollup_formula in the case where
	// another formula from rollup_formula.score_formulas makes the score outdated.

	// Convert artifact contents to map[string]interface{}
	artifactMap, err := getMap(artifact.GetContents(), artifact.GetMimeType())
	if err != nil {
		return scoreResult{
			value:       nil,
			needsUpdate: false,
			err:         err,
		}
	}

	// Apply the score_expression
	value, err := evaluateScoreExpression(formula.GetScoreExpression(), artifactMap)
	if err != nil {
		return scoreResult{
			value:       nil,
			needsUpdate: false,
			err:         err,
		}
	}
	return scoreResult{
		value:       value,
		needsUpdate: updateRequired,
		err:         nil,
	}
}

func processRollUpFormula(
	ctx context.Context,
	client connection.Client,
	formula *rpc.RollUpFormula,
	resource patterns.ResourceInstance,
	scoreArtifact *rpc.Artifact,
	takeAction bool) scoreResult {
	// Validate required fields
	if len(formula.GetScoreFormulas()) == 0 {
		return scoreResult{
			value:       nil,
			needsUpdate: false,
			err:         fmt.Errorf("missing rollup_formula.score_formulas in {%v}", formula),
		}
	}
	if formula.GetRollupExpression() == "" {
		return scoreResult{
			value:       nil,
			needsUpdate: false,
			err:         fmt.Errorf("missing rollup_formula.rollup_expression in {%v}", formula),
		}
	}

	// Update required tells the calling function if the score artifact needs to be updated
	updateRequired := takeAction
	rollUpMap := make(map[string]interface{}, 0)
	for _, f := range formula.GetScoreFormulas() {
		result := processScoreFormula(ctx, client, f, resource, scoreArtifact, takeAction)
		if result.err != nil {
			return scoreResult{
				value:       nil,
				needsUpdate: false,
				err:         fmt.Errorf("error processing rollup_formula.score_formulas: %s", result.err),
			}
		}

		refId := f.GetReferenceId()
		if refId == "" {
			return scoreResult{
				value:       nil,
				needsUpdate: false,
				err:         fmt.Errorf("missing reference_id for score_formula {%v}", f),
			}
		}
		if strings.Contains(refId, "-") {
			return scoreResult{
				value:       nil,
				needsUpdate: false,
				err:         fmt.Errorf("invalid reference_id for score_formula {%v}: cannot contain '-'", f),
			}
		}
		rollUpMap[refId] = result.value

		updateRequired = updateRequired || result.needsUpdate
	}

	// Apply the rollup_expression
	if updateRequired {
		value, err := evaluateScoreExpression(formula.GetRollupExpression(), rollUpMap)
		if err != nil {
			return scoreResult{
				value:       nil,
				needsUpdate: false,
				err:         err,
			}
		}
		return scoreResult{
			value:       value,
			needsUpdate: true,
			err:         nil,
		}
	}

	return scoreResult{
		value:       nil,
		needsUpdate: false,
		err:         nil,
	}
}

func processScoreType(definition *rpc.ScoreDefinition, scoreValue interface{}, project string) (*rpc.Score, error) {
	// Initialize Score proto
	score := &rpc.Score{
		Id:             fmt.Sprintf("score-%s", definition.GetId()),
		Kind:           "Score",
		DisplayName:    definition.GetDisplayName(),
		Description:    definition.GetDescription(),
		Uri:            definition.GetUri(),
		UriDisplayName: definition.GetUriDisplayName(),
		DefinitionName: fmt.Sprintf("%s/artifacts/%s", project, definition.GetId()),
	}

	// Set the Value field according to the type
	switch definition.GetType().(type) {
	case *rpc.ScoreDefinition_Integer:
		// Score proto expects int32 type
		var value int32

		// Convert scoreValue to appropriate type
		// evaluateScoreExpression can return either a float or int value.
		// Both are valid for an integer.
		switch v := scoreValue.(type) {
		case int64:
			value = int32(v)
		case float64:
			value = int32(v)
		default:
			return nil, fmt.Errorf("failed typecheck for output: expected either int64 or float64 got %s (type: %T)", v, v)
		}

		// Check that the scoreValue is within min/max limits
		configuredMin := definition.GetInteger().GetMinValue() // 0 if not set
		configuredMax := definition.GetInteger().GetMaxValue() // 0 if not set
		if value < configuredMin {
			return nil, fmt.Errorf("evaluated score value(%d) cannot be less than the configured min_value (%d)", value, configuredMin)
		}
		if value > configuredMax {
			return nil, fmt.Errorf("evaluated score value(%d) cannot be greater than the configured max_value (%d)", value, configuredMax)
		}

		// Populate Value field in Score proto
		score.Value = &rpc.Score_IntegerValue{
			IntegerValue: &rpc.IntegerValue{
				Value:    value,
				MinValue: configuredMin,
				MaxValue: configuredMax,
			},
		}

		// Populate the severity field according to Thresholds
		for _, t := range definition.GetInteger().GetThresholds() {
			if value >= t.GetRange().GetMin() && value <= t.GetRange().GetMax() {
				score.Severity = t.GetSeverity()
				break
			}
		}

	case *rpc.ScoreDefinition_Percent:
		// Score proto expects float32 type
		var value float32

		// Convert scoreValue to appropriate type
		// evaluateScoreExpression can return either a float or int value.
		// Both are valid for an integer.
		switch v := scoreValue.(type) {
		case int64:
			value = float32(v)
		case float64:
			value = float32(v)
		default:
			return nil, fmt.Errorf("failed typecheck for output: expected either int64 or float64 got %s (type: %T)", v, v)
		}

		// Check that the scoreValue is within min/max limits
		if value < 0 {
			return nil, fmt.Errorf("evaluated score value(%f) cannot be less than 0", value)
		}
		if value > 100 {
			return nil, fmt.Errorf("evaluated score value(%f) cannot be greater than 100", value)
		}

		// Populate Value field in Score proto
		score.Value = &rpc.Score_PercentValue{
			PercentValue: &rpc.PercentValue{
				Value: value,
			},
		}

		// Populate the severity field according to Thresholds
		for _, t := range definition.GetPercent().GetThresholds() {
			if value >= float32(t.GetRange().GetMin()) && value <= float32(t.GetRange().GetMax()) {
				score.Severity = t.GetSeverity()
				break
			}
		}

	case *rpc.ScoreDefinition_Boolean:
		// Convert scoreValue to appropriate type
		boolVal, ok := scoreValue.(bool)
		if !ok {
			return nil, fmt.Errorf("failed typecheck for output: expected bool")
		}

		var displayValue string
		if t := definition.GetBoolean().GetDisplayTrue(); boolVal && t != "" {
			displayValue = t
		} else if f := definition.GetBoolean().GetDisplayFalse(); !boolVal && f != "" {
			displayValue = f
		} else {
			displayValue = strconv.FormatBool(boolVal)
		}

		// Populate Value field in Score proto
		score.Value = &rpc.Score_BooleanValue{
			BooleanValue: &rpc.BooleanValue{
				Value:        boolVal,
				DisplayValue: displayValue,
			},
		}

		// Populate the severity field according to Thresholds
		for _, t := range definition.GetBoolean().GetThresholds() {
			if t.Value == boolVal {
				score.Severity = t.Severity
			}
		}
	}

	return score, nil
}

func uploadScore(ctx context.Context, client connection.Client, resource patterns.ResourceInstance, score *rpc.Score) error {
	artifactBytes, err := proto.Marshal(score)
	if err != nil {
		return err
	}
	artifact := &rpc.Artifact{
		Name:     fmt.Sprintf("%s/artifacts/%s", resource.ResourceName().String(), score.GetId()),
		Contents: artifactBytes,
		MimeType: patch.MimeTypeForKind("Score"),
	}
	log.Debugf(ctx, "Uploading %s", artifact.GetName())
	if err = core.SetArtifact(ctx, client, artifact); err != nil {
		return fmt.Errorf("failed to save artifact %s: %s", artifact.GetName(), err)
	}

	return nil
}

func getArtifact(ctx context.Context, client connection.Client, artifactPattern string, getContents bool) (*rpc.Artifact, error) {
	artifactName, err := names.ParseArtifact(artifactPattern)
	if err != nil {
		return nil, fmt.Errorf("invalid artifact pattern %q: %s", artifactPattern, err)
	}

	gotArtifact := &rpc.Artifact{}
	err = core.GetArtifact(ctx, client, artifactName, true, func(artifact *rpc.Artifact) error {
		gotArtifact = artifact
		return nil
	})
	if err != nil {
		return nil, err
	}
	return gotArtifact, nil
}