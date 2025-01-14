package virtual

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jfrog/terraform-provider-artifactory/v6/pkg/artifactory/resource/repository"
	"github.com/jfrog/terraform-provider-artifactory/v6/pkg/utils"
)

func ResourceArtifactoryVirtualNugetRepository() *schema.Resource {

	const packageType = "nuget"

	var nugetVirtualSchema = utils.MergeSchema(BaseVirtualRepoSchema, map[string]*schema.Schema{
		"force_nuget_authentication": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Force basic authentication credentials in order to use this repository.",
		},
	}, repository.RepoLayoutRefSchema("virtual", packageType))

	type NugetVirtualRepositoryParams struct {
		VirtualRepositoryBaseParams
		ForceNugetAuthentication bool `json:"forceNugetAuthentication"`
	}

	var unpackNugetVirtualRepository = func(s *schema.ResourceData) (interface{}, string, error) {
		d := &utils.ResourceData{s}

		repo := NugetVirtualRepositoryParams{
			VirtualRepositoryBaseParams: UnpackBaseVirtRepo(s, packageType),
			ForceNugetAuthentication:    d.GetBool("force_nuget_authentication", false),
		}
		repo.PackageType = packageType
		return &repo, repo.Key, nil
	}

	return repository.MkResourceSchema(nugetVirtualSchema, repository.DefaultPacker(nugetVirtualSchema), unpackNugetVirtualRepository, func() interface{} {
		return &NugetVirtualRepositoryParams{
			VirtualRepositoryBaseParams: VirtualRepositoryBaseParams{
				Rclass:      "virtual",
				PackageType: packageType,
			},
		}
	})
}
