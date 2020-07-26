package appd

import (
	"fmt"
	"github.com/HarryEMartland/appd-terraform-provider/appd/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
)

func resourceAction() *schema.Resource {
	return &schema.Resource{
		Create: resourceActionCreate,
		Read:   resourceActionRead,
		Update: resourceActionUpdate,
		Delete: resourceActionDelete,

		Schema: map[string]*schema.Schema{
			"application_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"action_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {

					validValues := []string{
						"SMS",
						"EMAIL",
						"CUSTOM_EMAIL",
						"THREAD_DUMP",
						"HTTP_REQUEST",
						"CUSTOM",
						"RUN_SCRIPT_ON_NODES",
						"DIAGNOSE_BUSINESS_TRANSACTIONS",
						"CREATE_UPDATE_JIRA",
					}

					strVal := val.(string)

					if !contains(validValues, strVal) {
						errs = append(errs, fmt.Errorf("%s is not a valid value for %s (%v)", strVal, key, validValues))
					}

					return
				},
			},
			"emails": {
				Type:     schema.TypeList,
				Required: true, //required untill we do other types
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}
func resourceActionCreate(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	applicationId := d.Get("application_id").(int)

	action := createAction(d)

	updatedHealthRule, err := appdClient.CreateAction(&action, applicationId)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(updatedHealthRule.ID))

	return resourceActionRead(d, m)
}

func createAction(d *schema.ResourceData) client.Action {

	name := d.Get("name").(string)
	actionType := d.Get("action_type").(string)
	emails := d.Get("emails").([]interface{})

	healthRule := client.Action{
		Name:       name,
		ActionType: actionType,
		Emails:     emails,
	}
	return healthRule
}

func updateAction(d *schema.ResourceData, action client.Action)  {
	d.Set("name", action.Name)
	d.Set("action_type", action.ActionType)
	d.Set("emails", action.Emails)
}

func resourceActionRead(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	applicationId := d.Get("application_id").(int)
	id := d.Id()

	actionId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	action, err := appdClient.GetAction(actionId, applicationId)
	if err != nil {
		return err
	}

	updateAction(d, *action)

	return nil
}

func resourceActionUpdate(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	applicationId := d.Get("application_id").(int)

	healthRule := createAction(d)

	healthRuleId, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	healthRule.ID = healthRuleId

	_, err = appdClient.UpdateAction(&healthRule, applicationId)
	if err != nil {
		return err
	}

	return resourceActionRead(d, m)
}

func resourceActionDelete(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	applicationId := d.Get("application_id").(int)
	id := d.Id()

	actionId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = appdClient.DeleteAction(applicationId, actionId)
	if err != nil {
		return err
	}

	return nil
}