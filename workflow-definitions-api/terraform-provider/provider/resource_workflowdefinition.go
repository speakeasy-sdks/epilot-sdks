package provider

import (
	"context"
	sdk "epilot-workflows-definition/sdk"
	operations "epilot-workflows-definition/sdk/pkg/models/operations"
	shared "epilot-workflows-definition/sdk/pkg/models/shared"
	"fmt"
	diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWorkflowDefinition() *schema.Resource {
	return &schema.Resource{CreateContext: resourceWorkflowDefinitionCreate, ReadContext: resourceWorkflowDefinitionRead, UpdateContext: resourceWorkflowDefinitionUpdate, DeleteContext: resourceWorkflowDefinitionDelete, Schema: map[string]*schema.Schema{
		"assigned_to":     &schema.Schema{Type: schema.TypeList, Optional: true, Elem: &schema.Schema{Type: schema.TypeString, Optional: true}},
		"closing_reasons": &schema.Schema{Type: schema.TypeSet, Optional: true, Elem: &schema.Resource{Schema: map[string]*schema.Schema{"id": &schema.Schema{Type: schema.TypeString, Required: true}}}},
		"creation_time":   &schema.Schema{Type: schema.TypeString, Optional: true, Description: "ISO String Date & Time"},
		"description":     &schema.Schema{Type: schema.TypeString, Optional: true},
		"due_date":        &schema.Schema{Type: schema.TypeString, Optional: true},
		"dynamic_due_date": &schema.Schema{Type: schema.TypeSet, MaxItems: 1, Optional: true, Description: "set a Duedate for a step then a specific", Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"action_type_condition": &schema.Schema{Type: schema.TypeString, Required: true},
			"number_of_units":       &schema.Schema{Type: schema.TypeFloat, Required: true},
			"step_id":               &schema.Schema{Type: schema.TypeString, Optional: true},
			"time_period":           &schema.Schema{Type: schema.TypeString, Required: true},
		}}},
		"enable_ecp_workflow": &schema.Schema{Type: schema.TypeBool, Optional: true, Description: "Indicates whether this workflow is available for End Customer Portal or not. By default it's not."},
		"flow":                &schema.Schema{Type: schema.TypeList, Required: true, Elem: &schema.Schema{Type: schema.TypeMap, Elem: &schema.Schema{Type: schema.TypeString}, Optional: true}},
		"id":                  &schema.Schema{Type: schema.TypeString, Optional: true},
		"last_update_time":    &schema.Schema{Type: schema.TypeString, Optional: true, Description: "ISO String Date & Time"},
		"name":                &schema.Schema{Type: schema.TypeString, Required: true},
		"update_entity_attributes": &schema.Schema{Type: schema.TypeSet, Optional: true, Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"source": &schema.Schema{Type: schema.TypeString, Required: true},
			"target": &schema.Schema{Type: schema.TypeSet, MaxItems: 1, Required: true, Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"entity_attribute": &schema.Schema{Type: schema.TypeString, Required: true},
				"entity_schema":    &schema.Schema{Type: schema.TypeString, Required: true},
			}}},
		}}},
		"user_ids": &schema.Schema{Type: schema.TypeList, Optional: true, Description: "This field is deprecated. Please use assignedTo", Elem: &schema.Schema{Type: schema.TypeFloat, Optional: true}},
	}}
}

func resourceWorkflowDefinitionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*sdk.SDK)

	assignedTos := d.Get("assigned_to").([]string)

	closingReasonsItems := d.Get("closing_reasons").([]map[string]interface{})

	closingReasons := make([]shared.ClosingReasonID, 0)

	for _, closingReasonsRawItem := range closingReasonsItems {
		id := closingReasonsRawItem["id"].(string)
		closingReasonsItem := shared.ClosingReasonID{ID: id}

		closingReasons = append(closingReasons, closingReasonsItem)

	}

	creationTime := d.Get("creation_time").(string)

	description := d.Get("description").(string)

	dueDate := d.Get("due_date").(string)

	actionTypeCondition := d.Get("action_type_condition").(shared.DynamicDueDateActionTypeConditionEnum)

	numberOfUnits := d.Get("number_of_units").(float64)

	stepId := d.Get("step_id").(string)

	timePeriod := d.Get("time_period").(shared.DynamicDueDateTimePeriodEnum)

	dynamicDueDate := shared.DynamicDueDate{
		ActionTypeCondition: actionTypeCondition,
		NumberOfUnits:       numberOfUnits,
		StepID:              &stepId,
		TimePeriod:          timePeriod,
	}

	enableECPWorkflow := d.Get("enable_ecp_workflow").(bool)

	flows := d.Get("flow").([]interface{})

	id2 := d.Get("id").(string)

	lastUpdateTime := d.Get("last_update_time").(string)

	name := d.Get("name").(string)

	updateEntityAttributesItems := d.Get("update_entity_attributes").([]map[string]interface{})

	updateEntityAttributes := make([]shared.UpdateEntityAttributes, 0)

	for _, updateEntityAttributesRawItem := range updateEntityAttributesItems {
		source := updateEntityAttributesRawItem["source"].(shared.UpdateEntityAttributesSourceEnum)
		entityAttribute := updateEntityAttributesRawItem["entity_attribute"].(string)
		entitySchema := updateEntityAttributesRawItem["entity_schema"].(string)
		target := shared.UpdateEntityAttributesTarget{
			EntityAttribute: entityAttribute,
			EntitySchema:    entitySchema,
		}
		updateEntityAttributesItem := shared.UpdateEntityAttributes{
			Source: source,
			Target: target,
		}

		updateEntityAttributes = append(updateEntityAttributes, updateEntityAttributesItem)

	}

	userIds := d.Get("user_ids").([]float64)

	Request := shared.WorkflowDefinition{
		AssignedTo:             assignedTos,
		ClosingReasons:         closingReasons,
		CreationTime:           &creationTime,
		Description:            &description,
		DueDate:                &dueDate,
		DynamicDueDate:         &dynamicDueDate,
		EnableEcpWorkflow:      &enableECPWorkflow,
		Flow:                   flows,
		ID:                     &id2,
		LastUpdateTime:         &lastUpdateTime,
		Name:                   name,
		UpdateEntityAttributes: updateEntityAttributes,
		UserIds:                userIds,
	}

	req := operations.CreateDefinitionRequest{Request: Request}

	res, err := c.Workflows.CreateDefinition(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}
	// not implemented yet: SDK call doesn't return object: TODO invoke additional GET to get entity WorkflowDefinition

	_ = res

	return nil
}

func resourceWorkflowDefinitionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*sdk.SDK)

	definitionId := d.Get("definition_id").(string)

	PathParams := operations.GetDefinitionPathParams{DefinitionID: definitionId}

	req := operations.GetDefinitionRequest{PathParams: PathParams}

	res, err := c.Workflows.GetDefinition(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}
	// not implemented yet: SDK call doesn't return object: TODO invoke additional GET to get entity WorkflowDefinition

	_ = res

	return nil
}

func resourceWorkflowDefinitionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*sdk.SDK)

	definitionId := d.Get("definition_id").(string)

	PathParams := operations.UpdateDefinitionPathParams{DefinitionID: definitionId}

	assignedTos := d.Get("assigned_to").([]string)

	closingReasonsItems := d.Get("closing_reasons").([]map[string]interface{})

	closingReasons := make([]shared.ClosingReasonID, 0)

	for _, closingReasonsRawItem := range closingReasonsItems {
		id := closingReasonsRawItem["id"].(string)
		closingReasonsItem := shared.ClosingReasonID{ID: id}

		closingReasons = append(closingReasons, closingReasonsItem)

	}

	creationTime := d.Get("creation_time").(string)

	description := d.Get("description").(string)

	dueDate := d.Get("due_date").(string)

	actionTypeCondition := d.Get("action_type_condition").(shared.DynamicDueDateActionTypeConditionEnum)

	numberOfUnits := d.Get("number_of_units").(float64)

	stepId := d.Get("step_id").(string)

	timePeriod := d.Get("time_period").(shared.DynamicDueDateTimePeriodEnum)

	dynamicDueDate := shared.DynamicDueDate{
		ActionTypeCondition: actionTypeCondition,
		NumberOfUnits:       numberOfUnits,
		StepID:              &stepId,
		TimePeriod:          timePeriod,
	}

	enableECPWorkflow := d.Get("enable_ecp_workflow").(bool)

	flows := d.Get("flow").([]interface{})

	id2 := d.Get("id").(string)

	lastUpdateTime := d.Get("last_update_time").(string)

	name := d.Get("name").(string)

	updateEntityAttributesItems := d.Get("update_entity_attributes").([]map[string]interface{})

	updateEntityAttributes := make([]shared.UpdateEntityAttributes, 0)

	for _, updateEntityAttributesRawItem := range updateEntityAttributesItems {
		source := updateEntityAttributesRawItem["source"].(shared.UpdateEntityAttributesSourceEnum)
		entityAttribute := updateEntityAttributesRawItem["entity_attribute"].(string)
		entitySchema := updateEntityAttributesRawItem["entity_schema"].(string)
		target := shared.UpdateEntityAttributesTarget{
			EntityAttribute: entityAttribute,
			EntitySchema:    entitySchema,
		}
		updateEntityAttributesItem := shared.UpdateEntityAttributes{
			Source: source,
			Target: target,
		}

		updateEntityAttributes = append(updateEntityAttributes, updateEntityAttributesItem)

	}

	userIds := d.Get("user_ids").([]float64)

	Request := shared.WorkflowDefinition{
		AssignedTo:             assignedTos,
		ClosingReasons:         closingReasons,
		CreationTime:           &creationTime,
		Description:            &description,
		DueDate:                &dueDate,
		DynamicDueDate:         &dynamicDueDate,
		EnableEcpWorkflow:      &enableECPWorkflow,
		Flow:                   flows,
		ID:                     &id2,
		LastUpdateTime:         &lastUpdateTime,
		Name:                   name,
		UpdateEntityAttributes: updateEntityAttributes,
		UserIds:                userIds,
	}

	req := operations.UpdateDefinitionRequest{
		PathParams: PathParams,
		Request:    Request,
	}

	res, err := c.Workflows.UpdateDefinition(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}
	// not implemented yet: SDK call doesn't return object: TODO invoke additional GET to get entity WorkflowDefinition

	_ = res

	return nil
}

func resourceWorkflowDefinitionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*sdk.SDK)

	definitionId := d.Get("definition_id").(string)

	PathParams := operations.DeleteDefinitionPathParams{DefinitionID: definitionId}

	req := operations.DeleteDefinitionRequest{PathParams: PathParams}

	res, err := c.Workflows.DeleteDefinition(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}
	if res.StatusCode != 204 {
		return diag.FromErr(fmt.Errorf("unexpected status code %d", res.StatusCode))
	}
	d.SetId("")
	return nil
}
