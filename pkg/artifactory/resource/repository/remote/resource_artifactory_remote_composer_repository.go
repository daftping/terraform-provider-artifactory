package remote

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/jfrog/terraform-provider-artifactory/v6/pkg/artifactory/resource/repository"
	"github.com/jfrog/terraform-provider-artifactory/v6/pkg/utils"
)

type ComposerRemoteRepo struct {
	RemoteRepositoryBaseParams
	RemoteRepositoryVcsParams
	ComposerRegistryUrl string `json:"composerRegistryUrl"`
}

func ResourceArtifactoryRemoteComposerRepository() *schema.Resource {
	const packageType = "composer"

	var composerRemoteSchema = utils.MergeSchema(BaseRemoteRepoSchema, VcsRemoteRepoSchema, map[string]*schema.Schema{
		"composer_registry_url": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "https://packagist.org",
			ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			Description:  `(Optional) Proxy remote Composer repository. Default value is "https://packagist.org".`,
		},
	}, repository.RepoLayoutRefSchema("remote", packageType))

	var unpackComposerRemoteRepo = func(s *schema.ResourceData) (interface{}, string, error) {
		d := &utils.ResourceData{s}
		repo := ComposerRemoteRepo{
			RemoteRepositoryBaseParams: UnpackBaseRemoteRepo(s, packageType),
			RemoteRepositoryVcsParams:  UnpackVcsRemoteRepo(s),
			ComposerRegistryUrl:        d.GetString("composer_registry_url", false),
		}
		return repo, repo.Id(), nil
	}

	return repository.MkResourceSchema(composerRemoteSchema, repository.DefaultPacker(composerRemoteSchema), unpackComposerRemoteRepo, func() interface{} {
		repoLayout, _ := utils.GetDefaultRepoLayoutRef("remote", packageType)()
		return &ComposerRemoteRepo{
			RemoteRepositoryBaseParams: RemoteRepositoryBaseParams{
				Rclass:              "remote",
				PackageType:         packageType,
				RemoteRepoLayoutRef: repoLayout.(string),
			},
		}
	})
}
