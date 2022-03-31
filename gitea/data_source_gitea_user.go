package gitea

import (
	"fmt"
	"log"
	"strings"

	"code.gitea.io/sdk/gitea"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGiteaUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGitlabUserRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"full_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_admin": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"avatar_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"language": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_login": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGitlabUserRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gitea.Client)

	var user *gitea.User
	var err error

	log.Printf("[INFO] Reading Gitea user")

	usernameData, usernameOk := d.GetOk("username")

	if !usernameOk {
		user, err = client.GetMyUserInfo()
	} else {
		username := strings.ToLower(usernameData.(string))

		user, err = client.GetUserInfo(username)
		if err != nil {
			return err
		}
	}

	d.Set("id", user.ID)
	d.Set("username", user.UserName)
	d.Set("email", user.Email)
	d.Set("full_name", user.FullName)
	d.Set("is_admin", user.IsAdmin)
	d.Set("created", user.Created)
	d.Set("avatar_url", user.AvatarURL)
	d.Set("last_login", user.LastLogin)
	d.Set("language", user.Language)

	d.SetId(fmt.Sprintf("%d", user.ID))

	return nil
}
